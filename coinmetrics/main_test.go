package coinmetrics_test

import (
	"bytes"
	"context"
	"encoding/gob"
	"net/http"
	"testing"

	api "github.com/Improwised/coinmetrics-go-sdk/api/v4"
	"github.com/Improwised/coinmetrics-go-sdk/api/v4/mock_v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetCatalogAssets(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	client := mock_v4.NewMockClientWithResponsesInterface(controller)

	ctx := context.Background()
	ctx.Value(`abcc`)
	message := `bad_parameter`
	errResponse := api.ErrorResponse{
		api.ErrorObject{
			Message: &message,
			Type:    "Bad parameter 'assets'. Value 'sdvwbtc' is not supported.",
		},
	}

	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	err := enc.Encode(errResponse)
	assert.Nil(t, err)
	param := api.GetCatalogAssetsParams{}
	response := api.GetCatalogAssetsResponse{
		Body:         network.Bytes(),
		HTTPResponse: &http.Response{StatusCode: 400, Status: `400`},
		JSON200:      nil,
		JSON400:      &errResponse,
		JSON401:      nil,
	}
	client.EXPECT().GetCatalogAssetsWithResponse(ctx, &param).Return(&response, nil)

	client.GetCatalogAssetsWithResponse(ctx, &param)
}
