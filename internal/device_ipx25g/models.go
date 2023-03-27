package device_ipx25g

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
	Message    string       `json:"message"`
	UUID       string       `json:"@UUID"`
	ECS        Ecs_version  `json:"ecs"`
	Event      EventCreated `json:"event"`
	Host       *Host        `json:"host,omitempty"`
	Annotation *Annotation  `json:"annotation,omitempty"`
	DataType   string       `json:"data_type"`
	Version    string       `json:"@version"`
	Tags       []string     `json:"tags"`

	Flow    *DeviceIPX_Flow    `json:"flow,omitempty"`
	Port    *DeviceIPX_Port    `json:"port,omitempty"`
	Chassis *DeviceIPX_Chassis `json:"chassis,omitempty"`
	LC      *DeviceIPX_LC      `json:"lc,omitempty"`
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
	IP *string `json:"ip,omitempty"`
}

// Old Schema
type DeviceIPX struct {
	Index  string           `json:"_index"`
	Type   string           `json:"_type"`
	Id     string           `json:"_id"`
	Score  int              `json:"_score"`
	Source DeviceIPX_source `json:"_source"`
}

type DeviceIPX_source struct {
	Timestamp  string      `json:"@timestamp"`
	Message    string      `json:"message"`
	UUID       string      `json:"@UUID"`
	Host       string      `json:"host"`
	ReceivedIP string      `json:"received_ip"`
	EXEIP      string      `json:"exe_ip"`
	Annotation *Annotation `json:"annotation"`
	DataType   string      `json:"data_type"`
	Version    string      `json:"@version"`
	Type       string      `json:"type"`
	Tags       []string    `json:"tags"`

	DeviceIPX_Flow
	DeviceIPX_Port
	DeviceIPX_LC
	DeviceIPX_Chassis
}

type DeviceIPX_Flow struct {
	EgressCardLabelList     []string `json:"egress_card_label_list,omitempty"`
	EgressCardPortNumList   []int    `json:"egress_card_port_num_list,omitempty"`
	EgressDevices           []string `json:"egress_devices,omitempty"`
	EgressDeviceSFPs        []string `json:"egress_device_sfps,omitempty"`
	EgressPorts             []int    `json:"egress_ports,omitempty"`
	EgressPortInterfaceList []string `json:"egress_port_interface_list,omitempty"`
	Egress                  []Egress `json:"egress,omitempty"`

	IngressCardLabel     *string `json:"ingress_card_label,omitempty"`
	IngressCardInterface *string `json:"ingress_card_interface,omitempty"`
	IngressCardPortNum   *int    `json:"ingress_card_port_num,omitempty"`
	IngressDevice        *string `json:"ingress_device,omitempty"`
	IngressDeviceSFP     *string `json:"ingress_device_sfp,omitempty"`
	IngressPortNumber    *int    `json:"ingress_port_number,omitempty"`

	MulticastDestIP *string `json:"multicast_dest_ip"`
	MeasuredBPS     *int    `json:"measured_bps,omitempty"`
	MeasuredKBPS    *int    `json:"measured_kbps,omitempty"`
	SignalGroup     *string `json:"signal_group,omitempty"`
	SignalType      *string `json:"signal_type,omitempty"`
	SignalLabel     *string `json:"signal_type_label,omitempty"`

	SourceMap *map[string]string `json:"source,omitempty"`
}

type Egress struct {
	CardLabel     *string `json:"card_label,omitempty"`
	PortInterface *string `json:"port_interface,omitempty"`
	CardPortNum   *int    `json:"card_port_num,omitempty"`
	DeviceSFP     *string `json:"device_sfp,omitempty"`
	EgressPort    *int    `json:"egress_port,omitempty"`
	Device        *string `json:"device,omitempty"`
}

type DeviceIPX_Port struct {
	CardLabel      *string `json:"card_label,omitempty"`
	CardPortNumber *string `json:"card_port_number,omitempty"`
	PortInterface  *string `json:"port_interface,omitempty"`
	PortNum        *int    `json:"port_num,omitempty"`
	PortNumber     *int    `json:"port_number,omitempty"`
	RXMeasured     *int    `json:"rx_measured,omitempty"`
	RXMeasuredUnit *string `json:"rx_measured_unit,omitempty"`
	TXMeasured     *int    `json:"tx_measured,omitempty"`
	TXMeasuredUnit *string `json:"tx_measured_unit,omitempty"`
	Device         *string `json:"device,omitempty"`
	DeviceSFP      *string `json:"device_sfp,omitempty"`
}

type DeviceIPX_LC struct {
	// v10 fields
	LCCardType *int `json:"lc_card_type,omitempty"`
	LCIndex    *int `json:"lc_index,omitempty"`

	// v11 fields
	CardType *int `json:"card_type,omitempty"`
	Index    *int `json:"index,omitempty"`
}

type DeviceIPX_Chassis struct {
	// v10 fields
	NumLCCards *int `json:"num_lc_cards,omitempty"`
	NumXCCards *int `json:"num_xc_cards,omitempty"`

	// v11 fields
	NumberOfLCCards *int `json:"number_of_lc_cards,omitempty"`
	NumberOfXCCards *int `json:"number_of_xc_cards,omitempty"`

	EXEType        *int `json:"exe_type,omitempty"`
	EXEVersion     *int `json:"exe_version,omitempty"`
	ChassisVersion *int `json:"version,omitempty"`
}

// Common Objects shared between the two
type Annotation struct {
	General *General `json:"general,omitempty"`
	EXE     *EXE     `json:"exe,omitempty"`
}

type General struct {
	DeviceType *string `json:"device_type,omitempty"`
	DeviceName *string `json:"device_name,omitempty"`
	SystemName *string `json:"system_name,omitempty"`
}

type EXE struct {
	Group *string `json:"group,omitempty"`
}
