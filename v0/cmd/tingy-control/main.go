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
	"fmt"
	tinyg "github.com/itschleemilch/tinyg-api/v0/tinyg/controller"
	"log"
	"net/http"
	"os"
)

var tgHandle *tinyg.TinygController

func main() {
	var err error
	tgHandle, err = tinyg.NewController()
	if err != nil {
		panic(err)
	}
	err = tgHandle.Open("/dev/ttyTinyg") // symbolic link to /dev/ttyUSBn, see https://unix.stackexchange.com/a/183492
	defer tgHandle.Close()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/api/", apiHome)
	http.HandleFunc("/api/state", apiState)
	http.HandleFunc("/api/exit", apiExit)
	http.HandleFunc("/api/gcode", apiGcode)

	log.Fatal(http.ListenAndServe(":8080", nil))
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

	tgHandle.SendData(req.URL.RawQuery)
	fmt.Fprintf(w, `{"ok": true}`)
}

func apiState(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Write(tgHandle.StateJson())
}
