package main

import (
	"github.com/dcjohnson/init/config"
	"github.com/dcjohnson/init/spawn"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Remounting rootfs")
	err := remountRootFs()
	if err != nil {
		fmt.Println("Error remounting rootfs:", err.Error())
		powerOff()
	}

	fmt.Println("Mounting /proc")
	err = proc()
	if err != nil {
		fmt.Println("Error mounting proc:", err.Error())
		powerOff()
	}

	fmt.Println("Parsing Configuration")
	c, err := config.ParseConfiguration()
	if err != nil {
		printlnAndShutdown("Error parsing configuration:", err.Error())
	}

	fmt.Println("Spawning user shell")
	spawn.ForkAndExec(c.Shell)

	for {
		time.Sleep(time.Second)
	}

	powerOff()
}
