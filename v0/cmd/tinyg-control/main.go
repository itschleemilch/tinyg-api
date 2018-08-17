// This demo app uses the tinyg-control library and opens an interactive shell.
// Copyright (C) 2018 Sebastian Schleemilch
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301  USA

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/itschleemilch/huanyango/v1/vfdio"
	tinyg "github.com/itschleemilch/tinyg-api/v0/tinyg/controller"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var tgHandle *tinyg.TinygController

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "huanyango-cli-demo -port=/dev/ttyUSB0")
		fmt.Fprintln(flag.CommandLine.Output())
		fmt.Fprintln(flag.CommandLine.Output(), "Use G-Codes M3, M4, M4 and Snnnn.")
		fmt.Fprintln(flag.CommandLine.Output(), "? prints the current RPM.")
		fmt.Fprintln(flag.CommandLine.Output(), "$ outputs if connected.")
		fmt.Fprintln(flag.CommandLine.Output())
		fmt.Fprintln(flag.CommandLine.Output())
		flag.PrintDefaults()
	}
	var serialDevice *string = flag.String("port", "/dev/ttyMotorspindel", "USB Port. Linux default: /dev/ttyUSB0. On Windows use COMx, e.g. COM3. On Linux a symbolic link can be created using udev rules, see https://unix.stackexchange.com/a/183492.")
	var pollRate *int64 = flag.Int64("interval", 750, "RPM status readout interval in milliseconds. Default: 750.")
	var rpmHertzConversation *float64 = flag.Float64("rpm2hz", 3.47222, "Unit conversation from RPM to Hz. May be determined experimentally.")
	var maxRpm *int64 = flag.Int64("maxrpm", 11520, "Maximum allowed RPM for your spindle.")
	flag.Parse()

	var err error
	tgHandle, err = tinyg.NewController()
	if err != nil {
		panic(err)
	}
	err = tgHandle.Open("/dev/ttyTinyg") // symbolic link to /dev/ttyUSBn, see https://unix.stackexchange.com/a/183492
	defer tgHandle.Close()
	if err != nil {
		fmt.Println("Could not open serial port for Tinyg communction.")
		panic(err)
	}
	tgHandle.VfdOutput = vfdio.NewVfd()
	defer tgHandle.VfdOutput.Close()
	if err != nil {
		fmt.Println("Could not open serial port for VFD communication.")
		panic(err)
	}
	tgHandle.VfdOutput.Open(*serialDevice, uint16(*maxRpm), *rpmHertzConversation, *pollRate)

	http.HandleFunc("/api/", apiHome)
	http.HandleFunc("/api/state", apiState)
	http.HandleFunc("/api/exit", apiExit)
	http.HandleFunc("/api/gcode", apiGcode)
	http.HandleFunc("/api/file", apiGCodeFile)
	http.HandleFunc("/api/halt", apiHalt)
	http.HandleFunc("/api/continue", apiContinue)
	http.HandleFunc("/api/stop", apiStop)
	http.HandleFunc("/api/reset", apiReset)
	http.HandleFunc("/api/vfd", apiSpindle)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	fmt.Println("Starting Webserver...")
	glog.Fatal(http.ListenAndServe(":8080", nil))
}

func apiHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Api home. Use /api/state.")
}

func apiExit(w http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func apiGcode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")

	cmd, _ := url.QueryUnescape(req.URL.RawQuery)
	tgHandle.Write(cmd)
	fmt.Fprintf(w, `{"ok": true}`)
}

func apiHalt(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")

	tgHandle.FeedHold()
	fmt.Fprintf(w, `{"ok": true}`)
}

func apiContinue(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")

	tgHandle.FeedResume()
	fmt.Fprintf(w, `{"ok": true}`)
}

func apiStop(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")

	tgHandle.Flush()
	fmt.Fprintf(w, `{"ok": true}`)
}

func apiReset(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	tgHandle.TinygReset()
	fmt.Fprintf(w, `{"ok": true}`)
}

func apiState(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Write(tgHandle.StateJson())
}

func apiGCodeFile(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")

	req.ParseForm()

	gcodes := strings.Split(req.Form.Get("gcode"), "\n")
	for n, gcode := range gcodes {
		glog.Infoln("Received GCode Line #", n, ": ", gcode)
	}
	tgHandle.WriteLines(gcodes)

	fmt.Fprintf(w, `{"ok": true}`)
}

func apiSpindle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")

	values := make(map[string]int)
	values["f_set"] = -1
	values["f_is"] = -1
	values["rpm"] = -1
	values["dir"] = 0

	if tgHandle.VfdOutput != nil {
		values["rpm"] = int(tgHandle.VfdOutput.OutputRpm())
		values["f_is"] = int(tgHandle.VfdOutput.OutputFrequency())
		values["f_set"] = int(tgHandle.VfdOutput.FrequencySet())
		values["dir"] = int(tgHandle.VfdOutput.DirectionSet())
	}

	jsonBytes, err := json.Marshal(values)
	if err != nil {
		panic(err)
	}
	w.Write(jsonBytes)
}
