package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/michaellindman/discoupload/upload"
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
		upload, err := upload.Upload(cfg.API.Key, cfg.API.Username, cfg.API.URL, *file)
		if err != nil {
			return errors.Wrap(err, "upload")
		}
		fmt.Printf("Uploaded %v (%v): %v\n", upload["original_filename"], upload["human_filesize"], upload["url"])
		return nil
	}
	flag.PrintDefaults()
	return nil
}
