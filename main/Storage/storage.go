package Storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"main/Model"
)

// load all cards from the all_CARDS json
func GetAllCards() Model.Cards {
	content, err := ioutil.ReadFile("../main/Storage/All_CARDS.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// unmarshall the data into `objmap`
	var all_cards Model.Cards
	err = json.Unmarshal(content, &all_cards.Cards)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return all_cards
	/*for i := 0; i < len(cards.cards); i++ {
		   fmt.Println("User Type: " + cards.cards[i].code)
	   }
	*/
}
// sorted cards  (keys)

var SortedCardsjkEYS [52]string = [52]string {
	"AS","2S","3S","4S","5S","6S","7S","8S","9S","10S","JS","QS","KS",
	"AD","2D","3D","4D","5D","6D","7D","8D","9D","10D","JD","QD","KD",
	"AC","2C","3C","4C","5C","6C","7C","8C","9C","10C","JC","QC","KC",
	"AH","2H","3H","4H","5H","6H","7H","8H","9H","10H","JH","QH","KH"}