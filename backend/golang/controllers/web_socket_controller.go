package controllers

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"net/url"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
	"encoding/json"
	"github.com/trade-lab-app/backend/golang/models"
)

var addr = flag.String("addr", "api.hitbtc.com", "http service address")
var CurrencyData = make([]models.Currency, 2)

func ClientWebSocket() {
	fmt.Println("Websocket Client Running !!!!")
	tempCurrency := new(models.Currency)
	//currency := []models.Currency{}


	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/api/2/ws"}
	log.Printf("connecting to %s", u.String())
	headerData1 := []byte(`{
  	"method": "getSymbol",
  	"params": {
   	 	"symbol": "ETHBTC"
  		},
  	"id": 123
	}`)

	headerData2 := []byte(`{
 	 "method": "getSymbol",
 	 "params": {
		"symbol": "BTCUSD"
		},
	 "id": 123
	}`)

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			//log.Printf("recv: %s", message)

			if err = json.Unmarshal(message, &tempCurrency); err!=nil{
				fmt.Println("Error : ",err)
			}

			inMemoryStorage(tempCurrency, CurrencyData)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, headerData1)
			if err != nil {
				log.Println("write:", err, t)
				return
			}

		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, headerData2)
			if err != nil {
				log.Println("write:", err, t)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func inMemoryStorage(tempData *models.Currency, currency []models.Currency) {
	if tempData.Result.Id == "ETHBTC" {
		currency[0] = *tempData
	} else {
		currency[1] = *tempData
	}
}
