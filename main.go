package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/michaellindman/request"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%s\n", err)
	}
}

func run() error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return errors.Wrap(err, "directory")
	}
	cfg, err := NewConfig(dir + "/config.yml")
	if err != nil {
		return errors.Wrap(err, "config")
	}
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	file := flag.String("file", "", "path for file to be uploaded")
	flag.Parse()
	if *file != "" {
		upload, err := Upload(cfg.API.Key, cfg.API.Username, cfg.API.URL, *file)
		if err != nil {
			return errors.Wrap(err, "upload")
		}
		fmt.Printf("Uploaded %v (%v): %v\n", upload["original_filename"], upload["human_filesize"], upload["url"])
		return nil
	}
	flag.PrintDefaults()
	return nil
}

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
