package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func runCmd(input string, games []game, byId map[int]game) bool {
	fmt.Println()

	cmd := strings.Fields(input)
	if len(cmd) == 0 {
		return true
	}
	switch cmd[0] {
	case "quit":
		return cmdQuit()
	case "list":
		return cmdList(games)
	case "id":
		return cmdByID(cmd, games, byId)
	case "save":
		return cmdSave(games)
	}
	return true

}

func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}

func cmdList(games []game) bool {
	for _, g := range games {
		printGame(g)
	}
	return true
}

func cmdByID(cmd []string, games []game, byId map[int]game) bool {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}
	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}
	g, ok := byId[id]
	if !ok {
		fmt.Println("Sorry, I don't have the game")
		return true
	}
	printGame(g)
	return true
}

func cmdSave(games []game) bool {
	var jsonEncode []jsonGame
	for _, g := range games {
		jsonEncode = append(jsonEncode, jsonGame{g.id, g.name, g.genre, g.price})
	}

	out, err := json.MarshalIndent(jsonEncode, "", "\t")
	if err != nil {
		fmt.Println("Sorry:", err)
		return true
	}

	fmt.Println(string(out))
	return false
}
