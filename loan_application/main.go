package main

import (
	"fmt"
	lead "loan_application/lead_table"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Inserted succesfully")
	http.HandleFunc("/", lead.InsertLead)
	http.HandleFunc("/update", lead.UpdateLead)
	http.ListenAndServe(":8080", nil)
}
