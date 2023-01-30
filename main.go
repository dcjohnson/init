package main

import (
	"fmt"
	"github.com/dcjohnson/init/config"
	"github.com/dcjohnson/init/spawn"
	// "log"
	"os"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Remounting rootfs")
	err := remountRootFs()
	if err != nil {
		printlnAndShutdown("Error remounting rootfs:", err.Error())
	}

	fmt.Println("Mounting /proc")
	err = proc()
	if err != nil {
		printlnAndShutdown("Error mounting proc:", err.Error())
	}

	fmt.Println("Mounting /dev/pts")
	err = devpts()
	if err != nil {
		printlnAndShutdown("Error mounting dev:", err.Error())
	}

	fmt.Println("Parsing Configuration")
	c, err := config.ParseConfiguration()
	if err != nil {
		printlnAndShutdown("Error parsing configuration:", err.Error())
	}

	fmt.Println("Spawning user shell")
	fmt.Println("Exec'ing path:", c.Shell)
	f, err := os.Open("/dev/ptmx")
	if err != nil {
		printlnAndShutdown("Failed opening /dev/ptmx:", err.Error())
	}

	spawn.ForkAndExec(c.Shell, &syscall.ProcAttr{
		Files: []uintptr{
			0, 1, 2,
		},
		Sys: &syscall.SysProcAttr{
			Ctty:       int(f.Fd()),
			Noctty:     false,
			Foreground: true,
			Setpgid:    true,
		},
	})

	for {
		time.Sleep(time.Second)
	}

	powerOff()
}
