package response

type UserExtDashboardResponse struct {
	Base  UserExtDashboardResponseBase  `json:"base"`
	Using UserExtDashboardResponseUsing `json:"using"`
}

type UserExtDashboardResponseBase struct {
	ValidStart   string `json:"validStart"`
	ValidEnd     string `json:"validEnd"`
	StandardData string `json:"standardData"`
	QRCode       string `json:"QRCode"`
	SubUrl       string `json:"subUrl"`
	Enable bool `json:"enable"`
}
type UserExtDashboardResponseUsing struct {
	Start        string `json:"start"`
	End          string `json:"end"`
	StandardData string `json:"standardData"`
	Remain       string `json:"remain"`
	Usage        string `json:"usage"`
	Percentage   int    `json:"percentage"`
}
