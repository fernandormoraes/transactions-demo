package remote

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

type Exchange struct {
	Country      string
	Currency     string
	ExchangeRate string
	RecordDate   string
}

type TreasuryRemote interface {
	FindAll() ([]Exchange, error)
}

type HttpTreasureRemote struct {
	httpClient *http.Client
}

func NewTreasuryRemote(httpClient *http.Client) *HttpTreasureRemote {
	return &HttpTreasureRemote{httpClient: httpClient}
}

func (r HttpTreasureRemote) FindAll() ([]Exchange, error) {
	var listExchanges []Exchange

	url := viper.GetString("BASE_URL_TREASURY") + viper.GetString("RATES_EXCHANGE_ENDPOINT")

	resp, err := r.httpClient.Get(url)

	if err != nil {
		return make([]Exchange, 0), err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&listExchanges)

	if err != nil {
		return make([]Exchange, 0), err
	}

	return listExchanges, nil
}
