package btcturk

type Balance struct {
	TRYBalance               float64 `json:"try_balance"`
	BTCBalance               float64 `json:"btc_balance"`
	ETHBalance               float64 `json:"eth_balance"`
	TRYReserved              float64 `json:"try_reserved"`
	BTCReversed              float64 `json:"btc_reversed"`
	ETHReversed              float64 `json:"eth_reversed"`
	TRYAvailable             float64 `json:"try_available"`
	BTCAvailable             float64 `json:"btc_available"`
	ETHAvailable             float64 `json:"eth_available"`
	BTCTRYTakerFeePercentage float64 `json:"btctry_taker_fee_percentage"`
	BTCTRYMakerFeePercentage float64 `json:"btctry_maker_fee_percentage"`
	ETHTRYTakerFeePercentage float64 `json:"ethtry_taker_fee_percentage"`
	ETHTRYMakerFeePercentage float64 `json:"ethtry_maker_fee_percentage"`
	ETHBTCTakerFeePercentage float64 `json:"ethbtc_taker_fee_percentage"`
	ETHBTCMakerFeePercentage float64 `json:"ethbtc_maker_fee_percentage"`
}

func (c *Client) Balance() (Balance, error) {
	req, err := c.newRequest("GET", "/api/balance", c.body)
	if err := c.auth(req); err != nil {
		return Balance{}, err
	}

	var response Balance
	if _, err = c.do(req, &response); err != nil {
		return Balance{}, err
	}

	return response, nil
}
