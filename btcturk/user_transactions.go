package btcturk

import (
	"fmt"
)

type UserTransactions struct {
	ID        string  `json:"id"`
	Date      string  `json:"date"`
	Operation string  `json:"operation"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Price     float64 `json:"price"`
	Funds     float64 `json:"funds"`
	Fee       float64 `json:"fee"`
	Tax       float64 `json:"tax"`
}

func (c *Client) UserTransactions() ([]UserTransactions, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/userTransactions?%s", c.params.Encode()), nil)
	if err := c.auth(req); err != nil {
		return []UserTransactions{}, err
	}

	var response []UserTransactions
	if _, err = c.do(req, &response); err != nil {
		return []UserTransactions{}, err
	}

	return response, nil
}
