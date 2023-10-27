package main
import (
	"gorm.io/gorm"
	"fmt"
	"db_conn"
	"paths"
	"net/http"
)

var DB *gorm.DB
const dsn string = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Rome"
const SRV_PORT string = ":7777"
func main() {
	//open server

	var error error
	DB, error = db_conn.DBConnect("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")
	if error != nil {
		fmt.Printf("problem connecting to db %v", error)
	}
	fmt.Printf("connected to db %v", &DB)

	http.Handle("/", http.FileServer(http.Dir("../static")))
	//http.HandleFunc("/clock", paths.GetClock)
	http.HandleFunc("/failed_logins", paths.InsertFailedLogins(DB))
	http.HandleFunc("/admin_logins", paths.InsertAdminLogins(DB))
	//http.HandleFunc("/network_traffic", paths.InsertNetworkTraffic(DB))

	fmt.Printf("server open on %s", SRV_PORT)

	err := http.ListenAndServe(SRV_PORT, nil)
	if err != nil {
		fmt.Printf("cant open server %v", err)
	}
	//DB CONNECTION
	fmt.Printf("server open on %s", SRV_PORT)
}
