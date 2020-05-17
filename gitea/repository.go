package gitea

import "encoding/json"

// SearchParams struct
type SearchParams struct {
	PackageName string
}

// SearchResponse struct
type SearchResponse struct {
	Data json.RawMessage `json:"data"`
}

const apiResource string = "/repos"

// Search request
func Search(params SearchParams) SearchResponse {
	path := apiResource + "/search/?q=" + params.PackageName
	body := Request(path)

	var response SearchResponse
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	return response
}
