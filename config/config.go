package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	BboxId    string // This is usually the Mac address of the raspberry pi in the BBox
	Bhive     []BHive
	Schedule  string  // Cron schedule from "github.com/robfig/cron/v3"
}

type BHive struct {
	BhiveId            string  //BHiveID is usually the Mac address of the raspberry pi in the bHive.
	RelayGpio          int     // The GPIO the relay is configured for.
	ScaleOffset        int     // The offset in grams we substract from the measurement to tare it.
	ScaleReferenceUnit float64 // The reference unit we divide the measurement by to get the desired unit.
	Cameras            int     // Number of cameras in the BHive
}

func (s *Config) String() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

func getMacAddr() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	a := interfaces[1].HardwareAddr.String()
	if a != "" {
		r := strings.Replace(a, ":", "", -1)
		return r, nil
	}
	return "", nil
}

func GetBHiveConfig(addr string) (*BHive, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, addr, nil)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Non-OK HTTP status: %s", res.StatusCode))
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = json.Unmarshal(body, &config)
	if err != nil {
		return nil, err
	}
	mac, err := getMacAddr()
	if err != nil {
		return nil, err
	}
	for _, bhive := range config.Bhive {
		if bhive.BhiveId == mac {
			return &bhive, nil
		}
	}

	return nil, errors.New("No bHive config found for this Raspberry Pi")

}

func Get(addr string, token string) (*Config, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, addr, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Q-Token", token)
	mac, err := getMacAddr()
	if err != nil {
		return nil, err
	}
	req.Header.Set("BBoxID", mac)

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Non-OK HTTP status: %s", res.StatusCode))
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = json.Unmarshal(body, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil

}
