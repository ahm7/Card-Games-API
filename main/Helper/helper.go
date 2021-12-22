package Helper

import (
	"fmt"
	"main/Model"
	"main/Storage"
	"math/rand"
	"strings"
	"time"
)


func ValidateCreateDeckReq(req_body map[string]interface{}) bool{
 

	var shuffledExist bool = false
	fmt.Println("in validate ")

	if(len(req_body) >2) {return false}

	keys := make([]string, len(req_body))

i := 0
for k := range req_body {
	keys[i] = k
	i++
	if(k != "shuffled" && k != "cards"){
		return false
	}
	if(k =="shuffled"){  
		shuffledExist = true
		switch v := req_body["shuffled"].(type) {
		case bool:
			_=v
		default:
			return false	
 		}
	
	}
	if(k =="Cards"){  

		switch v := req_body["Cards"].(type) {
		case string:
			_=v
		default:
			return false	
 		}
	
	}
}

if(!shuffledExist){ req_body["shuffled"]=false}

return true
}

func ValidateDrawCardReq(req_body map[string]interface{}) bool{



	return true;
}


func Construct_Cards_full_Deck ( ) []Model.Card{

	// getting all the 52 card from the local json file
	all_cards := Storage.GetAllCards();

	var cards = make([]Model.Card, len(all_cards.Cards))
	for i := 0; i < len(all_cards.Cards); i++ {
		cards[i] = all_cards.Cards[Storage.SortedCardsjkEYS[i]] // get card by key from all_cards
	}

return cards
}

func Construct_Cards_partial_Deck ( cards_code string) []Model.Card{
	
	      all_cards := Storage.GetAllCards();
          res := strings.Split(cards_code, ",")
		  var cards = make([]Model.Card, len(res))

		for i := 0; i < len(res); i++ {
			cards[i] = all_cards.Cards[res[i]] // get card by key from all_cards
		}


 
return cards
}

func Full_Deck(deck map[string]interface{}) bool {

	if val, ok := deck["cards"]; ok  {
		_ = val
		return false}
	return true
}


func ShuffleCards ( cards[]Model.Card) []Model.Card{


	rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })

 
return cards
}



