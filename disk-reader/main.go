package main

import (
	"flag"
	"fmt"
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

func init() {
	flag.StringVar(&vol, "volume", "", "Provide the name of the disk volume")
	flag.Parse()
}

// Todo: Index movies (@Movies) and series (@TV) through the disk into a single flat single entry
// Todo: Check encodings and quality
// Todo: Get details from IMDB api to rich up experience
func main() {
	fmt.Println(`
		  ______                                 ______             
 / _____) _                             (____  \            
( (____ _| |_  ____ _____ _____ ____     ____)  ) ___ _   _ 
 \____ (_   _)/ ___) ___ (____ |    \   |  __  ( / _ ( \ / )
 _____) )| |_| |   | ____/ ___ | | | |  | |__)  ) |_| ) X ( 
(______/  \__)_|   |_____)_____|_|_|_|  |______/ \___(_/ \_)
\`)

	var fp string

	if runtime.GOOS == "darwin" {
		fp = filepath.Join("/Volumes", vol)
	} else {
		panic("oops. Yet to map disks for this OS. Stay tuned")
	}

	vol, err := NewVolume(fp)

	if os.IsNotExist(err) {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error())
	}

	// Boot server
	BootServer(vol)
}
