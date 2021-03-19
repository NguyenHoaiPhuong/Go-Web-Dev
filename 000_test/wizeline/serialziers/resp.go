package serializers

// Resp : struct
type Resp struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}

// URLResp : struct
type URLResp struct {
	URL string `json:"url"`
}
