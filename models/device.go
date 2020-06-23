package models

type Device struct {
	Hostname     string `json:"hostname"`
	DeviceType   string `json:"device_type"`
	IPAddress    string `json:"ip_address"`
	SerialNumber string `json:"serial_number"`
	MacAddress   string `json:"mac_address"`
}
