package spawn

import (
	"fmt"
	"syscall"
)

func ForkAndExec(path string, pa *syscall.ProcAttr) {
	_, err := syscall.ForkExec(path, []string{}, pa)
	if err != nil {
		fmt.Println("error from ForkExec:", err.Error())
	}
}
