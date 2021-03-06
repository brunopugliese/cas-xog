package validate

import (
	"github.com/andreluzz/cas-xog/constant"
	"github.com/beevik/etree"
	"strings"
	"testing"
)

var packageMockFolder string

func init() {
	packageMockFolder = "../" + constant.FolderMock + "validate/"
}

func TestCheckToValidateSuccessOutput(t *testing.T) {
	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "mockSuccess.xml")

	output, err := Check(xog)

	if output.Code != constant.OutputSuccess || strings.Contains(output.Debug, "Elapsed time:") == false {
		t.Errorf("Expected output %s and recëived output %s", constant.OutputSuccess, output.Code)
	}
	if err != nil {
		t.Errorf("Encountered an error while checking the output. %s", err.Error())
	}
}

func TestCheckToValidateSuccessOutputWithoutElapsedTime(t *testing.T) {
	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "mockSuccessWithoutElapsedTime.xml")

	output, err := Check(xog)

	if output.Code != constant.OutputSuccess || strings.Contains(output.Debug, "Elapsed time:") {
		t.Errorf("Expected output %s and recëived output %s", constant.OutputSuccess, output.Code)
	}
	if err != nil {
		t.Errorf("Encountered an error while checking the output. %s", err.Error())
	}
}

func TestCheckToValidateErrorOutput(t *testing.T) {
	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "mockError.xml")

	output, err := Check(xog)

	if output.Code != constant.OutputError {
		t.Errorf("Expected output %s and recëived output %s", constant.OutputError, output.Code)
	}
	if err == nil {
		t.Error("Expected an error, but didn't receive one.")
	}
}

func TestCheckToValidateWarningOutput(t *testing.T) {
	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "mockWarning.xml")

	output, err := Check(xog)

	if output.Code != constant.OutputWarning {
		t.Errorf("Expected output %s and recëived output %s", constant.OutputWarning, output.Code)
	}
	if err != nil {
		t.Errorf("Encountered an error while checking the output. %s", err.Error())
	}
}

func TestCheckToValidateZeroTotalResultsOutput(t *testing.T) {
	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "mockZeroTotalResults.xml")

	output, err := Check(xog)

	if output.Code != constant.OutputError {
		t.Errorf("Expected output %s and recëived output %s", constant.OutputError, output.Code)
	}
	if err == nil || err.Error() != "output statistics totalNumberOfRecords = 0" {
		t.Errorf("Expected the error 'output statistics totalNumberOfRecords = 0' but instead received '%s'.", err.Error())
	}
}

func TestCheckToValidateOneFailureResultOutput(t *testing.T) {
	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "mockOneFailureResult.xml")

	output, err := Check(xog)

	if output.Code != constant.OutputError {
		t.Errorf("Expected output %s and recëived output %s", constant.OutputError, output.Code)
	}
	if err == nil || (strings.Contains(err.Error(), "output statistics failure on") == false) {
		t.Errorf("Expected an error containing the number of failure records but instead received '%s'.", err.Error())
	}
}

func TestCheckToValidateNoStatusTag(t *testing.T) {
	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "mockNoStatusTag.xml")

	output, err := Check(xog)

	if output.Code != constant.OutputError {
		t.Errorf("Expected output %s and recëived output %s", constant.OutputError, output.Code)
	}
	if err == nil || err.Error() != "no status tag defined" {
		t.Errorf("Expected the error 'no status tag defined' but instead received '%s'.", err.Error())
	}
}

func TestCheckToValidateNoXOGOutputTag(t *testing.T) {
	xog := etree.NewDocument()
	xog.ReadFromFile(packageMockFolder + "mockNoXOGOutputTag.xml")

	output, err := Check(xog)

	if output.Code != constant.OutputError {
		t.Errorf("Expected output %s and recëived output %s", constant.OutputError, output.Code)
	}
	if err == nil || err.Error() != "no output tag defined" {
		t.Errorf("Expected the error 'no output tag defined' but instead received '%s'.", err.Error())
	}
}

func TestCheckToValidateInvalidXml(t *testing.T) {
	output, err := Check(nil)

	if output.Code != constant.OutputError {
		t.Errorf("Expected output %s and recëived output %s", constant.OutputError, output.Code)
	}
	if err == nil || err.Error() != "invalid xog" {
		t.Errorf("Expected the error 'invalid xog' but instead received '%s'.", err.Error())
	}
}
