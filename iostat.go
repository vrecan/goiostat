package main

import (
   "os"
   "bufio"
   "log"
   "strings"
   "./diskStat"
   "fmt"

)

const linuxDiskStats = "/proc/diskstats"

func main() {
  file,err := os.Open(linuxDiskStats)
  if nil != err {
		log.Fatal(err)
	}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
  	// fmt.Println(scanner.Text())
  	line := strings.Fields(scanner.Text())
  	stat,err := diskStat.LineToStat(line)
  	if(nil != err) {
  		log.Fatal(err);
  	}
  	fmt.Println(stat);

  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
	}
}
