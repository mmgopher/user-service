// +build integration

package integration

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mmgopher/user-service/app"
	"github.com/mmgopher/user-service/app/api/response"
	"github.com/mmgopher/user-service/test/helpers"
)

var getUserListData = []struct {
	name            string
	responseSize    int
	queryParameters map[string]string
	firstUserID     int
	lastUserID      int
	linkPrev        string
	linkNext        string
}{
	{
		"SortAscByIDPage1FilterName",
		4,
		map[string]string{
			"limit": "4",
			"name":  "sorttest",
		},
		4,
		7,
		"",
		"/v1/users?after_id=7&limit=4&name=sorttest",
	},
	{
		"SortAscByIDPage2FilterName",
		4,
		map[string]string{
			"limit":    "4",
			"name":     "sorttest",
			"after_id": "7",
		},
		8,
		11,
		"/v1/users?before_id=8&limit=4&name=sorttest",
		"/v1/users?after_id=11&limit=4&name=sorttest",
	},
	{
		"SortAscByIDPage2PrevFilterName",
		4,
		map[string]string{
			"limit":     "4",
			"name":      "sorttest",
			"before_id": "8",
		},
		4,
		7,
		"",
		"/v1/users?after_id=7&limit=4&name=sorttest",
	},
	{
		"SortDescByIDPage1FilterName",
		4,
		map[string]string{
			"limit": "4",
			"name":  "sorttest",
			"sort":  ":desc",
		},
		13,
		10,
		"",
		"/v1/users?after_id=10&limit=4&name=sorttest&sort=%3Adesc",
	},
	{
		"SortDescByIDPage2FilterName",
		4,
		map[string]string{
			"limit":    "4",
			"name":     "sorttest",
			"sort":     ":desc",
			"after_id": "10",
		},
		9,
		6,
		"/v1/users?before_id=9&limit=4&name=sorttest&sort=%3Adesc",
		"/v1/users?after_id=6&limit=4&name=sorttest&sort=%3Adesc",
	},
	{
		"SortDescByIDPage2PrevFilterName",
		4,
		map[string]string{
			"limit":     "4",
			"name":      "sorttest",
			"sort":      ":desc",
			"before_id": "9",
		},
		13,
		10,
		"",
		"/v1/users?after_id=10&limit=4&name=sorttest&sort=%3Adesc",
	},
	{
		"SortAscByNamePage1FilterNameFilterGender",
		4,
		map[string]string{
			"limit":  "4",
			"name":   "sorttest",
			"sort":   "name:asc",
			"gender": "MaLe",
		},
		13,
		7,
		"",
		"/v1/users?after_id=7&gender=MaLe&limit=4&name=sorttest&sort=name%3Aasc",
	},
	{
		"SortDescByAgePage1FilterNameFilterGender",
		4,
		map[string]string{
			"limit":  "4",
			"name":   "sorttest",
			"sort":   "age:desc",
			"gender": "female",
		},
		12,
		11,
		"",
		"",
	},
	{
		"SortByIDPage1FilterNameFilterAge",
		3,
		map[string]string{
			"limit":   "4",
			"name":    "sorttest",
			"sort":    "id:asc",
			"min_age": "23",
			"max_age": "31",
		},
		4,
		13,
		"",
		"",
	},
	{
		"SortDescByAgePage1FilterNameFilterAge",
		3,
		map[string]string{
			"limit":   "4",
			"name":    "sorttest",
			"sort":    "age:desc",
			"min_age": "23",
			"max_age": "31",
		},
		13,
		8,
		"",
		"",
	},
	{
		"SortAscByAgePage1FilterNameFilterAge",
		4,
		map[string]string{
			"limit":   "4",
			"name":    "sorttest",
			"sort":    "age:asc",
			"min_age": "23",
			"max_age": "42",
		},
		8,
		6,
		"",
		"",
	},
}

func TestGetUserListOK(t *testing.T) {
	httpService := helpers.NewHTTPService(http.DefaultClient)
	statusCode, respBody, err := httpService.DoRequest(
		http.MethodGet,
		os.Getenv("APP_BASE_URL")+app.RootPath+app.GetUserListRoute,
		nil,
		nil,
		nil,
	)
	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, statusCode)
	var response response.UserListWithPagination
	assert.Nil(t, json.Unmarshal(respBody, &response))
	// I do not know exact size becasue other tests add and delete useres
	assert.True(t, len(response.Result) > 10)
}

func TestGetUserList_TableDriven(t *testing.T) {

	httpService := helpers.NewHTTPService(http.DefaultClient)
	for _, tt := range getUserListData {
		t.Run(tt.name, func(t *testing.T) {
			statusCode, respBody, err := httpService.DoRequest(
				http.MethodGet,
				os.Getenv("APP_BASE_URL")+app.RootPath+app.GetUserListRoute,
				tt.queryParameters,
				nil,
				nil,
			)
			require.Nil(t, err)
			assert.Equal(t, http.StatusOK, statusCode)
			var response response.UserListWithPagination
			assert.Nil(t, json.Unmarshal(respBody, &response))
			require.Equal(t, tt.responseSize, len(response.Result))
			assert.Equal(t, tt.firstUserID, response.Result[0].ID)
			assert.Equal(t, tt.lastUserID, response.Result[tt.responseSize-1].ID)
			assert.Equal(t, tt.linkPrev, response.Pagination.PrevLink)
			assert.Equal(t, tt.linkNext, response.Pagination.NextLink)

		})
	}

}
