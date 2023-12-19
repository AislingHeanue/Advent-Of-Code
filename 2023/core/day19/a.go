package day19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "19a",
		Short: "Day 19, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

type Instruction struct {
	id    string
	rules []Rule
	final string
}

type Rule struct {
	letter  string
	greater bool
	number  int
	sendTo  string
}

var re1 = regexp.MustCompile(`([a-z]+)\{([xmas])([<>])(\d+):([a-zAR]+)+,` +
	`(?:([xmas])([<>])(\d+):([a-zAR]+),)?` +
	`(?:([xmas])([<>])(\d+):([a-zAR]+),)?` +
	`(?:([xmas])([<>])(\d+):([a-zAR]+),)?` +
	`([a-zAR]+)\}`)
var re2 = regexp.MustCompile(`\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}`)

var instructionsMap map[string]Instruction

func partA(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	items := makeInstructionMap(challenge)
	counts := core.InputMap[int](core.FromLiteral(strings.Join(items, "\n")), processItem)
	total := 0
	for _, count := range counts {
		total += count
	}
	return total
}

func makeInstructionMap(challenge *core.Input) []string {
	lines := challenge.LineSlice()

	instructions := []string{}
	items := []string{}
	for s, line := range lines {
		if line == "" {
			instructions = lines[:s]
			items = lines[s+1:]
		}
	}
	doneInstructions := core.InputMap[Instruction](core.FromLiteral(strings.Join(instructions, "\n")), getInstruction)
	instructionsMap = make(map[string]Instruction)
	for _, inst := range doneInstructions {
		instructionsMap[inst.id] = inst
	}

	return items
}

func processItem(line string) int {
	item := make(map[string]int)
	regexRes := re2.FindStringSubmatch(line)
	item["x"], _ = strconv.Atoi(regexRes[1])
	item["m"], _ = strconv.Atoi(regexRes[2])
	item["a"], _ = strconv.Atoi(regexRes[3])
	item["s"], _ = strconv.Atoi(regexRes[4])
	currentState := "in"
	for {
		currentInst := instructionsMap[currentState]
		sent := false
		for _, rule := range currentInst.rules {
			if rule.greater {
				if item[rule.letter] > rule.number {
					currentState = rule.sendTo
					sent = true
					break
				}
			} else {
				if item[rule.letter] < rule.number {
					currentState = rule.sendTo
					sent = true
					break
				}
			}
		}
		if !sent {
			currentState = currentInst.final
		}
		if currentState == "A" {
			return item["x"] + item["m"] + item["a"] + item["s"]
		} else if currentState == "R" {
			return 0
		}
	}
}

func getInstruction(line string) Instruction {
	regexRes := re1.FindStringSubmatch(line)
	// fmt.Println(regexRes)
	rules := []Rule{}
	for i := 2; i <= 14 && regexRes[i] != ""; i += 4 {
		rules = append(rules, getRule(regexRes[i:i+4]))
	}

	return Instruction{
		id:    regexRes[1],
		rules: rules,
		final: regexRes[18],
	}
}

func getRule(lines []string) Rule {
	num, err := strconv.Atoi(lines[2])
	if err != nil {
		panic("death")
	}
	greater := false
	if lines[1] == ">" {
		greater = true
	}
	return Rule{
		letter:  lines[0],
		number:  num,
		greater: greater,
		sendTo:  lines[3],
	}

}
