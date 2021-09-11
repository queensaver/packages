package temperature

import (
	"encoding/json"
)

type Temperature struct {
	Temperature float64 `json:"temperature,omitempty"`
	BHiveID     string  `json:"bhiveId,omitempty"` //BHiveID is usually the Mac address of the raspberry pi in the bHive.
	SensorID    string  `json:"sensorId,omitempty"`
	Timestamp   int64   `json:"timestamp,omitempty"` //TODO: rename to epoch!
	Error       string  `json:"error,omitempty"`
}

func (t *Temperature) String() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}
