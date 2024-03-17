package synochat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func Send(url string, msg Message) error {
	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(&msg); err != nil {
		return err
	}

	result := struct {
		Error *struct {
			Code   int    `json:"code"`
			Errors string `json:"errors"`
		} `json:"error"`
		Success bool `json:"success"`
	}{}

	request := resty.New().R()
	request.SetFormData(map[string]string{
		"payload": buffer.String(),
	})
	request.SetResult(&result)
	_, err := request.Post(url)
	if err != nil {
		return err
	}
	if !result.Success {
		return fmt.Errorf("send error: code = %d, err = %s", result.Error.Code, result.Error.Errors)
	}
	return nil
}
