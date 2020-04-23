package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/igorjm/golang-finance/util"

	model "github.com/igorjm/golang-finance/model/transaction"
)

//GetTransactions func
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = model.Transactions{
		model.Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2000-20-20T15:00:00"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)

}

//CreateTransaction func
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = model.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)

}
