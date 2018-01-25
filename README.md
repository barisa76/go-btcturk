# go-btcturk

golang client for btcturk api

## documentation

the documentation is available on [godoc](http://godoc.org/github.com/vural/go-btcturk/btcturk)

## install

```sh
go get -u github.com/vural/go-btcturk/btcturk
```

## usage
```go
package main

import (
	"github.com/vural/go-btcturk/btcturk"
)

func main() {
    api := btcturk.NewBTCTurkClient()
    t, err := api.Ticker()
    if err != nil {
        print(err)
        os.Exit(1)
    }
    
    for _, v := range t {
        println(v.Ask)
    }

    // if you don't plan to call authenticated api methods. SetAuthKey not required.
    api.SetAuthKey("publicKey", "privateKey")
}

```

## Passing params

[Endpoint Params](https://github.com/vural/go-btcturk/blob/master/btcturk/params.go)
```go
package main

import (
	"github.com/vural/btcturk/btcturk"
)

func main() {
    api := btcturk.NewBTCTurkClient()
    api.SetAuthKey("publicKey", "privateKey")
    
    api.Price(3500).
        Amount(10).
        PairSymbol(btcturk.ETHTRY).
        PricePrecision(00).
        TotalPrecision(00).
        Buy()
}

```

## Notes
 - you can get your private/public key peer from your account

**[BTCTurk API documentation](https://github.com/BTCTrader/broker-api-docs/blob/master/README.md)**
