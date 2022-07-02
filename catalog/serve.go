// Copyright (C) 2022 Mya Pitzeruse
// The MIT License (MIT)

package catalog

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"code.pitz.tech/mya/emc/catalog/service"
)

//go:embed index.html.tpl
var catalog string

var funcs = map[string]any{
	"mod": func(mod, v int) int { return v % mod },
	"eq":  func(exp, act int) bool { return exp == act },
}

// Spec defines the requirements for hosting a catalog.
type Spec struct {
	Services []service.Spec
}

// Option defines an optional component of the spec.
type Option func(spec *Spec)

// Service registers a known service with the catalog.
func Service(label string, options ...service.Option) Option {
	return func(spec *Spec) {
		spec.Services = append(spec.Services, service.New(label, options...))
	}
}

// Serve provides command line functionality for running the service catalog.
func Serve(options ...Option) {
	addr := flag.String("bind_address", "127.0.0.1:8080", "the address the service should bind to when serving content")
	output := flag.String("output", "", "set to output the catalog to stdout (valid html, json)")
	flag.Parse()

	start := time.Now()

	spec := Spec{}
	for _, opt := range options {
		opt(&spec)
	}

	buffer := bytes.NewBuffer(nil)
	var err error

	switch {
	case output == nil || *output == "" || *output == "html":
		err = template.Must(template.New("catalog").Funcs(funcs).Parse(catalog)).Execute(buffer, spec)
	case *output == "json":
		err = json.NewEncoder(buffer).Encode(spec)
	}

	if err != nil {
		log.Fatal("failed to render output", err)
	}

	if output != nil && *output != "" {
		_, err = io.Copy(os.Stdout, bytes.NewReader(buffer.Bytes()))
		if err != nil {
			log.Fatal("failed to write buffer to stdout", err)
		}
	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeContent(w, r, "", start, bytes.NewReader(buffer.Bytes()))
		})

		log.Printf("serving on %s\n", *addr)
		err = http.ListenAndServe(*addr, http.DefaultServeMux)
		if err != nil {
			log.Fatal("failed to serve content", err)
		}
	}
}
