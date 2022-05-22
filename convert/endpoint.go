package convert

import (
	"context"
	"errors"

	"gitlab.com/mirror520/json/endpoint"
)

func SnakeCaseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(map[string]interface{})
		if !ok {
			return nil, errors.New("invalid params")
		}

		resp, err := svc.SnakeCase(ctx, req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}
