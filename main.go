package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/tty", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	_, err = s.Write([]byte("+++"))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	_, err = s.Write([]byte("ATH^M"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("KADSBUGGEL")

}
