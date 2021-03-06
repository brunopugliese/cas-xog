package xog

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/andreluzz/cas-xog/constant"
	"github.com/andreluzz/cas-xog/migration"
	"github.com/andreluzz/cas-xog/model"
	"github.com/andreluzz/cas-xog/transform"
	"github.com/andreluzz/cas-xog/util"
	"github.com/andreluzz/cas-xog/validate"
	"github.com/beevik/etree"
)

var driverXOG *model.Driver

//LoadDriver load an specific driver defined by a path
func LoadDriver(path string) (int, error) {
	driverXOG = &model.Driver{}
	driverXOG.Clear()
	xmlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, errors.New("Error loading driver file - " + err.Error())
	}

	driverXOGTypePattern := model.DriverTypesPattern{}
	xml.Unmarshal(xmlFile, &driverXOGTypePattern)

	v, err := strconv.ParseFloat(driverXOGTypePattern.Version, 64)
	if err != nil || v < constant.Version {
		return 0, fmt.Errorf("invalid driver(%s) version, expected version %.1f or greater", path, constant.Version)
	}

	if len(driverXOGTypePattern.Files) > 0 {
		return 0, fmt.Errorf("invalid driver(%s) tag <file> is no longer supported", path)
	}

	types := reflect.ValueOf(&driverXOGTypePattern).Elem()
	typeOfT := types.Type()
	for i := 0; i < types.NumField(); i++ {
		t := types.Field(i)
		if t.Kind() == reflect.Slice {
			for _, f := range t.Interface().([]model.DriverFile) {
				f.Type = typeOfT.Field(i).Name
				f.ExecutionOrder = -1
				driverXOG.Files = append(driverXOG.Files, f)
			}
		}
	}

	doc := etree.NewDocument()
	doc.ReadFromBytes(xmlFile)

	if doc.Root().Tag != "xogdriver" {
		return 0, fmt.Errorf("invalid driver(%s) tag <%s> is incorrect", path, doc.Root().Tag)
	}

	for i, e := range doc.FindElements("//xogdriver/*") {
		tag := e.Tag
		path := e.SelectAttrValue("path", constant.Undefined)
		code := e.SelectAttrValue("code", constant.Undefined)
		for y, f := range driverXOG.Files {
			if f.ExecutionOrder == -1 && (strings.ToLower(f.GetXMLType()) == strings.ToLower(tag)) && (f.Path == path) && (f.Code == code) {
				driverXOG.Files[y].ExecutionOrder = i
				break
			}
		}
	}

	sort.Sort(model.ByExecutionOrder(driverXOG.Files))

	driverXOG.Version = driverXOGTypePattern.Version
	driverXOG.AutomaticWrite = driverXOGTypePattern.AutomaticWrite
	driverXOG.FilePath = path

	return len(driverXOG.Files), nil
}

//GetLoadedDriver returns the pointer to the loaded driver
func GetLoadedDriver() *model.Driver {
	return driverXOG
}

//ValidateLoadedDriver verify if the driver was loaded successfully
func ValidateLoadedDriver() bool {
	if driverXOG == nil {
		return false
	}
	return len(driverXOG.Files) > 0
}

//GetDriversList returns a list of available drivers in the defined folder
func GetDriversList(folder string) ([]model.Driver, error) {
	var driversList []model.Driver

	err := filepath.Walk(folder, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return errors.New("invalid driver folder")
		}
		if !f.IsDir() && util.GetExtension(path) == ".driver" {
			driver := new(model.Driver)
			driver.Info = f
			driver.FilePath = path
			folder := util.GetDirectFolder(path)
			if folder != "drivers" {
				driver.Folder = folder
			}
			driversList = append(driversList, *driver)
		}
		return nil
	})

	if err != nil || len(driversList) == 0 {
		return nil, errors.New("driver folder not found or empty")
	}

	var driversListSorted []model.Driver
	for _, d := range driversList {
		if d.Folder == constant.Undefined {
			driversListSorted = append(driversListSorted, d)
		}
	}
	for _, d := range driversList {
		if d.Folder != constant.Undefined {
			driversListSorted = append(driversListSorted, d)
		}
	}
	return driversListSorted, nil
}

//ProcessDriverFile execute a xog files and return the output
func ProcessDriverFile(file *model.DriverFile, action, sourceFolder, outputFolder string, environments *model.Environments, soapFunc util.Soap) model.Output {
	output := model.Output{Code: constant.OutputSuccess, Debug: constant.Undefined}

	if action == constant.Migrate && file.Type != constant.TypeMigration {
		output.Code = constant.OutputWarning
		output.Debug = "Use action 'r' to this type(" + file.Type + ") of file"
		return output
	} else if action == constant.Read && file.Type == constant.TypeMigration {
		output.Code = constant.OutputWarning
		output.Debug = "Use action 'm' to this type(" + file.Type + ") of file"
		return output
	}

	if action == constant.Migrate {
		return processMigrate(file, outputFolder)
	}

	err := file.InitXML(action, sourceFolder)
	if err != nil {
		output.Code = constant.OutputError
		output.Debug = err.Error()
		return output
	}
	if action == constant.Write {
		iniTagRegexpStr, endTagRegexpStr := file.TagCDATA()
		if iniTagRegexpStr != constant.Undefined && endTagRegexpStr != constant.Undefined {
			transformedString := transform.IncludeCDATA(file.GetXML(), iniTagRegexpStr, endTagRegexpStr)
			file.SetXML(transformedString)
		}
	}
	err = file.RunXML(action, sourceFolder, environments, soapFunc)
	if err != nil {
		output.Code = constant.OutputError
		output.Debug = err.Error()
		return output
	}
	xogResponse := etree.NewDocument()
	xogResponse.ReadFromString(file.GetXML())
	output, err = validate.Check(xogResponse)
	if err != nil {
		output.Code = constant.OutputError
		output.Debug = err.Error()
		file.Write(constant.FolderDebug)
		return output
	}
	if action == constant.Read {
		output := processDriverFileRead(file, xogResponse, outputFolder)
		if output.Code != constant.OutputSuccess {
			file.Write(constant.FolderDebug)
			return output
		}
	}

	file.Write(outputFolder)
	return output
}

