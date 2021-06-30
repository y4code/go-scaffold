package reflect

import (
	"github.com/pelletier/go-toml"
	"reflect"
)

type ReflectOperator struct {
	target interface{}
	source interface{}
	tag    string
}

func NewReflectOperator(target interface{}, source interface{}, tag string) *ReflectOperator {
	return &ReflectOperator{target: target, source: source, tag: tag}
}

func (a ReflectOperator) AssignValueByTagFromToml() {
	t := reflect.TypeOf(a.target).Elem()
	v := reflect.ValueOf(a.target).Elem()
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get(a.tag)
		if tag != "" {
			//if s, ok := source.(*toml.Tree); ok {
			//	v.FieldByName(t.Field(i).Name).SetString(s.Get(tag).(string))
			//}
			gotValue := ((a.source).(*toml.Tree)).Get(tag)
			if gotValue == nil {
				gotValue = ""
			}
			v.FieldByName(t.Field(i).Name).SetString(gotValue.(string))
		}
	}
}
