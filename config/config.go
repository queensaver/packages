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
	"crypto/tls"

	"github.com/queensaver/packages/logger"
)

type Config struct {
	BboxId   string // This is usually the Mac address of the raspberry pi in the BBox
	Bhive    []BHive
	Schedule string // Cron schedule from "github.com/robfig/cron/v3"
	Local    bool   //If the bhive is the same box ad the bbox
}

type BHive struct {
	BhiveId            string  //BHiveID is usually the Mac address of the raspberry pi in the bHive.
	RelayGpio          int     // The GPIO the relay is configured for.
	ScaleOffset        float64 // The offset in grams we substract from the measurement to tare it.
	ScaleReferenceUnit float64 // The reference unit we divide the measurement by to get the desired unit.
	Local              bool    // if the bhive software runs locally this is set to true.
	WittyPi            bool    // If the bhive has a witty pi to wake it up
	RecordSound		bool	// If the bhive should record sound
	SoundRecordingDuration int32 // The duration of the sound recording in seconds
}

func (s Config) String() ([]byte, error) {
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
	// TODO: fix certificate an disable insecure skip verify
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	httpClient := http.Client{
		Timeout: time.Second * 180,
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
		return nil, fmt.Errorf("Non-OK HTTP status: %d", res.StatusCode)
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

	return nil, errors.New("no bHive config found for this Raspberry Pi")

}

func Get(addr string, token string) (*Config, error) {
	// TODO: fix certificate an disable insecure skip verify
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
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

	logger.Info("Requesting config for BBox", "bbox_id", mac)
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Non-OK HTTP status: %d", res.StatusCode)
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
