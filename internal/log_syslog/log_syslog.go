package log_syslog

import (
	"embed"
	"log"
	"strings"

	"github.com/thetherington/DataConverter/internal/helpers"
	"golang.org/x/exp/slices"
)

// var scorpion_json = `{"_index":"log-syslog-2022.03.31","_type":"syslog","_id":"AX_hT0gihHj2rCo_IAJg","_score":1,"_source":{"annotation":{"general":{"device_name":"1-PROC-04-APP","system_name":"SCORPION-X18-APP-IPG12G","device_type":"LA"}},"syslog_pid":"1347","syslog_severity_code":6,"syslog_process_id":"1347","syslog_facility":"user-level","device_meta":"{\"prod\":\"scorpionx18ipg\",\"ip\":\"10.129.132.160\",\"ip-control-1\":\"10.129.132.160\",\"ip-control-2\":\"10.129.133.160\",\"mac\":\"00:02:c5:2f:12:af\",\"alias\":\"la-proc-4\"}","syslog_facility_code":1,"syslog_program":"ibcService","message":"<14>2019-03-08T05:00:53-07:00 Evertz ibcService[1347]: readPTPPacket: sfp 1, filtered pkt count 1! device_meta=({\"PROD\":\"SCORPIONX18IPG\",\"IP\":\"10.129.132.160\",\"IP-CONTROL-1\":\"10.129.132.160\",\"IP-CONTROL-2\":\"10.129.133.160\",\"MAC\":\"00:02:c5:2f:12:af\",\"ALIAS\":\"LA-PROC-4\"})","type":"syslog","syslog_message":"readPTPPacket: sfp 1, filtered pkt count 1! device_meta=({\"PROD\":\"SCORPIONX18IPG\",\"IP\":\"10.129.132.160\",\"IP-CONTROL-1\":\"10.129.132.160\",\"IP-CONTROL-2\":\"10.129.133.160\",\"MAC\":\"00:02:c5:2f:12:af\",\"ALIAS\":\"LA-PROC-4\"})","syslog_severity":"informational","tags":["data-syslog","forward-to-rabbitmq","rfc-3164"],"received_from":"10.129.132.101","@timestamp":"2022-03-31T18:49:03.015Z","syslog_hostname":"Evertz","@UUID":"93c93478-d83b-4c3d-8670-50f40e5a2a6a","syslog_timestamp":"2019-03-08T05:00:53-07:00","received_at":"2022-03-31T18:49:03.015Z","device_timestamp":"2019-03-08T12:00:53.000Z","@version":"1","host":"10.129.132.160","syslog_priority":"14","device":{"prod":"scorpionx18ipg","ip":"10.129.132.160","alias":"la-proc-4","ip-control-2":"10.129.133.160","ip-control-1":"10.129.132.160","mac":"00:02:c5:2f:12:af"}}}`
// var ipg_json = `{"_index":"log-syslog-2022.03.31","_type":"syslog","_id":"AX_hWU96hHj2rCo_MK5w","_version":1,"_score":null,"_source":{"annotation":{"general":{"device_name":"1-CODEC-15","system_name":"EV670-X45-HW-U0E16D-3G","device_type":"LA"}},"syslog_pid":"6200","syslog_severity_code":7,"syslog_process_id":"6200","syslog_facility":"user-level","device_meta":"{\"prod\":\"ev670-x45-hw\",\"ip\":\"10.129.133.127\",\"mac\":\"00:02:c5:2f:70:14\"}","syslog_facility_code":1,"syslog_program":"ipgOutService","message":"<15>1971-04-14T23:44:08+00:00 670HBX-X45 ipgOutService[6200]: IPGOUT: jam audio tstamp, Ch 10  device_meta=({\"PROD\":\"EV670-X45-HW\",\"IP\":\"10.129.133.127\",\"MAC\":\"00:02:C5:2F:70:14\"}) OpID []","type":"syslog","syslog_message":"IPGOUT: jam audio tstamp, Ch 10  device_meta=({\"PROD\":\"EV670-X45-HW\",\"IP\":\"10.129.133.127\",\"MAC\":\"00:02:C5:2F:70:14\"}) OpID []","syslog_severity":"debug","tags":["data-syslog","forward-to-rabbitmq","rfc-3164"],"received_from":"10.129.132.101","@timestamp":"2022-03-31T18:59:59.983Z","syslog_hostname":"670HBX-X45","@UUID":"cb5aaeb8-7ca7-436f-b7dc-c595aaaaf864","syslog_timestamp":"1971-04-14T23:44:08+00:00","received_at":"2022-03-31T18:59:59.983Z","device_timestamp":"1971-04-14T23:44:08.000Z","@version":"1","host":"10.129.133.127","syslog_priority":"15","device":{"prod":"ev670-x45-hw","ip":"10.129.133.127","mac":"00:02:c5:2f:70:14"}}}`
// var natx_json = `{"_index":"log-syslog-2022.03.31","_type":"syslog","_id":"AX_hLw5-hHj2rCo_6e65","_version":1,"_score":null,"_source":{"annotation":{"general":{"device_name":"1-PLN-SW-2A-IPX","system_name":"IPX","device_type":"LA"}},"syslog_severity_code":6,"syslog_process_id":"25053","syslog_facility":"user-level","syslog_facility_code":1,"syslog_program":"cfgjsonrpc","message":"<14>1 2022-03-31T18:15:20.436497+00:00 LA-NATX-M cfgjsonrpc 25053 - - [158|main] return len:645 buf:{\"id\":683093,\"jsonrpc\":\"2.0\",\"result\":{\"parameters\":[{\"id\":\"18@s\",\"type\":\"string\",\"value\":\"+NATX+MC32000+PORT260+L2+CTRLPORT3+SFP260+DATAPORT260\"},{\"id\":\"9@s\",\"type\":\"string\",\"value\":\"EXE-VSR-A\"},{\"id\n","type":"syslog","syslog_message":"[158|main] return len:645 buf:{\"id\":683093,\"jsonrpc\":\"2.0\",\"result\":{\"parameters\":[{\"id\":\"18@s\",\"type\":\"string\",\"value\":\"+NATX+MC32000+PORT260+L2+CTRLPORT3+SFP260+DATAPORT260\"},{\"id\":\"9@s\",\"type\":\"string\",\"value\":\"EXE-VSR-A\"},{\"id\n","syslog_version":"1","syslog_severity":"informational","tags":["data-syslog","forward-to-rabbitmq","rfc-5424"],"received_from":"10.129.132.105","@timestamp":"2022-03-31T18:13:51.094Z","syslog_hostname":"LA-NATX-M","@UUID":"0cfc4dc2-a937-4c83-92fe-fbf07a557c56","syslog_timestamp":"2022-03-31T18:15:20.436497+00:00","received_at":"2022-03-31T18:13:51.094Z","device_timestamp":"2022-03-31T18:15:20.436Z","@version":"1","host":"10.129.132.105","syslog_priority":"14"}}`
// var magnum_json = `{"_index":"log-syslog-2022.03.31","_type":"syslog","_id":"AX_hWU96hHj2rCo_MK51","_version":1,"_score":null,"_source":{"annotation":{"general":{"device_name":"LA-PLN-MAGNUM","system_name":"LA","device_type":"Magnum Server"}},"syslog_severity_code":3,"syslog_facility":"user-level","syslog_facility_code":1,"syslog_program":"mdldrvudx4k","message":"<11>2022-03-31T12:01:29.349255-07:00 mag-sdvn-pln-la-m mdldrvudx4k: 48c30f3c-9761-48aa-8168-07cb77ad3d8e ERROR:mdldrvudx4k.driver:Get request failed. Response [{u'type': u'integer', u'id': u'475.0@i', u'name': u'Output Video Link Number', u'error': {u'mesasge': u'Failed to retrieve data.', u'level': u'error'}}]","type":"syslog","syslog_message":"48c30f3c-9761-48aa-8168-07cb77ad3d8e ERROR:mdldrvudx4k.driver:Get request failed. Response [{u'type': u'integer', u'id': u'475.0@i', u'name': u'Output Video Link Number', u'error': {u'mesasge': u'Failed to retrieve data.', u'level': u'error'}}]","syslog_severity":"error","tags":["data-syslog","forward-to-rabbitmq","rfc-3164"],"received_from":"10.129.134.20","@timestamp":"2022-03-31T18:59:59.998Z","syslog_hostname":"mag-sdvn-pln-la-m","@UUID":"2e0f8d36-d951-43ad-8029-60fac3f5669a","syslog_timestamp":"2022-03-31T12:01:29.349255-07:00","received_at":"2022-03-31T18:59:59.998Z","device_timestamp":"2022-03-31T19:01:29.349Z","@version":"1","host":"10.129.134.20","syslog_priority":"11"}}`
// var ipgx_json = `{"_index":"log-syslog-2021.11.01","_type":"syslog","_id":"AXzbzMShYyFQyDyePKxe","_version":1,"_score":null,"_source":{"annotation":{"general":{"device_name":"DR2-IPG-17-D-11","device_type":"570IPG-X19-25G"}},"syslog_pid":"1419","syslog_severity_code":3,"syslog_process_id":"1419","syslog_facility":"user-level","device_meta":"{\"prod\":\"ipg570x1925g\",\"ip\":\"100.103.217.220\",\"ip-control-1\":\"192.168.245.17\",\"ip-control-2\":\"100.103.217.220\",\"mac\":\"00:02:c5:29:90:1a\",\"alias\":\"dr2.ipg.17.d.11\"}","syslog_facility_code":1,"syslog_program":"fpgaIOService","message":"<11>2021-11-01T09:55:19-04:00 sitara-platform fpgaIOService[1419]: vpid_output[15]: unsupported picture rate:0 device_meta=({\"PROD\":\"IPG570X1925G\",\"IP\":\"100.103.217.220\",\"IP-CONTROL-1\":\"192.168.245.17\",\"IP-CONTROL-2\":\"100.103.217.220\",\"MAC\":\"00:02:c5:29:90:1a\",\"ALIAS\":\"DR2.IPG.17.D.11\"})","type":"syslog","syslog_message":"vpid_output[15]: unsupported picture rate:0 device_meta=({\"PROD\":\"IPG570X1925G\",\"IP\":\"100.103.217.220\",\"IP-CONTROL-1\":\"192.168.245.17\",\"IP-CONTROL-2\":\"100.103.217.220\",\"MAC\":\"00:02:c5:29:90:1a\",\"ALIAS\":\"DR2.IPG.17.D.11\"})","syslog_severity":"error","tags":["data-syslog","forward-to-rabbitmq","rfc-3164"],"received_from":"100.103.217.251","@timestamp":"2021-11-01T13:59:59.995Z","syslog_hostname":"sitara-platform","@UUID":"2c787def-054d-4b71-a0ce-5e487a001340","syslog_timestamp":"2021-11-01T09:55:19-04:00","received_at":"2021-11-01T13:59:59.995Z","device_timestamp":"2021-11-01T13:55:19.000Z","@version":"1","host":"100.103.217.220","syslog_priority":"11","device":{"prod":"ipg570x1925g","ip":"100.103.217.220","alias":"dr2.ipg.17.d.11","ip-control-2":"100.103.217.220","ip-control-1":"192.168.245.17","mac":"00:02:c5:29:90:1a"}}}`
// var exe_json = `{"_index":"log-syslog-2021.11.01","_type":"syslog","_id":"AXzbzLlc87PwCVccY0gj","_version":1,"_score":null,"_source":{"annotation":{"general":{"device_name":"7N.EXE.004.ACQ","device_type":"EXE"}},"syslog_severity_code":4,"syslog_process_id":"9178","syslog_facility":"user-level","syslog_facility_code":1,"syslog_program":"sc","message":"<12>1 2021-11-01T13:55:28.043535+00:00 ACQ-EXE-X-FCNCS-01 sc 9178 - - EXESTATD: Status fault asserted: EXESTAT_Root___Multicasts___MultiCast_1_to_4000___MultiCast_2201_to_2300___MultiCast_2284___RxUnderBandwidth\n","type":"syslog","syslog_message":"EXESTATD: Status fault asserted: EXESTAT_Root___Multicasts___MultiCast_1_to_4000___MultiCast_2201_to_2300___MultiCast_2284___RxUnderBandwidth\n","syslog_version":"1","syslog_severity":"warning","tags":["data-syslog","forward-to-rabbitmq","rfc-5424"],"received_from":"100.103.216.10","@timestamp":"2021-11-01T13:59:59.169Z","syslog_hostname":"ACQ-EXE-X-FCNCS-01","@UUID":"556099a5-1094-43de-8de5-40afa5f7433c","syslog_timestamp":"2021-11-01T13:55:28.043535+00:00","received_at":"2021-11-01T13:59:59.169Z","device_timestamp":"2021-11-01T13:55:28.043Z","@version":"1","host":"100.103.216.10","syslog_priority":"12"}}`
// var failed_parse = `{"_index":"log-syslog-2022.03.31","_type":"syslog","_id":"AX_hOCrzhHj2rCo_-W-B","_version":1,"_score":null,"_source":{"@timestamp":"2022-03-31T18:23:48.017Z","@UUID":"2483e592-f1f4-4a85-b771-e0e139f919dc","@version":"1","host":"10.129.132.101","message":"<13>2019-01-04T23:41:24+00:00 Evertz \\n\" device_meta=({\"PROD\":\"SCORPIONX18IPG\",\"IP\":\"10.129.132.220\",\"IP-CONTROL-1\":\"10.129.132.220\",\"IP-CONTROL-2\":\"10.129.133.220\",\"MAC\":\"00:02:c5:32:10:3d\",\"ALIAS\":\"\"})","type":"syslog","tags":["data-syslog","forward-to-rabbitmq","_grokparsefailure"]}}`
// var opid = `{"_index":"log-syslog-2022.03.31","_type":"syslog","_id":"AX_hOCrzhHj2rCo_-W-P","_version":1,"_score":null,"_source":{"syslog_pid":"1072","device_meta":"{\"prod\":\"scorpionx18ipg\",\"ip\":\"10.129.132.220\",\"ip-control-1\":\"10.129.132.220\",\"ip-control-2\":\"10.129.133.220\",\"mac\":\"00:02:c5:32:10:3d\",\"alias\":\"\"}","syslog_program":"fpgaIOService","type":"syslog","syslog_severity":"informational","received_from":"10.129.132.101","@UUID":"5c74cc99-878b-4b5c-8e92-bbd7abda8e8d","syslog_timestamp":"2019-01-04T23:41:24+00:00","@OPID":"c46cf009-e6ab-4365-9caa-99b6ab009ff9","device_timestamp":"2019-01-04T23:41:24.000Z","@version":"1","host":"10.129.132.220","syslog_priority":"14","annotation":{"general":{"device_name":"1-PROC-09-APP","device_type":"SCORPION-X18-APP-IPG12G"}},"syslog_severity_code":6,"syslog_process_id":"1072","syslog_facility":"user-level","syslog_facility_code":1,"message":"<14>2019-01-04T23:41:24+00:00 Evertz fpgaIOService[1072]: OpID [c46cf009-e6ab-4365-9caa-99b6ab009ff9] \"rpc_sdp_take_handler: sdi[2], media_type[0], payload_type[100], sample[34], depth[0], width[1886377448], hight[24], exactframerate[41], tcs[4], colorimetry[0], interface[23], segmented[0]\\n\" device_meta=({\"PROD\":\"SCORPIONX18IPG\",\"IP\":\"10.129.132.220\",\"IP-CONTROL-1\":\"10.129.132.220\",\"IP-CONTROL-2\":\"10.129.133.220\",\"MAC\":\"00:02:c5:32:10:3d\",\"ALIAS\":\"\"})","syslog_message":"OpID [c46cf009-e6ab-4365-9caa-99b6ab009ff9] \"rpc_sdp_take_handler: sdi[2], media_type[0], payload_type[100], sample[34], depth[0], width[1886377448], hight[24], exactframerate[41], tcs[4], colorimetry[0], interface[23], segmented[0]\\n\" device_meta=({\"PROD\":\"SCORPIONX18IPG\",\"IP\":\"10.129.132.220\",\"IP-CONTROL-1\":\"10.129.132.220\",\"IP-CONTROL-2\":\"10.129.133.220\",\"MAC\":\"00:02:c5:32:10:3d\",\"ALIAS\":\"\"})","tags":["data-syslog","forward-to-rabbitmq","rfc-3164"],"@timestamp":"2022-03-31T18:23:48.030Z","syslog_hostname":"Evertz","received_at":"2022-03-31T18:23:48.030Z","device":{"prod":"scorpionx18ipg","ip":"10.129.132.220","alias":"","ip-control-2":"10.129.133.220","ip-control-1":"10.129.132.220","mac":"00:02:c5:32:10:3d"}}}`

