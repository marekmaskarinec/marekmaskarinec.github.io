package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"runtime"
	"os/exec"

	"github.com/gen2brain/dlgs"
	"github.com/homedepot/flop"
)

var ws string = "none"
var version string = "check your tophat install"
var defOptions []string = []string{"run", "package", "init", "open workspace", "install", "update"}
var istophat bool = false
var tophatPath string = "/usr/share/tophat/"
var cpops flop.Options
var slash string = "/"

func wsinit() {
	err := flop.Copy(tophatPath + "umka", ws + slash + "tophat", cpops)
	if err != nil {
		dlgs.Error("Error", fmt.Sprintf("Could not copy files. %s", err.Error()))
		fmt.Println(err.Error())
		return
	}

	err = flop.SimpleCopy(tophatPath + "preset.um", ws + "game.um")
	fmt.Println(ws + "game.um")
	if err != nil {
		dlgs.Error("Error", fmt.Sprintf("Could not copy files. %s", err.Error()))
		fmt.Println(err.Error())
		return
	}

	dlgs.Info("tophat init", "done")
}

func run() {
	bin := "tophat-linux"

	if runtime.GOOS == "windows" {
		bin = "tophat-windows.exe"
	}

	out, err := exec.Command(tophatPath + "bin" + slash + bin).Output()

	if err != nil && err.Error() != "signal: segmentation fault" {
		dlgs.Error("Error", err.Error())
		return
	}

	sout := string(out)

	if sout == "" {
		return
	}

	dlgs.Info("output", sout)
}

func main() {

	if runtime.GOOS == "windows" {
		tophatPath = os.Getenv("HOMEPATH") + "\\Documents\\tophat\\"
		slash = "\\"
	}
	fmt.Println(tophatPath)
	fmt.Println("")

	dat, err := ioutil.ReadFile(tophatPath + "version")
	version = string(dat)
	if err != nil {
		version = "check your tophat install"
	}

	isMap := map[bool]string{
		true: "is a",
		false: "isn't a",
	}

	cpops = flop.Options {
		Recursive: true,
		MkdirAll: true,
	}

	for {
		options := []string{}
		for _, o := range defOptions {
			switch o {
			case "run", "package":
				if ws != "none" && istophat {
					options = append(options, o)
				}
			case "init":
				if ws != "none" && !istophat {
					options = append(options, o)
				}
			default:
				options = append(options, o)
			}
		}

		wssplit := strings.Split(ws, "/")
		val, success, err := dlgs.List(fmt.Sprintf("tophat engine - %s", version), fmt.Sprintf("Workspace: %s\n%s %s a tophat workspace", wssplit[len(wssplit)-1], wssplit[len(wssplit)-1], isMap[istophat]), options)
		if !success || err != nil {
			return
		}

		switch val {
		case "open workspace":
			ws, _, _ = dlgs.File("choose a worspace", "", true)
			os.Chdir(ws)
			if ws[len(ws) - 1] != '/' {
				ws += "/"
			}
			_, err = os.Stat(ws + "/game.um");

			istophat = (err == nil)
		case "init":
			wsinit()
			_, err = os.Stat(ws + "/game.um");

			istophat = (err == nil)
		case "run":
			run()
		}
	}
}
