package main

import (
	"log"
	"time"

	"github.com/Potewo/cobsserial"
	"github.com/tarm/serial"
)

var data1 = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
var data2 = []byte{0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff}

func main() {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 115200, ReadTimeout: 2 * time.Second}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	cobsserial.S = s
	for {
		err = cobsserial.Write(data1)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
		err = cobsserial.Write(data2)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
