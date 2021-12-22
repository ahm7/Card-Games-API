package Test

import (
	"bytes"
	//"fmt"
	"net/http"
 	"testing"
	//"io"
)

func TestCreateDeck_Right(t *testing.T) {

	var jsonStr = []byte(`{"shuffled":true,"cards":"AH,AS,AD"}`)
	resp, err := http.Post("http://localhost:8001/api/Deck", "application/json", bytes.NewBuffer(jsonStr))

    _=err
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
    }

    /*bodyBytes, err := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)*/
}

func TestCreateDeck_wrongType(t *testing.T) {

	var jsonStr = []byte(`{"shuffled":"true","cards":"AH,AS,AD"}`)
	resp, err := http.Post("http://localhost:8001/api/Deck", "application/json", bytes.NewBuffer(jsonStr))

    _=err
	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
    }
   
}

func TestCreateDeck_syntaxWrong(t *testing.T) {

	var jsonStr = []byte(`{"shofled":true,"cards":"AH,AS,AD"}`)
	resp, err := http.Post("http://localhost:8001/api/Deck", "application/json", bytes.NewBuffer(jsonStr))

    _=err
	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
    }
   
}

func TestGetDeck_existedID(t *testing.T) {

	resp, err := http.Get("http://localhost:8001/api/Deck/8d4fc95d-f378-4e0d-9843-f68cc11b4708")

    _=err
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
    }
   
}

func TestGetDeck_NotExistedID(t *testing.T) {

	resp, err := http.Get("http://localhost:8001/api/Deck/8d4fc95d-f378-4e0d-9843-f68cc11b4709")

    _=err
	if status := resp.StatusCode; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
    }
   
}


func TestDrawCard_Right(t *testing.T) {

	var jsonStr = []byte(`{"count":3}`)
	resp, err := http.Post("http://localhost:8001/api/Deck/1e5e0564-36c4-41dc-8e20-e3d7c0dd3885/draw", "application/json", bytes.NewBuffer(jsonStr))

    _=err
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
    }

    /*bodyBytes, err := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)*/
}

func TestDrawCard_wrongID(t *testing.T) {

	var jsonStr = []byte(`{"count":3}`)
	resp, err := http.Post("http://localhost:8001/api/Deck/8d4fc95d-f378-4e0d-9843-f68cc11b408/draw", "application/json", bytes.NewBuffer(jsonStr))

    _=err
	if status := resp.StatusCode; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
    }

    /*bodyBytes, err := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)*/
}
func TestDrawCard_CountBiggerThanRemaining(t *testing.T) {

	var jsonStr = []byte(`{"count":3}`)
	resp, err := http.Post("http://localhost:8001/api/Deck/8d4fc95d-f378-4e0d-9843-f68cc11b4708/draw", "application/json", bytes.NewBuffer(jsonStr))

    _=err
	if status := resp.StatusCode; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
    }

    /*bodyBytes, err := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)*/
}
func TestDrawCard_countSyntaxError(t *testing.T) {

	var jsonStr = []byte(`{"coud":2}`)
	resp, err := http.Post("http://localhost:8001/api/Deck/8d4fc95d-f378-4e0d-9843-f68cc11b4708/draw", "application/json", bytes.NewBuffer(jsonStr))

    _=err
	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
    }

    /*bodyBytes, err := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)*/
}

func TestDrawCard_NoBody(t *testing.T) {

	var jsonStr = []byte(`{}`)
	resp, err := http.Post("http://localhost:8001/api/Deck/8d4fc95d-f378-4e0d-9843-f68cc11b4708/draw", "application/json", bytes.NewBuffer(jsonStr))

    _=err
	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
    }

    /*bodyBytes, err := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)*/
}