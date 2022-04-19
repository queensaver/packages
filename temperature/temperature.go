package temperature

import (
	"encoding/json"
	"github.com/google/uuid"

  "github.com/queensaver/openapi/golang/proto/models"
)

type Temperature struct {
  models.Temperature
	Error       string  `json:"error,omitempty"`
	UUID        string  `json:"uuid,omitempty"`
}

func (t *Temperature) String() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

func (t *Temperature) GenerateUUID() {
	uuid := uuid.New()
	t.UUID = uuid.String()
}

func (t *Temperature) GetUUID() string {
	return t.UUID
}

func (s *Temperature) ClearUUID() {
	s.UUID = ""
}

func (s *Temperature) SetUUID(u string) {
  s.UUID = u
}

func (t *Temperature) IsMultipart() bool {
  return false
}

func (t *Temperature) Send(url string, token string) error {
  return nil
}
