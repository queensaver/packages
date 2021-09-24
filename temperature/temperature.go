package temperature

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Temperature struct {
	Temperature float64 `json:"temperature,omitempty"`
	BHiveID     string  `json:"bhiveId,omitempty"` //BHiveID is usually the Mac address of the raspberry pi in the bHive.
	SensorID    string  `json:"sensorId,omitempty"`
	Timestamp   int64   `json:"timestamp,omitempty"` //TODO: rename to epoch!
	Error       string  `json:"error,omitempty"`
	UUID        string  `json:"uuid,omitempty"`
}

func (t *Temperature) String() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

func (t *Temperature) SetUUID() {
	uuid := uuid.New()
	t.UUID = uuid.String()
}

func (t *Temperature) GetUUID() string {
	return t.UUID
}
