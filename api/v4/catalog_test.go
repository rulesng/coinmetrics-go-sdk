package v4_test

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"testing"

	api "github.com/Improwised/coinmetrics-go-sdk/api/v4"
	"github.com/Improwised/coinmetrics-go-sdk/coinmetrics"
	"github.com/jarcoal/httpmock"
)

func TestAssetNotFoundForGetCatalogAssetsWithResponseHttpMock(t *testing.T) {
	t.Parallel()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	c := coinmetrics.InitClient(`http://fake-endpoint.com/`)
	param := api.GetCatalogAssetsParams{
		Assets: &api.CatalogAssetId{`sdvwbtc`},
	}
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'assets'. Value 'sdvwbtc' is not supported.`)
	response := api.GetCatalogAssetsResponse{
		Body:         byteEncoder(errResponse),
		HTTPResponse: &http.Response{StatusCode: 400, Status: `400`},
		JSON200:      nil,
		JSON400:      &errResponse,
		JSON401:      nil,
	}
	httpmock.RegisterResponder(http.MethodGet, `http://fake-endpoint.com/v4/catalog/assets?assets='sdvwbtc'`,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(400, response)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
	c.GetCatalogAssetsWithResponse(context.Background(), &param)

}

// func TestAssetNotFoundForGetCatalogAssetsWithResponse(t *testing.T) {
// 	t.Parallel()
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()

// 	client := mock_v4.NewMockClientWithResponsesInterface(controller)

// 	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'assets'. Value 'sdvwbtc' is not supported.`)

// 	param := api.GetCatalogAssetsParams{
// 		Assets: &api.CatalogAssetId{`sdvwbtc`},
// 	}
// 	response := api.GetCatalogAssetsResponse{
// 		Body:         byteEncoder(errResponse),
// 		HTTPResponse: &http.Response{StatusCode: 400, Status: `400`},
// 		JSON200:      nil,
// 		JSON400:      &errResponse,
// 		JSON401:      nil,
// 	}
// 	client.EXPECT().GetCatalogAssetsWithResponse(context.Background(), &param).Return(&response, nil)

// 	res, err := client.GetCatalogAssetsWithResponse(context.Background(), &param)
// 	assert.Nil(t, err)
// 	assert.Equal(t, res.Body, byteEncoder(errResponse))
// }

// func TestGetCatalogAssetsWithoutParams(t *testing.T) {
// 	t.Parallel()
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()

// 	client := mock_v4.NewMockClientWithResponsesInterface(controller)

// 	param := api.GetCatalogAssetsParams{}
// 	data := getCatalogAssetResponse(`{"data":[{"asset":"100x","full_name":"100xCoin","exchanges":["gate.io"],"markets":["gate.io-100x-usdt-spot"]},{"asset":"10set","full_name":"Tenset","exchanges":["gate.io","lbank"],"markets":["gate.io-10set-usdt-spot","lbank-10set-usdt-spot"]},{"asset":"18c","full_name":"Block 18","exchanges":["huobi"],"markets":["huobi-18c-btc-spot","huobi-18c-eth-spot"]},{"asset":"1art","full_name":"ArtWallet","exchanges":["gate.io"],"markets":["gate.io-1art-usdt-spot"]},{"asset":"1box","full_name":"1BOX","exchanges":["zb.com"],"markets":["zb.com-1box-usdt-spot"]},{"asset":"1earth","full_name":"EarthFund","exchanges":["gate.io","kucoin"],"markets":["gate.io-1earth-usdt-spot","kucoin-1earth-usdt-spot"]}]}`)
// 	response := api.GetCatalogAssetsResponse{
// 		Body:         byteEncoder(data),
// 		HTTPResponse: &http.Response{StatusCode: 200, Status: `200`},
// 		JSON200:      data,
// 		JSON400:      nil,
// 		JSON401:      nil,
// 	}
// 	client.EXPECT().GetCatalogAssetsWithResponse(context.Background(), &param).Return(&response, nil)

// 	res, err := client.GetCatalogAssetsWithResponse(context.Background(), &param)
// 	assert.Nil(t, err)
// 	assert.Equal(t, res.Body, byteEncoder(data))
// }

// func TestGetCatalogAssetsWithParams(t *testing.T) {
// 	t.Parallel()
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()

// 	client := mock_v4.NewMockClientWithResponsesInterface(controller)

// 	param := api.GetCatalogAssetsParams{
// 		Assets: &api.CatalogAssetId{`100x`},
// 	}
// 	data := getCatalogAssetResponse(`{"data":[{"asset":"100x","full_name":"100xCoin","exchanges":["gate.io"],"markets":["gate.io-100x-usdt-spot"]}]}`)
// 	response := api.GetCatalogAssetsResponse{
// 		Body:         byteEncoder(data),
// 		HTTPResponse: &http.Response{StatusCode: 200, Status: `200`},
// 		JSON200:      data,
// 		JSON400:      nil,
// 		JSON401:      nil,
// 	}
// 	client.EXPECT().GetCatalogAssetsWithResponse(context.Background(), &param).Return(&response, nil)

// 	res, err := client.GetCatalogAssetsWithResponse(context.Background(), &param)
// 	assert.Nil(t, err)
// 	assert.Equal(t, res.Body, byteEncoder(data))
// }

func byteEncoder(e interface{}) []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(e)
	if err != nil {
		return []byte{}
	}

	return network.Bytes()
}

func buildErrorMessage(message, errorType string) api.ErrorResponse {
	errObject := api.ErrorObject{
		Message: &message,
		Type:    errorType,
	}
	errResponse := api.ErrorResponse{errObject}
	return errResponse
}

func getCatalogAssetResponse(res string) *api.AssetsResponse {
	responseStruct := api.AssetsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.AssetsResponse{}
	}
	return &responseStruct
}
