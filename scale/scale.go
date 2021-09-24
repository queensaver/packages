package scale

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Scale struct {
	Weight float64 `json:"weight,omitempty"`
	//BHiveID is usually the Mac address of the raspberry pi in the bHive.
	BhiveId string `json:"bhiveId,omitempty"`
	Epoch   int64  `json:"epoch,omitempty"`
	Error   string `json:"error,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

func (s *Scale) String() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

func (s *Scale) SetUUID() {
	uuid := uuid.New()
	s.UUID = uuid.String()
}

func (s *Scale) GetUUID() string {
	return s.UUID
}
