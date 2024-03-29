package day7

import (
	"cmp"
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "7a",
		Short: "Day 7, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

type HandType int

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	Pair
	HighCard
)

type Hand struct {
	Cards [5]string
	Type  HandType
	Bid   int
}

var powerList = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var powerListB = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
var handRe = regexp.MustCompile(`([2-9TJQKA])([2-9TJQKA])([2-9TJQKA])([2-9TJQKA])([2-9TJQKA]) (\d+)`)

func partA(challenge *core.Input) int {
	return solve(challenge, false)
}

func solve(challenge *core.Input, b bool) int {
	var handList []Hand
	var order []string
	if b {
		handList = core.InputMap[Hand](challenge, GetHandB)
		order = powerListB
	} else {
		handList = core.InputMap[Hand](challenge, GetHandA)
		order = powerList
	}

	sort.Slice(handList, func(i, j int) bool {
		hand1 := handList[i]
		hand2 := handList[j]
		if hand1.Type != hand2.Type {
			return hand1.Type < hand2.Type // higher value hands get ordered first
		}
		for k := 0; k < 5; k++ {
			if hand1.Cards[k] != hand2.Cards[k] {
				return sliceIndex(order, hand1.Cards[k]) > sliceIndex(order, hand2.Cards[k])
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

func sliceIndex[T cmp.Ordered](slice []T, v T) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}

	return -1
}

func GetHandA(line string) Hand {
	return GetHand(line, false)
}

func GetHand(line string, b bool) Hand {
	// []string{full match, 5 numbers, bid}
	regexRes := handRe.FindStringSubmatch(line)
	cards := [5]string{"", "", "", "", ""}
	for i := 0; i < 5; i++ {
		cards[i] = regexRes[i+1]
	}
	bid, _ := strconv.Atoi(regexRes[6])

	handCountsMap := make(map[string]int)
	jacks := 0
	if b {
		for _, card := range cards {
			if card != "J" {
				handCountsMap[card] += 1
			} else {
				jacks++
			}
		}
	} else {
		for _, card := range cards {
			handCountsMap[card] += 1
		}
	}
	handCountsList := []int{}
	for _, value := range handCountsMap {
		handCountsList = append(handCountsList, value)
	}
	sort.Slice(handCountsList, func(i, j int) bool {
		return handCountsList[i] > handCountsList[j]
	})

	// code for part B
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
