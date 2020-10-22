package main

import (
  "fmt"
  "os"
)

func createDir() {
	home := os.Getenv("HOME")
	dir := home + "/.trash"

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		if (err != nil) {
			panic(err)
		}
	}
}

//@@@ TODO: trash core functionality.
func trash() {}

func help() {
  fmt.Println(`
 _               _   
| |_ ___ ___ ___| |_ 
|  _|  _| .'|_ -|   |
|_| |_| |__,|___|_|_|

		~ 0xbiel

Usage: trash [file] [folder]
  `)
}

func main() {
	createDir()
	//@@@ TODO: add "-h" argument.
	if (len(os.Args) < 2) {
	  help()
	} else {
	  trash();
	}
}
