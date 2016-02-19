package unifi

import (
	"encoding/json"
	"fmt"
)

// Devices returns all of the Devices for a specified site name.
func (c *Client) Devices(siteName string) ([]*Device, error) {
	var v struct {
		Devices []*Device `json:"data"`
	}

	req, err := c.newRequest(
		"GET",
		fmt.Sprintf("/api/s/%s/stat/device", siteName),
		nil,
	)
	if err != nil {
		return nil, err
	}

	_, err = c.do(req, &v)
	return v.Devices, err
}

// A Device is a Ubiquiti UniFi device, such as a UniFi access point.
type Device struct {
	ID      string
	Adopted bool

	// TODO(mdlayher): add more fields from unexported device type
}

// UnmarshalJSON unmarshals the raw JSON representation of a Device.
func (d *Device) UnmarshalJSON(b []byte) error {
	var dev device
	if err := json.Unmarshal(b, &dev); err != nil {
		return err
	}

	*d = Device{
		ID:      dev.ID,
		Adopted: dev.Adopted,
	}

	return nil
}

// A device is the raw structure of a Device returned from the UniFi Controller
// API.
type device struct {
	// TODO(mdlayher): give all fields appropriate names and data types.
	ID            string `json:"_id"`
	Adopted       bool   `json:"adopted"`
	Bytes         int    `json:"bytes"`
	ConfigVersion string `json:"cfgversion"`
	ConfigNetwork struct {
		IP   string `json:"ip"`
		Type string `json:"type"`
	} `json:"config_network"`
	DeviceID      string `json:"device_id"`
	EthernetTable []struct {
		MAC     string `json:"mac"`
		Name    string `json:"name"`
		NumPort int    `json:"num_port"`
	} `json:"ethernet_table"`
	GuestNumSta   int         `json:"guest-num_sta"`
	HasSpeaker    bool        `json:"has_speaker"`
	InformIP      string      `json:"inform_ip"`
	InformURL     string      `json:"inform_url"`
	IP            string      `json:"ip"`
	LastSeen      int         `json:"last_seen"`
	MAC           string      `json:"mac"`
	Model         string      `json:"model"`
	NaGuestNumSta int         `json:"na-guest-num_sta"`
	NaNumSta      int         `json:"na-num_sta"`
	NaUserNumSta  int         `json:"na-user-num_sta"`
	NgGuestNumSta int         `json:"ng-guest-num_sta"`
	NgNumSta      int         `json:"ng-num_sta"`
	NgUserNumSta  int         `json:"ng-user-num_sta"`
	NumSta        int         `json:"num_sta"`
	RadioNa       interface{} `json:"radio_na"`
	RadioNg       struct {
		BuiltinAntGain int    `json:"builtin_ant_gain"`
		BuiltinAntenna bool   `json:"builtin_antenna"`
		MaxTxpower     int    `json:"max_txpower"`
		MinTxpower     int    `json:"min_txpower"`
		Name           string `json:"name"`
		Radio          string `json:"radio"`
	} `json:"radio_ng"`
	RadioTable []struct {
		BuiltinAntGain int    `json:"builtin_ant_gain"`
		BuiltinAntenna bool   `json:"builtin_antenna"`
		MaxTxpower     int    `json:"max_txpower"`
		MinTxpower     int    `json:"min_txpower"`
		Name           string `json:"name"`
		Radio          string `json:"radio"`
	} `json:"radio_table"`
	RxBytes int    `json:"rx_bytes"`
	Serial  string `json:"serial,omitempty"`
	SiteID  string `json:"site_id"`
	Stat    struct {
		Bytes            int    `json:"bytes"`
		GuestNgTxBytes   int    `json:"guest-ng-tx_bytes"`
		GuestNgTxDropped int    `json:"guest-ng-tx_dropped"`
		GuestNgTxPackets int    `json:"guest-ng-tx_packets"`
		GuestTxBytes     int    `json:"guest-tx_bytes"`
		GuestTxDropped   int    `json:"guest-tx_dropped"`
		GuestTxPackets   int    `json:"guest-tx_packets"`
		Mac              string `json:"mac"`
		NgRxBytes        int    `json:"ng-rx_bytes"`
		NgRxPackets      int    `json:"ng-rx_packets"`
		NgTxBytes        int    `json:"ng-tx_bytes"`
		NgTxDropped      int    `json:"ng-tx_dropped"`
		NgTxPackets      int    `json:"ng-tx_packets"`
		RxBytes          int    `json:"rx_bytes"`
		RxPackets        int    `json:"rx_packets"`
		TxBytes          int    `json:"tx_bytes"`
		TxDropped        int    `json:"tx_dropped"`
		TxPackets        int    `json:"tx_packets"`
		UplinkRxBytes    int    `json:"uplink-rx_bytes"`
		UplinkRxPackets  int    `json:"uplink-rx_packets"`
		UplinkTxBytes    int    `json:"uplink-tx_bytes"`
		UplinkTxPackets  int    `json:"uplink-tx_packets"`
		UserNgRxBytes    int    `json:"user-ng-rx_bytes"`
		UserNgRxPackets  int    `json:"user-ng-rx_packets"`
		UserNgTxBytes    int    `json:"user-ng-tx_bytes"`
		UserNgTxDropped  int    `json:"user-ng-tx_dropped"`
		UserNgTxPackets  int    `json:"user-ng-tx_packets"`
		UserRxBytes      int    `json:"user-rx_bytes"`
		UserRxPackets    int    `json:"user-rx_packets"`
		UserTxBytes      int    `json:"user-tx_bytes"`
		UserTxDropped    int    `json:"user-tx_dropped"`
		UserTxPackets    int    `json:"user-tx_packets"`
	} `json:"stat"`
	State         int           `json:"state"`
	TxBytes       int           `json:"tx_bytes"`
	Type          string        `json:"type"`
	UplinkTable   []interface{} `json:"uplink_table"`
	UserNumSta    int           `json:"user-num_sta"`
	Version       string        `json:"version"`
	VwireEnabled  bool          `json:"vwireEnabled"`
	VwireTable    []interface{} `json:"vwire_table"`
	WlangroupIDNg string        `json:"wlangroup_id_ng"`
	XAuthkey      string        `json:"x_authkey"`
	XFingerprint  string        `json:"x_fingerprint"`
	XVwirekey     string        `json:"x_vwirekey"`
}