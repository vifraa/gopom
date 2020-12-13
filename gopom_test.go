package gopom

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
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

func TestParsing_Project(t *testing.T) {
	if p.GroupID != "com.test" {
		t.Error("GroupID: expected 'com.test', got: " + p.GroupID)
	}
	if p.ArtifactID != "test-application" {
		t.Error("ArtifactID: expected 'test-application', got: " + p.ArtifactID)
	}
	if p.Version != "1.0.0" {
		t.Error("Version: expected '1.0.0', got: " + p.Version)
	}
	if p.Packaging != "war" {
		t.Error("Packaging: expected 'war', got: " + p.Packaging)
	}
	if p.Name != "gopom-test" {
		t.Error("Name: expected 'gopom-test', got: " + p.Name)
	}
	if p.Description != "pom parser in golang" {
		t.Error("Description: expected 'pom parser in golang', got: " + p.Description)
	}
	if p.URL != "testUrl" {
		t.Error("URL: expected 'testUrl', got: " + p.URL)
	}
	if p.InceptionYear != "2020" {
		t.Error("InceptionYear: expected '2020', got: " + p.InceptionYear)
	}
	if p.ModelVersion != "4.0.0" {
		t.Error("ModelVersion: expected '4.0.0', got: " + p.ModelVersion)
	}

	if len(p.Modules) != 2 {
		t.Error("Modules: expected len==2, got: " + strconv.Itoa(len(p.Modules)))
	} else {
		if p.Modules[0] != "module1" {
			t.Error("Modules[0]: expected 'module1', got: " + p.Modules[0])
		}
		if p.Modules[1] != "module2" {
			t.Error("Modules[1]: expected 'module1', got: " + p.Modules[1])
		}
	}
}

func TestParsing_Organization(t *testing.T) {
	if p.Organization.Name != "name" {
		t.Error("Organization.Name: expected 'name', got: " + p.Organization.Name)
	}
	if p.Organization.URL != "url" {
		t.Error("Organization.URL: expected 'url', got: " + p.Organization.URL)
	}
}

func TestParsing_Licenses(t *testing.T) {
	if len(p.Licenses) != 1 {
		t.Error("Licenses: expected len==1, got: " + strconv.Itoa(len(p.Licenses)))
	}
	if p.Licenses[0].Name != "name" {
		t.Error("Licenses.Name: expected 'name', got: " + p.Licenses[0].Name)
	}
	if p.Licenses[0].URL != "url" {
		t.Error("Licenses.URL: expected 'url', got: " + p.Licenses[0].URL)
	}
	if p.Licenses[0].Distribution != "manual" {
		t.Error("Licenses.Distribution: expected 'manual', got: " + p.Licenses[0].Distribution)
	}
	if p.Licenses[0].Comments != "comments" {
		t.Error("Licenses.Comments: expected 'comments', got: " + p.Licenses[0].Comments)
	}
}

func TestParsing_Developers(t *testing.T) {
	if len(p.Developers) != 1 {
		t.Error("Developers: expected len==1, got: " + strconv.Itoa(len(p.Developers)))
	}

	d := p.Developers[0]
	if d.ID != "id" {
		t.Error("Developers[0].ID: expected 'id', got: " + d.ID)
	}
	if d.Name != "name" {
		t.Error("Developers[0].Name: expected 'name', got: " + d.Name)
	}
	if d.Email != "email" {
		t.Error("Developers[0].Email: expected 'email', got: " + d.Email)
	}
	if d.URL != "url" {
		t.Error("Developers[0].URL: expected 'url', got: " + d.URL)
	}
	if d.Organization != "organization" {
		t.Error("Developers[0].Organization: expected 'organization', got: " + d.Organization)
	}
	if d.OrganizationURL != "organizationUrl" {
		t.Error("Developers[0].OrganizationUrl: expected 'organizationUrl', got: " + d.OrganizationURL)
	}
	if d.Timezone != "+1" {
		t.Error("Developers[0].Timezone: expected '+1', got: " + d.Timezone)
	}
	if len(d.Roles) != 2 {
		t.Error("Developers: expected len==2, got: " + strconv.Itoa(len(d.Roles)))
	}
	if d.Roles[0] != "role1" {
		t.Error("Developers[0].Roles[0]: expected 'role1', got: " + d.Roles[0])
	}
	if d.Roles[1] != "role2" {
		t.Error("Developers[0].Roles[1]: expected 'role2', got: " + d.Roles[1])
	}
}

