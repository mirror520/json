package convert

import (
	"context"
	"errors"
)

type Service interface {
	SnakeCase(ctx context.Context, in map[string]interface{}) (map[string]interface{}, error)
}

type service struct {
}

func NewService() Service {
	return new(service)
}

func (svc *service) SnakeCase(ctx context.Context, in map[string]interface{}) (map[string]interface{}, error) {
	out, ok := svc.snakeCase(in).(map[string]interface{})
	if !ok {
		return nil, errors.New("convert fail")
	}

	return out, nil
}

func (svc *service) snakeCase(in interface{}) interface{} {
	switch val := in.(type) {
	case []interface{}:
		out := make([]interface{}, len(val))
		for i, v := range val {
			out[i] = svc.snakeCase(v)
		}
		return out

	case map[string]interface{}:
		out := make(map[string]interface{})
		for k, v := range val {
			out[SnakeCase(k)] = svc.snakeCase(v)
		}

		return out

	default:
		return val
	}
}
