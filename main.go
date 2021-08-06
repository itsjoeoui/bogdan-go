package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(os.Getenv("DISCORD"))

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD"))

	if err != nil {
		fmt.Println(err)
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.ToLower(m.Content) == "bruh" {
		s.ChannelMessageSend(m.ChannelID, "bruh")
	}
}
