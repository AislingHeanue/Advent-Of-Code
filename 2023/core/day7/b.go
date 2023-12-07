package day7

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "7b",
		Short: "Day 7, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

var powerListB = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

func partB(challenge *core.Input) int {
	handList := core.InputMap[Hand](challenge, GetHandB)

	sort.Slice(handList, func(i, j int) bool {
		hand1 := handList[i]
		hand2 := handList[j]
		if hand1.Type != hand2.Type {
			return hand1.Type < hand2.Type // higher value hands get ordered first
		}
		for k := 0; k < 5; k++ {
			if hand1.Cards[k] != hand2.Cards[k] {
				return sliceIndex(powerListB, hand1.Cards[k]) > sliceIndex(powerListB, hand2.Cards[k])
			}
		}

		return false
	})

	total := 0
	for i, hand := range handList {
		total += (len(handList) - i) * hand.Bid
	}

	return total
}

func GetHandB(line string) Hand {
	// []string{full match, 5 numbers, bid}
	regexRes := handRe.FindStringSubmatch(line)
	cards := [5]string{"", "", "", "", ""}
	for i := 0; i < 5; i++ {
		cards[i] = regexRes[i+1]
	}
	bid, _ := strconv.Atoi(regexRes[6])

	handCountsMap := make(map[string]int)
	jacks := 0
	for _, card := range cards {
		if card != "J" {
			handCountsMap[card] += 1
		} else {
			jacks++
		}
	}
	handCountsList := []int{}
	for _, value := range handCountsMap {
		handCountsList = append(handCountsList, value)
	}
	sort.Slice(handCountsList, func(i, j int) bool {
		return handCountsList[i] > handCountsList[j]
	})
	if len(handCountsList) == 0 {
		return Hand{
			cards,
			FiveOfAKind,
			bid,
		} // five jacks
	}
	handCountsList[0] += jacks

	handType := HighCard
	if handCountsList[0] == 5 {
		handType = FiveOfAKind
	} else if handCountsList[0] == 4 {
		handType = FourOfAKind
	} else if handCountsList[0] == 3 && handCountsList[1] == 2 {
		handType = FullHouse
	} else if handCountsList[0] == 3 {
		handType = ThreeOfAKind
	} else if handCountsList[0] == 2 && handCountsList[1] == 2 {
		handType = TwoPair
	} else if handCountsList[0] == 2 {
		handType = Pair
	}

	return Hand{
		cards,
		handType,
		bid,
	}
}
