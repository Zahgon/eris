package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
)

var (
	// global error values can be useful when wrapping errors or inspecting error types
	errInternalServer = eris.New("error internal server")

	// declaring an error with pkg/errors for comparison
	errNotFound = errors.New("error not found")
)

type request struct {
	ID string
}

func (req *request) validate() error {
	_ = "STUB: not implemented"

	// create a new local error and wrap it with some context
	return nil
}

type resource struct {
	ID      string
	AbsPath string
}

func getResource(req request) (*resource, error) { _ = "STUB: not implemented"; return nil, nil }

// malformed absolute filepath to simulate a "bug"

func getRelPath(base string, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// it's generally useful to wrap external errors with a type that you know how to handle
// first (e.g. ErrInternalServer). this will help if/when you want to do error inspection
// via eris.Is(err, ErrInternalServer) or eris.Cause(err).

type response struct {
	RelPath string
}

func processResource(req request) (*response, error) { _ = "STUB: not implemented"; return nil, nil }

// simply return the error if there's no additional context

// do some processing on the data

// wrap the error if you want to add more context

type logReq struct {
	Method string
	Req    request
	Res    *response
	Err    error
}

func logRequest(logger *logrus.Logger, logReq logReq) { _ = "STUB: not implemented"; return }

// it's generally a good idea to contain error formatting logic inside a utility method like
// this one to ensure that all errors are logged uniformly. in this case, we're logging with
// the default format and stack traces enabled.

// This example demonstrates how to integrate eris with a JSON logger (e.g. logrus). It's broken
// into several methods to show the formatted output for wrapped errors, and it includes three
// failing cases to demonstrate how your error logs should look in different scenarios.
func main() {
	// setup JSON logger
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})

	// example requests
	reqs := []request{
		{
			ID: "", // bad request
		},
		{
			ID: "res1", // not found
		},
		{
			ID: "res2", // server error
		},
		{
			ID: "res3", // success
		},
	}

	// process the example requests and log the results
	for _, req := range reqs {
		res, err := processResource(req)
		if req.ID != "res1" {
			// log the eris error
			logRequest(logger, logReq{
				Method: "ProcessResource",
				Req:    req,
				Res:    res,
				Err:    err,
			})
		} else {
			// print the pkg/errors error
			fmt.Printf("%+v\n", err)
		}
	}
}
