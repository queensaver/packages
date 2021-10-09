package sound

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Sound struct {
	Sound []byte `json:"sound,omitempty"`
	//BHiveID is usually the Mac address of the raspberry pi in the bHive.
	BhiveId string `json:"bhiveId,omitempty"`
	Epoch   int64  `json:"epoch,omitempty"`
	Error   string `json:"error,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

func (s *Sound) String() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

func (s *Scale) SetUUID() {
	uuid := uuid.New()
	s.UUID = uuid.String()
}

func (s *Scale) GetUUID() string {
	return s.UUID
}

func (s *Sound) ClearUUID() string {
	s.UUID = ""
}
