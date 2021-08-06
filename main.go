package main

import (
	"fmt"
	"os"

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
	dg.Close()
}
