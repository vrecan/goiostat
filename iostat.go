package main

import (
   "fmt"
   "os"
   "log"
)

const linuxDiskStats = "/proc/diskstats"

func main() {
  
	file, err := os.Open(linuxDiskStats)
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
	log.Fatal(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])

}
