package log_ping

type Doc struct {
	Index  string     `json:"_index"`
	Type   string     `json:"_type"`
	Id     string     `json:"_id"`
	Score  int        `json:"_score"`
	Source Ecs_source `json:"_source"`
}

// New ECS (Elastic Common Schema)
type Ecs_source struct {
	Timestamp  string       `json:"@timestamp"`
	UUID       string       `json:"@UUID"`
	OPID       *string      `json:"@OPID,omitempty"`
	ECS        Ecs_version  `json:"ecs"`
	Event      EventCreated `json:"event"`
	Agent      Agent        `json:"agent"`
	Host       *Host        `json:"host,omitempty"`
	Annotation *Annotation  `json:"annotation,omitempty"`
	Version    string       `json:"@version"`
	Tags       []string     `json:"tags"`

	Service struct {
		Type string `json:"type"`
	} `json:"service"`

	Metricset struct {
		Name string `json:"name"`
	} `json:"metricset"`

	Ping struct {
		Ping PingDoc `json:"ping"`
	} `json:"ping"`
}

type Ecs_version struct {
	Version string `json:"version"`
}

type EventCreated struct {
	Created string `json:"created"`
	Module  string `json:"module"`
	Dataset string `json:"dataset"`
}

type Host struct {
	IP *string `json:"ip,omitempty"`
}

type Agent struct {
	Type     string `json:"type"`
	Version  string `json:"version"`
	Hostname string `json:"hostname"`
	ID       string `json:"id"`
}

type PingDoc struct {
	ResultCode         *int     `json:"i_result_code,omitempty"`
	PercentPacketLoss  *int     `json:"i_percent_packet_loss,omitempty"`
	PacketsReceived    *int     `json:"i_packets_received,omitempty"`
	PacketsTransmitted *int     `json:"i_packets_transmitted,omitempty"`
	MinimumResponseMs  *float32 `json:"d_minimum_response_ms,omitempty"`
	MaximumResponseMs  *float32 `json:"d_maximum_response_ms,omitempty"`
	AverageResponseMs  *float32 `json:"d_average_response_ms,omitempty"`
}

// Old Schema
type Pinglog struct {
	Index  string         `json:"_index"`
	Type   string         `json:"_type"`
	Id     string         `json:"_id"`
	Score  int            `json:"_score"`
	Source Pinglog_source `json:"_source"`
}

type Pinglog_source struct {
	Timestamp  string      `json:"@timestamp"`
	UUID       string      `json:"@UUID"`
	Host       string      `json:"host"`
	Annotation *Annotation `json:"annotation"`
	Version    string      `json:"@version"`
	Type       string      `json:"type"`
	Tags       []string    `json:"tags"`

	Poller struct {
		Ping struct {
			Ping PingDoc `json:"ping"`
		} `json:"ping"`
	} `json:"poller"`
}

type Annotation struct {
	General *General `json:"general,omitempty"`
}

type General struct {
	DeviceType *string `json:"device_type,omitempty"`
	DeviceName *string `json:"device_name,omitempty"`
	SystemName *string `json:"system_name,omitempty"`
}
