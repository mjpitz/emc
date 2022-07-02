// Copyright (C) 2022 Mya Pitzeruse
// The MIT License (MIT)

package service

import (
	"code.pitz.tech/mya/emc/catalog/linkgroup"
)

// New constructs a spec given a label and set of options.
func New(label string, options ...Option) Spec {
	spec := Spec{
		Label: label,
	}

	for _, opt := range options {
		opt(&spec)
	}

	return spec
}

// Option defines an optional component of the spec.
type Option func(spec *Spec)

// KV defines a metadata entry.
type KV struct {
	Key   string
	Value string
}

// Spec defines the elements needed to render a service.
type Spec struct {
	Label       string
	LogoURL     string
	Description string
	URL         string
	Metadata    []KV
	LinkGroups  []linkgroup.Spec
}

// LogoURL configures the icon for the service.
func LogoURL(url string) Option {
	return func(spec *Spec) {
		spec.LogoURL = url
	}
}

// Description specifies a short, brief description about the service.
func Description(description string) Option {
	return func(spec *Spec) {
		spec.Description = description
	}
}

// URL configures the services public facing URL that's presented to end users.
func URL(url string) Option {
	return func(spec *Spec) {
		spec.URL = url
	}
}

// Metadata allows additional metadata to be attached to a service.
func Metadata(kvs ...string) Option {
	// ensure even number of parameters
	if len(kvs)%2 > 0 {
		kvs = append(kvs, "")
	}

	return func(spec *Spec) {
		for i := 0; i < len(kvs); i += 2 {
			spec.Metadata = append(spec.Metadata, KV{
				Key:   kvs[i],
				Value: kvs[i+1],
			})
		}
	}
}

// LinkGroup appends a group of links to the provided service.
func LinkGroup(label string, options ...linkgroup.Option) Option {
	return func(spec *Spec) {
		spec.LinkGroups = append(spec.LinkGroups, linkgroup.New(label, options...))
	}
}
