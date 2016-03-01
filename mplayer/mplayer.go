package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/uuappli/gosamples/mplayer/library"
	"github.com/uuappli/gosamples/mplayer/mp"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Artist, e.Name, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})

		} else {
			fmt.Println("USAGE: lib add <name> <artist> <source> <type>")
		}
	case "remove":
		if len(tokens) == 3 {
			i, _ := strconv.Atoi(tokens[2])
			lib.Remove(i)

		} else {
			fmt.Println("USAGE: lib remove <id>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])

	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
	}

	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music ", tokens[1], "does not exist")
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println("Enter following commands to control the player : \n lib list -- View the existing music lib \n lib add <name> <artist> <source> <type> -- Add a music to music lib \n lib remove <name> -- Remove th specified music from the lib \n play <name> --Play the specified music")
	lib = library.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command->")

		rawline, _, _ := r.ReadLine()

		line := string(rawline)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)

		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)

		} else {
			fmt.Println("Unrecognized lib command:", tokens[0])
		}
	}
}
