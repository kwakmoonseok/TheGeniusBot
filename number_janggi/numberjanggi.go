package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var token string
var pre string="TEST"

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
	
	if (m.Content == pre+"ping"){
		s.ChannelMessageSend(m.ChannelID, "Pong!");
	}

	if (m.Content == pre+"pong"){
		s.ChannelMessageSend(m.ChannelID, "Ping!");
	}

	if(m.Content==pre+pre){
		s.ChannelMessageSend(m.ChannelID," :red_sang:735334778659668020 :x: :x: :green_jang:735336627651739658\n" +
			":red_wang:735332478344233000 :red_ja:735335582397497355 :green_ja:735337243601797233 :green_wang:735336627597082625\n " +
			":red_jang:735331743934054480 :x: :x: :green_sang:735336627874037790\n")
	}
}