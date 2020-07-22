package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"github.com/bwmarrin/discordgo"
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
		s.ChannelMessageSend(m.ChannelID,"현재 상황\n:regional_indicator_b: :x: :x: :three:\n" +
			":regional_indicator_a: :regional_indicator_d: :four: :one:\n" +
			":regional_indicator_c: :x: :x: :two:")
	}
}