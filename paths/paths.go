package paths

import (
	"fmt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"queries"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/hello.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}

func Migrate(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := queries.MigrateDB(DB)
		if err != nil {
			htmxFile, err := os.ReadFile("../static/components/error.html")
			if err != nil {
				fmt.Printf("error during db operation")
			}
			io.WriteString(w, string(htmxFile))
		} else {
			htmxFile, err := os.ReadFile("../static/components/insert.html")
			if err != nil {
				fmt.Printf("error reading file")

				io.WriteString(w, string(htmxFile))
			}
		}
	}
}

func InsertFailedLogins(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := queries.InsertFailedLogins(DB)
		if err != nil {
			htmxFile, err := os.ReadFile("../static/components/error.html")
			if err != nil {
				fmt.Printf("error reading file")
			}
			io.WriteString(w, string(htmxFile))
		} else {
			htmxFile, err := os.ReadFile("../static/components/insert.html")
			if err != nil {
				fmt.Printf("error reading file")

				io.WriteString(w, string(htmxFile))
			}
		}
	}
}


func InsertAdminLogins(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := queries.InsertAdminLogins(DB)
		if err != nil {
			htmxFile, err := os.ReadFile("../static/components/error.html")
			if err != nil {
				fmt.Printf("error reading file")
			}
			io.WriteString(w, string(htmxFile))
		} else {
			htmxFile, err := os.ReadFile("../static/components/insert.html")
			if err != nil {
				fmt.Printf("error reading file")

				io.WriteString(w, string(htmxFile))
			}
		}
	}
}
