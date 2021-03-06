package transform

import (
	"github.com/andreluzz/cas-xog/constant"
	"github.com/andreluzz/cas-xog/model"
	"github.com/beevik/etree"
	"testing"
)

func TestExecuteToReturnView(t *testing.T) {
	file := model.DriverFile{
		Code:    "*",
		ObjCode: "obj_sistema",
		Type:    constant.TypeView,
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_full_xog.xml")
	err := Execute(xog, nil, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewSourcePartition(t *testing.T) {
	file := model.DriverFile{
		Code:            "*",
		ObjCode:         "obj_sistema",
		Type:            constant.TypeView,
		SourcePartition: "partition10",
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, nil, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_source_partition_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewSourceTargetPartition(t *testing.T) {
	file := model.DriverFile{
		Code:            "*",
		ObjCode:         "obj_sistema",
		Type:            constant.TypeView,
		SourcePartition: "partition10",
		TargetPartition: "partition20",
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, nil, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_source_target_partition_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewCodePattern(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "*obj_sistemaCreate",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_code_pattern_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewSingle(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_single_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewSingleList(t *testing.T) {
	file := model.DriverFile{
		Type:    constant.TypeView,
		Code:    "odf.obj_sistemaList",
		ObjCode: "obj_sistema",
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_single_list_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewSingleNotInTarget(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_aux_without_view.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_result_target_without_code_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewSingleSection(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_single_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewReplaceSection(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionReplace,
				SourcePosition: "1",
				TargetPosition: "1",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_section_replace_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewRemoveSection(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionRemove,
				TargetPosition: "3",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_section_remove_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewUpdateSection(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		TargetPartition: "partition20",
		Sections: []model.Section{
			{
				Action:         constant.ActionUpdate,
				SourcePosition: "1",
				TargetPosition: "1",
				Fields: []model.SectionField{
					{
						Code:         "analista",
						Column:       constant.ColumnLeft,
						InsertBefore: "created_by",
					},
					{
						Code:         "status",
						Column:       constant.ColumnLeft,
						InsertBefore: "created_by",
					},
					{
						Code:   "status_novo",
						Column: constant.ColumnRight,
					},
					{
						Code:   "created_date",
						Remove: true,
					},
				},
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition20_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_section_update_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewUpdateSectionColumns(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		TargetPartition: "partition20",
		Sections: []model.Section{
			{
				Action:         constant.ActionUpdate,
				SourcePosition: "1",
				TargetPosition: "1",
				Fields: []model.SectionField{
					{
						Code:   "analista",
						Column: constant.ColumnLeft,
					},
					{
						Code:         "status",
						Column:       constant.ColumnLeft,
						InsertBefore: "created_by",
					},
					{
						Code:         "status_novo",
						Column:       constant.ColumnRight,
						InsertBefore: "created_date",
					},
					{
						Code:   "created_date",
						Remove: true,
					},
				},
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition20_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_section_update_other_columns_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewUpdateSectionTargetNoRightColumn(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		TargetPartition: "partition20",
		Sections: []model.Section{
			{
				Action:         constant.ActionUpdate,
				SourcePosition: "1",
				TargetPosition: "1",
				Fields: []model.SectionField{
					{
						Code:   "analista",
						Column: constant.ColumnLeft,
					},
					{
						Code:         "status",
						Column:       constant.ColumnLeft,
						InsertBefore: "created_by",
					},
					{
						Code:   "status_novo",
						Column: constant.ColumnRight,
					},
				},
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition20_no_right_column_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_section_update_other_columns_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewUpdateSectionTargetNoLeftColumn(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		TargetPartition: "partition20",
		Sections: []model.Section{
			{
				Action:         constant.ActionUpdate,
				SourcePosition: "1",
				TargetPosition: "1",
				Fields: []model.SectionField{
					{
						Code:   "analista",
						Column: constant.ColumnLeft,
					},
					{
						Code:   "status",
						Column: constant.ColumnLeft,
					},
					{
						Code:   "created_by",
						Column: constant.ColumnLeft,
					},
					{
						Code:         "status_novo",
						Column:       constant.ColumnRight,
						InsertBefore: "created_date",
					},
					{
						Code:   "created_date",
						Remove: true,
					},
				},
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition20_no_left_column_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_section_update_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnViewInsertSection(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionInsert,
				SourcePosition: "1",
				TargetPosition: "1",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_section_insert_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnErrorSectionsWithoutSingleView(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "*",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionInsert,
				SourcePosition: "1",
				TargetPosition: "1",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if code is * and sections defined")
	}
}

func TestExecuteToReturnErrorTargetWithoutSourcePartition(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "*",
		ObjCode:         "obj_sistema",
		TargetPartition: "partition10",
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if using target partition without source partition")
	}
}

func TestExecuteToReturnErrorSingleViewNotInTarget(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "INVALID_VIEW_CODE",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		TargetPartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionInsert,
				TargetPosition: "1",
				SourcePosition: "1",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if view exists in target")
	}
}

func TestExecuteToReturnErrorSectionSourcePositionNotDefined(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionInsert,
				TargetPosition: "1",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if section source position were defined")
	}
}

func TestExecuteToReturnErrorSectionSourcePositionIndexOutOfBounds(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionInsert,
				SourcePosition: "11",
				TargetPosition: "1",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if source has section in the defined position")
	}
}

func TestExecuteToReturnErrorSectionTargetPositionIndexOutOfBounds(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionInsert,
				SourcePosition: "1",
				TargetPosition: "11",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if target has section in the defined position")
	}
}

func TestExecuteToReturnErrorSectionReplaceWithoutTargetPosition(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionReplace,
				SourcePosition: "1",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if target position were defined to replace section")
	}
}

func TestExecuteToReturnErrorSectionRemoveWithoutTargetPosition(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionRemove,
				SourcePosition: "1",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if target position were defined to remove section")
	}
}

func TestExecuteToReturnErrorSectionUpdateWithoutField(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionUpdate,
				SourcePosition: "1",
				TargetPosition: "1",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if trying to update section without tag field")
	}
}

func TestExecuteToReturnErrorUpdateSectionRemoveInvalidFieldCode(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		TargetPartition: "partition20",
		Sections: []model.Section{
			{
				Action:         constant.ActionUpdate,
				SourcePosition: "1",
				TargetPosition: "1",
				Fields: []model.SectionField{
					{
						Code:   "INVALID_FIELD_CODE",
						Remove: true,
					},
				},
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition20_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if trying to remove field from section with invalid code")
	}
}

func TestExecuteToReturnErrorUpdateSectionInsertInvalidFieldCode(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		TargetPartition: "partition20",
		Sections: []model.Section{
			{
				Action:         constant.ActionUpdate,
				SourcePosition: "1",
				TargetPosition: "1",
				Fields: []model.SectionField{
					{
						Code:   "INVALID_FIELD_CODE",
						Column: constant.ColumnLeft,
					},
				},
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition20_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if trying to insert field from source section with invalid code")
	}
}

func TestExecuteToReturnErrorUpdateSectionInsertInvalidTargetInsertBeforeCode(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		TargetPartition: "partition20",
		Sections: []model.Section{
			{
				Action:         constant.ActionUpdate,
				SourcePosition: "1",
				TargetPosition: "1",
				Fields: []model.SectionField{
					{
						Code:         "analista",
						Column:       constant.ColumnLeft,
						InsertBefore: "INVALID_FIELD_CODE",
					},
				},
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition20_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if trying to insert before invalid code in target")
	}
}

func TestExecuteToReturnErrorUpdateSectionInvalidSectionAction(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		TargetPartition: "partition20",
		Sections: []model.Section{
			{
				Action:         "INVALID_SECTION_ACTION",
				SourcePosition: "1",
				TargetPosition: "1",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition20_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if trying use an invalid section action value")
	}
}

func TestExecuteToReturnErrorUpdateSectionInvalidFieldColumn(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistema.auditoria",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition10",
		TargetPartition: "partition20",
		Sections: []model.Section{
			{
				Action:         constant.ActionUpdate,
				SourcePosition: "1",
				TargetPosition: "1",
				Fields: []model.SectionField{
					{
						Code:   "analista",
						Column: "INVALID_COLUMN",
					},
				},
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition20_full_xog.xml")
	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if trying to use a invalid column value")
	}
}

func TestExecuteToReturnErrorSingleViewNotInSource(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "obj_sistemaList",
		ObjCode:         "obj_sistema",
		SourcePartition: "partition20",
		TargetPartition: "partition10",
		Sections: []model.Section{
			{
				Action:         constant.ActionInsert,
				TargetPosition: "1",
				SourcePosition: "1",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_partition20_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_partition_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err == nil {
		t.Fatalf("Error transforming view XOG file. Debug: not validating if view exists in source")
	}
}

func TestExecuteToInsertGroupAction(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:   constant.ElementTypeActionGroup,
				Code:   "actions_group_test",
				Action: constant.ActionInsert,
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_target_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_actions_insert_group_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToInsertBeforeGroupAction(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:         constant.ElementTypeActionGroup,
				Code:         "actions_group_test",
				Action:       constant.ActionInsert,
				InsertBefore: "general",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_target_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_actions_insert_before_group_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToInsertGroupActionWithExistingTarget(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:   constant.ElementTypeActionGroup,
				Code:   "actions_group_test",
				Action: constant.ActionInsert,
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_actions_insert_with_existing_group_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToRemoveGroupAction(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:   constant.ElementTypeActionGroup,
				Code:   "actions_group_test",
				Action: constant.ActionRemove,
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_actions_remove_group_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToInsertAction(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:   constant.ElementTypeAction,
				Code:   "rally_full_sync",
				Action: constant.ActionInsert,
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_target_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_actions_insert_action_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToRemoveAction(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:   constant.ElementTypeAction,
				Code:   "rally_full_sync",
				Action: constant.ActionRemove,
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_actions_remove_action_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToInsertBeforeAction(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:         constant.ElementTypeAction,
				Code:         "rally_full_sync",
				Action:       constant.ActionInsert,
				InsertBefore: "odf_XMLExportcas_environment",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_target_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_actions_insert_before_action_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToInsertActionWithExistingTarget(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:   constant.ElementTypeAction,
				Code:   "rally_full_sync",
				Action: constant.ActionInsert,
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err != nil {
		t.Fatalf("Error transforming view XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "view_actions_insert_with_existing_group_result.xml") == false {
		t.Errorf("Error transforming view XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnErrorInvalidSourceGroupCodeInsert(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:         constant.ElementTypeActionGroup,
				Code:         "invalid_code",
				Action:       constant.ActionInsert,
				InsertBefore: "odf_XMLExportcas_environment",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_target_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err == nil {
		t.Errorf("Error transforming view XOG file. Not validating when inserting invalid group source code.")
	}
}

func TestExecuteToReturnErrorInvalidInsertBeforeGroupCode(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:         constant.ElementTypeActionGroup,
				Code:         "actions_group_test",
				Action:       constant.ActionInsert,
				InsertBefore: "invalid_code",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_target_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err == nil {
		t.Errorf("Error transforming view XOG file. Not validating invalid insertBefore code.")
	}
}

func TestExecuteToReturnErrorInvalidTargetGroupCodeRemove(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:         constant.ElementTypeActionGroup,
				Code:         "invalid_code",
				Action:       constant.ActionRemove,
				InsertBefore: "odf_XMLExportcas_environment",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_target_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err == nil {
		t.Errorf("Error transforming view XOG file. Not validating when removing invalid group target code.")
	}
}

func TestExecuteToReturnErrorInvalidSourceActionCodeInsert(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:         constant.ElementTypeAction,
				Code:         "invalid_code",
				Action:       constant.ActionInsert,
				InsertBefore: "odf_XMLExportcas_environment",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_target_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err == nil {
		t.Errorf("Error transforming view XOG file. Not validating when inserting invalid group source code.")
	}
}

func TestExecuteToReturnErrorInvalidInsertBeforeActionCode(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:         constant.ElementTypeAction,
				Code:         "rally_full_sync",
				Action:       constant.ActionInsert,
				InsertBefore: "invalid_code",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_target_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err == nil {
		t.Errorf("Error transforming view XOG file. Not validating invalid insertBefore code.")
	}
}

func TestExecuteToReturnErrorInvalidTargetActionCodeRemove(t *testing.T) {
	file := model.DriverFile{
		Type:            constant.TypeView,
		Code:            "cas_environmentProperties",
		ObjCode:         "cas_environment",
		SourcePartition: "NIKU.ROOT",
		Elements: []model.Element{
			{
				Type:         constant.ElementTypeAction,
				Code:         "invalid_code",
				Action:       constant.ActionRemove,
				InsertBefore: "odf_XMLExportcas_environment",
			},
		},
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "view_actions_source_full_xog.xml")
	aux := etree.NewDocument()
	aux.ReadFromFile(packageMockFolder + "view_actions_target_full_xog.xml")

	err := Execute(xog, aux, &file)

	if err == nil {
		t.Errorf("Error transforming view XOG file. Not validating when removing invalid group Target code.")
	}
}