func TestParse_Contributors(t *testing.T) {
	if len(p.Contributors) != 1 {
		t.Error("Contributors: expected len==1, got: " + strconv.Itoa(len(p.Contributors)))
	}

	c := p.Contributors[0]
	if c.Name != "name" {
		t.Error("Contributors[0].Name: expected 'name', got: " + c.Name)
	}
	if c.Email != "email" {
		t.Error("Contributors[0].Email: expected 'email', got: " + c.Email)
	}
	if c.URL != "url" {
		t.Error("Contributors[0].URL: expected 'url', got: " + c.URL)
	}
	if c.Organization != "organization" {
		t.Error("Contributors[0].Organization: expected 'organization', got: " + c.Organization)
	}
	if c.OrganizationURL != "organizationUrl" {
		t.Error("Contributors[0].OrganizationUrl: expected 'organizationUrl', got: " + c.OrganizationURL)
	}
	if c.Timezone != "+1" {
		t.Error("Contributors[0].Timezone: expected '+1', got: " + c.Timezone)
	}
	if len(c.Roles) != 2 {
		t.Error("Contributors: expected len==2, got: " + strconv.Itoa(len(c.Roles)))
	}
	if c.Roles[0] != "role1" {
		t.Error("Contributors[0].Roles[0]: expected 'role1', got: " + c.Roles[0])
	}
	if c.Roles[1] != "role2" {
		t.Error("Developers[0].Roles[1]: expected 'role2', got: " + c.Roles[1])
	}
}

func TestParse_MailingLists(t *testing.T) {
	if len(p.MailingLists) != 1 {
		t.Error("MailingLists: expected len==1, got: " + strconv.Itoa(len(p.MailingLists)))
	}

	m := p.MailingLists[0]
	if m.Name != "name" {
		t.Error("MailingLists[0].Name: expected 'name', got: " + m.Name)
	}
	if m.Subscribe != "subscribe" {
		t.Error("MailingLists[0].Subscribe: expected 'subscribe', got: " + m.Subscribe)
	}
	if m.Unsubscribe != "unsubscribe" {
		t.Error("MailingLists[0].Unsubscribe: expected 'unsubscribe', got: " + m.Unsubscribe)
	}
	if m.Post != "post" {
		t.Error("MailingLists[0].Post: expected 'post', got: " + m.Post)
	}
	if m.Archive != "archive" {
		t.Error("MailingLists[0].Archive: expected 'archive', got: " + m.Archive)
	}

	if len(m.OtherArchives) != 2 {
		t.Error("MailingLists.OtherArchives: expected len==2, got: " + strconv.Itoa(len(m.OtherArchives)))
	}
	if m.OtherArchives[0] != "archive1" {
		t.Error("MailingLists[0].OtherArchives[0]: expected 'archive1', got: " + m.OtherArchives[0])
	}
	if m.OtherArchives[1] != "archive2" {
		t.Error("MailingLists[0].OtherArchives[1]: expected 'archive2', got: " + m.OtherArchives[1])
	}
}

func TestParsing_Prerequisites(t *testing.T) {
	if p.Prerequisites.Maven != "2.0" {
		t.Error("Prerequisites.Maven: expected '2.0', got: " + p.Prerequisites.Maven)
	}
}

func TestParsing_SCM(t *testing.T) {
	if p.SCM.URL != "url" {
		t.Error("SCM.URL: expected 'url', got: " + p.SCM.URL)
	}
	if p.SCM.Connection != "connection" {
		t.Error("SCM.Connection: expected 'connection', got: " + p.SCM.Connection)
	}
	if p.SCM.DeveloperConnection != "developerConnection" {
		t.Error("SCM.DeveloperConnection: expected 'developerConnection', got: " + p.SCM.DeveloperConnection)
	}
	if p.SCM.Tag != "tag" {
		t.Error("SCM.Tag: expected 'tag', got: " + p.SCM.Tag)
	}
}

func TestParsing_IssueManagement(t *testing.T) {
	if p.IssueManagement.URL != "url" {
		t.Error("IssueManagement.URL: expected 'url', got: " + p.IssueManagement.URL)
	}
	if p.IssueManagement.System != "system" {
		t.Error("IssueManagement.System: expected 'system', got: " + p.IssueManagement.System)
	}
}
