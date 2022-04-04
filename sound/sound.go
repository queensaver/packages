package sound

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Sound struct {
	Sound    string `json:"sound,omitempty"`
	BhiveId  string `json:"bhiveId,omitempty"` //BHiveID is usually the Mac address of the raspberry pi in the bHive.
	Epoch    int64  `json:"epoch,omitempty"`
	Error    string `json:"error,omitempty"`
	UUID     string `json:"uuid,omitempty"`
	Duration int    `json:"duration,omitempty"`
}

func (s *Sound) String() ([]byte, error) {
	copy := *s
	copy.Sound = ""
	return json.MarshalIndent(copy, "", "  ")
}

func (s *Sound) GenerateUUID() {
	uuid := uuid.New()
	s.UUID = uuid.String()
}

func (s *Sound) GetUUID() string {
	return s.UUID
}

func (s *Sound) ClearUUID() {
	s.UUID = ""
}

func (s *Sound) SetUUID(u string) {
	s.UUID = u
}

func (s *Sound) IsMultipart() bool {
  return false
}

func (s *Sound) Send(url string, token string) error {
  return nil
}
