package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	token string = "ODIxMzE0NTg2ODg3MzIzNjQ4.YFB60A.0JuGRgw3J6ayDCwQdb8GGj63vGA"
	//BotID is to
	BotID string
)

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	dg.AddHandler(messageHandler)
	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
