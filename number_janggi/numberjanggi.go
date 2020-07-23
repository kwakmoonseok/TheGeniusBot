package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const boardRow = 3
const boardColumn = 4

var token string
var pre string="TEST"
var gameBoard = [boardRow][boardColumn]string {
	{"red_sang", "", "", "green_jang"},
	{"red_wang", "red_ja", "green_ja", "green_wang"},
	{"red_jang", "", "", "green_jang"},
}

func init(){
	flag.StringVar(&token, "t", "NzM1MDExOTY1NzYzNTg0MDYx.XxbUbA.xMiDXYUbjixzVTgHN-O9WdVlEOc", "Bot Token")
	flag.Parse()
}

func main(){
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord Session, ", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()
	
	if err != nil {
		fmt.Println("error opening connection, ", err)
		return
	}

	fmt.Println("Bot is now running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<- sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	if m.Author.ID == s.State.User.ID {
		return
	}

	if (m.Content==pre+pre) {
		var command string
		for i := 0; i < boardRow; i++ {
			for j := 0; j < boardColumn; j++ {
				switch {
					case gameBoard[i][j] == "red_sang":
						command += "<:red_sang:735334778659668020> "
					case gameBoard[i][j] == "red_wang":
						command += "<:red_wang:735332478344233000> "
					case gameBoard[i][j] == "red_jang":
						command += "<:red_jang:735331743934054480> "
					case gameBoard[i][j] == "red_ja":
						command += "<:red_ja:735335582397497355> "
					case gameBoard[i][j] == "green_sang":
						command += "<:green_sang:735336627874037790> "
					case gameBoard[i][j] == "green_wang":
						command += "<:green_wang:735336627597082625> "
					case gameBoard[i][j] == "green_jang":
						command += "<:green_jang:735336627651739658> "
					case gameBoard[i][j] == "green_ja":
						command += "<:green_ja:735337243601797233> "
					default:
						command += ":x: "
				}
			}
			command += "\n"
		}
		s.ChannelMessageSend(m.ChannelID, command)
	}
	if (m.Content == "timer") {
		 t:=time.Now()

		 const timeInterval = 30
		 const nanoToSec = 1000000000
		 const limitSec = 120

		 secondMap := make(map[int]int)
		 elapsedTime := int(time.Since(t) / nanoToSec)

		 for elapsedTime < limitSec {
		 	elapsedTime = int(time.Since(t) / nanoToSec)
		 	if elapsedTime % timeInterval == 0 {
				_, exist := secondMap[elapsedTime / timeInterval]
				if !exist && elapsedTime < limitSec{
					s.ChannelMessageSend(m.ChannelID, strconv.Itoa(limitSec - elapsedTime)+"초 남았습니다.\n")
					secondMap[elapsedTime / timeInterval] = elapsedTime
				}
			}
		 }
		s.ChannelMessageSend(m.ChannelID, "시간이 종료되었습니다.\n")
	}
}