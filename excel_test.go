package goxcel

import (
	"fmt"
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
	//for i := 0; i < 50; i++ {
	//	students = append(students,Student{
	//		StuName: strconv.Itoa(rand.Intn(400)),
	//		StuNum:  rand.Intn(50000),
	//	})
	//}
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

func TestOrigin(t *testing.T)  {
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
	table := InitTable("学生",students)
	for i, i2 := range students {
		table.Insert(i,i2)
	}
	err := table.StoreFile("")
	if err!=nil{
		fmt.Println("error in store file")
	}
}


func TestExcelStructsNotStore2(t *testing.T)  {
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
	table := ExcelStructsNotStore(students)
	
	table.MultiInsert(students)

	err := table.StoreFile("").Error()

	if err!=nil{
		fmt.Println("error in store file")
	}
}