package flyapi

type AddOnPlan struct {
	DisplayName              string `json:"displayName"`
	Id                       string `json:"id"`
	MaxCommandsPerSec        int    `json:"maxCommandsPerSec"`
	MaxConcurrentConnections int    `json:"maxConcurrentConnections"`
	MaxDailyBandwidth        string `json:"maxDailyBandwidth"`
	MaxDailyCommands         int    `json:"maxDailyCommands"`
	MaxDataSize              string `json:"maxDataSize"`
	MaxRequestSize           string `json:"maxRequestSize"`
	Name                     string `json:"name"`
	PricePerMonth            int    `json:"pricePerMonth"`
}
