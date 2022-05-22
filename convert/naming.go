package convert

import (
	"reflect"
)

type NamingStyle string

const (
	SNAKE_CASE       NamingStyle = "snake_case"
	LOWER_CAMEL_CASE NamingStyle = "camelCase"
	UPPER_CAMEL_CASE NamingStyle = "CamelCase"
)

func ConvertNaming(origin interface{}) interface{} {
	rt := reflect.TypeOf(origin)
	rv := reflect.ValueOf(origin)
	if rv.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = rv.Elem()
	}

	values := make(map[string]reflect.Value)
	var structFields []reflect.StructField
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)

		switch value.Kind() {
		case reflect.Slice:
			rs := reflect.MakeSlice(value.Type(), value.Len(), value.Cap())
			for j := 0; j < value.Len(); j++ {
				index := value.Index(j)
				if index.Kind() == reflect.Ptr {
					index = index.Elem()
				}

				switch index.Kind() {
				case reflect.Struct:
					// iface := ConvertNaming(index.Interface())
					// collection[j] = iface

				case reflect.String:
					rs.Index(j).Set(index)
				}

			}

			// value = rs
		}

		values[field.Name] = value

		naming, ok := field.Tag.Lookup("naming")
		if ok {
			switch NamingStyle(naming) {
			case SNAKE_CASE:
				field.Tag = reflect.StructTag(`json:"` + SnakeCase(field.Name) + `"`)

			case LOWER_CAMEL_CASE:
				// TODO

			case UPPER_CAMEL_CASE:
				// TODO

			}
		}

		structFields = append(structFields, field)
	}

	newStruct := reflect.StructOf(structFields)
	impl := reflect.New(newStruct)
	newRV := impl.Elem()

	for i := 0; i < newRV.NumField(); i++ {
		field := structFields[i]
		value := newRV.Field(i)

		v, ok := values[field.Name]
		if !ok {
			continue
		}

		value.Set(v)
	}

	return impl.Interface()
}

func SnakeCase(camel string) string {
	var snake string

	for i, c := range camel {
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				snake += "_"
			}

			snake += string(c + 32)
			continue
		}

		snake += string(c)
	}

	return snake
}
