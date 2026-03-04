package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

const serverPort int = 3333

func Main() {
	errChan := make(chan error)
	// A goroutine which starts a HTTP server
	go func(channel chan error) {
		mux := http.NewServeMux()

		mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s\n", r.Method)
		})

		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		fmt.Println("starting the server on port: ", serverPort)
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
				channel <- err
			} else {
				fmt.Println("shutting down the server")
				channel <- nil
			}
		}
	}(errChan)

	err := <-errChan

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com",
		"http://junkyurll.com",
	}
	g, ctx := errgroup.WithContext(context.Background())

	for _, url := range urls {
		url := url
		g.Go(func() error {
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

			if err != nil {
				return err
			}

			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				return err
			}
			defer resp.Body.Close()

			fmt.Printf("fetch url %s status %s\n", url, resp.Status)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
