package main

import (
	"log"
	"time"

	"github.com/tarm/serial"
	"github.com/Potewo/cobsserial"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600, ReadTimeout: 2 * time.Second}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	cobsserial.S = s
	err = cobsserial.Write([]byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	if err != nil {
		log.Fatal(err)
	}
	d, err := cobsserial.Read()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bytes: %#v", d)
}
