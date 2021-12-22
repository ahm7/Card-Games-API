package Model

import (
	//"encoding/json"
	//"fmt"

)
type Deck struct {
	ID     string  `json:"id"`
	Shuffled   bool  `json:"shuffled"`
	Remaining  int  `json:"remaining"`
	Cards []Card `json:"cards"`
}
// there is another struct for the requested deck it has a strings of cards like "AS,AH,..."
type Req_Deck_Body struct {
	Shuffled   bool  `json:"shuffled"`
	Cards string `json:"cards"`
}


// card struct
type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}
type Cards struct {
   Cards map[string]Card `json:"cards"`
}

type Draw_card_req_body struct {
	Count     int  `json:"count"`
}





 