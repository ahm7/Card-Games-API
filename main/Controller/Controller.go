package Controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"main/ApiHelper"
	"main/DataBase"
	"main/Helper"
	"main/Model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Create_Deck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("in create deck")
	// creating unique uuid for the Deck
	id := uuid.New()

	// getting request body parameters

	decoder := json.NewDecoder(r.Body)
	var req_body map[string]interface{}
	err := decoder.Decode(&req_body)
	if err != nil {panic(err)}

	//validate request
	valid := Helper.ValidateCreateDeckReq(req_body)

	if !valid {
		ApiHelper.BadRequest(w)
		return
	}

	// getting the deck properties from the request body
	

	Is_FUll_Deck := Helper.Full_Deck(req_body) // return true for full

	if Is_FUll_Deck {

		var this_deck_Cards = Helper.Construct_Cards_full_Deck()

		// check if shuffled
		if req_body["shuffled"].(bool){
			this_deck_Cards = Helper.ShuffleCards(this_deck_Cards)
		}

		newdeck := Model.Deck{ID: id.String(), Shuffled: req_body["shuffled"].(bool), Remaining: len(this_deck_Cards), Cards: this_deck_Cards}
		DataBase.Insert_Deck(newdeck)
		json.NewEncoder(w).Encode(newdeck)

	} else {

		var this_deck_Cards = Helper.Construct_Cards_partial_Deck(req_body["cards"].(string))

		if req_body["shuffled"].(bool) {
			this_deck_Cards = Helper.ShuffleCards(this_deck_Cards)
		}

		newdeck := Model.Deck{ID: id.String(), Shuffled: req_body["shuffled"].(bool), Remaining: len(this_deck_Cards), Cards: this_deck_Cards}
		DataBase.Insert_Deck(newdeck)
		json.NewEncoder(w).Encode(newdeck)
	}

}

func Open_Deck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in open deck")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
     
	data := DataBase.Get_Deck(w, params["id"]) // get the deck from db
	var deck Model.Deck
	json.Unmarshal([]byte(data), &deck)
	json.NewEncoder(w).Encode(deck)

	fmt.Println(deck)

}

func Draw_Card(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body Model.Draw_card_req_body
	json.NewDecoder(r.Body).Decode(&body)
	params := mux.Vars(r)
	data := DataBase.Get_Deck(w, params["id"]) // get the deck from db
	var deck Model.Deck
	json.Unmarshal([]byte(data), &deck)

	if body.Count == 0 {
		ApiHelper.BadRequest(w)
		return
	}
	if deck.Remaining < body.Count {
		ApiHelper.ContentNotFOund(w)
		return
	} else {

		drawn_cards := DataBase.Delete_cards(w, body.Count, deck.ID)

		json.NewEncoder(w).Encode(drawn_cards)
	}

}
