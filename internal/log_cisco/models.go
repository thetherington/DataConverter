package log_cisco

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

	CiscoNX struct {
		CPU       *CPU           `json:"cpu,omitempty"`
		CPU_CORE  *CPU_CORE      `json:"cpu_core,omitempty"`
		Memory    *Memory        `json:"memory,omitempty"`
		Processes *Processes     `json:"processes,omitempty"`
		Hardware  *Hardware      `json:"hardware,omitempty"`
		PSU       *PSU           `json:"psu,omitempty"`
		Fan       *Fan           `json:"fan,omitempty"`
		Temp      *Temp          `json:"temp,omitempty"`
		Uptime    *Uptime        `json:"uptime,omitempty"`
		Port      map[string]any `json:"ports,omitempty"`
		MRoute    map[string]any `json:"mroute,omitempty"`
	} `json:"cisconx"`
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

// Old Schema
type Ciscolog struct {
	Index  string          `json:"_index"`
	Type   string          `json:"_type"`
	Id     string          `json:"_id"`
	Score  int             `json:"_score"`
	Source Ciscolog_source `json:"_source"`
}

type Ciscolog_source struct {
	Timestamp  string      `json:"@timestamp"`
	UUID       string      `json:"@UUID"`
	Host       string      `json:"host"`
	Annotation *Annotation `json:"annotation"`
	Version    string      `json:"@version"`
	Type       string      `json:"type"`
	Tags       []string    `json:"tags"`

	Poller struct {
		CiscoNX struct {
			CPU       *CPU           `json:"cpu,omitempty"`
			CPU_CORE  *CPU_CORE      `json:"cpu_core,omitempty"`
			Memory    *Memory        `json:"memory,omitempty"`
			Processes *Processes     `json:"processes,omitempty"`
			Hardware  *Hardware      `json:"hardware,omitempty"`
			PSU       *PSU           `json:"psu,omitempty"`
			Fan       *Fan           `json:"fan,omitempty"`
			Temp      *Temp          `json:"temp,omitempty"`
			Uptime    *Uptime        `json:"uptime,omitempty"`
			Port      map[string]any `json:"ports,omitempty"`
			MRoute    map[string]any `json:"mroute,omitempty"`
		} `json:"cisconx"`
	} `json:"poller"`
}

// Common
type Annotation struct {
	General *General `json:"general,omitempty"`
}

type General struct {
	DeviceType *string `json:"device_type,omitempty"`
	DeviceName *string `json:"device_name,omitempty"`
	SystemName *string `json:"system_name,omitempty"`
}

type Device struct {
	Device     *string `json:"s_device,omitempty"`
	DeviceName *string `json:"s_device_name,omitempty"`
	DeviceType *string `json:"s_device_type,omitempty"`
	DeviceSize *string `json:"s_device_size,omitempty"`
}
