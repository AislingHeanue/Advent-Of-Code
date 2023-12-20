package day20

import (
	"fmt"
	"regexp"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "20a",
		Short: "Day 20, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(core.FromFile()))
		},
	}
}

type ConnectionType int

const (
	Broadcaster ConnectionType = iota
	FlipFlop
	Conjunction
	None
)

type Connection struct {
	id              string
	connectionType  ConnectionType
	on              bool
	mostRecentHigh  bool
	connectToString []string
	connectTo       []*Connection
	connectFrom     []*Connection
}

// do we also need "from sting" here? I'm guessing not
type Pulse struct {
	to   string
	high bool
}

var re = regexp.MustCompile(`(?:(broadcaster)|(\%)([a-z]+)|(\&)([a-z]+)) -> (?:([a-z]+), )?(?:([a-z]+), )?(?:([a-z]+), )?(?:([a-z]+), )?(?:([a-z]+), )?(?:([a-z]+), )?([a-z]+)`)

func partA(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	connections := core.InputMap[Connection](challenge, getConnection)
	connectionMap := getMap(connections)
	highCount := 0
	lowCount := 0

	for i := 0; i < 1000; i++ {
		pressTheButton(connectionMap, &highCount, &lowCount, "")
	}

	return highCount * lowCount
}

func getMap(connections []Connection) map[string]*Connection {
	connectionMap := make(map[string]*Connection)
	for i := range connections {
		connectionMap[connections[i].id] = &connections[i]
	}
	for _, pointer := range connectionMap {
		for _, childId := range pointer.connectToString {
			c2, ok := connectionMap[childId]
			if !ok {
				// none-type connections (eg output) are put here
				connections = append(connections, Connection{id: childId, connectionType: None})
				c2 = &connections[len(connections)-1]
				connectionMap[childId] = c2
			}

			c2.connectFrom = append(c2.connectFrom, pointer)
			pointer.connectTo = append(pointer.connectTo, c2)
		}
	}

	return connectionMap
}

func getConnection(line string) Connection {
	// regex indices:
	// 0: full string
	// 1: broadcaster or ""
	// 2 and 3: % or ""
	// 4 and 5: & or ""
	// 6-12: indices (filled in order 12,6,7,8,9,10,11)
	regexRes := re.FindStringSubmatch(line)
	var conType ConnectionType
	var id string
	if regexRes[1] != "" {
		conType = Broadcaster
		id = "broadcaster"
	} else if regexRes[2] != "" {
		conType = FlipFlop
		id = regexRes[3]
	} else if regexRes[4] != "" {
		conType = Conjunction
		id = regexRes[5]
	} else {
		panic("what type is this??")
	}
	connectTo := []string{}
	for i := 6; i < 12; i++ {
		if regexRes[i] == "" {
			break
		}
		connectTo = append(connectTo, regexRes[i])
	}
	connectTo = append(connectTo, regexRes[12])
	return Connection{
		id:              id,
		connectionType:  conType,
		connectToString: connectTo,
		connectTo:       []*Connection{},
		connectFrom:     []*Connection{},
		on:              false,
		mostRecentHigh:  false,
	}
}

func pressTheButton(connectionMap map[string]*Connection, highCount *int, lowCount *int, awaitingLowTo string) bool {
	lowHitAtThePlaceWhereYouWantItTo := false // I have passed the point of insanity
	queue := []Pulse{{to: "broadcaster", high: false}}
	for len(queue) > 0 {
		pulse := queue[0]
		if pulse.high {
			*highCount++ // this is a pulse, we are consuming said pulse, all is good
		} else {
			if pulse.to == awaitingLowTo {
				lowHitAtThePlaceWhereYouWantItTo = true
			}
			*lowCount++
		}
		to, ok := connectionMap[pulse.to]
		if !ok {
			continue
		}
		queue = queue[1:]
		switch to.connectionType {
		// default or none: do nothing
		case Broadcaster:
			for _, connectionId := range to.connectTo {
				queue = append(queue, Pulse{to: connectionId.id, high: pulse.high})
			}
			to.mostRecentHigh = pulse.high
			connectionMap[pulse.to] = to
		case FlipFlop:
			if pulse.high {
				continue // nothing happens with this pulse
			}
			to.on = !to.on

			if to.on { // off -> on: send a high pulse
				for _, connectionId := range to.connectTo {
					queue = append(queue, Pulse{to: connectionId.id, high: true})
				}
				to.mostRecentHigh = true
			} else {
				for _, connectionId := range to.connectTo {
					queue = append(queue, Pulse{to: connectionId.id, high: false})
				}
				to.mostRecentHigh = false
			}
			connectionMap[pulse.to] = to
		case Conjunction:
			sendLow := true
			for _, from := range to.connectFrom {
				if !connectionMap[from.id].mostRecentHigh {
					sendLow = false
					break
				}
			}
			for _, connectionId := range to.connectTo {
				queue = append(queue, Pulse{to: connectionId.id, high: !sendLow})
			}
			to.mostRecentHigh = !sendLow
			connectionMap[pulse.to] = to
		}
	}

	return lowHitAtThePlaceWhereYouWantItTo
}
