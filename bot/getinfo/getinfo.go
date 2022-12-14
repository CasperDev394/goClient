package getinfo

import (
	"encoding/json"
	"fmt"
	"log"

	"io/ioutil"
	"net/http"

	"github.com/CasperDev394/goClient/getinfo/types"
)

const BaseURL = "https://api.coingecko.com/api/v3"

type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{httpClient: httpClient}
}

func (c *Client) newRequest(method, url string, body interface{}) ([]byte, error) {
	URL := fmt.Sprintf("%s%s", BaseURL, url)

	req, err := http.NewRequest(method, URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req, c.httpClient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) do(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (c *Client) Ping() (*types.Ping, error) {
	resp, err := c.newRequest("GET", "/ping", nil)
	if err != nil {
		return nil, err
	}
	var data *types.Ping
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) SimpleSupportedVSCurrencies() (*types.SimpleSupportedVSCurrencies, error) {
	resp, err := c.newRequest("GET", "/simple/supported_vs_currencies", nil)
	if err != nil {
		return nil, err
	}
	var data *types.SimpleSupportedVSCurrencies
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CoinList() (*types.CoinList, error) {
	resp, err := c.newRequest("GET", "/coins/list", nil)
	if err != nil {
		return nil, err
	}
	var data *types.CoinList
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CoinHistory(id string, date string) (*types.CoinsIDHistory, error) {
	log.Println("C date - ", date)
	url := "/coins/" + id + "/history?date=" + date
	resp, err := c.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var data *types.CoinsIDHistory
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CoinOhlc(id string, days string, VsCurrency string) (*types.CoinOhlc, error) {
	log.Println("C date - ", days)
	url := "/coins/" + id + "/ohlc?days=" + days + "&vs_currency=" + VsCurrency
	resp, err := c.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var data *types.CoinOhlc
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
