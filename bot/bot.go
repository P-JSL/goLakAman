package bot

import (
	"botimgs/config"
	"botimgs/database"
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

var BotId string
var goBot *discordgo.Session
var DB, _ = database.Init()

const PREFIX string = "!"

func Start() {

	// Create a new Discord session using the provided bot token.
	goBot, err := discordgo.New("Bot " + config.Token)
	setPrefix("!")

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	BotId = u.ID

	defer goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	fmt.Println("bot Running")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.ChannelID == "913654107154284604" {
		// !
		var prefixString string = substring(m.Content, 0, 1)
		if PREFIX == prefixString {
			var ids = []string{m.Author.Username, m.Author.Discriminator}
			//사용자아이디
			var authorId string = substring(m.Content, 1, len(strings.Join(ids, "#"))+1)
			//#XXXX
			var another string = substring(m.Content, 1+len(strings.Join(ids, "#")), len(m.Content))

			if strings.Join(ids, "#") == authorId {
				if another == "/거래소 인증" {
					_, _ = s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+"님 거래소 인증이 되었습니다.")
					//아만디스코드 역할 Trade : 913356476544864276
					err := s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, "913356476544864276")
					if err != nil {
						fmt.Println("error", err.Error())
					}
				} else if another == "/삭제" {
					err := s.GuildMemberRoleRemove(m.GuildID, m.Author.ID, "913356476544864276")
					if err != nil {
						fmt.Println("error", err.Error())
					}
					_, _ = s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+"님 거래소 역할이 삭제되었습니다.")
				}
				DB.QueryRow("insert into userInfo (userName, )")
			}
		}
		defer multipleOutputs(m)
	}

}

func substring(str string, firstIndex int, lenIndex int) string {
	subStr2 := str[firstIndex:lenIndex]

	fmt.Println(subStr2)
	return subStr2

}

func setPrefix(pf string) string {
	PREFIX := pf
	return PREFIX
}
func multipleOutputs(m *discordgo.MessageCreate) {
	logFile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiWriter)
	log.Println(m.Author.Mention() + ": " + m.Content)
}
