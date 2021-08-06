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
	// Load the token using godotenv.
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	// Create a dg session.
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD"))
	if err != nil {
		fmt.Println(err)
	}

	dg.AddHandler(messageCreate)

	// Open a connection to Discord.
	err = dg.Open()
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	// Listen for the CTRL-C command.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Close down everything.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore everythong from the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.ToLower(m.Content) == "bruh" {
		s.ChannelMessageSend(m.ChannelID, "bruh")
	}
}
