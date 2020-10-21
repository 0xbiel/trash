package main

import (
  "fmt"
  "os"
)

func trash() {}

func help() {
  fmt.Println(`
 _               _   
| |_ ___ ___ ___| |_ 
|  _|  _| .'|_ -|   |
|_| |_| |__,|___|_|_|

		-> 0xbiel


Usage: trash [file] [folder]
  `)
}

func main() {
  if (len(os.Args) < 2) {
	help()
  } else {
	trash();
  }
}
