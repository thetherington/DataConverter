package log_cisco

import (
	"embed"
	"log"
	"strings"

	"github.com/thetherington/DataConverter/internal/helpers"
)

var (
	//go:embed tmpl
	folder embed.FS

	templates = map[string]string{
		"alias":    "tmpl/cisco-index-alias.json",
		"mapping":  "tmpl/cisco-index-mapping.json",
		"settings": "tmpl/cisco-index-settings.json",
	}
)

type LogCisco struct {
	num_errors       int
	SourceIndex      string
	DestinationIndex string
	Templates        map[string]string
	Count            chan int
}

func New(index string) *LogCisco {
	return &LogCisco{
		num_errors:       0,
		SourceIndex:      index,
		DestinationIndex: strings.Replace(index, "log-metric-poller-cisconx", "log-metric-p-cisconx", -1),
		Templates:        templates,
		Count:            make(chan int),
	}
}

func (cisco *LogCisco) IncrementErrors() {
	cisco.num_errors++
}

func (cisco *LogCisco) ReturnErrors() int {
	return cisco.num_errors
}

func (cisco *LogCisco) SendToCountChan(v int) {
	cisco.Count <- v
}

func (cisco *LogCisco) GetCountChan() chan int {
	return cisco.Count
}

func (cisco *LogCisco) GetTemplate(name string) string {
	f, err := folder.ReadFile(cisco.Templates[name])
	if err != nil {
		log.Println(err)
	}

	tmpl := strings.Replace(string(f), "cisco-index", cisco.DestinationIndex, -1)

	return tmpl
}

func (cisco *LogCisco) GetSourceIndexName() string {
	return cisco.SourceIndex
}

func (cisco *LogCisco) GetDestinationIndexName() string {
	return cisco.DestinationIndex
}

func (cisco *LogCisco) Convert(i string) (string, error) {
	var s Ciscolog

	err := helpers.ReadJSON(i, &s)
	if err != nil {
		return "", err
	}

	doc := ConverCiscologToDoc(s)

	j, err := helpers.WriteJSON(doc)
	if err != nil {
		return "", err
	}

	// j = strings.TrimSuffix(j, "\n")

	return j, nil
}

func ConverCiscologToDoc(s Ciscolog) Doc {
	doc := Doc{
		Index: strings.Replace(s.Index, "log-metric-poller-cisconx", "log-metric-p-cisconx", -1),
		Type:  "_doc",
		Id:    s.Id,
		Score: s.Score,
		Source: Ecs_source{
			Timestamp: s.Source.Timestamp,
			UUID:      s.Source.UUID,
			Version:   s.Source.Version,
			Event: EventCreated{
				Created: s.Source.Timestamp,
				Module:  "cisconx",
				Dataset: "",
			},
			Agent: Agent{
				Type:     "collector-poller",
				Version:  "2.0.15",
				Hostname: "insite",
				ID:       "pll-11111111-2222-3333-4444-555555555555",
			},
			ECS: Ecs_version{
				Version: "1.3.0",
			},
			Tags: s.Source.Tags,
		},
	}

	doc.Source.Service = struct {
		Type string "json:\"type\""
	}{
		Type: "cisconx",
	}

	doc.Source.Metricset = struct {
		Name string "json:\"name\""
	}{
		Name: "",
	}

	// Populate host object
	doc.Source.Host = &Host{
		IP: &s.Source.Host,
	}

	// Populate annotation object if it exists.
	if s.Source.Annotation != nil {
		doc.Source.Annotation = s.Source.Annotation
	}

	event := s.Source.Poller.CiscoNX

	switch {

	case event.CPU != nil:
		doc.Source.CiscoNX.CPU = event.CPU
		doc.Source.Event.Dataset = "cisconx.cpu"
		doc.Source.Metricset.Name = "cpu"

	case event.CPU_CORE != nil:
		doc.Source.CiscoNX.CPU_CORE = event.CPU_CORE
		doc.Source.Event.Dataset = "cisconx.cpu_core"
		doc.Source.Metricset.Name = "cpu_core"

	case event.Memory != nil:
		doc.Source.CiscoNX.Memory = event.Memory
		doc.Source.Event.Dataset = "cisconx.memory"
		doc.Source.Metricset.Name = "memory"

	case event.Processes != nil:
		doc.Source.CiscoNX.Processes = event.Processes
		doc.Source.Event.Dataset = "cisconx.processes"
		doc.Source.Metricset.Name = "processes"

	case event.Hardware != nil:
		doc.Source.CiscoNX.Hardware = event.Hardware
		doc.Source.Event.Dataset = "cisconx.hardware"
		doc.Source.Metricset.Name = "hardware"

	case event.PSU != nil:
		doc.Source.CiscoNX.PSU = event.PSU
		doc.Source.Event.Dataset = "cisconx.psu"
		doc.Source.Metricset.Name = "psu"

	case event.Fan != nil:
		doc.Source.CiscoNX.Fan = event.Fan
		doc.Source.Event.Dataset = "cisconx.fan"
		doc.Source.Metricset.Name = "fan"

	case event.Temp != nil:
		doc.Source.CiscoNX.Temp = event.Temp
		doc.Source.Event.Dataset = "cisconx.temp"
		doc.Source.Metricset.Name = "fan"

	case event.Uptime != nil:
		doc.Source.CiscoNX.Uptime = event.Uptime
		doc.Source.Event.Dataset = "cisconx.uptime"
		doc.Source.Metricset.Name = "uptime"

	case event.Port != nil:
		doc.Source.CiscoNX.Port = event.Port
		doc.Source.Event.Dataset = "cisconx.ports"
		doc.Source.Metricset.Name = "ports"

	case event.MRoute != nil:
		doc.Source.CiscoNX.MRoute = event.MRoute

		if val, ok := event.MRoute["s_odev_name"]; ok {
			doc.Source.CiscoNX.MRoute["as_odev_name"] = strings.Split(val.(string), ", ")
		}

		if val, ok := event.MRoute["s_oif_name"]; ok {
			doc.Source.CiscoNX.MRoute["as_oif_name"] = strings.Split(val.(string), ", ")
		}

		doc.Source.Event.Dataset = "cisconx.mroute"
		doc.Source.Metricset.Name = "mroute"

	}

	return doc
}
