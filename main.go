package main

import (
	// "bufio"
	// "strings"
	"fmt"
	// "os"
	"time"
)

const (
	fstab = "/etc/fstab"
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

	// byts, err := os.ReadFile("/proc/self/mounts")
	// if err != nil {
	// 	fmt.Println("err:", err.Error())
	// } else {
	// 	fmt.Printf("\n/proc/self/mounts:\n%s", string(byts))
	// }

	// writeCrap()

	for {
		time.Sleep(time.Second)
	}

	powerOff()
}

// func writeCrap() {
// 	err := os.Mkdir("/crap", 0666)
// 	if err != nil {
// 		fmt.Println("error making crapdir:", err.Error())
// 	}

// 	f, err := os.Create("/crap/crapfile")
// 	if err != nil {
// 		fmt.Println("error making crapfile:", err.Error())
// 	}

// 	f.Write([]byte("This is crap!"))
// }
