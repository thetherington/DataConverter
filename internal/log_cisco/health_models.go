package log_cisco

// "d_load_avg_15min": 0.54,
// "d_cpu_state_kernel": 0.024,
// "d_cpu_state_user": 0.016,
// "d_load_avg_1min": 0.44,
// "d_cpu_state_idle": 0.959,
// "d_load_avg_5min": 0.47,

// "cpu"
type CPU struct {
	CPULoadAvgOneMin     *float32 `json:"d_load_avg_1min,omitempty"`
	CPULoadAvgFiveMin    *float32 `json:"d_load_avg_5min,omitempty"`
	CPULoadAvgFifteenMin *float32 `json:"d_load_avg_15min,omitempty"`
	CPUStateKernel       *float32 `json:"d_cpu_state_kernel,omitempty"`
	CPUStateUser         *float32 `json:"d_cpu_state_user,omitempty"`
	CPUStateIdle         *float32 `json:"d_cpu_state_idle,omitempty"`
	Device
}

// "d_kernel": 0.05,
// "d_idle": 0.858,
// "d_user": 0.091,
// "i_cpuid": 3

// "cpu_core"
type CPU_CORE struct {
	CoreKernel *float32 `json:"d_kernel,omitempty"`
	CoreIdle   *float32 `json:"d_idle,omitempty"`
	CoreUser   *float32 `json:"d_user,omitempty"`
	CoreId     *int     `json:"i_cpuid,omitempty"`
	Device
}

// "d_memory_used_pct": 0.335,
// "l_memory_usage_free": 16327300000,
// "s_current_memory_status": "OK",
// "l_memory_usage_used": 8242668000,
// "l_memory_usage_total": 24569968000

// "memory"
type Memory struct {
	MemoryUsedPct    *float32 `json:"d_memory_used_pct,omitempty"`
	MemoryUsageFree  *int     `json:"l_memory_usage_free,omitempty"`
	MemoryUsageUsed  *int     `json:"l_memory_usage_used,omitempty"`
	MemoryUsageTotal *int     `json:"l_memory_usage_total,omitempty"`
	MemoryStatus     *string  `json:"s_current_memory_status,omitempty"`
	Device
}

// "i_processes_running": 2,
// "i_processes_total": 778

// "processes"
type Processes struct {
	ProcRunning *int `json:"i_processes_running,omitempty"`
	ProcTotal   *int `json:"i_processes_total,omitempty"`
	Device
}

// "s_status_ok_empty": "PS1 ok",
// "s_model_num": "NXA-PAC-1100W-PE2",
// "s_part_num": "341-1799-01",
// "s_clei_code": "CMUPAFMCAA",
// "s_type": "1100.00W 220v AC",
// "i_id": 2,
// "s_part_revision": "A0",
// "s_serial_num": "ART2226F8K7",
// "s_manuf_date": "Year 2018 Week 26",
// "s_hw_ver": "160"

// "hardware"
type Hardware struct {
	HWStatusOK      *string `json:"s_status_ok_empty,omitempty"`
	HWModelNum      *string `json:"s_model_num,omitempty"`
	HWPartNum       *string `json:"s_part_num,omitempty"`
	HWCleiCode      *string `json:"s_clei_code,omitempty"`
	HWType          *string `json:"s_type,omitempty"`
	HWID            *int    `json:"i_id,omitempty"`
	HWPartRevision  *string `json:"s_part_revision,omitempty"`
	HWSerialNum     *string `json:"s_serial_num,omitempty"`
	HWManufDate     *string `json:"s_manuf_date,omitempty"`
	HWhwVers        *string `json:"s_hw_ver,omitempty"`
	HWNumSubModules *string `json:"s_num_submods,omitempty"`
	Device
}

// "i_psnum": 1,
// "i_actual_input": 178,
// "s_ps_status": "Ok",
// "s_psmodel": "NXA-PAC-1100W-PE2",
// "i_actual_out": 156,
// "i_tot_capa": 1100

// "psu"
type PSU struct {
	PSUnum           *int    `json:"i_psnum,omitempty"`
	PSUActualInput   *int    `json:"i_actual_input,omitempty"`
	PSUStatus        *string `json:"s_ps_status,omitempty"`
	PSUModel         *string `json:"s_psmodel,omitempty"`
	PSUActualOut     *int    `json:"i_actual_out,omitempty"`
	PSUTotalCapactiy *int    `json:"i_tot_capa,omitempty"`
	Device
}

// "s_fanhwver": "--",
// "s_fanmodel": "NXA-FAN-65CFM-PE",
// "s_fanname": "Fan1(sys_fan1)",
// "s_fanstatus": "Ok",
// "s_fandir": "back-to-front"

// "fan"
type Fan struct {
	FanHWver  *string `json:"s_fanhwver,omitempty"`
	FanModel  *string `json:"s_fanmodel,omitempty"`
	FanName   *string `json:"s_fanname,omitempty"`
	FanStatus *string `json:"s_fanstatus,omitempty"`
	FanDir    *string `json:"s_fandir,omitempty"`
	Device
}

// "s_sensor": "BACK",
// "s_alarmstatus": "Ok",
// "i_minthres": 42,
// "i_tempmod": 1,
// "i_curtemp": 25,
// "i_majthres": 70

// "temp"
type Temp struct {
	TempSensor         *string `json:"s_sensor,omitempty"`
	TempAlarmStatus    *string `json:"s_alarmstatus,omitempty"`
	TempMinThreshold   *int    `json:"i_minthres,omitempty"`
	TempMajorThreshold *int    `json:"i_majthres,omitempty"`
	TempMod            *int    `json:"i_tempmod,omitempty"`
	TempCurTemperature *int    `json:"i_curtemp,omitempty"`
	Device
}

// "i_kern_uptm_days": 34,
// "i_kern_uptm_hrs": 20,
// "i_kern_uptm_mins": 0,
// "i_kern_uptm_secs": 15,

// "processes"
type Uptime struct {
	KernUptimeDays  *int `json:"i_kern_uptm_days,omitempty"`
	KernUptimeHours *int `json:"i_kern_uptm_hrs,omitempty"`
	KernUptimeMins  *int `json:"i_kern_uptm_mins,omitempty"`
	KernUptimeSecs  *int `json:"i_kern_uptm_secs,omitempty"`
	Device
}
