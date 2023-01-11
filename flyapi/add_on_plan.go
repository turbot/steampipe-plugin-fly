package flyapi

type AddOnPlan struct {
	Name                     string `json:"name"`
	Id                       string `json:"id"`
	DisplayName              string `json:"displayName"`
	MaxCommandsPerSec        int    `json:"maxCommandsPerSec"`
	MaxConcurrentConnections int    `json:"maxConcurrentConnections"`
	MaxDailyCommands         int    `json:"maxDailyCommands"`
	MaxDailyBandwidth        string `json:"maxDailyBandwidth"`
	MaxDataSize              string `json:"maxDataSize"`
	MaxRequestSize           string `json:"maxRequestSize"`
	PricePerMonth            int    `json:"pricePerMonth"`
}
