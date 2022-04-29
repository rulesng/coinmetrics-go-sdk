# Coinmetrics SDK

## Install

- To use this sdk you need to have [golang](https://go.dev/doc/install) installed on your system.

```bash
go install github.com/Improwised/coinmetrics-go-sdk
```

- sync vendor files by running `go mod vendor`

## How to use

- After installing sdk, it requires credentials to be passed which can be further used to access api of coinmetrics.

```go
import "github.com/Improwised/coinmetrics-go-sdk"

func main() {
    client, err := coinmetrics.Init(`api-endpoint.com`, `api-key`)
    if err != nil {
        // Change how you want to handle an error
        panic(err)
    }
}
```

- Above snippet will initialize the client, which will be use to access api endpoint. Make sure to use same object which calling other methods.

## Usage

- All api which is listed at [Coinmetrics](https://docs.coinmetrics.io/api/v4) are implemented.

- This methods name are defined similar as per documentation.

- If you to get **STRUCT AS RESPONSE** then use method which name are `ending` with `WithResponse`

    Example : 
    ```go
    import (
        api "github.com/Improwised/coinmetrics-go-sdk/api/v4"
        "github.com/Improwised/coinmetrics-go-sdk"
    )
    func main() {
        client, err := coinmetrics.Init(`api-endpoint.com`, `api-key`)
        if err != nil {
        // Change how you want to handle an error
            panic(err)
        }

        // Prepare query parameter that you want to pass
        param := api.GetCatalogAssetsParams{
            Assets:  &api.CatalogAssetId{`100x`},
            Include: &api.CatalogAssetIncludeFields{`markets`, `exchanges`},
            Exclude: &api.CatalogAssetExcludeFields{`metrics`},
	    }
        response, err := client.GetCatalogAssetsWithResponse(context.Background(), &param)
        // further code handling
    }
    ```
### Custom Wrapper

- As you seen in usage how you can access method of api, on that api we have written some of the methods which will help to eliminate usecase of handling pagination.

- Custom method are ending with `WithResponseSync` will return slice and error in return.
