package lead

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type lead_info struct {
	Created_at    time.Time
	Last_modified time.Time
	Status        string
}

func Conn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/loan_application")
	if err != nil {
		panic(err)
	}
	return db
}
func InsertLead(w http.ResponseWriter, r *http.Request) {
	db := Conn()
	result, err := db.Exec("INSERT INTO lead_table(created_at, last_modified, status) Values(NOW(),NOW(), 'Approved')")
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	fmt.Fprintf(w, "New record created with ID %d", id)
}

func UpdateLead(w http.ResponseWriter, r *http.Request) {
	db := Conn()
	//id := r.FormValue("id")
	stmt, err := db.Prepare("UPDATE lead_table SET last_modified = NOW() WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	fmt.Fprintf(w, "Lead with ID updated successfully")
}

func DeleteLead(w http.ResponseWriter, r *http.Request) {
	db := Conn()
	id := r.FormValue("id")
	stmt, err := db.Prepare("DELETE FROM lead_table WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Lead with ID %s deleted successfully", id)
}

func LeadIndex(w http.ResponseWriter, r *http.Request) {
	db := Conn()
	_, err := db.Query("SELECT * FROM lead_table")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Lead Table", r)
}
