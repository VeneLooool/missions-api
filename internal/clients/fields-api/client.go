package fields_api

import (
	"context"
	"fmt"
	"github.com/VeneLooool/missions-api/internal/config"
	"github.com/VeneLooool/missions-api/internal/model"
	"github.com/VeneLooool/missions-api/internal/pb/fields-api/api/v1/fields"
	proto_model "github.com/VeneLooool/missions-api/internal/pb/fields-api/api/v1/model"
	"github.com/VeneLooool/missions-api/internal/pkg/error_hub"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Client struct {
	fieldsApi fields.FieldsClient
}

func New(ctx context.Context, cfg *config.FieldsApiClientConfig) (*Client, error) {
	conn, err := grpc.NewClient(fmt.Sprint("%s:%s", cfg.Host, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "grpc.NewClient()")
	}

	return &Client{
		fieldsApi: fields.NewFieldsClient(conn),
	}, nil
}

func (c *Client) GetFieldByID(ctx context.Context, fieldID uint64) (model.Field, error) {
	resp, err := c.fieldsApi.GetFieldByID(ctx, &fields.GetFieldByID_Request{Id: fieldID})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return model.Field{}, error_hub.ErrFieldNotFound
		}
		return model.Field{}, errors.Wrap(err, "c.fieldsApi.GetFieldByID()")
	}
	return transformFieldToModel(resp.GetField()), nil
}

func transformFieldToModel(protoField *proto_model.Field) model.Field {
	coordinates := make(model.Coordinates, 0, len(protoField.GetCoordinates()))
	for _, coordinate := range protoField.GetCoordinates() {
		coordinates = append(coordinates, model.Coordinate{
			Latitude:  coordinate.GetLatitude(),
			Longitude: coordinate.GetLongitude(),
		})
	}

	return model.Field{
		ID:          protoField.GetId(),
		Name:        protoField.GetName(),
		Culture:     protoField.GetCulture(),
		Coordinates: coordinates,
	}
}
