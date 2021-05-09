package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type cowinResponse struct {
	Centers []Center `json:"centers"`
}

type Center struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	FeeType  string `json:"fee_type"`
	Sessions []Session
}

type Session struct {
	AvaliableCapacity int `json:"available_capacity"`
	MinAgeLimit       int `json:"min_age_limit"`
}

func CheckCenters() {
	fmt.Println("CHECKING CENTERS")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?district_id=363&date=11-05-2021", nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	cowinResp := cowinResponse{}
	err = json.Unmarshal(responseData, &cowinResp)
	if err != nil {
		log.Fatal(err)
	}
	for _, center := range cowinResp.Centers {
		for _, session := range center.Sessions {
			if session.AvaliableCapacity > 0 && session.MinAgeLimit == 18 {
				fmt.Println(fmt.Sprintf("FOUND AVAILABLITY IN HOSPITAL: %s AND ADDRESS: %s", center.Name, center.Address))
				//Notify(fmt.Sprintf("FOUND AVAILABLITY IN HOSPITAL: %s AND ADDRESS: %s", center.Name, center.Address))
			}
		}
	}
}
