package goxcel

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExcelHelper_Analyze(t *testing.T) {
	type User struct {
		Name string `helper:"姓名"`
		PhoneNum string `helper:"电话"`
	}
	var a []User
	b := User{
		Name:     "12",
		PhoneNum: "1241241",
	}
	a = append(a,b)
	a = append(a,b)
	//for i, user := range a {
	//
	//}

	table := InitTable("testing", User{})
	for i, i2 := range a {
		table.Insert(i,i2)
	}
	err := table.StoreFile("")


	if err != nil{
		fmt.Println(err)
	}

}

func TestIndex2Chara(t *testing.T) {
	fmt.Println(Index2Chara(3))
}

func TestSlice(t *testing.T) {
	type User struct {
		Name string `helper:"姓名"`
		PhoneNum string `helper:"电话"`
	}
	var a []User
	b := User{
		Name:     "12",
		PhoneNum: "1241241",
	}
	a = append(a,b)
	a = append(a,b)

	test(a)
}

func test(t interface{}) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		for i := 0; i < s.Len(); i++ {
			elem := s.Index(i).Elem().Interface()
			num := reflect.TypeOf(elem).NumField()
			for j := 0; j < num; j++ {
				va :=reflect.ValueOf(elem).Field(j)
				ty := reflect.TypeOf(elem).Field(j)
				fmt.Println("我的内容是",va,"我的类型是",ty.Type,"我属于结构体",reflect.TypeOf(elem).Name())
			}
		}
	case reflect.Struct:
		f := reflect.TypeOf(t).NumField()
		for j := 0; j < f; j++ {
			fmt.Println(reflect.ValueOf(t).Field(j))
		}
	}
}

func TestFinal(t *testing.T)  {
	type Dog struct {
		Name string
	}
	type User struct {
		Name string `helper:"姓名"`
		PhoneNum string
		id []int `helper:"序号"`
		d []Dog
	}
	var a []User
	for i := 0; i < 50; i++ {
		b := User{
			Name:     "12",
			PhoneNum: "1241241",
			id: []int{1,2,3},
			d: []Dog{{"yyx"},{"oyhn"}},
		}
		a = append(a,b)
	}

	err := ExcelStructs(a, "test", "",User{})

	if err != nil{
		fmt.Println(err)
	}

}

func TestSth(t *testing.T)  {
	type User struct {
		Name string
		PhoneNum string
		id int
	}
	type Dog struct {
		Name string
	}
	type Cake struct {
		Prise int
	}

	var a []interface{}
	a= append(a,User{
		Name:     "我是打工人",
		PhoneNum: "114514",
		id:       123,
	})
	a = append(a,Dog{Name: "我是狗"})
	a =append(a,Cake{999})

	test(a)

}