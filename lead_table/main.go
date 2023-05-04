package main

import (
	lead "loan_application/lead_table"
	"net/http"
)

func main() {
	http.HandleFunc("/insert", lead.InsertLead)
	http.HandleFunc("/update", lead.UpdateLead)
	http.HandleFunc("/delete", lead.DeleteLead)
	http.HandleFunc("/", lead.LeadIndex)
	http.ListenAndServe(":8080", nil)
}
