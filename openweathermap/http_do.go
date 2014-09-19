package openweathermap

import (
	"code.google.com/p/go.net/context"
	"net/http"
)

type ResponseHandlerFunc func(*http.Response, error) error

func httpDo(ctx context.Context, req *http.Request, f ResponseHandlerFunc) error {
	// Run the HTTP request in a go routine and pass the response to f.
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)

	go func() { c <- f(client.Do(req)) }()
	select {
	case <-ctx.Done():
		debug("httpDo - CancelRequest")
		tr.CancelRequest(req)
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		debug("httpDo - ok")
		return err
	}
}