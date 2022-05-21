// Copyright (C) 2022 Mya Pitzeruse
// The MIT License (MIT)

package catalog

import (
	"bytes"
	_ "embed"
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/mjpitz/emc/catalog/service"
)

//go:embed index.html.tpl
var catalog string

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
	flag.Parse()

	start := time.Now()

	t := template.Must(template.New("catalog").
		Funcs(map[string]any{
			"mod": func(mod, v int) int {
				return v % mod
			},
			"eq": func(exp, act int) bool {
				return exp == act
			},
		}).
		Parse(catalog))

	spec := Spec{}
	for _, opt := range options {
		opt(&spec)
	}

	html := bytes.NewBuffer(nil)
	err := t.Execute(html, spec)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, "", start, bytes.NewReader(html.Bytes()))
	})

	log.Printf("serving on %s\n", *addr)
	err = http.ListenAndServe(*addr, http.DefaultServeMux)
	if err != nil {
		panic(err)
	}
}
