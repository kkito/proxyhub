package main

import (
	"fmt"
	"testing"
)

func TestIsCNHost(t *testing.T) {
	isCNHost("www.baidu.com")
}

func TestGetIPFromHost(t *testing.T) {
	result := getIPFromHost("www.baidu.com")
	t.Logf(result.String())
	t.Logf(getCountryCodeByHostOrIP("www.baidu.com"))
	result = getIPFromHost("114.114.114.114")
	t.Logf(result.String())
	t.Logf(getCountryCodeByHostOrIP("114.114.114.114"))
	t.Logf(getCountryCodeByHostOrIP("kkito.cn"))
	if isCNHost("kkito.cn") {
		t.FailNow()
	}
}

/// ========= test embed struct

type embStr struct {
	name string
	age  int
}

func (em *embStr) printName() string {
	return em.name + " " + em.getAge()
}

func (em *embStr) getAge() string {
	return "empty age"
}

type student struct {
	embStr
	no string
}

func (st *student) getAge() string {
	return "age from student"
}

func TestStduent(t *testing.T) {
	stu := student{}
	stu.name = "test name"
	// fmt.Println(stu.name)
	ret := stu.printName()
	fmt.Println(ret)
	fmt.Println(stu.getAge())
}

// ===== override

type I interface {
	Foo() string
}

type A struct {
	i I
}

func (a *A) Foo() string {
	return "A.Foo()"
}

func (a *A) Bar() string {
	return a.i.Foo()
}

type B struct {
	A
}

func (b *B) Foo() string {
	return "B.Foo()"
}

func TestOverride(t *testing.T) {
	// 借助一个接口把父子关系串联起来
	inst := B{}
	inst.i = &inst
	t.Log(inst.i)
	t.Log(inst.Bar())
}
