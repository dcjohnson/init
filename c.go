package main

import (
	"fmt"
	"os"
	"syscall"
)

func printlnAndShutdown(args ...any) {
	fmt.Println(args...)
	powerOff()
}

func powerOff() {
	syscall.Sync()
	syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
}

func proc() error {
	err := os.Mkdir("/proc", 0666)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return syscall.Mount("proc", "/proc", "proc", 0, "")
}

func remountRootFs() error {
	return syscall.Mount("/dev/root", "/", "", syscall.MS_REMOUNT, "")
}
