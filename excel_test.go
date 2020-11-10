package goxcel

import (
	"math/rand"
	"strconv"
	"testing"
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
	_ = ExcelStructs(students, "学生", "", Student{})

}

