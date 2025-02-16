package main

import (
	"fmt"
	binance_connector "github.com/binance/binance-connector-go"
	"time"
)

func main() {
	WsBookTickerExample()
}

func WsBookTickerExample() {
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false)
	wsBookTickerHandler := func(event *binance_connector.WsBookTickerEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, stopCh, err := websocketStreamClient.WsBookTickerServe("DOGEUSDT", wsBookTickerHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopCh to exit
	go func() {
		time.Sleep(10 * time.Second)
		stopCh <- struct{}{}
	}()
	<-doneCh
}
