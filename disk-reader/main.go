package main

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
func main() {
	// 1. Read volume
	// 2. Index the volume of movies (@Movies) and series (@TV) through the disk into a single flat single entry 
	// 3. Check encodings and quality
	// 4. Get details from IMDB api to rich up experience
	// 5. Boot up the server
	// 6. Client connects and gets index page
	// 7. Choses item
	// 8. Open stream
}
