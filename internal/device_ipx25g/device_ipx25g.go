package device_ipx25g

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
		"alias":    "tmpl/device-ipx25g-alias.json",
		"mapping":  "tmpl/device-ipx25g-mapping.json",
		"settings": "tmpl/device-ipx25g-settings.json",
	}
)

type DeviceIPX25g struct {
	num_errors int
	Index      string
	Templates  map[string]string
	Count      chan int
}

func New(index string) *DeviceIPX25g {
	return &DeviceIPX25g{
		num_errors: 0,
		Index:      index,
		Templates:  templates,
		Count:      make(chan int),
	}
}

func (ipx *DeviceIPX25g) IncrementErrors() {
	ipx.num_errors++
}

func (ipx *DeviceIPX25g) ReturnErrors() int {
	return ipx.num_errors
}

func (ipx *DeviceIPX25g) SendToCountChan(v int) {
	ipx.Count <- v
}

func (ipx *DeviceIPX25g) GetCountChan() chan int {
	return ipx.Count
}

func (ipx *DeviceIPX25g) GetTemplate(name string) string {
	f, err := folder.ReadFile(ipx.Templates[name])
	if err != nil {
		log.Println(err)
	}

	tmpl := strings.Replace(string(f), "device-ipx25g", ipx.Index, -1)

	return tmpl
}

func (ipx *DeviceIPX25g) GetSourceIndexName() string {
	return ipx.Index
}

func (ipx *DeviceIPX25g) GetDestinationIndexName() string {
	return ipx.Index
}

func (sys *DeviceIPX25g) Convert(i string) (string, error) {
	var s DeviceIPX

	err := helpers.ReadJSON(i, &s)
	if err != nil {
		return "", err
	}

	doc := ConverDeviceIPXToDoc(s)

	j, err := helpers.WriteJSON(doc)
	if err != nil {
		return "", err
	}

	// j = strings.TrimSuffix(j, "\n")

	return j, nil
}

func ConverDeviceIPXToDoc(s DeviceIPX) Doc {
	// basic setup of the new document
	doc := Doc{
		Index: s.Index,
		Type:  "_doc",
		Id:    s.Id,
		Score: s.Score,
		Source: Ecs_source{
			Timestamp: s.Source.Timestamp,
			Message:   s.Source.Message,
			UUID:      s.Source.UUID,
			Version:   s.Source.Version,
			Event: EventCreated{
				Created: s.Source.Timestamp,
			},
			ECS: Ecs_version{
				Version: "1.3.0",
			},
			DataType: s.Source.DataType,
			Tags:     s.Source.Tags,
		},
	}

	// Populate host object
	doc.Source.Host = &Host{
		IP: &s.Source.Host,
	}

	// Populate annotation object if it exists.
	if s.Source.Annotation != nil {
		doc.Source.Annotation = s.Source.Annotation
	}

	switch doc.Source.DataType {

	case "lc":
		lc := &DeviceIPX_LC{
			Index:    s.Source.LCIndex,
			CardType: s.Source.LCCardType,
		}

		doc.Source.LC = lc

	case "chas_info":
		// some default values
		exetype := 3
		exeversion := 1
		version := 22888

		if s.Source.ChassisVersion != nil {
			version = *s.Source.ChassisVersion
		}

		if s.Source.EXEType != nil {
			exetype = *s.Source.EXEType
		}

		if s.Source.EXEVersion != nil {
			exeversion = *s.Source.EXEVersion
		}

		chassis := &DeviceIPX_Chassis{
			EXEType:         &exetype,
			EXEVersion:      &exeversion,
			NumberOfLCCards: s.Source.NumLCCards,
			NumberOfXCCards: s.Source.NumXCCards,
			ChassisVersion:  &version,
		}

		doc.Source.Chassis = chassis
		doc.Source.DataType = "chassis"

	case "port":
		port := &DeviceIPX_Port{
			CardLabel:      s.Source.CardLabel,
			CardPortNumber: s.Source.CardPortNumber,
			PortInterface:  s.Source.PortInterface,
			PortNumber:     s.Source.PortNum,
			RXMeasured:     s.Source.RXMeasured,
			RXMeasuredUnit: s.Source.RXMeasuredUnit,
			TXMeasured:     s.Source.TXMeasured,
			TXMeasuredUnit: s.Source.TXMeasuredUnit,
			Device:         s.Source.Device,
			DeviceSFP:      s.Source.DeviceSFP,
		}

		doc.Source.Port = port

	case "flow":
		flow := &DeviceIPX_Flow{
			EgressCardLabelList:     s.Source.EgressCardLabelList,
			EgressCardPortNumList:   s.Source.EgressCardPortNumList,
			EgressDevices:           s.Source.EgressDevices,
			EgressDeviceSFPs:        s.Source.EgressDeviceSFPs,
			EgressPorts:             s.Source.EgressPorts,
			EgressPortInterfaceList: s.Source.EgressPortInterfaceList,
			Egress:                  s.Source.Egress,
			IngressCardLabel:        s.Source.IngressCardLabel,
			IngressCardInterface:    s.Source.IngressCardInterface,
			IngressCardPortNum:      s.Source.IngressPortNumber,
			IngressDevice:           s.Source.IngressDevice,
			IngressDeviceSFP:        s.Source.IngressDeviceSFP,
			IngressPortNumber:       s.Source.IngressPortNumber,
			MulticastDestIP:         s.Source.MulticastDestIP,
			MeasuredBPS:             s.Source.MeasuredBPS,
			MeasuredKBPS:            s.Source.MeasuredKBPS,
			SignalGroup:             s.Source.SignalGroup,
			SignalType:              s.Source.SignalType,
			SignalLabel:             s.Source.SignalLabel,
			SourceMap:               s.Source.SourceMap,
		}

		doc.Source.Flow = flow
	}

	return doc
}
