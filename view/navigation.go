package view

import (
	"fmt"
	"github.com/andreluzz/cas-xog/constant"
	"github.com/andreluzz/cas-xog/log"
	"github.com/andreluzz/cas-xog/model"
	"github.com/andreluzz/cas-xog/xog"
	"os"
	"strings"
)

var startInstallingPackage int
var environments *model.Environments

func Home() {
	var err error

	log.InitLog()

	log.Info("\n")
	log.Info("--------------------------------------------\n")
	log.Info("##### CAS XOG Automation - Version %.1f #####\n", constant.VERSION)
	log.Info("--------------------------------------------\n")

	startInstallingPackage = 0

	model.LoadXMLReadList("xogRead.xml")

	environments, err = model.LoadEnvironmentsList("xogEnv.xml")
	if err != nil {
		log.Info("\n[CAS-XOG][red[Error]]: %s\n", err.Error())
	}

	Drivers()
}

func Interface() bool {
	var inputAction string

	if startInstallingPackage == 1 {
		startInstallingPackage = -1
		inputAction = "p"
	} else {
		log.Info("\nChoose action")
		log.Info("\n(l = Load XOG Driver, r = Read XOGs, w = Write XOGs, m = Create Migration, p = Install Package or x = eXit): ")
		fmt.Scanln(&inputAction)
	}

	action := strings.ToLower(inputAction)
	switch action {
	case "w", "r", "m":
		if xog.ValidateLoadedDriver() == false {
			log.Info("\n[CAS-XOG][red[ERROR]] - XOG driver not loaded. Try action 'l' to load a valid driver.\n")
			return false
		}
		if !Environments(action, environments) {
			return false
		}
		ProcessDriverFiles(xog.GetLoadedDriver(), action, environments)
	case "p":
		xog.LoadPackages(constant.FOLDER_PACKAGE, "packages/")
		output, selectedPackage, selectedVersion := Packages()
		if !output {
			return false
		}
		if !Environments(action, environments) {
			return false
		}
		err := InstallPackage(environments, selectedPackage, selectedVersion)
		if err != nil {
			log.Info("\n[CAS-XOG][red[PACKAGE]]: %s\n", err.Error())
			return false
		}
	case "l":
		Drivers()
	case "x":
		log.Info("\n[CAS-XOG][blue[Action exit selected]] - Press enter key to exit...\n")
		scanExit := ""
		fmt.Scanln(&scanExit)
		os.Exit(0)
		return true
	default:
		log.Info("\n[CAS-XOG][red[ACTION]] - Action not implemented!\n")
	}

	return false
}