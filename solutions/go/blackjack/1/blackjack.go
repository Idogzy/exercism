package blackjack



// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
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
        case "ten", "jack", "queen", "king":
        	return 10
        default:
        	return 0
    }
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    fCard := ParseCard(card1)
    sCard := ParseCard(card2)
    dCard := ParseCard(dealerCard)
    score := fCard + sCard
    if score == 22 {
        return "P"
    } else if score == 21 {
        if dCard >= 10 {
            return "S"
        }
        return "W"
    } else if score >= 17 && score <= 20 {
        return "S"
    } else if score >= 12 && score <= 16 {
        if dCard >= 7 {
            return "H"
        }
        return "S"
    } else {
        return "H"
    }
	panic("Please implement the FirstTurn function")
}

