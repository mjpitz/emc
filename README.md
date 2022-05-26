# emc

_The minimal, declarative service catalog._

## Background

In the last three jobs I've worked at, it's always been a hassle trying to locate the various dashboards, documentation,
and support for a given project. I joined effx to try and help address just that. As I've been spinning up a new cluster,
I found myself wanting a landing page for the systems that I use regularly.

## Building your catalog

The `emc` service catalog is defined using a simple Golang script. This makes it easy for engineers to drop in their
own functionality for rendering links, link groups, or services. For an example, see the provided `grafana` package
which includes several of my personal dashboards for different systems.

```go
// catalog.go

//go:build ignore
// +build ignore

package main

import (
	"github.com/mjpitz/emc/catalog"
	"github.com/mjpitz/emc/catalog/grafana"
	"github.com/mjpitz/emc/catalog/linkgroup"
	"github.com/mjpitz/emc/catalog/service"
)

func main() {
	catalog.Serve(
		catalog.Service(
			"Drone",
			service.LogoURL("https://path/to/drone-logo.png"),
			service.URL("https://drone.example.com"),
			service.Description("Drone is a self-service Continuous Integration platform for busy development teams."),
			service.Metadata("Contact", "drone@example.com"),
			service.LinkGroup(
				"Dashboards",
				linkgroup.Link("Drone", grafana.Drone("cicd", "drone")),
				linkgroup.Link("Golang", grafana.Golang("cicd", "drone")),
				linkgroup.Link("Litestream", grafana.Litestream("cicd", "drone")),
				linkgroup.Link("Redis Queue", grafana.Redis("cicd", "drone-redis-queue")),
			),
			service.LinkGroup(
				"Documentation",
				linkgroup.Link("docs.drone.io", "https://docs.drone.io/"),
			),
		),
		// ...
	)
}
```

## Hosting your catalog

Once you've built your catalog, you can easily run a landing page by executing the catalog file.

```sh
$ go run ./catalog.go
```

This starts a web server for you to interact with on `localhost:8080`. If `:8080` is already in use, you can configure
the bind address by passing the `-bind_address` flag with the desired host and port.

<center>
  <img src="screenshot.png" alt="Screenshot" width="72%"/>
</center>

### Exporting your catalog

Instead of needing to compile a binary or host your catalog using `go run`, you can export your catalog to HTML or JSON.
This makes it easy to drop into existing self-host platforms or leverage with other popular systems.

```sh
$ go run ./catalog.go -output html > index.html
$ go run ./catalog.go -output json > catalog.json
```

### Protecting your catalog using oauth-proxy

Regardless of how you host your catalog, you'll likely want to protect access to it. An easy way to do this is using the
[oauth-proxy][] project. This project provides common OAuth2 client functionality to any project, making it easy to
require authentication in order to access a system / project.

<!-- TODO: write up guide and link to it from here -->

Until I have more of a concrete guide, you can follow my setup [here](https://github.com/mjpitz/mjpitz/blob/main/infra/helm/catalog/values.yaml).
A simple analogy to this deployment would be a docker compose file with two services, one for the oauth-proxy and the
other for the catalog (bound to 127.0.0.1). Using the new `-output` functionality, this deployment could definitely
be simplified.

[oauth-proxy]: https://oauth2-proxy.github.io/oauth2-proxy
