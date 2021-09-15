package temperature

import (
	"encoding/json"
)

type Temperature struct {
	Temperature float64 `json:"temperature,omitempty"`

	BHiveID   string `json:"bhiveid,omitempty"`
//BHiveID is usually the Mac address of the raspberry pi in the bHive.
	SensorID  string `json:"sensorid,omitempty"`

	Timestamp int64 `json:"epoch,omitempty"`
	UUID string `json:"uuid,omitempty"`

}

func (t *Temperature) String() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}
