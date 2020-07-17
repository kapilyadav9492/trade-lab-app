package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/trade-lab-app/backend/golang/models"
	"log"
)

//Currency Controller

func (c *Controller) Currency(w http.ResponseWriter, r *http.Request) {
	var requestData models.CurrencySymbol

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Fatal(err)
		return
	}

	if CurrencyData[0].Result.Id == requestData.Symbol {
		resp, err := json.Marshal(CurrencyData[0])
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

	if CurrencyData[1].Result.Id == requestData.Symbol {
		resp, err := json.Marshal(CurrencyData[1])
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}
	return
}

func (c *Controller) GetAllCurrency(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(CurrencyData)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	return
}