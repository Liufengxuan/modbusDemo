package main

import (
	"fmt"
	"time"

	"github.com/goburrow/modbus"
)

func main() {
	handler := modbus.NewRTUClientHandler("COM1")
	handler.BaudRate = 57600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 0x0A
	handler.Timeout = 5 * time.Second
	err := handler.Connect()
	if err != nil {
		fmt.Println("error :", err)
	}
	defer handler.Close()
	//--------------------------------------------------------------
	var val uint16
	i := 1
	for {
		time.Sleep(time.Millisecond * 300)
		i++
		if i%2 == 0 {
			val = uint16(0x0002)
		} else {
			val = uint16(0x0001)
		}
		client := modbus.NewClient(handler)
		results, err := client.WriteSingleRegister(uint16(0x0221), val)
		if err != nil {
			fmt.Println("error :", err)
		} else {
			fmt.Println(results)
			for _, v := range results {
				fmt.Printf("%#x\n", v)
			}
		}

		if i%2 == 0 {
			val = uint16(0x0001)
		} else {
			val = uint16(0x0003)
		}
		client = modbus.NewClient(handler)
		results, err = client.WriteSingleRegister(uint16(0x0220), val)
		if err != nil {
			fmt.Println("error :", err)
		} else {
			fmt.Println(results)
			for _, v := range results {
				fmt.Printf("%#x\n", v)
			}
		}

	}

}
