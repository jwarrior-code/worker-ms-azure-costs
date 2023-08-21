package azure_pricing

type AzurePriceItem struct {
	CurrencyCode       string  `json:"currencyCode"`
	TierMinimumUnits   float64 `json:"tierMinimumUnits"`
	RetailPrice        float64 `json:"retailPrice"`
	ArmRegionName      string  `json:"armRegionName"`
	Location           string  `json:"location"`
	EffectiveStartDate string  `json:"effectiveStartDate"`
	ProductName        string  `json:"productName"`
	SkuName            string  `json:"skuName"`
	SkuId              string  `json:"SkuId"`
	ServiceName        string  `json:"serviceName"`
	ArmSkuName         string  `json:"armSkuName"`
}
