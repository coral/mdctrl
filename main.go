package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 19200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	_, err = s.Write([]byte{0x2B, 0x2B, 0x2B, 0x0D, 0x0A})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	_, err = s.Write([]byte("ATH^M"))
	[]byte{0x41, 0x54, 0x48, 0x5E, 0x4D, 0x0D, 0x0A}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("KADSBUGGEL")

}
