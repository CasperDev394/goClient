package types

type Ping struct {
	GeckoSays string `json:"gecko_says"`
}

type SimpleSinglePrice struct {
	ID          string
	Currency    string
	MarketPrice float32
}

type SimpleSupportedVSCurrencies []string

type CoinList []CoinsListItem

type CoinsMarket []CoinsMarketItem

type CoinsID struct {
	coinBaseStruct
	BlockTimeInMin      int32               `json:"block_time_in_minutes"`
	Categories          []string            `json:"categories"`
	Localization        LocalizationItem    `json:"localization"`
	Description         DescriptionItem     `json:"description"`
	Links               *LinksItem          `json:"links"`
	Image               ImageItem           `json:"image"`
	CountryOrigin       string              `json:"country_origin"`
	GenesisDate         string              `json:"genesis_date"`
	MarketCapRank       uint16              `json:"market_cap_rank"`
	CoinGeckoRank       uint16              `json:"coingecko_rank"`
	CoinGeckoScore      float32             `json:"coingecko_score"`
	DeveloperScore      float32             `json:"developer_score"`
	CommunityScore      float32             `json:"community_score"`
	LiquidityScore      float32             `json:"liquidity_score"`
	PublicInterestScore float32             `json:"public_interest_score"`
	MarketData          *MarketDataItem     `json:"market_data"`
	CommunityData       *CommunityDataItem  `json:"community_data"`
	DeveloperData       *DeveloperDataItem  `json:"developer_data"`
	PublicInterestStats *PublicInterestItem `json:"public_interest_stats"`
	StatusUpdates       *[]StatusUpdateItem `json:"status_updates"`
	LastUpdated         string              `json:"last_updated"`
	Tickers             *[]TickerItem       `json:"tickers"`
}

type CoinsIDTickers struct {
	Name    string       `json:"name"`
	Tickers []TickerItem `json:"tickers"`
}

type CoinsIDHistory struct {
	coinBaseStruct
	Localization   LocalizationItem    `json:"localization"`
	Image          ImageItem           `json:"image"`
	MarketData     *MarketDataItem     `json:"market_data"`
	CommunityData  *CommunityDataItem  `json:"community_data"`
	DeveloperData  *DeveloperDataItem  `json:"developer_data"`
	PublicInterest *PublicInterestItem `json:"public_interest_stats"`
}

type CoinsIDMarketChart struct {
	coinBaseStruct
	Prices       *[]ChartItem `json:"prices"`
	MarketCaps   *[]ChartItem `json:"market_caps"`
	TotalVolumes *[]ChartItem `json:"total_volumes"`
}

type CoinOhlc [][]float64

type EventsCountries struct {
	Data []EventCountryItem `json:"data"`
}

type EventsTypes struct {
	Data  []string `json:"data"`
	Count uint16   `json:"count"`
}

type ExchangeRatesResponse struct {
	Rates ExchangeRatesItem `json:"rates"`
}

type GlobalResponse struct {
	Data Global `json:"data"`
}
