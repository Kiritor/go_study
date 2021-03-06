package main

import (
    "bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"musiclib"
	"music"
)
var lib *musiclib.MusicManager
var id int =1
var ctr,signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
		case "list":
		     for i:=0;i<lib.Len();i++ {
				 e,_:=lib.Get(i)
				 fmt.Println(i+1,":",e.Name,"-",e.Artist,"-",e.Source,"-",e.Type)
			 }
		case "add":{
			if len(tokens) == 6 {
				id++
				lib.Add(&musiclib.Music{
					strconv.Itoa(id),tokens[2],tokens[3],tokens[4],tokens[5]})
			}else {
				fmt.Println("USAGE:lib add <name><artist><source><type>")
			}
		}
		case "remove":{
			if len(tokens) ==3 {
				lib.RemoveByName(tokens[2])
			}else {
				fmt.Println("USAGE:lib remove <name>")
			}
		}
		default:
		    fmt.Println("Unrecognized lib command:",tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) !=2 {
		fmt.Println("USAGE:play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e==nil {
		fmt.Println("the music ",tokens[1]," not exist.")
		return
	}
	music.Play(e.Source,e.Type)
}

func main() {
	fmt.Println(`Enter following commands to control the player:
                 lib list -- view the existing msic lib
                 lib add <name><artist><source><type> --add music to the lib
                 lib remove <name>--remove the specified music from the lib
                 play <name> --play the specified music
   `)
	lib =musiclib.NewMusicManager()
	r :=bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command-> ")
		rawLine,_,_:=r.ReadLine()
		line :=string(rawLine)
		if line =="q"||line == "e" {
			break
		}
		tokens :=strings.Split(line," ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		}else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		}else {
			fmt.Println("Unrecognized command:",tokens[0])
		}
	}
}
