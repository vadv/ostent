package ostent

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/justinas/alice"
	"github.com/ostrost/ostent/share/assets"
)

type bindValue struct {
	string
	defport string // const
	Host    string // available after flag.Parse()
	Port    string // available after flag.Parse()
}

func newBind(defstring, defport string) bindValue {
	bv := bindValue{defport: defport}
	bv.Set(defstring)
	return bv
}

// satisfying flag.Value interface
func (bv bindValue) String() string { return string(bv.string) }
func (bv *bindValue) Set(input string) error {
	if input == "" {
		bv.Port = bv.defport
	} else {
		if !strings.Contains(input, ":") {
			input = ":" + input
		}
		var err error
		bv.Host, bv.Port, err = net.SplitHostPort(input)
		if err != nil {
			return err
		}
		if bv.Host == "*" {
			bv.Host = ""
		} else if bv.Port == "127" {
			bv.Host = "127.0.0.1"
			bv.Port = bv.defport
		}
		if _, err = net.LookupPort("tcp", bv.Port); err != nil {
			if bv.Host != "" {
				return err
			}
			bv.Host, bv.Port = bv.Port, bv.defport
		}
	}

	bv.string = bv.Host + ":" + bv.Port
	return nil
}

// OstentBindFlag is a bindValue hoding the ostent bind address.
var OstentBindFlag = newBind(":8050", "8050")

// CollectdBindFlag is a bindValue hoding the ostent collectd bind address.
// var CollectdBindFlag = newBind("", "8051") // "" by default meaning DO NOT BIND
func init() {
	flag.Var(&OstentBindFlag, "b", "short for bind")
	flag.Var(&OstentBindFlag, "bind", "Bind address")
	// flag.Var(&CollectdBindFlag, "collectdb",    "short for collectdbind")
	// flag.Var(&CollectdBindFlag, "collectdbind", "Bind address for collectd receiving")
}

// Muxmap is a type of a map of pattern to HandlerFunc.
type Muxmap map[string]http.HandlerFunc

type Server struct {
	http.Server
	Access *logger
	MUX    *TrieServeMux
	Chain  alice.Chain
}

func NewServer(listener net.Listener, production bool) *Server {
	access := newLogged(production, log.New(os.Stdout, "", 0))
	recovery := recovery(production)
	chain := alice.New(
		access.Constructor,
		recovery.Constructor,
	)
	mux := NewMux(recovery, chain.Then)
	return &Server{
		Server: http.Server{Addr: listener.Addr().String(), Handler: mux},
		Access: access,
		Chain:  chain,
		MUX:    mux,
	}
}

// ServeContentFunc does http.ServeContent the asset by path
func ServeContentFunc(prefix, path string, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		text, err := assets.Uncompressedasset(path)
		if err != nil {
			panic(err)
		}
		modtime, err := assets.ModTime(prefix, path)
		if err != nil {
			logger.Println(err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			// return
		}

		http.ServeContent(w, r, path, modtime, bytes.NewReader(text))
	}
}

type loggerPrint interface {
	Print(v ...interface{})
}

func Banner(listenaddr string, logger loggerPrint) {
	hostname, _ := getHostname()
	var addrsp *[]net.Addr
	if addrs, err := net.InterfaceAddrs(); err == nil {
		addrsp = &addrs
	}
	bannerText(listenaddr, hostname, addrsp, logger)
}

func bannerText(listenaddr, hostname string, addrsp *[]net.Addr, logger loggerPrint) {
	logger.Print(fmt.Sprintf("   %s\n", strings.Repeat("-", len(hostname)+7)))
	if len(hostname) > 19 {
		hostname = hostname[:16] + "..."
	}
	logger.Print(fmt.Sprintf(" / %s ostent \\\n", hostname))
	logger.Print("+------------------------------+\n")

	if h, port, err := net.SplitHostPort(listenaddr); err == nil && h == "::" && addrsp != nil {
		// wildcard bind
		fst := true
		for _, a := range *addrsp {
			ip := a.String()
			if ipnet, ok := a.(*net.IPNet); ok {
				ip = ipnet.IP.String()
			}
			if strings.Contains(ip, ":") { // IPv6, skip for now
				continue
			}
			f := fmt.Sprintf("http://%s:%s", ip, port)
			if len(f) < 28 {
				f += strings.Repeat(" ", 28-len(f))
			}
			if !fst {
				logger.Print("|------------------------------|\n")
			}
			fst = false
			logger.Print(fmt.Sprintf("| %s |\n", f))
		}
	} else {
		f := fmt.Sprintf("http://%s", listenaddr)
		if len(f) < 28 {
			f += strings.Repeat(" ", 28-len(f))
		}
		logger.Print(fmt.Sprintf("| %s |\n", f))
	}
	logger.Print("+------------------------------+\n")
}

// VERSION of the latest known release.
// Unused in non-production mode.
// Compared with in main.production.go.
const VERSION = "0.1.9"
