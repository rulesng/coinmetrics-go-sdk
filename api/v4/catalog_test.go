package v4_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	api "github.com/Improwised/coinmetrics-go-sdk/api/v4"
	"github.com/Improwised/coinmetrics-go-sdk/coinmetrics"
	"github.com/Improwised/coinmetrics-go-sdk/constants"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var _coinmetrics coinmetrics.CoinMetrics

func TestMain(m *testing.M) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	_coinmetrics = coinmetrics.InitClient(constants.TEST_ENDPOINT)
	os.Exit(m.Run())
}

func TestAssetNotFoundForGetCatalogAssetsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'assets'. Value 'sdvwbtc' is not supported.`)
	param := api.GetCatalogAssetsParams{
		Assets: &api.CatalogAssetId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/assets?assets=sdvwbtc`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogAssetsWithoutParams(t *testing.T) {
	data := getCatalogAssetResponse(`{"data":[{"asset":"100x","full_name":"100xCoin","exchanges":["gate.io"],"markets":["gate.io-100x-usdt-spot"]},{"asset":"10set","full_name":"Tenset","exchanges":["gate.io","lbank"],"markets":["gate.io-10set-usdt-spot","lbank-10set-usdt-spot"]},{"asset":"18c","full_name":"Block 18","exchanges":["huobi"],"markets":["huobi-18c-btc-spot","huobi-18c-eth-spot"]},{"asset":"1art","full_name":"ArtWallet","exchanges":["gate.io"],"markets":["gate.io-1art-usdt-spot"]},{"asset":"1box","full_name":"1BOX","exchanges":["zb.com"],"markets":["zb.com-1box-usdt-spot"]},{"asset":"1earth","full_name":"EarthFund","exchanges":["gate.io","kucoin"],"markets":["gate.io-1earth-usdt-spot","kucoin-1earth-usdt-spot"]}]}`)
	param := api.GetCatalogAssetsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogAssetsWithParams(t *testing.T) {
	data := getCatalogAssetResponse(`{"data":[{"asset":"100x","full_name":"100xCoin","exchanges":["gate.io"],"markets":["gate.io-100x-usdt-spot"]}]}`)
	param := api.GetCatalogAssetsParams{
		Assets:  &api.CatalogAssetId{`100x`},
		Include: &api.CatalogAssetIncludeFields{`markets`, `exchanges`},
		Exclude: &api.CatalogAssetExcludeFields{`metrics`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogAssetsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogAssetsParams{
		Assets: &api.CatalogAssetId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/assets?assets=sdvwbtc`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
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
