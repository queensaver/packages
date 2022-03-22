package varroa

import (
	"bytes"
	"fmt"
	"io"
	"strings"
  "time"

	"mime/multipart"
	"net/http"
)

func postImage(scan []byte, url string, bhiveId string, epoch int64, token string) error {

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
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		fmt.Errorf("Request failed with response code: %d", rsp.StatusCode)
	}
	return nil
}
