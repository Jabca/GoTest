package main_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func assertEqual(t *testing.T, expected, actual interface{}, message string) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Assertion failed: %s\nExpected: %v\nActual: %v", message, expected, actual)
	}
}

func readFile(path string) ([4]string, error) {
	file, err := os.Open(path)
	var lines [4]string
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines[index] = line
		index++
	}

	// Check if there was any error during scanning
	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, nil
}

type ReqBody struct {
	Image string `json:"image"`
}

type ResBody struct {
	Data ReqBody `json:"data"`
}

type StoredImage struct {
	gorm.Model
	NEGATIVE string `json:"negative_image"gorm:"not null;default:null"`
	POSITIVE string `json:"positive_image"gorm:"not null;default:null"`
}

type ThreeIMages struct {
	Data []StoredImage `json:"data"`
}

func fillDb() error {
	loadEnv()
	serverUrl := "http://" + os.Getenv("BACKEND_HOST") + ":" + os.Getenv("BACKEND_PORT") + "/negative_image"
	files, err := readFile("src.txt")
	if err != nil {
		return err
	}
	for i := 0; i < 3; i++ {
		cur_image := files[i]

		payload := ReqBody{Image: cur_image}

		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		req, err := http.NewRequest("POST", serverUrl, bytes.NewBuffer(payloadBytes))
		if err != nil {
			return err
		}

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
	}

	return nil

}

func TestNegativeImage(t *testing.T) {
	loadEnv()

	serverUrl := "http://" + os.Getenv("BACKEND_HOST") + ":" + os.Getenv("BACKEND_PORT") + "/negative_image"
	files, err := readFile("src.txt")

	if err != nil {
		t.Fatalf("Couldn't open file: %v", err)
	}

	upload_image := files[0]
	expected_return := files[1]

	payload := ReqBody{Image: upload_image}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", serverUrl, bytes.NewBuffer(payloadBytes))

	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var data ResBody

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assertEqual(t, expected_return, data.Data.Image, "Response is unexpected")

}

func TestNegativeIMageFail(t *testing.T) {
	loadEnv()

	serverUrl := "http://" + os.Getenv("BACKEND_HOST") + ":" + os.Getenv("BACKEND_PORT") + "/negative_image"

	files, err := readFile("src.txt")
	if err != nil {
		t.Fatalf("Couldn't open file: %v", err)
	}

	wrong_type := files[3]

	payload := ReqBody{Image: wrong_type}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", serverUrl, bytes.NewBuffer(payloadBytes))

	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to perform HTTP : %v", err)
	}
	if resp.StatusCode == 201 {
		t.Fatalf("Failed to detect wrong file type")
	}
}

func TestGetLastImages(t *testing.T) {
	loadEnv()
	fillDb()
	files, err := readFile("src.txt")

	serverUrl := "http://" + os.Getenv("BACKEND_HOST") + ":" + os.Getenv("BACKEND_PORT") + "/get_last_images"

	req, err := http.NewRequest("GET", serverUrl, nil)

	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var data ThreeIMages

	err = json.Unmarshal(body, &data)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	for i := 0; i < 3; i++ {
		assertEqual(t, files[i], data.Data[2-i].POSITIVE, "Response is unexpected")
	}

}
