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
		values[field.Name] = rv.Field(i)

		naming, ok := field.Tag.Lookup("naming")
		if ok {
			switch NamingStyle(naming) {
			case SNAKE_CASE:
				field.Tag = reflect.StructTag(`json:"` + SnakeCase(field.Name) + `"`)

			case LOWER_CAMEL_CASE:
				// TODO

			case UPPER_CAMEL_CASE:
				// TODO

			default:
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

	for _, c := range camel {
		if c >= 'A' && c <= 'Z' {
			snake += "_"
			snake += string(c + 32)
			continue
		}

		snake += string(c)
	}

	return snake[1:]
}
