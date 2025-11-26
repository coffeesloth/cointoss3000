package main

import (
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/pflag"
)

const (
	defaultTosses   = 3000
	defaultVerbose  = false
	maxServerTosses = 10000
)

func main() {
	port := pflag.IntP("server", "s", 0,
		"run as an HTTP server listening for GET / on the specified port")
	help := pflag.BoolP("help", "h", false, "show help")
	tosses := pflag.Uint64P("tosses", "t", defaultTosses,
		"number of tosses to perform")
	verbose := pflag.BoolP("verbose", "v", defaultVerbose,
		"print total heads and tails generated")
	pflag.Parse()

	if *help {
		pflag.Usage()
		return
	}

	if *port == 0 {
		if *verbose && *tosses != defaultTosses {
			fmt.Printf("tosses: %d\n", *tosses)
		}
		fmt.Println(toss(*tosses, *verbose))
		return
	}

	if *port < 0 {
		log.Fatalf("bad port: %d\n", *port)
	}

	if *tosses != defaultTosses || *verbose != defaultVerbose {
		log.Fatalln("other flags may not be used with the server flag")
	}

	server(*port)
}

func toss(tosses uint64, verbose bool) string {
	results := []uint64{0, 0}
	for i := uint64(0); i < tosses; i++ {
		results[rand.Int()%2]++
	}

	var winner string
	if results[0] > results[1] {
		winner = "heads"
	} else if results[1] > results[0] {
		winner = "tails"
	} else {
		winner = "draw"
	}

	if verbose {
		return fmt.Sprintf("heads: %d\ntails: %d\nwinner: %s",
			results[0], results[1], winner)
	}

	return winner
}

func server(port int) {
	maxTossesExceededMsg := fmt.Sprintf("max tosses is %d", maxServerTosses)

	writeOrLog := func(w http.ResponseWriter, body string) {
		if _, err := w.Write([]byte(body)); err != nil {
			slog.Error("failed to write http response", "err", err)
		}
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		verbose := r.URL.Query().Has("v")

		var tosses int
		if !verbose {
			tosses = 1
		} else {
			tosses = defaultTosses
		}
		if r.URL.Query().Has("t") {
			tossesStr := r.URL.Query().Get("t")
			var err error
			tosses, err = strconv.Atoi(tossesStr)
			if err != nil {
				w.WriteHeader(400)
				writeOrLog(w, fmt.Sprintf("what is that -> '%s'?", tossesStr))
				return
			}
			if tosses < 0 {
				w.WriteHeader(400)
				writeOrLog(w, "how do you propose we do a negative number of tosses?")
				return
			}
			if tosses > maxServerTosses {
				w.WriteHeader(400)
				writeOrLog(w, maxTossesExceededMsg)
				return
			}
			if !verbose {
				tosses = 1
			}
		}

		writeOrLog(w, fmt.Sprintln(toss(uint64(tosses), verbose)))
	})

	fmt.Printf("listening on %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		slog.Error("http listen and serve failed", "err", err)
	}
}
