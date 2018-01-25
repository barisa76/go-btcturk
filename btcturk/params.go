package btcturk

import "strconv"

// Transactions Limit
func (c *Client) Limit(v int) *Client {
	return addParamsAndReturnClient(c, "limit", v)
}

// Transactions Offset
func (c *Client) Offset(v int) *Client {
	return addParamsAndReturnClient(c, "offest", v)
}

// Cancel Order ID
func (c *Client) OrderID(v int) *Client {
	return addParamsAndReturnClient(c, "id", v)
}

// Transactions Sort
func (c *Client) Sort(v string) *Client {
	c.params.Add("sort", v)
	return c
}

// Buy or Sell OrderMethod
func (c *Client) OrderMethod(orderMethod int) *Client {
	c.params.Add("OrderMethod", strconv.Itoa(orderMethod))
	return c
}

// Buy or Sell PricePrecision
func (c *Client) PricePrecision(v int) *Client {
	c.params.Add("PricePrecision", strconv.Itoa(v))
	return c
}

// Buy or Sell Amount
func (c *Client) Amount(v int) *Client {
	c.params.Add("Amount", strconv.Itoa(v))
	return c
}

// Buy or Sell AmountPrecision
func (c *Client) AmountPrecision(v int) *Client {
	c.params.Add("AmountPrecision", strconv.Itoa(v))
	return c
}

// Buy or Sell Total
func (c *Client) Total(v int) *Client {
	c.params.Add("Total", strconv.Itoa(v))
	return c
}

// Buy or Sell TotalPrecision
func (c *Client) TotalPrecision(v int) *Client {
	c.params.Add("TotalPrecision", strconv.Itoa(v))
	return c
}

// Buy or Sell TriggerPrice
func (c *Client) TriggerPrice(v int) *Client {
	c.params.Add("TriggerPrice", strconv.Itoa(v))
	return c
}

// Buy or Sell TriggerPricePrecision
func (c *Client) TriggerPricePrecision(v int) *Client {
	c.params.Add("TriggerPricePrecision", strconv.Itoa(v))
	return c
}

// (Buy or Sell), Open Orders, Trades, Order Book must be PairSymbol Params
func (c *Client) PairSymbol(v string) *Client {
	c.params.Add("PairSymbol", v)
	return c
}

// Buy or Sell Price
func (c *Client) Price(v int) *Client {
	c.params.Add("Price", strconv.Itoa(v))
	return c
}

// Trades, OHCL Data Params (Not required)
func (c *Client) Last(v int) *Client {
	c.params.Add("last", strconv.Itoa(v))
	return c
}

func addParamsAndReturnClient(c *Client, key string, value int) *Client {
	c.params.Add(key, strconv.Itoa(value))
	return c
}
