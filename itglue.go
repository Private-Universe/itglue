package itglue

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//TBD: Implement
type Links struct{}

type Metadata struct {
	CurrentPage int         `json:"current-page"`
	NextPage    interface{} `json:"next-page"`
	PrevPage    interface{} `json:"prev-page"`
	TotalPages  int         `json:"total-pages"`
	TotalCount  int         `json:"total-count"`
	Filters     struct {
	} `json:"filters"`
}

//ITGAPI contains the ITG API URL for North America, as well as the API key.
//Think of it as an instance of the API client
type ITGAPI struct {
	Site   string //Full URL ITG API
	APIKey string //API Key
}

type Request struct {
	ITG          *ITGAPI
	RestAction   string
	Method       string     //GET, POST, DELETE, etc
	Body         []byte     //In a GET request, this is an instance of the struct that the expected json data is to be unmarshaled into
	URLValues    url.Values //Parameters sent to ITG for filtering by conditions that utilize strings
	RawURLValues string     //URL values must be a string.  Raw URL values allow an int to be used, for instance in a filter
	Page         int
	PageSize     int
}

//NewITGAPI expects the API key to be passed to it
//Returns a pointer to an ITGAPI struct
func NewITGAPI(apiKey string) *ITGAPI {
	return &ITGAPI{Site: "https://api.itglue.com", APIKey: apiKey}
}

//NewRequest is a function which takes the mandatory fields to perform a request to the CW API and returns a pointer to a Request struct
func (itg *ITGAPI) NewRequest(restAction, method string, body []byte) *Request {
	req := Request{ITG: itg, RestAction: restAction, Method: method, Body: body}
	req.URLValues = url.Values{}
	return &req
}

//Do is a method of the Request struct which uses the data contained within the Request instance to perform an HTTP request to ConnectWise
func (req *Request) Do() error {
	itgurl, err := req.ITG.BuildURL(req.RestAction)
	if err != nil {
		return fmt.Errorf("could not build url %s: %s", req.RestAction, err)
	}

	//If pagination parameters are not being specified, then don't include them in the URL - not all endpoints will accept page and pagesize, they will 400 - bad request
	if req.Page == 0 && req.PageSize == 0 {
		itgurl.RawQuery = fmt.Sprintf("%s%s", req.URLValues.Encode(), req.RawURLValues)
	} else if req.Page != 0 && req.PageSize == 0 {
		itgurl.RawQuery = fmt.Sprintf("%s%s&page[number]=%d", req.URLValues.Encode(), req.RawURLValues, req.Page)
	} else if req.Page == 0 && req.PageSize != 0 {
		itgurl.RawQuery = fmt.Sprintf("%s%s&page[size]=%d", req.URLValues.Encode(), req.RawURLValues, req.PageSize)
	} else {
		itgurl.RawQuery = fmt.Sprintf("%s%s&page[number]=%d&page[size]=%d", req.URLValues.Encode(), req.RawURLValues, req.Page, req.PageSize)
	}

	client := &http.Client{}
	jsonBuffer := bytes.NewReader(req.Body)
	httpreq, err := http.NewRequest(req.Method, itgurl.String(), jsonBuffer)
	if err != nil {
		return fmt.Errorf("could not construct http request: %s", err)
	}
	httpreq.Header.Set("Content-Type", "application/vnd.api+json")
	httpreq.Header.Set("x-api-key", req.ITG.APIKey)
	httpreq.Header.Set("cache-control", "no-cache")

	resp, err := client.Do(httpreq)
	if err != nil {
		return fmt.Errorf("could not perform http %s request: %s", req.Method, err)
	}
	req.Body, err = getHTTPResponseBody(resp)
	if err != nil {
		return fmt.Errorf("failed to get http response body: %s", err)
	}

	return nil
}

func getHTTPResponseBody(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ITG returned HTTP status code %s\n%s", resp.Status, resp.Body)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read http response body: %s", err)
	}

	return body, nil
}

//BuildURL expects a restaction to be passed to it
//Returns the full request URL containing the ITG API domain prepended to the rest action
func (itg *ITGAPI) BuildURL(restAction string) (*url.URL, error) {
	var itgurl *url.URL
	itgurl, err := url.Parse(itg.Site)
	if err != nil {
		return nil, fmt.Errorf("could not parse url %s: %s", itg.Site, err)
	}
	itgurl.Path += restAction
	return itgurl, nil
}
