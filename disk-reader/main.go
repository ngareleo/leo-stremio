package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

// I am building the foundation of the "stream-box" (not final name)
//
// The big picture is to create a peer-to-peer streaming platform, for movies and tv-shows
// Users however must contribute to the network by having movies of their own
// and to enforce this, we will distribute hardware (raspberry-pies) that run special software
// The hardware will control seeking and authentication to make it simple
// A user will then wire the pie to a disk, and they can stream movies from that disk off a browser
// as well as from other users in the network
//
//
// This here is a stab at an element of all the ultimate goal.
// What I aim to achieve from this one, is simply being able to wire up your machine with a disk and stream on browser
// via network. By opening up the connection to the internet, means you can stream your movies on the disk from anywhere
// the peer-to-peer network will come later.

var vol string

func volumeExists(vol string) bool {
	_, err := os.Stat(vol)

	if os.IsNotExist(err) {
		return false
	}

	return err == nil
}

func init() {
	flag.StringVar(&vol, "volume", "", "Provide the name of the disk volume")
	flag.Parse()
}

func main() {
	// 1. Read volume
	// 2. Index the volume of movies (@Movies) and series (@TV) through the disk into a single flat single entry
	// 3. Check encodings and quality
	// 4. Get details from IMDB api to rich up experience
	// 5. Boot up the server
	// 6. Client connects and gets index page
	// 7. Choses item
	// 8. Open stream

	// walk the volume and look for "@Movies" and "@Tv"

	filePath := filepath.Join("/Volumes", vol)

	// Guards
	if !volumeExists(filePath) {
		panic("This volume doesn't exist")
	}

	volInfo, _ := os.Stat(filePath)

	fmt.Printf("Volume name : %s\nVolume size: %d\n", volInfo.Name(), volInfo.Size())
	validFiles := make([]string, 0, volInfo.Size()/8) // estimate start size
	if runtime.GOOS == "darwin" {
		last3 := func(w string) string {
			return w[len(w)-3:]
		}

		filepath.Walk(filePath, func(path string, info fs.FileInfo, err error) error {
			if last3(path) == "mp4" || last3(path) == "mkv" {
				validFiles = append(validFiles, path)
			}
			return nil
		})
	}

	files := make([]File, 0, len(validFiles))

	for i, f := range validFiles {
		files = append(files, File{
			Id: i,
			Label: f,
		})
	}
	
	// Boot server
	BootServer(Dir{
		Files: files,
	})
}
