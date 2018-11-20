package driversensor

import (
	"encoding/json"
)

const (
	DbName    = "status"
	TableName = "sensors"

	UrlHello   = "setup/hello"
	UrlStatus  = "status/dump"
	UrlSetup   = "setup/config"
	UrlSetting = "update/settings"
)

//Sensor driver representation
type Sensor struct {
	ID                         string  `gorethink:"id,omitempty" json:"ID"`
	Mac                        string  `gorethink:"mac" json:"mac"`
	Group                      int     `gorethink:"group" json:"group"`
	Protocol                   string  `gorethink:"protocol" json:"protocol"`
	Topic                      string  `gorethink:"topic" json:"topic"`
	SwitchMac                  string  `gorethink:"switch_mac" json:"switchMac"`
	IsConfigured               bool    `gorethink:"is_configured" json:"isConfigured"`
	Version                    float32 `gorethink:"version" json:"version"`
	IsBleEnabled               bool    `gorethink:"is_ble_enabled" json:"isBleEnabled"`
	Temperature                int     `gorethink:"temperature" json:"temperature"`
	Error                      int     `gorethink:"error" json:"error"`
	ResetNumbers               int     `gorethink:"reset_numbers" json:"resetNumbers"`
	InitialSetupDate           float64 `gorethink:"initial_setup_date" json:"initialSetupDate"`
	LastResetDate              float64 `gorethink:"last_reset_date" json:"lastResetDate"`
	Brigthness                 int     `gorethink:"brightness" json:"brightness"`
	Presence                   bool    `gorethink:"presence" json:"presence"`
	BrigthnessCorrectionFactor int     `gorethink:"brigthness_correction_factor" json:"brigthnessCorrectionFactor"`
	ThresoldPresence           int     `gorethink:"thresold_presence" json:"thresoldPresence"`
	TemperatureOffset          int     `gorethink:"temperature_offset" json:"temperatureOffset"`
	BrigthnessRaw              int     `gorethink:"brightness_raw" json:"brigthnessRaw"`
	LastMovment                int     `gorethink:"last_movement" json:"lastMovement"`
	VoltageInput               int     `gorethink:"voltage_input" json:"voltageInput"`
	TemperatureRaw             int     `gorethink:"temperature_raw" json:"temperatureRaw"`
}

//SensorSetup initial setup send by the server when the driver is authorized
type SensorSetup struct {
	Mac                        string `json:"mac"`
	Group                      *int   `json:"group"`
	BrigthnessCorrectionFactor *int   `json:"brigthnessCorrectionFactor"`
	ThresoldPresence           *int   `json:"thresoldPresence"`
	TemperatureOffset          *int   `json:"temperatureOffset"`
	IsBleEnabled               *bool  `json:"isBleEnabled"`
}

//SensorConf customizable configuration by the server
type SensorConf struct {
	Mac                        string `json:"mac"`
	Group                      *int   `json:"group"`
	BrigthnessCorrectionFactor *int   `json:"brigthnessCorrectionFactor"`
	IsConfigured               *bool  `json:"isConfigured"`
	ThresoldPresence           *int   `json:"thresoldPresence"`
	TemperatureOffset          *int   `json:"temperatureOffset"`
	IsBleEnabled               *bool  `json:"isBleEnabled"`
}

// ToMapInterface convert sensor struct in Map[string] interface{}
func (sensor Sensor) ToMapInterface() map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(sensor)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

//ToSensor convert interface to Sensor object
func ToSensor(val interface{}) (*Sensor, error) {
	var cell Sensor
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &cell)
	return &cell, err
}

// ToJSON dump sensor struct
func (sensor Sensor) ToJSON() (string, error) {
	inrec, err := json.Marshal(sensor)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}

// ToJSON dump sensor struct
func (sensor SensorSetup) ToJSON() (string, error) {
	inrec, err := json.Marshal(sensor)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}

//ToJSON dump struct in json
func (sensor SensorConf) ToJSON() (string, error) {
	inrec, err := json.Marshal(sensor)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}
