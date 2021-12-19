package inmemory

type responseGet struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type responsePost struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
