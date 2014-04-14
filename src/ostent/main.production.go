// +build production

package main
import (
	"ostential"

	"os"
	"log"
	"fmt"
	"flag"
	"time"
 	"syscall"
	"runtime"
	"strings"
	"net/url"
	"net/http"
	"math/rand"
	"path/filepath"

	"github.com/codegangsta/martini"
	"github.com/rcrowley/goagain"

	"github.com/inconshreveable/go-update"
)

func init() {
// 	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix(fmt.Sprintf("[%d] ", os.Getpid()))
}

func update_loop() {
	const dlimit = time.Hour
	delta := time.Duration(rand.Int63n(int64(dlimit)))
	for {
		select {
		case <-time.After(time.Hour * 1 + delta): // 1.5 +- 0.5 h
			if update_once(true) {
				break
			}
		}
	}
}

func newer_version() string {
	// 1. https://github.com/rzab/ostent/releases/latest # redirects, NOT followed
	// 2. https://github.com/rzab/ostent/releases/vX.Y.Z # Redirect location
	// 3. return "vX.Y.Z" # basename of the location

	type redirected struct {
		error
		url url.URL
	}
	checkRedirect := func(req *http.Request, _via []*http.Request) error {
		return redirected{url: *req.URL,}
	}

	client := &http.Client{CheckRedirect: checkRedirect,}
	resp, err := client.Get("https://github.com/rzab/ostent/releases/latest")
	if err == nil {
		resp.Body.Close()
		return ""
	}
	urlerr, ok := err.(*url.Error)
	if !ok {
		return ""
	}
	resp.Body.Close()
	redir, ok := urlerr.Err.(redirected)
	if !ok {
		return ""
	}
	return filepath.Base(redir.url.Path)
}

const current_version = "v0.1.3"
// const update_from = current_version[1:]
func update_once(kill bool) bool {

	mach := runtime.GOARCH
	if mach == "amd64" {
		mach = "x86_64"
	} else if mach == "386" {
		mach = "i686"
	}
	new_version := newer_version()
	if new_version == "" || new_version == current_version {
		return false
	}
// 	url := fmt.Sprintf("https://ostrost.com"+ "/ostent/releases/%s/%s %s/newer",    update_from, strings.Title(runtime.GOOS), mach) // before v0.1.3
	url := fmt.Sprintf("https://github.com/rzab/ostent/releases/download/%s/%s.%s", new_version, strings.Title(runtime.GOOS), mach)

	err, _ := update.FromUrl(url) // , snderr
	if err != nil ||  err != nil {
		// log.Printf("Update failed: %v|%v\n", err, snderr)
		return false
	}
	log.Println("Successfull UPDATE, relaunching")
	if kill {
		syscall.Kill(os.Getpid(), syscall.SIGUSR2)
	}
	return true
}

func main() {
	updatelater := flag.Bool("updatelater", false, "Update later")
	flag.Parse()

	had_update := false
	if !*updatelater && os.Getenv("GOAGAIN_PPID") == "" { // not after gone again
		log.Println("Initial check for updates; run with -updatelater to disable")
		had_update = update_once(false)
	}

	martini.Env = martini.Prod
	if !had_update { // start the background routine unless just had an update and gonna relaunch anyway
		go ostential.Loop()
	}

	listen, err := goagain.Listener()
	if err != nil {
		listen, err = ostential.Listen()
		if err != nil {
			if _, ok := err.(ostential.FlagError); ok {
				flag.Usage()
				os.Exit(2)
				return
			}
			log.Fatalln(err)
		}

		if had_update { // goagain
			go func() {
				time.Sleep(time.Second) // not before goagain.Wait
				syscall.Kill(os.Getpid(), syscall.SIGUSR2)
				// goagain.ForkExec(listen)
			}()
		} else {
			go update_loop()
			go ostential.Serve(listen, ostential.LogOne, nil)
		}

	} else {
		go update_loop()
		go ostential.Serve(listen, ostential.LogOne, nil)

		if err := goagain.Kill(); err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := goagain.Wait(listen); err != nil { // signals before won't be catched
		log.Fatalln(err)
	}
	if err := listen.Close(); err != nil {
		log.Fatalln(err)
	}
	time.Sleep(time.Second)
}
