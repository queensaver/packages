package telemetry

import (
	"encoding/json"
	"github.com/google/uuid"
  "github.com/queensaver/openapi/golang/proto/models"
)

type Telemetry struct {
  *models.Telemetry
	BhiveId string `json:"m,omitempty"`
	Epoch   int64  `json:"e,omitempty"`
	UUID    string `json:"-"`
}

func (t *Telemetry) String() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

func (t *Telemetry) GenerateUUID() {
	uuid := uuid.New()
	t.UUID = uuid.String()
}

func (t *Telemetry) GetUUID() string {
	return t.UUID
}

func (t *Telemetry) ClearUUID() {
	t.UUID = ""
}

func (t *Telemetry) SetUUID(u string) {
	t.UUID = u
}

func (t *Telemetry) IsMultipart() bool {
  return false
}

func (t *Telemetry) Send(url string, token string) error {
  return nil
}
