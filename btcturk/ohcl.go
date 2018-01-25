package btcturk

import "fmt"

type Ohcl struct {
	DateTime              string  `json:"Date"`
	Open                  float64 `json:"Open"`
	High                  float64 `json:"High"`
	Low                   float64 `json:"Low"`
	Close                 float64 `json:"Close"`
	Volume                float64 `json:"Volume"`
	Average               float64 `json:"Average"`
	DailyChangeAmount     float64 `json:"DailyChangeAmount"`
	DailyChangePercentage float64 `json:"DailyChangePercentage"`
}

func (c *Client) OhclData() ([]Ohcl, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/ohlcdata?%s", c.params.Encode()), nil)

	var response []Ohcl
	if _, err = c.do(req, &response); err != nil {
		return []Ohcl{}, err
	}

	return response, nil
}
