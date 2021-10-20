package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	myjpeg "github.com/dlecorfec/image/jpeg"
)

var names = []string{"gray8.png"}
var quality = 75

func main() {
	for _, f := range names {
		img, err := readPng(f)
		if err != nil {
			log.Fatalf("read png err: %s", err)
		}
		var buf bytes.Buffer
		err = myjpeg.Encode(&buf, img, &myjpeg.Options{Quality: quality, Progressive: true})
		if err != nil {
			log.Fatalf("encode jpeg err: %s", err)
		}
		suffix := fmt.Sprintf(".q%d.jpg", quality)
		err = os.WriteFile(f+suffix, buf.Bytes(), 0644)
		if err != nil {
			log.Printf("write jpeg err: %s", err)
		}

	}
}

func readPng(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}
