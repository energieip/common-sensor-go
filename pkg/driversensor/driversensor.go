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
	SoftwareVersion            float32 `json:"softwareVersion"`
	HardwareVersion            string  `json:"hardwareVersion"`
	IsBleEnabled               bool    `json:"isBleEnabled"`
	Temperature                int     `json:"temperature"`
	Humidity                   int     `json:"humidity"`
	Error                      int     `json:"error"`
	ResetNumbers               int     `json:"resetNumbers"`
	InitialSetupDate           float64 `json:"initialSetupDate"`
	LastResetDate              float64 `json:"lastResetDate"`
	Brightness                 int     `json:"brightness"`
	Presence                   bool    `json:"presence"`
	BrightnessCorrectionFactor int     `json:"brightnessCorrectionFactor"`
	ThresholdPresence          int     `json:"thresholdPresence"`
	TemperatureOffset          int     `json:"temperatureOffset"`
	BrightnessRaw              int     `json:"brightnessRaw"`
	LastMovement               int     `json:"lastMovement"`
	VoltageInput               int     `json:"voltageInput"`
	TemperatureRaw             int     `json:"temperatureRaw"`
	FriendlyName               string  `json:"friendlyName"`
	DumpFrequency              int     `json:"dumpFrequency"`
}

//SensorSetup initial setup send by the server when the driver is authorized
type SensorSetup struct {
	Mac                        string  `json:"mac"`
	Group                      *int    `json:"group,omitempty"`
	BrightnessCorrectionFactor *int    `json:"brightnessCorrectionFactor,omitempty"`
	ThresholdPresence          *int    `json:"thresholdPresence,omitempty"`
	TemperatureOffset          *int    `json:"temperatureOffset,omitempty"`
	IsBleEnabled               *bool   `json:"isBleEnabled,omitempty"`
	FriendlyName               *string `json:"friendlyName,omitempty"`
	SwitchMac                  string  `json:"switchMac"`
	IsConfigured               *bool   `json:"isConfigured,omitempty"`
	DumpFrequency              int     `json:"dumpFrequency"`
}

//SensorConf customizable configuration by the server
type SensorConf struct {
	Mac                        string  `json:"mac"`
	Group                      *int    `json:"group,omitempty"`
	BrightnessCorrectionFactor *int    `json:"brightnessCorrectionFactor,omitempty"`
	IsConfigured               *bool   `json:"isConfigured,omitempty"`
	ThresholdPresence          *int    `json:"thresholdPresence,omitempty"`
	TemperatureOffset          *int    `json:"temperatureOffset,omitempty"`
	IsBleEnabled               *bool   `json:"isBleEnabled,omitempty"`
	FriendlyName               *string `json:"friendlyName,omitempty"`
	DumpFrequency              *int    `json:"dumpFrequency,omitempty"`
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

//ToSensorSetup convert interface to SensorSetup object
func ToSensorSetup(val interface{}) (*SensorSetup, error) {
	var cell SensorSetup
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &cell)
	return &cell, err
}

//ToSensorConf convert map interface to Sensor object
func ToSensorConf(val interface{}) (*SensorConf, error) {
	var light SensorConf
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &light)
	return &light, err
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