func processMigrate(file *model.DriverFile, outputFolder string) model.Output {
	output := model.Output{Code: constant.OutputSuccess, Debug: constant.Undefined}
	resp, err := migration.ReadDataFromExcel(file)
	if err != nil {
		output.Code = constant.OutputError
		output.Debug = err.Error()
		return output
	}
	if file.GetInstanceTag() != constant.Undefined && file.InstancesPerFile > 0 {
		xogResponse := etree.NewDocument()
		xogResponse.ReadFromString(resp)
		splitInstancesIntoMultipleFiles(file, xogResponse, outputFolder)
	}
	file.SetXML(resp)
	file.Write(outputFolder)
	return output
}

func processDriverFileRead(file *model.DriverFile, xogResponse *etree.Document, outputFolder string) model.Output {
	output := model.Output{Code: constant.OutputSuccess, Debug: constant.Undefined}

	var auxResponse *etree.Document
	if file.NeedAuxXML() {
		auxResponse = etree.NewDocument()
		auxResponse.ReadFromString(file.GetAuxXML())
		output, err := validate.Check(auxResponse)
		if err != nil {
			output.Code = constant.OutputError
			output.Debug = "aux validation - " + err.Error()
			return output
		}
	}
	err := transform.Execute(xogResponse, auxResponse, file)

	if file.GetInstanceTag() != constant.Undefined && file.InstancesPerFile > 0 {
		splitInstancesIntoMultipleFiles(file, xogResponse, outputFolder)
	}

	str, _ := xogResponse.WriteToString()
	file.SetXML(str)
	if err != nil {
		output.Code = constant.OutputError
		output.Debug = err.Error()
		return output
	}
	iniTagRegexpStr, endTagRegexpStr := file.TagCDATA()
	if iniTagRegexpStr != constant.Undefined && endTagRegexpStr != constant.Undefined {
		transformedString := transform.IncludeCDATA(file.GetXML(), iniTagRegexpStr, endTagRegexpStr)
		file.SetXML(transformedString)
	}

	if file.ExportToExcel {
		migration.ExportInstancesToExcel(xogResponse, file, constant.FolderMigration)
	}

	return output
}

func splitInstancesIntoMultipleFiles(file *model.DriverFile, xogResponse *etree.Document, outputFolder string) {

	instanceTagPath := "//" + file.GetInstanceTag()

	parentTagPath := "//" + xogResponse.FindElement(instanceTagPath).Parent().Tag

	splitXogResponse := xogResponse.Copy()
	for _, e := range splitXogResponse.FindElements(instanceTagPath) {
		e.Parent().RemoveChild(e)
	}
	instances := xogResponse.FindElements(instanceTagPath)

	totalFiles := math.Ceil(float64(len(instances)) / float64(file.InstancesPerFile))
	path := util.GetPathWithoutExtension(file.Path)
	for i := 0; i < int(totalFiles); i++ {
		file.Path = fmt.Sprintf("%s_%03d.xml", path, i+1)
		xog := splitXogResponse.Copy()
		elementParent := xog.FindElement(parentTagPath)
		for z := file.InstancesPerFile * i; z < file.InstancesPerFile*(i+1); z++ {
			if z < len(instances) {
				elementParent.AddChild(instances[z].Copy())
			}
		}
		xog.IndentTabs()
		s, _ := xog.WriteToString()
		file.SetXML(s)
		file.Write(outputFolder)
	}
	file.Path = "complete_" + path + ".xml"
}

//CreateFileFolder creates the folders according to the action (read, write or migrate)
func CreateFileFolder(action, fileType, path string) (string, string) {
	sourceFolder := constant.Undefined
	outputFolder := constant.Undefined

	switch action {
	case constant.Read:
		sourceFolder = constant.FolderRead
		outputFolder = constant.FolderWrite
		util.ValidateFolder(constant.FolderRead + fileType + util.GetPathFolder(path))
	case constant.Write:
		sourceFolder = constant.FolderWrite
		outputFolder = constant.FolderDebug
	case constant.Migrate:
		sourceFolder = constant.FolderWrite
		outputFolder = constant.FolderMigration
	}

	util.ValidateFolder(outputFolder + fileType + util.GetPathFolder(path))

	return sourceFolder, outputFolder
}
