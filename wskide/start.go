package wskide

import (
	"time"

	"github.com/pkg/browser"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	startCmd    = kingpin.Command("start", "Start Development Enviroment")
	startDirArg = startCmd.Arg("dir", "Project dir").Default("project").String()
	initCmd     = kingpin.Command("init", "Initialise SDK Repository")
	stopCmd     = kingpin.Command("stop", "Stop Development Environment")
)

func startParse(cmd string) bool {
	switch cmd {
	case startCmd.FullCommand():
		err := Start(*startDirArg)
		ShowError(err)
		if err == nil {
			time.Sleep(2 * time.Second)
			browser.OpenURL(BrowserURL)
		}
		return true
	case stopCmd.FullCommand():
		Stop()
		return true
	case initCmd.FullCommand():
		err := Init()
		ShowError(err)
		return true
	}
	return false
}

// Start openwhisk-ide
func Start(dir string) error {
	err := Preflight(dir)
	if err != nil {
		return err
	}
	err = WhiskDeploy()
	if err != nil {
		return err
	}
	err = IdeDeploy(dir)
	if err != nil {
		return err
	}
	return nil
}

// Stop openwhisk-ide
func Stop() error {
	IdeDestroy()
	WhiskDestroy()
	return nil
}
