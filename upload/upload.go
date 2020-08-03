package upload

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"git.0cd.xyz/michael/request"
)

// Upload file to discourse server
func Upload(key, username, url, filepath string) (response map[string]interface{}, err error) {
	params := map[string]string{
		"type":        "upload",
		"synchronous": "true",
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		file.Close()
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, err
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", stat.Name())
	part.Write(contents)

	for key, val := range params {
		writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Api-Key":      key,
		"Api-Username": username,
		"Content-Type": writer.FormDataContentType(),
	}

	resp, err := request.API(http.MethodPost, url+"/uploads", headers, body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp.Body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
