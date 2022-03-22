package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var timezone string

func init() {
	t := os.Getenv("TIMEZONE")
	timezone = t
}

/**

  "abbreviation": "+0545",
  "client_ip": "103.163.182.164",
  "datetime": "2021-09-26T19:11:19.968921+05:45",
  "day_of_week": 0,
  "day_of_year": 269,
  "dst": false,
  "dst_from": null,
  "dst_offset": 0,
  "dst_until": null,
  "raw_offset": 20700,
  "timezone": "Asia/Kathmandu",
  "unixtime": 1632662779,
  "utc_datetime": "2021-09-26T13:26:19.968921+00:00",
  "utc_offset": "+05:45",
  "week_number": 38
*/

type Response struct {
	Time     string `json:"time"`
	Err      string `json:"error"`
	ClientIp string `json:"client_ip"`
	Timezone string `json:"timezone"`
}

type TimeResponse struct {
	Datetime string `json:"datetime"`
	Timezone string `json:"timezone"`
	ClientIp string `json:"client_ip"`
}

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		reqUri := fmt.Sprintf("http://worldtimeapi.org/api/timezone/%s", timezone)
		resp, err := http.Get(reqUri)

		w.Header().Add("Content-Type", "application/json")

		if err != nil {
			enc := json.NewEncoder(w)
			enc.Encode(Response{
				Err: err.Error(),
			})
			return
		}

		jsonResponse := TimeResponse{}

		jsonBody, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			enc := json.NewEncoder(w)
			enc.Encode(Response{
				Err: err.Error(),
			})
			return
		}

		err = json.Unmarshal(jsonBody, &jsonResponse)

		if err != nil {
			enc := json.NewEncoder(w)
			enc.Encode(Response{
				Err: err.Error(),
			})
			return
		}

		enc := json.NewEncoder(w)
		enc.Encode(Response{
			Time:     jsonResponse.Datetime,
			Timezone: jsonResponse.Timezone,
			ClientIp: jsonResponse.ClientIp,
		})
		return
	})

	err := http.ListenAndServe(":4000", server)

	if err != nil {
		panic(err)
	}
}
