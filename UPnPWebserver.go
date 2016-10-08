// UPnPWebserver project UPnPWebserver.go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NebulousLabs/go-upnp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
func closeport(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Closing Port...")
	d, _ := upnp.Discover()
	_ = d.Clear(80)
	os.Exit(0)
}

func webServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/end", closeport)
	http.ListenAndServe(":80", nil)
}

func openPort(port int) {
	prt := uint16(port)
	d, _ := upnp.Discover()
	ip, _ := d.ExternalIP()
	fmt.Println("Your external IP is:", ip)
	_ = d.Forward(prt, "Server")
}

func main() {
	openPort(80)
	webServer()
	d, _ := upnp.Discover()
	//_ = d.Clear(80)
}
