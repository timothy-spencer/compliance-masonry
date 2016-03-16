package renderers

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/opencontrol/compliance-masonry-go/models"
)

type exportControlTest struct {
	opencontrolDir    string
	certificationPath string
	standardKey       string
	controlKey        string
	expectedPath      string
	expectedText      string
}

var exportControlTests = []exportControlTest{
	{
		"../fixtures/opencontrol_fixtures/",
		"../fixtures/opencontrol_fixtures/certifications/LATO.yaml",
		"NIST-800-53",
		"AC-2",
		"NIST-800-53-AC-2.md",
		"#NIST-800-53-AC-2  \n##Account Management  \n",
	},
}

func TestExportControl(t *testing.T) {
	for _, example := range exportControlTests {
		dir, err := ioutil.TempDir("", "example")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(dir)
		openControl := OpenControlGitBook{
			models.LoadData(example.opencontrolDir, example.certificationPath),
			dir,
		}
		control := openControl.Standards.Get(example.standardKey).Controls[example.controlKey]
		actualPath, actualText := openControl.exportControl(&ControlGitbook{&control, dir, example.standardKey, example.controlKey})
		expectedPath := filepath.Join(dir, example.expectedPath)
		if expectedPath != actualPath {
			t.Errorf("Expected %s, Actual: %s", example.expectedPath, actualPath)
		}
		if example.expectedText != actualText {
			t.Errorf("Expected %s, Actual: %s", example.expectedText, actualText)
		}
	}
}
