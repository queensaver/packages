package scale

import (
	"encoding/json"
)

type Scale struct {
	Weight    float64
	BhiveId   string //BHiveID is usually the Mac address of the raspberry pi in the bHive.
	Epoch int64
}

func (s *Scale) String() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}
