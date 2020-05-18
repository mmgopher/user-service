package controller

import (
	"net/url"
	"strconv"
)

func getPaginationURLs(reqURL *url.URL, beforeID int, afterID int) (prevURL, nextURL string) {

	// keep all query params, except pagination related
	queryParams := reqURL.Query()
	queryParams.Del("before_id")
	queryParams.Del("after_id")

	// append before_id if needed
	if beforeID > 0 {
		queryParams["before_id"] = []string{strconv.Itoa(beforeID)}
		reqURL.RawQuery = queryParams.Encode()
		prevURL = reqURL.String()
		queryParams.Del("before_id")
	}

	// append after_id if needed
	if afterID > 0 {
		queryParams["after_id"] = []string{strconv.Itoa(afterID)}
		reqURL.RawQuery = queryParams.Encode()
		nextURL = reqURL.String()
	}

	return
}
