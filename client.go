package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/smartinrub/accounts"
)

func create(w http.ResponseWriter, r *http.Request) {
	account := accounts.Account{
		Data: accounts.Data{
			Type:           "accounts",
			Id:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			OrganizationId: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Attributes: accounts.Attributes{
				Name:         []string{"sergiobank"},
				Country:      "GB",
				BaseCurrency: "GBP",
				BankId:       "400300",
				BankIdCode:   "GBDSC",
				Bic:          "NWBKGB22",
			},
		},
	}

	accountResponse, err := accounts.Create(account)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, accountResponse.Data.Id)
	fmt.Println(accountResponse.Data.Id)
}

func fetch(w http.ResponseWriter, r *http.Request) {
	account, err := accounts.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, account.Data.Id)
	fmt.Println(account.Data.Id)
}

func delete(w http.ResponseWriter, r *http.Request) {
	err := accounts.Delete("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/fetch", fetch)
	http.HandleFunc("/delete", delete)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
