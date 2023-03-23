package log_syslog

type Doc struct {
	Index  string     `json:"_index"`
	Type   string     `json:"_type"`
	Id     string     `json:"_id"`
	Score  int        `json:"_score"`
	Source Ecs_source `json:"_source"`
}

// New ECS (Elastic Common Schema)
type Ecs_source struct {
	Timestamp       string          `json:"@timestamp"`
	Message         string          `json:"message"`
	UUID            string          `json:"@UUID"`
	OPID            *string         `json:"@OPID,omitempty"`
	ECS             Ecs_version     `json:"ecs"`
	Event           EventCreated    `json:"event"`
	DeviceTimestamp DeviceTimestamp `json:"device"`
	Host            *Host           `json:"host,omitempty"`
	Annotation      *Annotation     `json:"annotation,omitempty"`
	Log             *Log            `json:"log,omitempty"`
	Process         *Process        `json:"process,omitempty"`
	Version         string          `json:"@version"`
	Tags            []string        `json:"tags"`
}

type Ecs_version struct {
	Version string `json:"version"`
}

type EventCreated struct {
	Created string `json:"created"`
}

type DeviceTimestamp struct {
	Timestamp string `json:"timestamp"`
}

type Host struct {
	IP     *string `json:"ip,omitempty"`
	Name   *string `json:"name,omitempty"`
	Device *Device `json:"device,omitempty"`
}

type Log struct {
	Original string        `json:"original"`
	Syslog   *SyslogStruct `json:"syslog,omitempty"`
}

type SyslogStruct struct {
	Message  *string `json:"message,omitempty"`
	Priority *string `json:"priority,omitempty"`

	Facility *struct {
		Name *string `json:"name,omitempty"`
		Code *int    `json:"code,omitempty"`
	} `json:"facility,omitempty"`

	Severity *struct {
		Name *string `json:"name,omitempty"`
		Code *int    `json:"code,omitempty"`
	} `json:"severity,omitempty"`
}

type Process struct {
	Name *string `json:"name,omitempty"`

	PID *struct {
		Number   *string `json:"number,omitempty"`
		Original *string `json:"original,omitempty"`
	}
}

// Old Schema
type Syslog struct {
	Index  string        `json:"_index"`
	Type   string        `json:"_type"`
	Id     string        `json:"_id"`
	Score  int           `json:"_score"`
	Source Syslog_source `json:"_source"`
}

type Syslog_source struct {
	Timestamp          string      `json:"@timestamp"`
	Message            string      `json:"message"`
	UUID               string      `json:"@UUID"`
	OPID               *string     `json:"@OPID"`
	Host               string      `json:"host"`
	ReceivedFrom       *string     `json:"received_from"`
	ReceivedAt         *string     `json:"received_at"`
	Annotation         *Annotation `json:"annotation"`
	SyslogTimestamp    *string     `json:"syslog_timestamp"`
	SyslogPID          *string     `json:"syslog_pid"`
	SyslogProcessId    *string     `json:"syslog_process_id"`
	SyslogFacility     *string     `json:"syslog_facility"`
	SyslogFacilityCode *int        `json:"syslog_facility_code"`
	SyslogProgram      *string     `json:"syslog_program"`
	SyslogMessage      *string     `json:"syslog_message"`
	SyslogSeverity     *string     `json:"syslog_severity"`
	SyslogHostname     *string     `json:"syslog_hostname"`
	SyslogPriority     *string     `json:"syslog_priority"`
	SyslogSeverityCode *int        `json:"syslog_severity_code"`
	DeviceMeta         *string     `json:"device_meta"`
	Device             *Device     `json:"device"`
	DeviceTimestamp    *string     `json:"device_timestamp"`
	Version            string      `json:"@version"`
	Type               string      `json:"type"`
	Tags               []string    `json:"tags"`
}

// Common Objects shared between the two
type Device struct {
	Prod          *string `json:"prod,omitempty"`
	IP            *string `json:"ip,omitempty"`
	Alias         *string `json:"alias,omitempty"`
	IPControl1    *string `json:"ip-control-1,omitempty"`
	IPControl2    *string `json:"ip-control-2,omitempty"`
	MAC           *string `json:"mac,omitempty"`
	MulticastIP   *string `json:"multicast-ip,omitempty"`
	MulticastPort *int    `json:"multicast-port,omitempty"`
}

type Annotation struct {
	General *General `json:"general,omitempty"`
}

type General struct {
	DeviceType *string `json:"device_type,omitempty"`
	DeviceName *string `json:"device_name,omitempty"`
	SystemName *string `json:"system_name,omitempty"`
}
