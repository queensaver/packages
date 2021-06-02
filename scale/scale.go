package scale

import (
	"encoding/json"
)

type Scale struct {
  Weight    float64 `json:"weight,omitempty"`
  //BHiveID is usually the Mac address of the raspberry pi in the bHive.
	BhiveId   string `json:"bhiveId,omitempty"`
	Epoch int64 `json:"epoch,omitempty"`
}

func (s *Scale) String() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}
