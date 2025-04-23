package main

import (
	"fmt"
	"os/exec"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

func main() {
	devName := "tap0"
	config := water.Config{
		DeviceType: water.TAP,
	}
	config.Name = devName

	ifce, err := water.New(config)
	if err != nil {
		panic(err)
	}

	//add ip: sudo ip addr add 10.xxx/14 dev xx
	ip := "10.1.0.10/24"
	if _, err := exec.Command("ip", "addr", "add", ip, "dev", devName).CombinedOutput(); err != nil {
		panic(err)
	}

	//link up: sudo ip link set  devName up
	if _, err := exec.Command("ip", "link", "set", devName, "up").CombinedOutput(); err != nil {
		panic(err)
	}

	frame := ethernet.Frame{}
	for {
		frame.Resize(1500)
		n, err := ifce.Read([]byte(frame))
		if err != nil {
			panic(err)
		}

		//write
		frame = frame[:n]
		fmt.Printf("dst:%s\n", frame.Destination())
		fmt.Printf("src:%s\n", frame.Source())
		fmt.Printf("EthernetType:%s\n", frame.Ethertype())
		fmt.Printf("Payload:%x\n", frame.Payload())
	}
}