var (
	//go:embed tmpl
	folder embed.FS

	templates = map[string]string{
		"alias":    "tmpl/syslog-index-alias.json",
		"mapping":  "tmpl/syslog-index-mapping.json",
		"settings": "tmpl/syslog-index-settings.json",
	}
)

type LogSyslog struct {
	num_errors int
	Index      string
	Templates  map[string]string
	Count      chan int
}

func New(index string) *LogSyslog {
	return &LogSyslog{
		num_errors: 0,
		Index:      index,
		Templates:  templates,
		Count:      make(chan int),
	}
}

func (sys *LogSyslog) IncrementErrors() {
	sys.num_errors++
}

func (sys *LogSyslog) ReturnErrors() int {
	return sys.num_errors
}

func (sys *LogSyslog) SendToCountChan(v int) {
	sys.Count <- v
}

func (sys *LogSyslog) GetCountChan() chan int {
	return sys.Count
}

func (sys *LogSyslog) GetTemplate(name string) string {
	f, err := folder.ReadFile(sys.Templates[name])
	if err != nil {
		log.Println(err)
	}

	tmpl := strings.Replace(string(f), "syslog-index", sys.Index, -1)

	return tmpl
}

func (sys *LogSyslog) GetSourceIndexName() string {
	return sys.Index
}

