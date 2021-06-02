package scale

import (
	"encoding/json"
)

type Scale struct {
	weight    float64
	bhiveId   string //BHiveID is usually the Mac address of the raspberry pi in the bHive.
	epoch int64
}

func (s *Scale) String() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}
