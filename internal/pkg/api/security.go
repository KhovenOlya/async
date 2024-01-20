package api

import (
	"encoding/json"
	"fmt"
	"bytes"
	"math/rand"
	"net/http"
	"time"
)

var token = "H#12H$EdEi^9"

type checkReq struct {
	ID int `json:"id"`
}

func Check(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	req := &checkReq{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(req); err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	go worker(req.ID)
}

func worker(id int) {
	fmt.Println("Start")
	time.Sleep(10 * time.Second)
	fmt.Println("Done")

	var result string
	if (rand.Intn(10 - 1) > 3) {
		result = "success"
	} else {
		result = "fail"
	}

	putURL := fmt.Sprintf("http://localhost:8000/api/permits/%d/update_security_decision", id)

	// Create the JSON payload
	requestData := map[string]interface{}{
		"security_decision": result,
		"token": token,
	}

	// Convert data to JSON
	putBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("PUT", putURL, bytes.NewBuffer(putBody))
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making PUT request:", err)
		return
	}
	defer resp.Body.Close()

	// Проверьте код статуса ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Unexpected status code %d\n", resp.StatusCode)
		return
	}

	fmt.Println("PUT request successful")
}
