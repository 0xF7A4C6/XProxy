package utils

import (
	"fmt"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/gookit/color"
	"github.com/inancgumus/screen"
)

func PrintLogo() {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println(`
	__  ______                      
	\ \/ /  _ \ _ __ _____  ___   _ 
	 \  /| |_) | '__/ _ \ \/ / | | |
	 /  \|  __/| | | (_) >  <| |_| |
	/_/\_\_|   |_|  \___/_/\_\\__, |
                                   |___/ 
   `)
}

func Log(Content string) {
	date := strings.ReplaceAll(time.Now().Format("15:04:05"), ":", "<fg=353a3b>:</>")
	content := fmt.Sprintf("[%s] [%d] %s.", date, Valid, Content)

	content = strings.ReplaceAll(content, "DEAD", "<fg=f5291b>DEAD</>")
	content = strings.ReplaceAll(content, "ALIVE", "<fg=61eb42>ALIVE</>")

	for _, element := range []string{"(", ")", "[", "]", "#"} {
		content = strings.ReplaceAll(content, element, fmt.Sprintf("<fg=3d3d3d>%s</>", element))
	}

	color.Println(content)
}

func HandleError(Err error) bool {
	if Err != nil {
		if Config.Dev.Debug {
			fmt.Println(Err)
		}
		return true
	}

	return false
}

func SetTitle(title string) {
	if runtime.GOOS == "windows" {
		handle, err := syscall.LoadLibrary("Kernel32.dll")
		if HandleError(err) {
			return
		}
		
		defer syscall.FreeLibrary(handle)

		proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
		if HandleError(err) {
			return
		}

		syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(fmt.Sprintf("XProxy - github.com/its-vichy | %s", title)))), 0, 0)
	} else {
		fmt.Printf("\033]0;%s\007", title)
	}
}