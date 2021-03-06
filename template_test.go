package main

import (
	"bytes"
	"testing"
	"text/template"
)

// see https://www.jianshu.com/p/05671bab2357

func tplWithString(tpl string, content string) string {
	tmpl, err := template.New("test").Parse(tpl)
	if err != nil {
		panic(err)
	}
	var result bytes.Buffer
	err = tmpl.Execute(&result, content)
	if err != nil {
		panic(err)
	}
	return result.String()
}

func TestDotTpl(t *testing.T) {
	result := tplWithString("value={{.}}", "result")
	if result != "value=result" {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	dataMap := make(map[string]string)
	dataMap["Name"] = "theName"
	dataMap["Age"] = "theAge"

	tpl := "value={{.Name}}-{{.Age}}"

	tmpl, err := template.New("test").Parse(tpl)
	if err != nil {
		panic(err)
	}
	var result bytes.Buffer
	err = tmpl.Execute(&result, dataMap)
	if err != nil {
		panic(err)
	}
	t.Log(result.String())
	if result.String() != "value=theName-theAge" {
		t.Fail()
	}
}
func TestTempalteDemo(t *testing.T) {

	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, sweaters)
	if err != nil {
		panic(err)
	}
	if tpl.String() != "17 items are made of wool" {
		t.Fail()
	}
}
