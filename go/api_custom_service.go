package openapi

import (
    "context"
)

type CustomApiService struct {
    DefaultApiService
}

func (s *DefaultApiService) UsersGet(ctx context.Context) (ImplResponse, error) {
	// TODO - update UsersGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []string{}) or use other options such as http.Ok ...
	return Response(200, []string{"foo","bar"}), nil

	//return Response(http.StatusNotImplemented, nil), errors.New("UsersGet method not implemented")
}
