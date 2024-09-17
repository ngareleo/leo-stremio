package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

type File struct {
	Id       int
	Label    string
	ImageUrl string
}

type Volume struct {
	Location string
	Files    []File
}

func volumeExists(vol string) bool {
	_, err := os.Stat(vol)

	return err == nil
}

func NewVolume(vol string) (Volume, error) {

	info, err := os.Stat(vol)

	if os.IsNotExist(err) {
		return Volume{}, err
	}

	if err != nil {
		return Volume{}, errors.New("unknown error")
	}

	fmt.Printf("Volume name : %s\nVolume size: %d\n", info.Name(), info.Size())
	validFiles := make([]string, 0, info.Size()/8) // estimate start size

	if runtime.GOOS == "darwin" {
		last3 := func(w string) string {
			return w[len(w)-3:]
		}

		filepath.Walk(vol, func(path string, info fs.FileInfo, err error) error {
			if last3(path) == "mp4" || last3(path) == "mkv" {
				validFiles = append(validFiles, path)
			}
			return nil
		})
	}

	files := make([]File, 0, len(validFiles))

	for i, f := range validFiles {
		files = append(files, File{
			Id:    i,
			Label: f,
		})
	}

	return Volume{
		Location: vol,
		Files:    files,
	}, nil
}
