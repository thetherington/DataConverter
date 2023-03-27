package log_ping

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
		"alias":    "tmpl/ping-index-alias.json",
		"mapping":  "tmpl/ping-index-mapping.json",
		"settings": "tmpl/ping-index-settings.json",
	}
)

type LogPing struct {
	num_errors       int
	SourceIndex      string
	DestinationIndex string
	Templates        map[string]string
	Count            chan int
}

func New(index string) *LogPing {
	return &LogPing{
		num_errors:       0,
		SourceIndex:      index,
		DestinationIndex: strings.Replace(index, "log-metric-poller-ping", "log-metric-p-ping", -1),
		Templates:        templates,
		Count:            make(chan int),
	}
}

func (ping *LogPing) IncrementErrors() {
	ping.num_errors++
}

func (ping *LogPing) ReturnErrors() int {
	return ping.num_errors
}

func (ping *LogPing) SendToCountChan(v int) {
	ping.Count <- v
}

func (ping *LogPing) GetCountChan() chan int {
	return ping.Count
}

func (ping *LogPing) GetTemplate(name string) string {
	f, err := folder.ReadFile(ping.Templates[name])
	if err != nil {
		log.Println(err)
	}

	tmpl := strings.Replace(string(f), "ping-index", ping.DestinationIndex, -1)

	return tmpl
}

func (ping *LogPing) GetSourceIndexName() string {
	return ping.SourceIndex
}

func (ping *LogPing) GetDestinationIndexName() string {
	return ping.DestinationIndex
}

func (ping *LogPing) Convert(i string) (string, error) {
	var s Pinglog

	err := helpers.ReadJSON(i, &s)
	if err != nil {
		return "", err
	}

	doc := ConverPinglogToDoc(s)

	j, err := helpers.WriteJSON(doc)
	if err != nil {
		return "", err
	}

	// j = strings.TrimSuffix(j, "\n")

	return j, nil
}

func ConverPinglogToDoc(s Pinglog) Doc {
	doc := Doc{
		Index: strings.Replace(s.Index, "log-metric-poller-ping", "log-metric-p-ping", -1),
		Type:  "_doc",
		Id:    s.Id,
		Score: s.Score,
		Source: Ecs_source{
			Timestamp: s.Source.Timestamp,
			UUID:      s.Source.UUID,
			Version:   s.Source.Version,
			Event: EventCreated{
				Created: s.Source.Timestamp,
				Module:  "ping",
				Dataset: "ping.ping",
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
		Type: "ping",
	}

	doc.Source.Metricset = struct {
		Name string "json:\"name\""
	}{
		Name: "ping",
	}

	// Populate host object
	doc.Source.Host = &Host{
		IP: &s.Source.Host,
	}

	// Populate annotation object if it exists.
	if s.Source.Annotation != nil {
		doc.Source.Annotation = s.Source.Annotation
	}

	// Copy ping object as-is to new doc
	doc.Source.Ping = struct {
		Ping PingDoc "json:\"ping\""
	}{
		Ping: s.Source.Poller.Ping.Ping,
	}

	return doc
}
