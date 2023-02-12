package main

import (
	"fmt"
	"github.com/dcjohnson/init/config"
	// "github.com/dcjohnson/init/spawn"
	// "log"
	"os"
	"os/exec"
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

	// err = printMounts()
	// if err != nil {
	// 	printlnAndShutdown("Error printing mounts:", err.Error())
	// }

	fmt.Println("Mounting /dev/pts")
	err = devpts()
	if err != nil {
		printlnAndShutdown("Error mounting /dev/pts:", err.Error())
	}

	printFilesDev()

	fmt.Println("Parsing Configuration")
	c, err := config.ParseConfiguration()
	if err != nil {
		printlnAndShutdown("Error parsing configuration:", err.Error())
	}

	fmt.Println("Spawning user shell")
	fmt.Println("Exec'ing path:", c.Shell)

	cmd := exec.Command(c.Shell)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Ctty:       0,
		// Foreground: true,
		Setsid: true,
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err = cmd.Start()
	if err != nil {
		printlnAndShutdown("failed to start shell:", c.Shell, err.Error())
	}

	for {
		time.Sleep(time.Second)
	}

	powerOff()
}

func printFiles(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	fmt.Println("Listing:", dir)
	for _, e := range entries {
		i, err := e.Info()
		if err != nil {
			return err
		}
		fmt.Println(e.Name(), i.Mode().String())
	}

	return nil
}

func printFilesDev() {
	err := printFiles("/dev")
	if err != nil {
		printlnAndShutdown("Error getting files:", err.Error())
	}
}

func printMounts() error {
	byts, err := os.ReadFile("/proc/mounts")
	if err != nil {
		return err
	}

	fmt.Println(string(byts))
	return nil
}
