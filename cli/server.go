package main

import (
    "io"
    "log"
    "net/http"
    "sync"
	"strconv"
	"net/http/httputil"
)


func startHttpServer(p int, endpoint string, wg *sync.WaitGroup) *http.Server {
	port := strconv.Itoa(p)
    srv := &http.Server{Addr: ":" + port}

    http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "OK")
		printRequest(r)
    })

    go func() {
        defer wg.Done() // let main know we are done cleaning up

        // always returns error. ErrServerClosed on graceful close
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            // unexpected error. port in use?
            log.Fatalf("ListenAndServe(): %v", err)
        }
    }()

    // returning reference so caller can call Shutdown()
    return srv
}

func printRequest(r *http.Request){
	reqDump, err := httputil.DumpRequest(r, true)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("REQUEST:\n%s", string(reqDump))
}