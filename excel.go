package goxcel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"reflect"
	"strconv"
	"time"
)

/**
 * user: LCYH
 * Date: 2020/10/13
 * Update: 2020/11/10
 */

//todo
/*
*1.优化掉model
*2.支持嵌套结构体
*3.支持数组等更多类型(已经支持了。。)
*/

type ExcelHelper struct {
	Object interface{}
	TableName string
	TableHeader []string
	File *excelize.File
}

//ExcelStructs 参数分别为希望储存的结构体切片，.xslx表名，保存路径，（将来会被优化掉的源结构体），返回一个error
func ExcelStructs(v interface{},tableName string,filePath string)(err error) {
	table := InitTable(tableName,v)
	table.MultiInsert(v)
	return table.StoreFile(filePath)
}

func ExcelStructsNotStore(v interface{})*ExcelHelper  {
	table := InitTable("",v)
	table.MultiInsert(v)
	return table
}

func InitTable(tableName string, v interface{})*ExcelHelper  {
	var e ExcelHelper
	if tableName == ""{
		name := reflect.ValueOf(v).Index(0).Type().Name()
		e.TableName = name + strconv.Itoa(int(time.Now().Unix()))
	}else {
		e.TableName = tableName
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		typ := reflect.ValueOf(v).Index(0).Type()
		Num := reflect.ValueOf(v).Index(0).NumField()
		for i := 0; i < Num; i++ {
			var tag string
			tag = typ.Field(i).Tag.Get("helper")
			if tag == ""{
				tag  = typ.Field(i).Name
			}
			e.TableHeader = append(e.TableHeader,tag)
		}
	}
	e.InsertHeader()
	return &e
}

func (e *ExcelHelper) AnalyzeTableHeader() *ExcelHelper {
	obj := reflect.ValueOf(e.Object)
	if obj.Kind() == reflect.Ptr {
		obj = obj.Elem()
	}
	typ := reflect.TypeOf(e.Object)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	Num := obj.NumField()

	for i := 0; i < Num; i++ {
		var tag string
		tag = typ.Field(i).Tag.Get("helper")
		if tag == ""{
			tag  = typ.Field(i).Name
		}
		e.TableHeader = append(e.TableHeader,tag)
	}

	return e
}

func (e *ExcelHelper) AnalyzeTableValue(v interface{})(field []string)   {
	obj := reflect.ValueOf(v)
	if obj.Kind() == reflect.Ptr {
		obj = obj.Elem()
	}
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	Num := obj.NumField()

	for i := 0; i < Num; i++ {
		v := obj.Field(i)
		field = append(field,v.String())
	}
	return field
}


func (e *ExcelHelper) InsertHeader()*ExcelHelper  {
	f := excelize.NewFile()
	for i,v := range e.TableHeader{
		f.SetCellValue("Sheet1", id2index(i,1),v)
	}
	e.File = f
	return e
}

func (e *ExcelHelper) Insert(index int,v interface{})*ExcelHelper  {
	value := e.AnalyzeTableValue(v)
	for i,j := range value{
		e.File.SetCellValue("Sheet1", id2index(i,index+2),j)
	}
	return e
}

func (e *ExcelHelper) MultiInsert(v interface{})*ExcelHelper {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		for i := 0; i < s.Len(); i++ {
			elem := s.Index(i)
			num := elem.NumField()
			for j := 0; j < num; j++ {
				e.File.SetCellValue("Sheet1", id2index(j,i+2),elem.Field(j))
			}
		}
	}
	return e
}

func (e *ExcelHelper) StoreFile(filepath string) error {
	finalFile := filepath + e.TableName + ".xlsx"
	if err := e.File.SaveAs(finalFile); err != nil {
		println(err.Error())
		return  err
	}
	return nil
}

func index2Chara(i int)string  {
	if i >= 24{
		return "nil"
	}
	return string(rune(65+i))
}

func id2index(charaID int,i int)string  {
	s := strconv.Itoa(i)
	return index2Chara(charaID)+s
}