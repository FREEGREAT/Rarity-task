package entity

type Creator struct {
	Account string `json:"account"`
	Value   int    `json:"value"`
}

type Ownership struct {
	ID            string    `json:"id"`
	Blockchain    string    `json:"blockchain"`
	ItemID        string    `json:"itemId"`
	Contract      string    `json:"contract"`
	Collection    string    `json:"collection"`
	TokenID       string    `json:"tokenId"`
	Owner         string    `json:"owner"`
	Value         string    `json:"value"`
	CreatedAt     string    `json:"createdAt"`
	LastUpdatedAt string    `json:"lastUpdatedAt"`
	Creators      []Creator `json:"creators"`
	LazyValue     string    `json:"lazyValue"`
	Pending       []string  `json:"pending"`
	OriginOrders  []string  `json:"originOrders"`
	Version       int       `json:"version"`
}
