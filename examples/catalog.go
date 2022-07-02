package main

import (
	"code.pitz.tech/mya/emc/catalog"
	"code.pitz.tech/mya/emc/catalog/linkgroup"
	"code.pitz.tech/mya/emc/catalog/service"
)

func main() {
	catalog.Serve(
		catalog.Service(
			"CI/CD",
			service.LogoURL("https://th.bing.com/th/id/OIP.wd0WnO0MF56eQ23LR8XzRAAAAA?pid=ImgDet&rs=1"),
			service.URL("https://deploy.example.com"),
			service.Description("Continuous integration and delivery platform."),
			service.Metadata(
				"Key1", "Key1Value1",
				"Key1", "Key1Value2",
				"Key2", "Key2Value1",
			),
			service.LinkGroup(
				"Dashboards",
				linkgroup.Link("System", "#"),
				linkgroup.Link("Queue", "#"),
			),
			service.LinkGroup(
				"Documentation",
				linkgroup.Link("Developing", "#"),
				linkgroup.Link("Releasing", "#"),
				linkgroup.Link("Contributing", "#"),
			),
		),
		catalog.Service(
			"Version Control",
			service.LogoURL(""),
			service.URL("https://code.example.com"),
			service.Description("System source code."),
			service.Metadata(
				"Key1", "Key1Value1",
				"Key1", "Key1Value2",
				"Key2", "Key2Value1",
			),
			service.LinkGroup(
				"Dashboards",
				linkgroup.Link("System", "#"),
				linkgroup.Link("Database", "#"),
				linkgroup.Link("Queue", "#"),
			),
			service.LinkGroup(
				"Documentation",
				linkgroup.Link("Developing", "#"),
				linkgroup.Link("Releasing", "#"),
			),
		),
		catalog.Service(
			"Monitoring",
			service.Metadata(
				"Key1", "Key1Value1",
				"Key1", "Key1Value2",
				"Key2", "Key2Value1",
			),
			service.LinkGroup(
				"Dashboards",
				linkgroup.Link("System", "#"),
				linkgroup.Link("Database", "#"),
				linkgroup.Link("Queue", "#"),
			),
			service.LinkGroup(
				"Documentation",
				linkgroup.Link("Developing", "#"),
				linkgroup.Link("Releasing", "#"),
			),
		),
		catalog.Service(
			"Alerting",
			service.Metadata(
				"Key1", "Key1Value1",
				"Key1", "Key1Value2",
				"Key2", "Key2Value1",
			),
			service.LinkGroup(
				"Dashboards",
				linkgroup.Link("System", "#"),
				linkgroup.Link("Database", "#"),
				linkgroup.Link("Queue", "#"),
			),
			service.LinkGroup(
				"Documentation",
				linkgroup.Link("Developing", "#"),
				linkgroup.Link("Releasing", "#"),
			),
		),
	)
}
