package print

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/zpeters/speedtest/internal/debug"
	"github.com/zpeters/speedtest/internal/settings"
	"github.com/zpeters/speedtest/internal/sthttp"

	"github.com/codegangsta/cli"
	"github.com/dchest/uniuri"
)

// Server prints the results in "human" format
func Server(server sthttp.Server) {
	fmt.Printf("%-4s | %s (%s, %s)\n", server.ID, server.Sponsor, server.Name, server.Country)
}

// ServerReport prints restults in a machine useable format
func ServerReport(server sthttp.Server) {
	fmt.Printf("%s%s%s%s%s(%s,%s)%s", time.Now(), settings.REPORTCHAR, server.ID, settings.REPORTCHAR, server.Sponsor, server.Name, server.Country, settings.REPORTCHAR)
}

// EnvironmentReport is a debugging report helpful for debugging
func EnvironmentReport(c *cli.Context) {
	log.Printf("Env Report")
	log.Printf("-------------------------------\n")
	log.Printf("[User Environment]\n")
	log.Printf("Arch: %v\n", runtime.GOARCH)
	log.Printf("OS: %v\n", runtime.GOOS)
	log.Printf("IP: %v\n", sthttp.CONFIG.IP)
	log.Printf("Lat: %v\n", sthttp.CONFIG.Lat)
	log.Printf("Lon: %v\n", sthttp.CONFIG.Lon)
	log.Printf("ISP: %v\n", sthttp.CONFIG.Isp)
	log.Printf("Config: %s\n", sthttp.SpeedtestConfigURL)
	log.Printf("Servers: %s\n", sthttp.SpeedtestServersURL)
	r := uniuri.New()
	log.Printf("TEST: %v\n", r)
	log.Printf("-------------------------------\n")
	log.Printf("[Settings]\n")
	if c.Bool("debug") {
		log.Printf("Debug (user): %v\n", debug.DEBUG)
	} else {
		log.Printf("Debug (default): %v\n", debug.DEBUG)
	}
	if c.Bool("quiet") {
		log.Printf("Quiet (user): %v\n", debug.QUIET)
	} else {
		log.Printf("Quiet (default): %v\n", debug.QUIET)
	}
	if c.Int("numclosest") == 0 {
		log.Printf("NUMCLOSEST (default): %v\n", settings.NUMCLOSEST)
	} else {
		log.Printf("NUMCLOSEST (user): %v\n", settings.NUMCLOSEST)

	}
	if c.Int("numlatency") == 0 {
		log.Printf("NUMLATENCYTESTS (default): %v\n", settings.NUMLATENCYTESTS)
	} else {
		log.Printf("NUMLATENCYTESTS (user): %v\n", settings.NUMLATENCYTESTS)
	}
	if c.String("server") == "" {
		log.Printf("server (default none specified)\n")
	} else {
		log.Printf("server (user): %s\n", c.String("server"))
	}
	if c.String("reportchar") == "" {
		log.Printf("reportchar (default): %s\n", settings.REPORTCHAR)
	} else {
		log.Printf("reportchar (user): %s\n", c.String("reportchar"))
	}
	if c.String("algo") == "" {
		log.Printf("algo (default): %s\n", settings.ALGOTYPE)
	} else {
		log.Printf("algo (user): %s\n", c.String("algo"))
	}
	log.Printf("--------------------------------\n")
	log.Printf("[Mode]\n")
	log.Printf("Report: %v\n", c.Bool("report"))
	log.Printf("List: %v\n", c.Bool("list"))
	log.Printf("Ping: %v\n", c.Bool("Ping"))
	log.Printf("-------------------------------\n")
}
