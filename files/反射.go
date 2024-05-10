package files

import (
	"fmt"
	"reflect"
	"strconv"
)

type Dts struct {
	Name int
	Data int `json:"data"`
}

func (d Dts) CALL() {
	fmt.Println("Call:" + strconv.Itoa(d.Name))
}

//func main5() {
//	dt := Dts{
//		Name: 1,
//		Data: 1,
//	}
//	t := reflect.TypeOf(dt)
//	v := reflect.ValueOf(&dt).Elem() //需要该数据 传指针
//
//	// 反射获取结构体中参数、标签、修改数据
//	for i := 0; i < t.NumField(); i++ {
//		field := t.Field(i)
//		jsonField := field.Tag.Get("json")
//		if jsonField == "" {
//			v.Field(i).SetInt(123)
//		}
//	}
//
//	//反射获取结构体方法并调用
//	for i := 0; i < t.NumMethod(); i++ {
//		methodType := t.Method(i)
//		if methodType.Name == "CAll" {
//			methodValue := v.Method(i)
//			methodValue.Call([]reflect.Value{})
//		}
//	}
//
//	fmt.Println(dt)
//}

// 反射判断obj的类型
func refType(obj any) {
	typeObj := reflect.TypeOf(obj)
	fmt.Println(typeObj, typeObj.Kind())
	// 去判断具体的类型
	switch typeObj.Kind() {
	case reflect.Slice:
		fmt.Println("切片")
	case reflect.Map:
		fmt.Println("map")
	case reflect.Struct:
		fmt.Println("结构体")
	case reflect.String:
		fmt.Println("字符串")
	}
}

// 反射获取值
func refValue(obj any) {
	value := reflect.ValueOf(obj)
	fmt.Println(value, value.Type())
	switch value.Kind() {
	case reflect.Int:
		fmt.Println(value.Int())
	case reflect.Struct:
		fmt.Println(value.Interface())
	case reflect.String:
		fmt.Println(value.String())

	}
}

// 反射修改值  必须要传指针
func refSetValue(obj any) {
	value := reflect.ValueOf(obj)
	elem := value.Elem()
	// 专门取指针反射的值
	switch elem.Kind() {
	case reflect.String:
		elem.SetString("1234")
	}
}
