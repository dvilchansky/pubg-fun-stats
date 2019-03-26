package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"os"
	"pubg-fun-stats/parser"
)

func init() {
	API = gopubg.NewAPI(os.Getenv(`PUBG_API_KEY`))
}

var (
	API *gopubg.API
)

func main() {
	dg, err := discordgo.New("Bot " + "DISCORD_TOKEN")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//u, err := dg.User("@me")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	dg.AddHandler(MessageHandler)
	err = dg.Open()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("Bot is running")
}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Content)
}
