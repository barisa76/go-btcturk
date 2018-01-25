package btcturk

import "fmt"

type Trade struct {
	TimeStamp float64 `json:"date"`
	TID       string  `json:"tid"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
}

func (c *Client) Trades() ([]Trade, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/trades?%s", c.params.Encode()), nil)

	var response []Trade
	if _, err = c.do(req, &response); err != nil {
		return []Trade{}, err
	}

	return response, nil
}
