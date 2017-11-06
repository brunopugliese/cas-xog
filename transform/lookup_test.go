package transform

import (
	"testing"
	"github.com/beevik/etree"
	"github.com/andreluzz/cas-xog/common"
)

func TestExecuteToReturnStaticLookupTransformed(t *testing.T) {
	file := common.DriverFile{
		Code: "LOOKUP_CAS_XOG",
		Type: common.LOOKUP,
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "lookup_static_full_xog.xml")
	err := Execute(xog, nil, file)

	if err != nil {
		t.Fatalf("Error transforming static lookup XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "lookup_static_result.xml") == false {
		t.Errorf("Error transforming static lookup XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnStaticLookupTargetPartition(t *testing.T) {
	file := common.DriverFile{
		Code: "LOOKUP_CAS_XOG",
		Type: common.LOOKUP,
		TargetPartition: "NIKU.ROOT",
		Path: "testTarget.xml",
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "lookup_static_full_xog.xml")
	err := Execute(xog, nil, file)

	if err != nil {
		t.Fatalf("Error transforming static lookup XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "lookup_static_partitions_result.xml") == false {
		t.Errorf("Error transforming static lookup XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnStaticLookupSourceAndTargetPartition(t *testing.T) {
	file := common.DriverFile{
		Code: "LOOKUP_CAS_XOG",
		Type: common.LOOKUP,
		SourcePartition: "NIKU.ROOT",
		TargetPartition: "partition10",
		Path: "testSourceAndTarget.xml",
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "lookup_static_full_xog.xml")
	err := Execute(xog, nil, file)

	if err != nil {
		t.Fatalf("Error transforming static lookup XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "lookup_static_source_target_partition_result.xml") == false {
		t.Errorf("Error transforming static lookup XOG file. Invalid result XML.")
	}
}

func TestExecuteToReturnDynamicLookupPartitionsTransformed(t *testing.T) {
	file := common.DriverFile{
		Code: "LOOKUP_CAS_XOG",
		Type: common.LOOKUP,
	}

	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "lookup_dynamic_full_xog.xml")
	err := Execute(xog, nil, file)

	if err != nil {
		t.Fatalf("Error transforming static lookup XOG file. Debug: %s", err.Error())
	}

	if readMockResultAndCompare(xog, "lookup_dynamic_result.xml") == false {
		t.Errorf("Error transforming static lookup XOG file. Invalid result XML.")
	}
}

