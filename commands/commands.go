package commands

import (
	"flag"
	"io"
	"log"
	"sync"
)

type sub func()
type deferred func()
type commandLineHandler func() bool

type deferrerMaker interface {
	MakeDeferrer() deferred
}

type makeSub func(*flag.FlagSet, []string) (sub, error, []string)
type setupFunc func(*flag.FlagSet) (sub, io.Writer)

type submap struct {
	submap              map[string]makeSub
	setups              map[string]setupFunc
	commandLineHandlers []commandLineHandler
	keys                []string
}

func (sm submap) Len() int {
	return len(sm.keys)
}
func (sm submap) Swap(i, j int) {
	sm.keys[i], sm.keys[j] = sm.keys[j], sm.keys[i]
}
func (sm submap) Less(i, j int) bool {
	return sm.keys[i] < sm.keys[j]
}

var (
	commands = struct {
		mutex  sync.Mutex
		mapsub submap
	}{
		mapsub: submap{
			submap: make(map[string]makeSub),
			setups: make(map[string]setupFunc),
		},
	}

	defaults = struct {
		mutex  sync.Mutex
		mapdef map[string]deferrerMaker
	}{
		mapdef: make(map[string]deferrerMaker),
	}
)

func AddCommandLine(hfunc func(*flag.FlagSet) commandLineHandler) {
	s := hfunc(flag.CommandLine) // NB global flag.CommandLine
	if s == nil {
		return
	}
	commands.mutex.Lock()
	defer commands.mutex.Unlock()
	commands.mapsub.commandLineHandlers = append(
		commands.mapsub.commandLineHandlers, s)
}

func AddFlaggedCommand(name string, sfunc setupFunc) {
	commands.mutex.Lock()
	defer commands.mutex.Unlock()
	commands.mapsub.setups[name] = sfunc
	commands.mapsub.keys = append(commands.mapsub.keys, name)
}

func setupFlagset(name string, sfunc setupFunc) (*flag.FlagSet, sub, io.Writer) {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	run, output := sfunc(fs)
	return fs, run, output
}

func setup(name string, sfunc setupFunc, arguments []string) (sub, error, []string) {
	fs, run, output := setupFlagset(name, sfunc)
	err := fs.Parse(arguments)
	if err == nil && output != nil {
		fs.SetOutput(output)
	}
	return run, err, fs.Args()
}

func AddCommand(name string, makes makeSub) {
	commands.mutex.Lock()
	defer commands.mutex.Unlock()
	commands.mapsub.submap[name] = makes
	commands.mapsub.keys = append(commands.mapsub.keys, name)
}

func AddDefault(name string, def deferrerMaker) {
	defaults.mutex.Lock()
	defer defaults.mutex.Unlock()
	defaults.mapdef[name] = def
}

func Defaults() deferred {
	defaults.mutex.Lock()
	defer defaults.mutex.Unlock()
	finish := []deferred{}
	for _, ding := range defaults.mapdef {
		if fin := ding.MakeDeferrer(); fin != nil {
			finish = append(finish, fin)
		}
	}
	return func() {
		for _, fin := range finish {
			fin()
		}
	}
}

func parseCommand(subs []sub, args []string) ([]sub, bool) {
	if len(args) == 0 || args[0] == "" {
		return subs, false
	}
	name := args[0]
	if ctor, ok := commands.mapsub.setups[name]; ok {
		if sub, err, nextargs := setup(name, ctor, args[1:]); err == nil {
			return parseCommand(append(subs, sub), nextargs)
		}
	} else if ctor, ok := commands.mapsub.submap[name]; ok {
		fs := flag.NewFlagSet(name, flag.ContinueOnError)
		if sub, err, nextargs := ctor(fs, args[1:]); err == nil {
			return parseCommand(append(subs, sub), nextargs)
		}
		// else { /* log.Printf("%s: %s\n", name, err)
		// printed already by flag package // */ }
	} else {
		log.Fatalf("%s: No such command\n", name)
	}
	return subs, true
}

func parseCommands() ([]sub, bool) {
	commands.mutex.Lock()
	defer commands.mutex.Unlock()
	return parseCommand([]sub{}, flag.Args())
}

// true is when to abort
func ArgCommands() bool {
	subs, errd := parseCommands()
	if errd {
		return true
	}
	commandLineHandlers := commands.mapsub.commandLineHandlers
	if len(commandLineHandlers) > 0 {
		stop := false
		for _, clh := range commandLineHandlers {
			if clh() {
				stop = true
			}
		}
		if stop {
			return true
		}
	}
	if len(subs) == 0 {
		return false
	}
	for _, sub := range subs {
		sub()
	}
	return true
}

type loggerWriter struct {
	*log.Logger
}

func (lw *loggerWriter) Write(p []byte) (int, error) {
	lw.Logger.Printf("%s", p)
	return len(p), nil
}
