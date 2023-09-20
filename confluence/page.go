package confluence

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"smart_confluence_label/constant"
	"smart_confluence_label/lib"
)

func GetPageContentByID(pageID string) PageContent {
	// confluenceEndpoint := os.Getenv("CONFLUENCE_ENDPOINT")
	confluenceEndpoint := ""
	url := confluenceEndpoint + fmt.Sprintf(constant.GetPageContenAPI, pageID) + "?body-format=view"

	req := getClient(url, "GET", nil)

	res, err := http.DefaultClient.Do(req)
	lib.HandleErr(err)
	defer res.Body.Close()

	var PageContent PageContent
	err = json.NewDecoder(res.Body).Decode(&PageContent)
	if err != nil {
		panic(err)
	}

	return PageContent
}

func UpdateTagsToPage(pageID string, tagsString string) *http.Response {
	// confluenceEndpoint := os.Getenv("CONFLUENCE_ENDPOINT")
	confluenceEndpoint := ""
	url := confluenceEndpoint + fmt.Sprintf(constant.PostPageLabelAPI, pageID)

	input := Label{
		Prefix: "global",
		Name:   tagsString,
	}

	jsonData, err := json.Marshal(input)
	lib.HandleErr(err)

	req := getClient(url, "POST", bytes.NewBuffer(jsonData))

	res, err := http.DefaultClient.Do(req)
	lib.HandleErr(err)

	return res
}

func getClient(url string, method string, body io.Reader) *http.Request {
	// username := os.Getenv("CONFLUENCE_USERNAME")
	// password := os.Getenv("CONFLUENCE_PASSWORD")
	username := ""
	password := ""

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	return req
}
