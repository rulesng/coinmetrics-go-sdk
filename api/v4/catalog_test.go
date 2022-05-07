package v4_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	api "github.com/rulesng/coinmetrics-go-sdk/api/v4"
	"github.com/rulesng/coinmetrics-go-sdk/coinmetrics"
	"github.com/rulesng/coinmetrics-go-sdk/constants"
	"github.com/stretchr/testify/assert"
)

var _coinmetrics coinmetrics.CoinMetrics

func TestMain(m *testing.M) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	var _err error
	_coinmetrics, _err = coinmetrics.InitClient(constants.TEST_ENDPOINT, constants.TEST_KEY)
	if _err != nil {
		fmt.Println(_err)
		//panic(_err)
	}
	os.Exit(m.Run())
}

func TestAssetNotFoundForGetCatalogAssetsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'assets'. Value 'sdvwbtc' is not supported.`)
	param := api.GetCatalogAssetsParams{
		Assets: &api.CatalogAssetId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
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

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
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

func TestAssetNotFoundForGetCatalogMetricsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'metrics'. Value 'asdgwav' is not supported.`)
	param := api.GetCatalogMetricsParams{
		Metrics: &api.CatalogMetric{`asdgwav`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMetricsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogMetricsWithoutParamsResponse(t *testing.T) {
	data := getCatalogMetricsResponse(`{"data":[{"metric":"AdrActCnt","full_name":"Addresses, active, count","description":"The sum count of unique addresses that were active in the network (either as a recipient or originator of a ledger change) that interval. All parties in a ledger change action (recipients and originators) are counted. Individual addresses are not double-counted if previously active.","category":"Addresses","subcategory":"Active","unit":"Addresses","data_type":"bigint","type":"Sum","frequencies":[{}],"display_name":"Active Addr Cnt"}]}`)
	param := api.GetCatalogMetricsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMetricsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogMetricsWithParamsResponse(t *testing.T) {
	data := getCatalogMetricsResponse(`{"data":[{"metric":"AdrActCnt","full_name":"Addresses, active, count","description":"The sum count of unique addresses that were active in the network (either as a recipient or originator of a ledger change) that interval. All parties in a ledger change action (recipients and originators) are counted. Individual addresses are not double-counted if previously active.","category":"Addresses","subcategory":"Active","unit":"Addresses","data_type":"bigint","type":"Sum","frequencies":[{}],"display_name":"Active Addr Cnt"}]}`)
	param := api.GetCatalogMetricsParams{
		Metrics: &api.CatalogMetric{`AdrActCnt`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMetricsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogMetricsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogMetricsParams{
		Metrics: &api.CatalogMetric{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMetricsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Exchange

func TestExchangeNotFoundForGetCatalogExchangesWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'id'. Exchange 'sdvwbtc' is not supported.`)
	param := api.GetCatalogExchangesParams{
		Exchanges: &api.CatalogExchangeId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchanges`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogExchangesWithoutParams(t *testing.T) {
	data := getCatalogExchangeResponse(`{"data":[{"exchange":"bitbank","markets":["bitbank-bch-btc-spot","bitbank-bch-jpy-spot","bitbank-btc-jpy-spot","bitbank-eth-btc-spot","bitbank-eth-jpy-spot","bitbank-ltc-btc-spot","bitbank-ltc-jpy-spot","bitbank-mona-btc-spot","bitbank-mona-jpy-spot","bitbank-xrp-btc-spot","bitbank-xrp-jpy-spot"],"min_time":"2017-02-14T03:26:18.512000000Z","max_time":"2022-04-25T07:01:12.566000000Z","metrics":[{"metric":"volume_reported_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2017-02-14T04:00:00.000000000Z","max_time":"2022-04-25T06:00:00.000000000Z"}]},{"metric":"volume_reported_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2017-02-15T00:00:00.000000000Z","max_time":"2022-04-24T00:00:00.000000000Z"}]}]},{"exchange":"itbit","markets":["itbit-aave-usd-spot","itbit-bch-usd-spot","itbit-btc-eur-spot","itbit-btc-sgd-spot","itbit-btc-usd-spot","itbit-eth-eur-spot","itbit-eth-sgd-spot","itbit-eth-usd-spot","itbit-link-usd-spot","itbit-ltc-usd-spot","itbit-matic-usd-spot","itbit-paxg-usd-spot","itbit-uni-usd-spot"],"min_time":"2019-03-13T07:53:05.963000000Z","max_time":"2022-04-25T07:01:02.866000000Z","metrics":[{"metric":"volume_reported_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2019-03-13T08:00:00.000000000Z","max_time":"2022-04-25T06:00:00.000000000Z"}]},{"metric":"volume_reported_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2019-03-14T00:00:00.000000000Z","max_time":"2022-04-24T00:00:00.000000000Z"}]}]}]}
	`)
	param := api.GetCatalogExchangesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchanges`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogExchangesWithParams(t *testing.T) {
	data := getCatalogExchangeResponse(`{"data":[{"exchange":"bitbank","markets":["bitbank-bch-btc-spot","bitbank-bch-jpy-spot","bitbank-btc-jpy-spot","bitbank-eth-btc-spot","bitbank-eth-jpy-spot","bitbank-ltc-btc-spot","bitbank-ltc-jpy-spot","bitbank-mona-btc-spot","bitbank-mona-jpy-spot","bitbank-xrp-btc-spot","bitbank-xrp-jpy-spot"],"min_time":"2017-02-14T03:26:18.512000000Z","max_time":"2022-04-25T07:01:12.566000000Z","metrics":[{"metric":"volume_reported_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2017-02-14T04:00:00.000000000Z","max_time":"2022-04-25T06:00:00.000000000Z"}]},{"metric":"volume_reported_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2017-02-15T00:00:00.000000000Z","max_time":"2022-04-24T00:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogExchangesParams{
		Exchanges: &api.CatalogExchangeId{`bitbank`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchanges`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogExchangesWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogExchangesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchanges`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Exchange Asset Pairs

func TestExchangeAssetsNotFoundForGetCatalogExchangeAssetsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'pairs'. Value 'sdvwbtc' is not supported.`)
	param := api.GetCatalogExchangeAssetsParams{
		ExchangeAssets: &api.CatalogExchangeAssetId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchange-assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangeAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogExchangeAssetsWithoutParams(t *testing.T) {
	data := getCatalogExchangeAssetsResponse(`{"data":[{"exchange_asset":"binance-btc","metrics":[{"metric":"volume_trusted_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2020-10-16T00:00:00.000000000Z","max_time":"2021-01-05T00:00:00.000000000Z"}]},{"metric":"volume_trusted_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-10-15T03:00:00.000000000Z","max_time":"2021-01-06T12:00:00.000000000Z"}]}]},{"exchange_asset":"coinbase-eth","metrics":[{"metric":"volume_trusted_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2020-10-11T00:00:00.000000000Z","max_time":"2021-01-05T00:00:00.000000000Z"}]},{"metric":"volume_trusted_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-10-10T19:00:00.000000000Z","max_time":"2021-01-06T12:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogExchangeAssetsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchange-assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangeAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogExchangeAssetsWithParams(t *testing.T) {
	data := getCatalogExchangeAssetsResponse(`{"data":[{"exchange_asset":"binance-btc","metrics":[{"metric":"volume_trusted_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2020-10-16T00:00:00.000000000Z","max_time":"2021-01-05T00:00:00.000000000Z"}]},{"metric":"volume_trusted_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-10-15T03:00:00.000000000Z","max_time":"2021-01-06T12:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogExchangeAssetsParams{
		ExchangeAssets: &api.CatalogExchangeAssetId{`binance-btc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchange-assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangeAssetsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogExchangeAssetsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogExchangesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/exchange-assets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Exchange Asset Pairs

func TestPairsNotFoundForGetCatalogAssetPairsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'pairs'. Value 'sdvwbtc' is not supported.`)
	param := api.GetCatalogAssetPairsParams{
		Pairs: &api.CatalogPairId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/pairs`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetPairsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogAssetPairsWithoutParams(t *testing.T) {
	data := getCatalogAssetPairsResponse(`{"data":[{"pair": "1inch-btc", "metrics": [{ "metric": "volume_reported_spot_usd_1d", "frequencies": [{ "frequency": "1d", "min_time": "2020-12-30T00:00:00.000000000Z", "max_time": "2022-04-25T00:00:00.000000000Z" }] }, { "metric": "volume_reported_spot_usd_1h", "frequencies": [{ "frequency": "1h", "min_time": "2020-12-30T00:00:00.000000000Z", "max_time": "2022-04-26T10:00:00.000000000Z" }] }, { "metric": "volume_trusted_spot_usd_1d", "frequencies": [{ "frequency": "1d", "min_time": "2020-12-30T00:00:00.000000000Z", "max_time": "2022-04-25T00:00:00.000000000Z" }] }, { "metric": "volume_trusted_spot_usd_1h", "frequencies": [{ "frequency": "1h", "min_time": "2020-12-30T00:00:00.000000000Z", "max_time": "2022-04-26T10:00:00.000000000Z" }] }] }, { "pair": "1inch-busd", "metrics": [{ "metric": "volume_reported_spot_usd_1h", "frequencies": [{ "frequency": "1h", "min_time": "2021-02-23T12:00:00.000000000Z", "max_time": "2022-04-26T10:00:00.000000000Z" }] }, { "metric": "volume_reported_spot_usd_1d", "frequencies": [{ "frequency": "1d", "min_time": "2021-02-24T00:00:00.000000000Z", "max_time": "2022-04-25T00:00:00.000000000Z" }] }, { "metric": "volume_trusted_spot_usd_1d", "frequencies": [{ "frequency": "1d", "min_time": "2021-02-24T00:00:00.000000000Z", "max_time": "2022-04-25T00:00:00.000000000Z" }] }, { "metric": "volume_trusted_spot_usd_1h", "frequencies": [{ "frequency": "1h", "min_time": "2021-02-23T12:00:00.000000000Z", "max_time": "2022-04-26T10:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogAssetPairsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/pairs`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetPairsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogAssetPairsWithParams(t *testing.T) {
	data := getCatalogAssetPairsResponse(`{"data":[{"exchange_asset":"binance-btc","metrics":[{"metric":"volume_trusted_spot_usd_1d","frequencies":[{"frequency":"1d","min_time":"2020-10-16T00:00:00.000000000Z","max_time":"2021-01-05T00:00:00.000000000Z"}]},{"metric":"volume_trusted_spot_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-10-15T03:00:00.000000000Z","max_time":"2021-01-06T12:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogAssetPairsParams{
		Pairs: &api.CatalogPairId{`1inch-btc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/pairs`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetPairsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogAssetPairsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogExchangesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/pairs`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogExchangesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Available Institution

func TestInstitutionNotFoundForGetCatalogInstitutionsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'pairs'. Value 'sdvwbtc' is not supported.`)
	param := api.GetCatalogInstitutionsParams{
		Institutions: &api.CatalogInstitutionId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/institutions`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogInstitutionsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogInstitutionsWithoutParams(t *testing.T) {
	data := getCatalogInstitutionsResponse(`{"data":[{"institution":"grayscale","metrics":[{"metric":"batfund_net_asset_value","frequencies":[{"frequency":"1d","min_time":"2021-03-17T00:00:00.000000000Z","max_time":"2022-02-11T00:00:00.000000000Z"}]},{"metric":"xlmfund_net_asset_value","frequencies":[{"frequency":"1d","min_time":"2018-12-06T00:00:00.000000000Z","max_time":"2021-10-19T00:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogInstitutionsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/institutions`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogInstitutionsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogInstitutionsWithParams(t *testing.T) {
	data := getCatalogInstitutionsResponse(`{"data":[{"institution":"grayscale","metrics":[{"metric":"batfund_net_asset_value","frequencies":[{"frequency":"1d","min_time":"2021-03-17T00:00:00.000000000Z","max_time":"2022-02-11T00:00:00.000000000Z"}]},{"metric":"xlmfund_net_asset_value","frequencies":[{"frequency":"1d","min_time":"2018-12-06T00:00:00.000000000Z","max_time":"2021-10-19T00:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogInstitutionsParams{
		Institutions: &api.CatalogInstitutionId{`grayscale`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/institutions`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogInstitutionsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogInstitutionsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogInstitutionsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/institutions`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogInstitutionsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Available indexes

func TestIndexesNotFoundForGetCatalogIndexesWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'indexes'. Value 'asdgwav' is not supported.`)
	param := api.GetCatalogIndexesParams{
		Indexes: &api.CatalogIndexId{`asdgwav`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/indexes`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogIndexesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogIndexesWithoutParams(t *testing.T) {
	data := getCatalogIndexesResponse(`{"data":[{"index":"CMBI10","description":"'CMBI10' index.","frequencies":[{"frequency":"15s","min_time":"2020-06-08T20:12:40.000000000Z","max_time":"2020-06-08T20:29:30.000000000Z"}]},{"index":"CMBIBTC","description":"'CMBIBTC' index.","frequencies":[{"frequency":"15s","min_time":"2010-07-18T20:00:00.000000000Z","max_time":"2020-06-08T20:29:45.000000000Z"},{"frequency":"1d","min_time":"2010-07-19T00:00:00.000000000Z","max_time":"2020-06-08T00:00:00.000000000Z"},{"frequency":"1d-ny-close","min_time":"2010-07-18T20:00:00.000000000Z","max_time":"2020-06-08T20:00:00.000000000Z"},{"frequency":"1d-sg-close","min_time":"2010-07-19T08:00:00.000000000Z","max_time":"2020-06-08T08:00:00.000000000Z"},{"frequency":"1h","min_time":"2010-07-18T20:00:00.000000000Z","max_time":"2020-06-08T20:00:00.000000000Z"}]}]}`)
	param := api.GetCatalogIndexesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/indexes`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogIndexesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogIndexesWithParams(t *testing.T) {
	data := getCatalogIndexesResponse(`{"data":[{"index":"CMBI10","description":"'CMBI10' index.","frequencies":[{"frequency":"15s","min_time":"2020-06-08T20:12:40.000000000Z","max_time":"2020-06-08T20:29:30.000000000Z"}]}]}`)
	param := api.GetCatalogIndexesParams{
		Indexes: &api.CatalogIndexId{`CMBI10`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/indexes`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogIndexesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogIndexesWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogIndexesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/indexes`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogIndexesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Available assert alerts

func TestAlertNotFoundForGetCatalogAssetAlertRulesWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'assets'. Value 'sdvwbtc' is not supported.`)
	param := api.GetCatalogAssetAlertRulesParams{
		Assets: &api.CatalogAssetId{`sdvwbtc`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/alerts`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetAlertRulesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogAssetAlertWithoutParams(t *testing.T) {
	data := getCatalogAssetAlertRulesResponse(`{"data":[{"asset":"btc","name":"block_count_empty_6b_hi","conditions":[{"description":"The last 4 blocks were empty.","threshold":"4","constituents":["block_count_empty_6b"]}]}]}`)
	param := api.GetCatalogAssetAlertRulesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/alerts`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetAlertRulesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogAssetAlertWithParams(t *testing.T) {
	data := getCatalogAssetAlertRulesResponse(`{"data":[{"asset":"btc","name":"block_count_empty_6b_hi","conditions":[{"description":"The last 4 blocks were empty.","threshold":"4","constituents":["block_count_empty_6b"]}]}]}`)
	param := api.GetCatalogAssetAlertRulesParams{
		Assets: &api.CatalogAssetId{`btc`},
		Alerts: &api.CatalogAssetAlertId{`block_count_empty_6b`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/alerts`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetAlertRulesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogAssetAlertWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogAssetAlertRulesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/alerts`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogAssetAlertRulesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Available markets

func TestMetricsNotFoundForGetCatalogMarketsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'markets'. Invalid format for market: 'asdw'.`)
	param := api.GetCatalogMarketsParams{
		Markets: &api.CatalogMarketId{`asdw`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/markets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogMarketsWithoutParams(t *testing.T) {
	data := getCatalogMarketsResponse(`{"data":[{"market":"bitmex-XBTF15-future","min_time":"2014-11-24T13:05:32.850000000Z","max_time":"2015-01-30T12:00:00.000000000Z","trades":{"min_time":"2014-11-24T13:05:32.850000000Z","max_time":"2015-01-30T12:00:00.000000000Z"},"exchange":"bitmex","type":"future","symbol":"XBTF15","base":"btc","quote":"usd","size_asset":"XBT","margin_asset":"USD","contract_size":"1","tick_size":"0.1","listing":"2014-11-24T13:05:32.850000000Z","expiration":"2015-01-30T12:00:00.000000000Z"},{"market":"bitfinex-agi-btc-spot","min_time":"2018-04-07T16:25:55.000000000Z","max_time":"2020-03-25T20:12:09.639000000Z","trades":{"min_time":"2018-04-07T16:25:55.000000000Z","max_time":"2020-03-25T20:12:09.639000000Z"},"exchange":"bitfinex","type":"spot","base":"agi","quote":"btc"}]}`)
	param := api.GetCatalogMarketsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/markets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogMarketsWithResponse(t *testing.T) {
	exchange := `bitmex`
	var marketType api.GetCatalogMarketsParamsType = `future`
	var base api.MarketBase = `btc`
	var asset api.MarketAsset = `btc`
	var symbol api.MarketSymbol = `XBTF15`
	var limit api.CatalogMarketLimit = `1`
	var format api.CatalogMarketFormat = `json`
	var quote api.MarketQuoteParameter = `usd`
	data := getCatalogMarketsResponse(`{"data":[{"market":"bitmex-XBTF15-future","min_time":"2014-11-24T13:05:32.850000000Z","max_time":"2015-01-30T12:00:00.000000000Z","trades":{"min_time":"2014-11-24T13:05:32.850000000Z","max_time":"2015-01-30T12:00:00.000000000Z"},"exchange":"bitmex","type":"future","symbol":"XBTF15","base":"btc","quote":"usd","size_asset":"XBT","margin_asset":"USD","contract_size":"1","tick_size":"0.1","listing":"2014-11-24T13:05:32.850000000Z","expiration":"2015-01-30T12:00:00.000000000Z"}]}`)
	param := api.GetCatalogMarketsParams{
		Markets:  &api.CatalogMarketId{`bitmex-XBTF15-future`},
		Exchange: &exchange,
		Type:     &marketType,
		Base:     &base,
		Asset:    &asset,
		Symbol:   &symbol,
		Quote:    &quote,
		Include:  &api.CatalogMarketIncludeFields{`trades`, `orderbooks`, `quotes`, `candles`, `liquidations`},
		Exclude:  &api.CatalogMarketExcludeFields{`funding_rates`, `openinterest`},
		Format:   &format,
		Limit:    &limit,
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/markets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogMarketsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogMarketsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/markets`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Available markets candles

func TestMetricsNotFoundForGetCatalogMarketCandlesWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'markets'. Invalid format for market: 'asdw'.`)
	param := api.GetCatalogMarketCandlesParams{
		Markets: &api.CatalogMarketId{`asdw`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/market-candles`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketCandlesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogMarketCandlesWithoutParams(t *testing.T) {
	data := getCatalogMarketCandlesResponse(`{"data":[{"market":"binance-BTCUSDT-future","metrics":[{"metric":"liquidations_reported_future_buy_usd_5m","frequencies":[{"frequency":"5m","min_time":"2020-01-01T01:25:00.000000000Z","max_time":"2022-01-21T00:30:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-01-01T01:00:00.000000000Z","max_time":"2022-01-20T23:00:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_units_1d","frequencies":[{"frequency":"1d","min_time":"2020-01-01T00:00:00.000000000Z","max_time":"2022-01-20T00:00:00.000000000Z"}]}]},{"market":"bybit-BTCUSDT-future","metrics":[{"metric":"liquidations_reported_future_buy_usd_5m","frequencies":[{"frequency":"5m","min_time":"2021-04-30T12:35:00.000000000Z","max_time":"2022-01-21T00:25:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_usd_1h","frequencies":[{"frequency":"1h","min_time":"2021-04-30T12:00:00.000000000Z","max_time":"2022-01-20T23:00:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_units_1d","frequencies":[{"frequency":"1d","min_time":"2021-04-30T00:00:00.000000000Z","max_time":"2022-01-20T00:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogMarketCandlesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/market-candles`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketCandlesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogMarketCandlesWithResponse(t *testing.T) {
	exchange := `bitmex`
	var marketType api.GetCatalogMarketCandlesParamsType = `future`
	var base api.MarketBase = `btc`
	var asset api.MarketAsset = `btc`
	var symbol api.MarketSymbol = `XBTF15`
	var limit api.CatalogMarketLimit = `1`
	var format api.CatalogMarketFormat = `json`
	var quote api.MarketQuoteParameter = `usd`
	data := getCatalogMarketCandlesResponse(`{"data":[{"market":"binance-BTCUSDT-future","metrics":[{"metric":"liquidations_reported_future_buy_usd_5m","frequencies":[{"frequency":"5m","min_time":"2020-01-01T01:25:00.000000000Z","max_time":"2022-01-21T00:30:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-01-01T01:00:00.000000000Z","max_time":"2022-01-20T23:00:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_units_1d","frequencies":[{"frequency":"1d","min_time":"2020-01-01T00:00:00.000000000Z","max_time":"2022-01-20T00:00:00.000000000Z"}]}]},{"market":"bybit-BTCUSDT-future","metrics":[{"metric":"liquidations_reported_future_buy_usd_5m","frequencies":[{"frequency":"5m","min_time":"2021-04-30T12:35:00.000000000Z","max_time":"2022-01-21T00:25:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_usd_1h","frequencies":[{"frequency":"1h","min_time":"2021-04-30T12:00:00.000000000Z","max_time":"2022-01-20T23:00:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_units_1d","frequencies":[{"frequency":"1d","min_time":"2021-04-30T00:00:00.000000000Z","max_time":"2022-01-20T00:00:00.000000000Z"}]}]}]}`)
	param := api.GetCatalogMarketCandlesParams{
		Markets:  &api.CatalogMarketId{`bitmex-XBTF15-future`},
		Exchange: &exchange,
		Type:     &marketType,
		Base:     &base,
		Asset:    &asset,
		Symbol:   &symbol,
		Format:   &format,
		Quote:    &quote,
		Limit:    &limit,
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/market-candles`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketCandlesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogMarketCandlesWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogMarketCandlesParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/market-candles`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketCandlesWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Nil(t, actualResponse.JSON400)
	assert.Equal(t, *actualResponse.JSON401, errResponse)
}

// Catalog Available market metrics

func TestMetricsNotFoundForGetCatalogMarketMetricsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`bad_request`, `Bad parameter 'markets'. Invalid format for market: 'asdw'.`)
	param := api.GetCatalogMarketMetricsParams{
		Markets: &api.CatalogMarketId{`asdw`},
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/market-metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusBadRequest, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketMetricsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Nil(t, actualResponse.JSON200)
	assert.Equal(t, *actualResponse.JSON400, errResponse)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogMarketMetricsWithoutParams(t *testing.T) {
	data := getCatalogMarketMetricsResponse(`{"data":[{"market":"binance-BTCUSDT-future","metrics":[{"metric":"liquidations_reported_future_buy_usd_5m","frequencies":[{"frequency":"5m","min_time":"2020-01-01T01:25:00.000000000Z","max_time":"2022-01-21T00:30:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-01-01T01:00:00.000000000Z","max_time":"2022-01-20T23:00:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_units_1d","frequencies":[{"frequency":"1d","min_time":"2020-01-01T00:00:00.000000000Z","max_time":"2022-01-20T00:00:00.000000000Z"}]}]},{"market":"bybit-BTCUSDT-future","metrics":[{"metric":"liquidations_reported_future_buy_usd_5m","frequencies":[{"frequency":"5m","min_time":"2021-04-30T12:35:00.000000000Z","max_time":"2022-01-21T00:25:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_usd_1h","frequencies":[{"frequency":"1h","min_time":"2021-04-30T12:00:00.000000000Z","max_time":"2022-01-20T23:00:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_units_1d","frequencies":[{"frequency":"1d","min_time":"2021-04-30T00:00:00.000000000Z","max_time":"2022-01-20T00:00:00.000000000Z"}]}]}]}
	`)
	param := api.GetCatalogMarketMetricsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/market-metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketMetricsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestGetCatalogMarketMetricsWithResponse(t *testing.T) {
	exchange := `binance`
	var base api.MarketBase = `btc`
	var asset api.MarketAsset = `btc`
	var symbol api.MarketSymbol = `BTCUSDT`
	var limit api.CatalogMarketLimit = `1`
	var format api.CatalogMarketFormat = `json`
	data := getCatalogMarketMetricsResponse(`{"data":[{"market":"binance-BTCUSDT-future","metrics":[{"metric":"liquidations_reported_future_buy_usd_5m","frequencies":[{"frequency":"5m","min_time":"2020-01-01T01:25:00.000000000Z","max_time":"2022-01-21T00:30:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_usd_1h","frequencies":[{"frequency":"1h","min_time":"2020-01-01T01:00:00.000000000Z","max_time":"2022-01-20T23:00:00.000000000Z"}]},{"metric":"liquidations_reported_future_buy_units_1d","frequencies":[{"frequency":"1d","min_time":"2020-01-01T00:00:00.000000000Z","max_time":"2022-01-20T00:00:00.000000000Z"}]}]}]}
	`)
	param := api.GetCatalogMarketMetricsParams{
		Markets:  &api.CatalogMarketId{`binance-BTCUSDT-future`},
		Exchange: &exchange,
		Base:     &base,
		Asset:    &asset,
		Symbol:   &symbol,
		Format:   &format,
		Limit:    &limit,
	}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/market-metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusOK, data)
			if err != nil {
				return nil, err
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketMetricsWithResponse(context.Background(), &param)
	assert.Nil(t, err)
	assert.Equal(t, *actualResponse.JSON200, *data)
	assert.Nil(t, actualResponse.JSON400)
	assert.Nil(t, actualResponse.JSON401)
}

func TestFailAuthenticationForGetCatalogMarketMetricsWithResponse(t *testing.T) {
	errResponse := buildErrorMessage(`unauthorized`, `Requested resource requires authorization.`)
	param := api.GetCatalogMarketMetricsParams{}

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(`%s%s/catalog/market-metrics`, constants.TEST_ENDPOINT, constants.API_VERSION),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(http.StatusUnauthorized, errResponse)
			if err != nil {
				return httpmock.NewStringResponse(http.StatusInternalServerError, `Unable to return mock response`), nil
			}
			return resp, nil
		},
	)
	actualResponse, err := _coinmetrics.GetCatalogMarketMetricsWithResponse(context.Background(), &param)
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

func getCatalogMetricsResponse(res string) *api.MetricsResponse {
	responseStruct := api.MetricsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.MetricsResponse{}
	}
	return &responseStruct
}

func getCatalogAssetResponse(res string) *api.AssetsResponse {
	responseStruct := api.AssetsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.AssetsResponse{}
	}
	return &responseStruct
}

func getCatalogExchangeResponse(res string) *api.ExchangesResponse {
	responseStruct := api.ExchangesResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.ExchangesResponse{}
	}
	return &responseStruct
}

func getCatalogExchangeAssetsResponse(res string) *api.ExchangeAssetsResponse {
	responseStruct := api.ExchangeAssetsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.ExchangeAssetsResponse{}
	}
	return &responseStruct
}

func getCatalogAssetPairsResponse(res string) *api.PairsResponse {
	responseStruct := api.PairsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.PairsResponse{}
	}
	return &responseStruct
}

func getCatalogInstitutionsResponse(res string) *api.InstitutionsResponse {
	responseStruct := api.InstitutionsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.InstitutionsResponse{}
	}
	return &responseStruct
}

func getCatalogIndexesResponse(res string) *api.IndexesResponse {
	responseStruct := api.IndexesResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.IndexesResponse{}
	}
	return &responseStruct
}

func getCatalogAssetAlertRulesResponse(res string) *api.AssetAlertRulesResponse {
	responseStruct := api.AssetAlertRulesResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.AssetAlertRulesResponse{}
	}
	return &responseStruct
}

func getCatalogMarketsResponse(res string) *api.MarketsResponse {
	responseStruct := api.MarketsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.MarketsResponse{}
	}
	return &responseStruct
}

func getCatalogMarketCandlesResponse(res string) *api.CatalogMarketCandlesResponse {
	responseStruct := api.CatalogMarketCandlesResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.CatalogMarketCandlesResponse{}
	}
	return &responseStruct
}

func getCatalogMarketMetricsResponse(res string) *api.CatalogMarketMetricsResponse {
	responseStruct := api.CatalogMarketMetricsResponse{}
	err := json.Unmarshal([]byte(res), &responseStruct)
	if err != nil {
		return &api.CatalogMarketMetricsResponse{}
	}
	return &responseStruct
}
