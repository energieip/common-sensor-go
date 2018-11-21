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
	ID                         string  `json:"ID,omitempty"`
	IP                         string  `json:"ip"`
	Mac                        string  `json:"mac"`
	Group                      int     `json:"group"`
	Protocol                   string  `json:"protocol"`
	Topic                      string  `json:"topic"`
	SwitchMac                  string  `json:"switchMac"`
	IsConfigured               bool    `json:"isConfigured"`
	Version                    float32 `json:"version"`
	IsBleEnabled               bool    `json:"isBleEnabled"`
	Temperature                int     `json:"temperature"`
	Error                      int     `json:"error"`
	ResetNumbers               int     `json:"resetNumbers"`
	InitialSetupDate           float64 `json:"initialSetupDate"`
	LastResetDate              float64 `json:"lastResetDate"`
	Brigthness                 int     `json:"brightness"`
	Presence                   bool    `json:"presence"`
	BrigthnessCorrectionFactor int     `json:"brigthnessCorrectionFactor"`
	ThresoldPresence           int     `json:"thresoldPresence"`
	TemperatureOffset          int     `json:"temperatureOffset"`
	BrigthnessRaw              int     `json:"brigthnessRaw"`
	LastMovment                int     `json:"lastMovement"`
	VoltageInput               int     `json:"voltageInput"`
	TemperatureRaw             int     `json:"temperatureRaw"`
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
