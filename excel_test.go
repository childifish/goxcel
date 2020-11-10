package goxcel

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestExcelHelper(t *testing.T) {
	type Student struct {
		StuName string `helper:"姓名"`
		StuNum  int `helper:"学号"`
	}
	var students []Student
	for i := 0; i < 50; i++ {
		students = append(students,Student{
			StuName: strconv.Itoa(rand.Intn(400)),
			StuNum:  rand.Intn(50000),
		})
	}
	_ = ExcelStructs(students, "学生", "")

}

func TestExcelStructsNotStore(t *testing.T) {
	type Student struct {
		StuName string `helper:"姓名"`
		StuNum  int `helper:"学号"`
	}
	var students []Student
	for i := 0; i < 50; i++ {
		students = append(students,Student{
			StuName: strconv.Itoa(rand.Intn(400)),
			StuNum:  rand.Intn(50000),
		})
	}
	_ = ExcelStructsNotStore(students).StoreFile("")

}

func TestExcelStructsLite(t *testing.T) {
	type Student struct {
		StuName string `helper:"姓名"`
		StuNum  int `helper:"学号"`
	}
	var students []Student
	for i := 0; i < 50; i++ {
		students = append(students,Student{
			StuName: strconv.Itoa(rand.Intn(400)),
			StuNum:  rand.Intn(50000),
		})
	}
	ExcelStructsLite(students).DeleteTimer(time.Second*1)
	//ExcelStructsLite(students,"学生")
	//ExcelStructsLite(students,"学生","./storage")
	time.Sleep(time.Second*2)

}