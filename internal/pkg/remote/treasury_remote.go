package remote

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fernandormoraes/transaction-demo/pkg/helpers"
	"github.com/spf13/viper"
)

type Exchange struct {
	Country      string `json:"country"`
	Currency     string `json:"currency"`
	ExchangeRate string `json:"exchange_rate"`
	RecordDate   string `json:"record_date"`
}

type TreasuryRemote interface {
	FindAll(transactionDate string, currencyDesc string) ([]Exchange, error)
}

type HttpTreasureRemote struct {
	httpClient *http.Client
}

func NewTreasuryRemote(httpClient *http.Client) *HttpTreasureRemote {
	return &HttpTreasureRemote{httpClient: httpClient}
}

func (r HttpTreasureRemote) FindAll(transactionDate string, currencyDesc string) ([]Exchange, error) {
	url := viper.GetString("BASE_URL_TREASURY") + viper.GetString("RATES_EXCHANGE_ENDPOINT")

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return make([]Exchange, 0), err
	}

	dateTime, err := time.Parse("2006-01-02", transactionDate)

	if err != nil {
		return make([]Exchange, 0), err
	}

	transactionDateMinusSixMonths := dateTime.Add(-(4380 * time.Hour)).Format(helpers.LayoutDate)

	query := req.URL.Query()
	query.Add("fields", "country_currency_desc,exchange_rate,record_date")
	query.Add("filter", "country_currency_desc:in:("+currencyDesc+"),record_date:gte:"+transactionDateMinusSixMonths)
	query.Add("sort", "record_date")
	req.URL.RawQuery = query.Encode()

	resp, err := r.httpClient.Do(req)

	if err != nil {
		return make([]Exchange, 0), err
	}

	defer resp.Body.Close()

	responseParser := helpers.DefaultRemoteResponse[Exchange]{}

	err = json.NewDecoder(resp.Body).Decode(&responseParser)

	if err != nil {
		return make([]Exchange, 0), err
	}

	return responseParser.Data, nil
}
