// This the only the answer, it will not run since I haven't created main func. I have just solved it.

package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card{
	case "ace":
    	return 11
    case "two":
        return 2
    case "three":
    	return 3
	case "four":
    	return 4
    case "five":
    	return 5
    case "six":
    	return 6
    case "seven":
    	return 7
    case "eight":
    	return 8
    case "nine":
    	return 9
    case "ten":
    	return 10
    case "jack":
    	return 10
    case "queen":
    	return 10
    case "king":
    	return 10
    default:
    	return 0
        }
}

func FirstTurn(card1, card2, dealerCard string) string {
	if (card1 == "ace" && card2 == "ace"){
        return "P"
    }
	cardValues := map[string]int{
		"ace":   11,
		"eight": 8,
		"two":   2,
		"nine":  9,
		"three": 3,
		"ten":   10,
		"four":  4,
		"jack":  10,
		"five":  5,
		"queen": 10,
		"six":   6,
		"king":  10,
		"seven": 7,
		"other": 0,
	}

	// Get the values of the player's cards and the dealer's card
	card1Value := cardValues[card1]
	card2Value := cardValues[card2]
	dealerCardValue := cardValues[dealerCard]

	if (card1Value == 10 && card2Value == 11) || (card1Value == 11 && card2Value == 10) {
		if dealerCardValue != 11 && dealerCardValue < 10 {
			return "W" 
		}
		return "S" 
	}

	if card1 == "ace" && card2 == "ace" {
		return "P" 
	}

	sum := card1Value + card2Value
    
	if sum >= 17 && sum <= 20 {
		return "S"
	} else if sum >= 12 && sum <= 16 {
		if dealerCardValue >= 7 {
			return "H"
		}
		return "S"
	}
	return "H"
}