func (sys *LogSyslog) GetDestinationIndexName() string {
	return sys.Index
}

func (sys *LogSyslog) Convert(i string) (string, error) {
	var s Syslog

	err := helpers.ReadJSON(i, &s)
	if err != nil {
		return "", err
	}

	doc := ConverSyslogToDoc(s)

	j, err := helpers.WriteJSON(doc)
	if err != nil {
		return "", err
	}

	// j = strings.TrimSuffix(j, "\n")

	return j, nil
}

func ConverSyslogToDoc(s Syslog) Doc {
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
			Tags: s.Source.Tags,
		},
	}

	// check if device_timestamp is there and use it, otherwise use @timestamp
	if s.Source.DeviceTimestamp != nil {
		doc.Source.DeviceTimestamp = DeviceTimestamp{Timestamp: *s.Source.DeviceTimestamp}
	} else {
		doc.Source.DeviceTimestamp = DeviceTimestamp{Timestamp: s.Source.Timestamp}
	}

	// Populate host object
	host := Host{
		IP: &s.Source.Host,
	}
	if s.Source.SyslogHostname != nil {
		host.Name = s.Source.SyslogHostname
	}
	if s.Source.Device != nil {
		host.Device = s.Source.Device
	}
	doc.Source.Host = &host

	// Populate annotation object if it exists.
	if s.Source.Annotation != nil {
		doc.Source.Annotation = s.Source.Annotation
	}

	// Create log object but only if _grokparsefailure is not in tags.
	// Safe bet that if this exists then no syslog fields were created.
	if !slices.Contains(doc.Source.Tags, "_grokparsefailure") {
		log := Log{
			Original: s.Source.Message,
			Syslog: &SyslogStruct{
				Message: s.Source.SyslogMessage,
			},
		}

		if s.Source.SyslogFacility != nil && s.Source.SyslogFacilityCode != nil {
			log.Syslog.Facility = &struct {
				Name *string "json:\"name,omitempty\""
				Code *int    "json:\"code,omitempty\""
			}{
				Name: s.Source.SyslogFacility,
				Code: s.Source.SyslogFacilityCode,
			}

		}

		if s.Source.SyslogSeverity != nil && s.Source.SyslogSeverityCode != nil {
			log.Syslog.Severity = &struct {
				Name *string "json:\"name,omitempty\""
				Code *int    "json:\"code,omitempty\""
			}{
				Name: s.Source.SyslogSeverity,
				Code: s.Source.SyslogSeverityCode,
			}
		}

		doc.Source.Log = &log

		if s.Source.SyslogProgram != nil {
			var pid *string

			if s.Source.SyslogPID != nil {
				pid = s.Source.SyslogPID
			} else if s.Source.SyslogProcessId != nil {
				pid = s.Source.SyslogProcessId
			}

			process := Process{
				Name: s.Source.SyslogProgram,
				PID: &struct {
					Number   *string "json:\"number,omitempty\""
					Original *string "json:\"original,omitempty\""
				}{
					Number:   pid,
					Original: pid,
				},
			}

			doc.Source.Process = &process
		}
	}

	// Check if OPID exists
	if s.Source.OPID != nil {
		doc.Source.OPID = s.Source.OPID
	}

	return doc
}
