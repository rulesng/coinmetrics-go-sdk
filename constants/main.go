package constants

const (
	API_VERSION = "v4"

	// API query params
	PARAMS_API_KEY = `api_key`

	// Error message
	NO_DATA_FOUND = `no data found`

	// Test
	TEST_ENDPOINT = `http://fake-endpoint.com/`
	TEST_KEY      = `abc`
	// DEFAULT_PAGE_SIZE Every api call would get this legnth of data to avoid delay
	DEFAULT_PAGE_SIZE int32 = 100
	// DEFAULT_PAGE_LIMIT Represents pagination how many records you want to fetch
	DEFAULT_PAGE_LIMIT int32 = -1
)
