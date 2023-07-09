package main

import (
	"flag"
	"fmt"
	"github.com/Bermos/EV3-gRPC/internal/server"
	"github.com/Bermos/EV3-gRPC/internal/status"
	"github.com/Bermos/EV3-gRPC/internal/util"
	"log"
	"os"
	"time"
)

var (
	getHostname = flag.Bool("get-hostname", false, "only return hostname for this device")
	noMonitor   = flag.Bool("no-monitor", false, "do not create a display overlay")
	verify      = flag.Bool("verify", false, "exit with status code 0, check if executable")
	update      = flag.Bool("update", false, "check if new versions are available")
	address     = flag.String("address", ":9000", "address to bind to, can also be just the port ':9000'")
)

func main() {
	flag.Parse()

	if *verify {
		log.Printf("INFO - Verify mode, exiting...")
		os.Exit(0)
	}

	if *getHostname {
		fmt.Print(util.GetHostname())
		os.Exit(0)
	}

	if *update {
		util.CheckForNewVersion()
	}

	if !*noMonitor {
		status.Start(time.Second * 4)
	}

	server.StartServer(*address)
}
