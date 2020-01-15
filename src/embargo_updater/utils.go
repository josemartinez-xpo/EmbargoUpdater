package main

import (
	"net/http"
	"net/url"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"os"
	"encoding/csv"
	"errors"
)

func RequestToken(username string, password string) string {
	apiUrl := "https://api.ltl.xpo.com"
	path := "/token"

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Add("username", username)
	// data.Add("password", "mDweb18test")
	data.Add("password", password)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = path
	urlStr := fmt.Sprintf("%v", u)
	fmt.Sprintf("request: %v\n", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", "Bearer eXlSazFmMzJRZnlNc3ZvR01LY0FaRG82VENNYTo2cVFXdFVzc0ZxUVRPa040S0lkQTZXX19pQjRh")

	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	respStr := string(body)

	// fmt.Println(resp.Status)
	// fmt.Println(respStr)

	return GetTokenFromResponse(respStr)	
}

// API response for embargo, local example taken from running PR's prototyper
func RequestEmbargo(token string) string {
	apiUrl := "https://api.ltl.xpo.com"

	path := "/freightflow/1.0/embargo-locations"

	// apiUrl := "http://127.0.0.1:9875"

	// path := "/mock.json"

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = path
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("GET", urlStr, nil)
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	respStr := string(body)

	// fmt.Println(resp.Status)
	// fmt.Println(respStr)
	return respStr
}

// Map the response to our complex response structs
// If all goes well, returns the array of embargo items
func ParseEmbargoResponse(jsonResponse string) ([]EmbargoItem, error){
	var response APIResponse
	json.Unmarshal([]byte(jsonResponse), &response)
	// fmt.Printf("Response code: %s, Item1:  %v\n", response.Code, response.Data.EmbargoLocation[0])
	if response.Code == "" || len(response.Data.EmbargoLocation) == 0 {
		return nil, errors.New("No data was loaded")
	}
	return response.Data.EmbargoLocation, nil
}

// Writes to a new CSV file (include .csv in the name)
func CreateCSV(data []EmbargoItem, filename string) error {
	// attempt to create a file, return error if it fails
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// create a new writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// read through the data and write to the file
	for _, item := range data {
		row := []string{item.DestZip, GetEmbargoType(item)}
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}
	return nil
}

// Just to mimick the old code, plus it reads better up there
// returns a more readable embargo type based on the single-letter response
func GetEmbargoType(item EmbargoItem) string {
	switch item.EmbargoType {
	case "F":
		return "Freezable"
	case "L":
		return "Full Embargo"
	default:
		return ""
	}
}

// Weird string magic to get the access token
// Regex might've been better, but this is kinda more stable and easier
func GetTokenFromResponse(response string) string {
	// First we split by commas and grab the first ("access_token":"blabla")
	fields := strings.SplitN(response, ",", -1)
	// fmt.Printf("sliced: %s\n", fields[0])

	// Then we split the first result from earlier by ":" and grab the SECOND ("blabla")
	values := strings.SplitN(fields[0], ":", -1)
	// fmt.Printf("2nd slice: %s\n", values[1])

	// Just for kicks, remove the extra quotes
	return strings.Replace(values[1], "\"", "", -1)
}

