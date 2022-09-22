/*
example:

func main() {
	p := &Person{
		Name: "xx",
		Age: 12,
	}
	m := StructToMapViaReflect(p)
	fmt.Println(m)
}
*/
package v2

import (
	"reflect"
)

// 将结构体转为 map，用以通过 for 循环同时获取 key 和 value，以便将这些信息添加到请求体中或者 url 的 query 中
func StructToMapStr(obj interface{}) map[string]string {
	data := make(map[string]string)

	objV := reflect.ValueOf(obj)
	v := objV.Elem()
	typeOfType := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tField := typeOfType.Field(i)
		tFieldTag := string(tField.Tag.Get("request"))
		if len(tFieldTag) > 0 {
			data[tFieldTag] = field.String()
		} else {
			data[tField.Name] = field.String()
		}
	}
	return data
}
