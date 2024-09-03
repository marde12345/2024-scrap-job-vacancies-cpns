package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	FormationPublic   string = "1"
	FormationCumlaude string = "4"

	DegreeInformaticEngineering string = "5101087"
)

var (
	FormationList = []string{FormationPublic, FormationCumlaude}
)

type InstansiResponse struct {
	Result []Instansi `json:"result"`
}

// Define a struct to match the JSON structure
type Instansi struct {
	ID   string `json:"id"`
	Info struct {
		ID   string `json:"id"`
		Name string `json:"nama"`
		Code string `json:"kode"`
	} `json:"instansi"`
}

type PositionResponse struct {
	Result []Position `json:"result"`
}

type Position struct {
	Code                 string  `json:"kode"`
	Name                 string  `json:"nama"`
	JenisJabatanId       int     `json:"jenisJabatanId"`
	MintaSertifikasiGuru *string `json:"mintaSertifikasiGuru"`
	MintaSTR             *string `json:"mintaSTR"`
	IsDiplomat           string  `json:"isDiplomat"`
	KodeBkn              *string `json:"kodeBkn"`
	JenisKelamin         *string `json:"jenisKelamin"`
	KodeDapodik          *string `json:"kodeDapodik"`
	AsalData             *string `json:"asalData"`
	IsDisabilitas        bool    `json:"isDisabilitas"`
}

type JobLocationResponse struct {
	Result []JobLocation `json:"result"`
}

type JobLocation struct {
	Code                 string  `json:"kode"`
	Name                 string  `json:"nama"`
	KeteranganPendidikan *string `json:"keteranganPendidikan"` // Nullable string
	Toefl                bool    `json:"toefl"`
	Npsn                 *string `json:"npsn"` // Nullable string
	JumlahKebutuhan      int     `json:"jumlahKebutuhan"`
}

func main() {
	// Create a CSV file
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	agencyList := fetchAgencyList()

	// Print the parsed data
	for _, agency := range agencyList {
		for _, formation := range FormationList {
			positionList := fetchPositionByAgencyIDFormationIDDegreeID(agency.ID, formation, DegreeInformaticEngineering)
			for _, position := range positionList {
				jobLocationList := fetchPositionByAgencyIDFormationIDDegreeIDPositionID(agency.ID, formation, DegreeInformaticEngineering, position.Code)
				for _, job := range jobLocationList {
					if err := writer.Write([]string{agency.ID, agency.Info.Name, position.Code, position.Name, job.Code, job.Name, fmt.Sprintf("%d", job.JumlahKebutuhan)}); err != nil {
						log.Fatalf("Failed to write record to file: %v", err)
					}
				}
			}
		}
	}

	fmt.Println("DONEEE")
}

func fetchAgencyList() (result []Instansi) {
	// URL to fetch the JSON data from
	url := "https://daftar-sscasn.bkn.go.id/daftar/instansi.json"

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set headers as specified in the curl command
	// TODO: Set session here

	// Send the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Parse the JSON response
	var respInstansiList InstansiResponse
	err = json.Unmarshal(body, &respInstansiList)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	return respInstansiList.Result
}

func fetchPositionByAgencyIDFormationIDDegreeID(agencyID, formationID, degreeID string) (result []Position) {
	// URL with query parameters
	url := fmt.Sprintf(
		"https://daftar-sscasn.bkn.go.id/daftar/jabatan.json?instansi=%s&isCpnsNakes=false&jenisFormasi=%s&pendidikan=%s",
		agencyID,
		formationID,
		degreeID,
	)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set headers as specified in the curl command
	// TODO: Set session here

	// Send the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Parse the JSON response
	var positionResp PositionResponse
	err = json.Unmarshal(body, &positionResp)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	return positionResp.Result
}

func fetchPositionByAgencyIDFormationIDDegreeIDPositionID(agencyID, formationID, degreeID, positionID string) (result []JobLocation) {
	// URL with query parameters
	url := fmt.Sprintf(
		"https://daftar-sscasn.bkn.go.id/daftar/lokasiKerja.json?instansi=%s&isCpnsNakes=false&isJabatanDisabilitas=0&jabatan=%s&jenisFormasi=%s&jenisKelamin=&pendidikan=%s",
		agencyID,
		positionID,
		formationID,
		degreeID,
	)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set headers as specified in the curl command
	// TODO: Set session here

	// Send the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading response body: %v", err)
		return
	}

	// Parse the JSON response
	var jobLocationResp JobLocationResponse
	err = json.Unmarshal(body, &jobLocationResp)
	if err != nil {
		fmt.Printf("error parsing JSON: %v", err)
		return
	}

	return jobLocationResp.Result
}
