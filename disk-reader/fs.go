package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
)

type File struct {
	Id       int
	Label    string
	ImageUrl string
}

func (file File) String() string {
	return fmt.Sprintf("Label [%s], Id [%d], ImageUrl [%s]", file.Label, file.Id, file.ImageUrl)
}

type Volume struct {
	Location string
	Files    []File
	FileMap  map[int]File
}

func NewVolume(vol string) (Volume, error) {

	info, err := os.Stat(vol)

	if os.IsNotExist(err) {
		return Volume{}, err
	}

	if err != nil {
		return Volume{}, errors.New("unknown error")
	}

	slog.Info("Volume found", "name", info.Name(), "size", info.Size())

	validfiles := make([]string, 0)

	last3 := func(w string) string {
		return w[len(w)-3:]
	}

	filepath.Walk(vol, func(path string, info fs.FileInfo, err error) error {
		if last3(path) == "mp4" || last3(path) == "mkv" {
			validfiles = append(validfiles, path)
		}
		return nil
	})

	files := make([]File, 0, len(validfiles))
	filemap := make(map[int]File, len(validfiles))

	for i, f := range validfiles {
		temp := File{
			Id:    i,
			Label: f,
		}
		files = append(files, temp)
		filemap[i] = temp
	}

	return Volume{
		Location: vol,
		Files:    files,
		FileMap:  filemap,
	}, nil
}

func (vol Volume) FindFileById(id int) (File, bool) {
	f, ok := vol.FileMap[id]
	return f, ok
}
