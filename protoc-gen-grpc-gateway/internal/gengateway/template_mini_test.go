package gengateway

import (
	"bytes"
	"strings"
	"testing"
	"text/template"
)

type TestBinding struct {
	FuncToCall func(string) string
	FieldName string
}

func TestTemplate(t *testing.T) {
	localTestTemplate, err := template.New("test-local").Parse(
		"{{call .FuncToCall .FieldName}}",
	)
	if err != nil {
		t.Fatal(err)
	}
	w := bytes.NewBuffer(nil)
	if err := localTestTemplate.Execute(w, TestBinding{
		FuncToCall: getTypeFromName,
		FieldName: "epoch",
	}); err != nil {
		t.Fatal(err)
	}
	t.Log(w.String())
}

func getTypeFromName(name string) string {
	if strings.Contains(name, "epoch") {
		return "types.Epoch"
	}
	return ""
}
