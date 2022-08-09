package varroa

import (
	"bytes"
	"fmt"
	"io"
	"strings"
  "time"

	"mime/multipart"
  "encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/queensaver/openapi/golang/proto/services"
)

type Varroa struct {
  services.VarroaScanImagePostRequest
	UUID        string  `json:"uuid,omitempty"`
}

func (v *Varroa) String() ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

func (v *Varroa) GenerateUUID() {
	uuid := uuid.New()
	v.UUID = uuid.String()
}

func (v *Varroa) GetUUID() string {
	return v.UUID
}

func (v *Varroa) ClearUUID() {
	v.UUID = ""
}

func (v *Varroa) SetUUID(u string) {
  v.UUID = u
}

func (v *Varroa) IsMultipart() bool {
  return true
}

func (v *Varroa) Send(url string, token string) error {
  return PostImage(v.Scan, url, v.BhiveId, v.Epoch, token)
}

func PostImage(scan []byte, url string, bhiveId string, epoch int64, token string) error {

	client := &http.Client{
		Timeout: 300 * time.Second,
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	f, err := writer.CreateFormFile("scan", "scan")
	if err != nil {
		return err
	}
	_, err = io.Copy(f, bytes.NewReader(scan))
	if err != nil {
		return err
	}
	f, err = writer.CreateFormField("bhiveId")
	if err != nil {
		return err
	}
	_, err = io.Copy(f, strings.NewReader(bhiveId))
	if err != nil {
		return err
	}
	f, err = writer.CreateFormField("epoch")
	if err != nil {
		return err
	}
	_, err = io.Copy(f, strings.NewReader(fmt.Sprintf("%d", epoch)))
	if err != nil {
		return err
	}
	writer.Close()
	req, err := http.NewRequest("POST", url, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
  req.Header.Set("Q-Token", token)
	rsp, err := client.Do(req)
  if err != nil {
    return err
  }
	if rsp.StatusCode != http.StatusOK {
		fmt.Errorf("Request failed with response code: %d", rsp.StatusCode)
	}
	return nil
}
