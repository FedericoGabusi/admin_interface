package queries

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io"
	"models"
	"net/http"
	"net/url"
	"time"
)

func qradarQuery(query string) (json_res []byte) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	qradar_addr := "https://10.1.1.24/api/"
	qradar_token := "d639f2f6-9c55-4cb3-b95b-c10eee3ad34a"
	req, err := http.NewRequest("POST", qradar_addr+"ariel/searches?query_expression="+query, nil)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	req.Header.Set("SEC", qradar_token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	var QueryResult struct {
		Status    string
		Search_Id string
	}
	json.NewDecoder(resp.Body).Decode(&QueryResult)
	for QueryResult.Status != "COMPLETED" {
		fmt.Printf("STATUS: %v\n", QueryResult.Status)
		fmt.Printf("SEARCH ID: %v\n", QueryResult.Search_Id)
		req, err := http.NewRequest("GET", qradar_addr+"ariel/searches/"+QueryResult.Search_Id, nil)
		if err != nil {
			fmt.Printf("err %v", err)
		}
		req.Header.Set("SEC", qradar_token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("err %v", err)
		}
		if QueryResult.Status == "ERROR" {
			fmt.Printf("error with qradar query")
			return nil
		}
		json.NewDecoder(resp.Body).Decode(&QueryResult)

		time.Sleep(10 * time.Second)
	}
	result_url := qradar_addr + "ariel/searches/" + QueryResult.Search_Id + "/results"

	fmt.Printf("result url: %s", result_url)

	req, err = http.NewRequest("GET", result_url, nil)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	req.Header.Set("SEC", qradar_token)
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	json_res = body
	return json_res
}

func InsertFailedLogins(db *gorm.DB) (err error) {
	query := `select QIDNAME(qid) as 'Event Name',"eventCount" as 'event_count',"startTime" as 'start_time',"sourceIP" as 'source_ip',"destinationIP" as 'destination_ip',"userName" as 'username' from events where ( "creEventList"='100063' AND "userName" != 'null' ) order by "startTime" desc LIMIT 10000 last 2592000 minutes`
	json_res := qradarQuery(url.QueryEscape(query))

	var data models.Data
	fmt.Printf("%v\n", string(json_res))
	err = json.Unmarshal(json_res, &data)

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", data)

	for _, event := range data.Events {
		currentEvent := models.FailedLoginModel{
			Username:      event.Username,
			EventCount:    event.EventCount,
			StartTime:     event.StartTime,
			SourceIP:      event.SourceIP,
			EndTime:       event.EndTime,
			DestinationIP: event.DestinationIP,
			
		}

		/*result := db.Debug().Exec("INSERT INTO failed_logins (username, start_time, end_time, event_count, source_ip, destination_ip) VALUES (?,?,?,?,?,?)",
			currentEvent.Username,
			currentEvent.StartTime,
			currentEvent.EndTime,
			currentEvent.EventCount,
			currentEvent.SourceIP,
			currentEvent.DestinationIP)*/
		result := db.Create(&currentEvent)
		if result.Error != nil {
			var error error
			error = result.Error
			return error
		}
		fmt.Printf("result: %v\n", currentEvent.ID)
		fmt.Printf("errors: %v\n", result.Error)
		fmt.Printf("rows affected %v\n", result.RowsAffected)
	}
	return
}

func InsertAdminLogins(db *gorm.DB) (err error) {
	query := `SELECT "eventCount" as 'event_count',"startTime" as 'start_time',"sourceIP" as 'source_ip',"destinationIP" as 'destination_ip',"userName" as 'username' from events where ( ( ( category='3014' AND qid='5000858' ) AND "isUnparsed"='false' ) AND ("userName"='Admin') or ("userName"='Administrator') or ("userName"='Root') or ("userName"='admin') or ("userName"='administrator') or ("userName"='root') ) GROUP BY "userName" order by "userName" desc LIMIT 10012 last 3 days`
	json_res := qradarQuery(url.QueryEscape(query))

	var data models.Data
	fmt.Printf("%v\n", string(json_res))
	err = json.Unmarshal(json_res, &data)

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", data)

	for _, event := range data.Events {
		currentEvent := models.AdminLoginModel{
			Username:      event.Username,
			EventCount:    event.EventCount,
			StartTime:     event.StartTime,
			SourceIP:      event.SourceIP,
			EndTime:       event.EndTime,
			DestinationIP: event.DestinationIP,
			
		}

		/*result := db.Debug().Exec("INSERT INTO failed_logins (username, start_time, end_time, event_count, source_ip, destination_ip) VALUES (?,?,?,?,?,?)",
			currentEvent.Username,
			currentEvent.StartTime,
			currentEvent.EndTime,
			currentEvent.EventCount,
			currentEvent.SourceIP,
			currentEvent.DestinationIP)*/
		result := db.Create(&currentEvent)
		if result.Error != nil {
			var error error
			error = result.Error
			return error
		}
		fmt.Printf("result: %v\n", currentEvent.ID)
		fmt.Printf("errors: %v\n", result.Error)
		fmt.Printf("rows affected %v\n", result.RowsAffected)
	}
	return
}

func MigrateDB(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		&models.FailedLoginModel{},
		&models.AdminLoginModel{},
		&models.BlockedAccountsModel{},
		&models.FirewallDenyModel{},
		&models.NetworkTrafficModel{},
	)
	if err != nil {
		fmt.Printf("error during db migration. %v\n", err)
		return err
	}
	fmt.Printf("Migration executed correctly!\n")
	return nil
}
