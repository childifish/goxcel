# goxcel📑

幼儿园小朋友也能轻松使用的golang structs -->.xslx小工具

- 使用方便
- 学习成本低（大概）
- 源码简单易懂

##### 使用例（懒人版）

```go
goxcel.ExcelStructs(想储存的结构体数组,表名,储存路径,结构体)
```

##### 使用例（比较详细版）

```go
go get github.com/childifish/goxcel
```

我们需要一个的结构体

```go
type Student struct {
	StuName string
	StuNum  int
}
```

在生产环境中，一般都会用到结构体切片

```go
var students []Student
```

这里我们写一个随机生成结构体用于初始化（下方的students是50个Student结构体，过程不做赘述）

ExcelStructs 参数分别为希望储存的结构体切片，.xslx表名，保存路径，（将来会被优化掉的源结构体），返回一个error。

```go
goxcel.ExcelStructs(students, "学生", "", Student{})
```

在根目录下创建了"学生.xslx"

| StuName | StuNum |
| ------- | ------ |
| 81      | 27887  |
| 247     | 34059  |
| 81      | 41318  |
| 25      | 22540  |
| 56      | 3300   |

如果对表头有特殊要求，不希望使用结构体字段名的话，可以使用tag：

```go
type Student struct {
	StuName string `helper:"姓名"`
    StuNum  int `helper:"学号"`
}
```

| 姓名 | 学号  |
| ---- | ----- |
| 81   | 27887 |
| 247  | 34059 |
| 81   | 41318 |
| 25   | 22540 |
| 56   | 3300  |

字段类型目前支持数组，切片，结构体，具有嵌套结构的结构体（其他的俺还没试）；