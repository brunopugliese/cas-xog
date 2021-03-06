package migration

import (
	"errors"
	"strconv"
	"strings"

	"github.com/andreluzz/cas-xog/constant"
	"github.com/andreluzz/cas-xog/model"
	"github.com/andreluzz/cas-xog/util"
	"github.com/beevik/etree"
	"github.com/tealeg/xlsx"
)

//ReadDataFromExcel used to create xog file from data in excel format. Only accept .xlsx extension
func ReadDataFromExcel(file *model.DriverFile) (string, error) {

	excelStartRowIndex, xog, templateInstanceElement, err := validateReadDataFromExcelDriverAttributes(file)
	if err != nil {
		return constant.Undefined, err
	}

	parent := templateInstanceElement.Parent()
	instanceCopy := templateInstanceElement.Copy()
	templateInstanceElement.Parent().RemoveChild(templateInstanceElement)

	xlFile, err := xlsx.OpenFile(util.ReplacePathSeparatorByOS(file.ExcelFile))
	if err != nil {
		return constant.Undefined, errors.New("migration - error opening excel. Debug: " + err.Error())
	}

	for index, row := range xlFile.Sheets[0].Rows {
		if index >= excelStartRowIndex {
			element := instanceCopy.Copy()
			for _, match := range file.MatchExcel {
				var e *etree.Element
				if match.XPath == constant.Undefined {
					e = element
				} else {
					e = element.FindElement(match.XPath)
				}

				if e == nil {
					return constant.Undefined, errors.New("migration - invalid xpath (" + match.XPath + "), element not found in template file")
				}

				value := constant.Undefined
				if match.Col-1 < len(row.Cells) {
					value = row.Cells[match.Col-1].String()
				}

				if match.RemoveIfNull && match.XPath != constant.Undefined && value == constant.Undefined {
					e.Parent().RemoveChild(e)
				} else {
					if match.AttributeName != constant.Undefined {
						e.CreateAttr(match.AttributeName, value)
					} else {
						if match.MultiValued && value != constant.Undefined {
							fillMultiValued(match, value, e)
						} else {
							e.SetText(value)
						}
					}
				}
			}
			parent.AddChild(element)
		}
	}
	xog.IndentTabs()
	return xog.WriteToString()
}

func fillMultiValued(match model.MatchExcel, value string, e *etree.Element) {
	separator := ";"
	if match.Separator != constant.Undefined {
		separator = match.Separator
	}
	elemName := "Value"
	if match.Element != constant.Undefined {
		elemName = match.Element
	}
	for _, val := range strings.Split(value, separator) {
		v := e.CreateElement(elemName)
		if match.Attr != constant.Undefined {
			v.CreateAttr(match.Attr, strings.TrimSpace(val))
		} else {
			v.SetText(strings.TrimSpace(val))
		}
		for _, a := range match.Attrs {
			v.CreateAttr(a.Name, a.Value)
		}
	}
}

func validateReadDataFromExcelDriverAttributes(file *model.DriverFile) (int, *etree.Document, *etree.Element, error) {

	xog := etree.NewDocument()
	err := xog.ReadFromFile(util.ReplacePathSeparatorByOS(file.Template))
	if err != nil {
		return 0, nil, nil, errors.New("migration - invalid template file. Debug: " + err.Error())
	}

	excelStartRowIndex := 0
	if file.ExcelStartRow != constant.Undefined {
		excelStartRowIndex, err = strconv.Atoi(file.ExcelStartRow)
		if err != nil {
			return 0, nil, nil, errors.New("migration - tag 'startRow' not a number. Debug:  " + err.Error())
		}
		excelStartRowIndex--
	}

	instance := constant.DefaultInstanceTag
	if file.InstanceTag != constant.Undefined {
		instance = file.InstanceTag
	}
	templateInstanceElement := xog.FindElement("//" + instance)
	if templateInstanceElement == nil {
		return 0, nil, nil, errors.New("migration - template invalid no instance element found")
	}

	return excelStartRowIndex, xog, templateInstanceElement, nil
}

//ExportInstancesToExcel used to create excel file with the data from xog file
func ExportInstancesToExcel(xog *etree.Document, file *model.DriverFile, folder string) error {
	util.ValidateFolder(folder)
	xlsxFile := xlsx.NewFile()
	sheet, _ := xlsxFile.AddSheet("Instances")

	for _, i := range xog.FindElements("//" + file.InstanceTag) {
		instance := etree.NewDocument()
		instance.AddChild(i)

		row := sheet.AddRow()

		for _, match := range file.MatchExcel {
			var e *etree.Element

			if match.XPath == constant.Undefined {
				e = instance.ChildElements()[0]
			} else {
				e = instance.FindElement(match.XPath)
			}

			cell := row.AddCell()

			if e == nil {
				cell.Value = ""
				continue
			}

			value := ""
			if match.AttributeName != "" {
				value = e.SelectAttrValue(match.AttributeName, "")
			} else {
				if match.MultiValued {
					separator := ";"
					if match.Separator != "" {
						separator = match.Separator
					}
					for _, v := range e.FindElements("//Value") {
						value += v.Text() + separator
					}
					value = value[:len(value)-1]
				} else {
					value = e.Text()
				}
			}
			cell.Value = value
		}
	}

	util.ValidateFolder(folder + util.GetPathFolder(util.ReplacePathSeparatorByOS(file.ExcelFile)))
	err := xlsxFile.Save(folder + util.ReplacePathSeparatorByOS(file.ExcelFile))
	if err != nil {
		return errors.New("migration - ExportInstancesToExcel saving excel error. Debug: " + err.Error())
	}

	return nil
}
