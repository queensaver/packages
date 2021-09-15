package temperature

import (
	"encoding/json"
)

type Temperature struct {
	Temperature float64 `json:"temperature,omitempty"`
<<<<<<< HEAD

	BHiveID   string `json:"bhiveid,omitempty"`
//BHiveID is usually the Mac address of the raspberry pi in the bHive.
	SensorID  string `json:"sensorid,omitempty"`

	Timestamp int64 `json:"epoch,omitempty"`
	UUID string `json:"uuid,omitempty"`

=======
	BHiveID     string  `json:"bhiveId,omitempty"` //BHiveID is usually the Mac address of the raspberry pi in the bHive.
	SensorID    string  `json:"sensorId,omitempty"`
	Timestamp   int64   `json:"timestamp,omitempty"` //TODO: rename to epoch!
	Error       string  `json:"error,omitempty"`
>>>>>>> 8bd665514270354e00ea71da56b200d1efe31eed
}

func (t *Temperature) String() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}
