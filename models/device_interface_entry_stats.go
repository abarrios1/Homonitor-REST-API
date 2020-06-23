package models

type DeviceInterfaceEntryStats struct {
	IFIndex string `json:"if_index"` // 1.3.6.1.2.1.2.2.1.1 (ifIndex)
	IFDesc string `json:"if_desc"` // 1.3.6.1.2.1.2.2.1.2 (ifDescr)
	IFType string `json:"if_type"` // 1.3.6.1.2.1.2.2.1.3 (ifType)
	IFSpeed string `json:"if_speed"` // 1.3.6.1.2.1.2.2.1.5 (ifSpeed)
	IFAdminStatus string `json:"if_admin_status"` // 1.3.6.1.2.1.2.2.1.7 (ifAdminStatus)
	IFOperStatus string `json:"if_oper_status"` // 1.3.6.1.2.1.2.2.1.8 (ifOperStatus)
	IFLastChange string `json:"if_last_change"` // 1.3.6.1.2.1.2.2.1.9 (ifLastChange)
	IFInOctets string `json:"if_in_octets"` // 1.3.6.1.2.1.2.2.1.10 (ifInOctets)
	IFOutOctets string `json:"if_out_octets"` // 1.3.6.1.2.1.2.2.1.16 (ifOutOctets)
}
