package gopom

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"
)

var p Project

func init() {
	err := xml.Unmarshal([]byte(examplePom), &p)
	if err != nil {
		fmt.Println("unable to parse xml ", err)
		os.Exit(1)
		return
	}
}

func TestParsing_Parent(t *testing.T) {
	if p.Parent.ArtifactID != "test-application" {
		t.Error("Parent.ArtifactID: expected 'test-application', got: " + p.Parent.ArtifactID)
	}
	if p.Parent.GroupID != "com.test" {
		t.Error("Parent.GroupID: expected 'com.test', got: " + p.Parent.GroupID)
	}
	if p.Parent.Version != "1.0.0" {
		t.Error("Parent.Version: expected '1.0.0', got: " + p.Parent.Version)
	}
	if p.Parent.RelativePath != "../pom.xml" {
		t.Error("Parent.RelativePath: expected '../pom.xml', got: " + p.Parent.RelativePath)
	}
}
