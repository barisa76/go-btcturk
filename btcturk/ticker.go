package btcturk

type Ticker struct {
	Pair      string  `json:"pair"`
	High      float64 `json:"high"`
	Last      float64 `json:"last"`
	Timestamp float32 `json:"timestamp"`
	Bid       float64 `json:"bid"`
	Volume    float64 `json:"volume"`
	Low       float64 `json:"low"`
	Ask       float64 `json:"ask"`
	Open      float64 `json:"open"`
	Average   float64 `json:"average"`
}

func (c *Client) Ticker() ([]Ticker, error) {
	req, err := c.newRequest("GET", "/api/ticker", nil)

	var response []Ticker
	if _, err = c.do(req, &response); err != nil {
		return []Ticker{}, err
	}

	return response, nil
}
