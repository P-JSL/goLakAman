package util

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func RemoveRole(s *discordgo.Session, GuildID string, UserID string, arg string) string {
	fmt.Println("removeRole", arg)

	g, err := s.State.Guild(GuildID)
	if err != nil {
		fmt.Println(err)
		return "You fucking broke it~"
	}

	member, err := s.GuildMember(GuildID, UserID)
	if err != nil {
		fmt.Println(err)
	}

	pos := -1
	for i, r := range member.Roles {
		if r == "913632866238332978" {
			pos = i
			break
		}
	}

	if pos < 0 {
		return fmt.Sprintf("You're already not subscribed to %s~", arg)
	}

	member.Roles = append(member.Roles[:pos], member.Roles[pos+1:]...)
	err = s.GuildMemberEdit(GuildID, UserID, member.Roles)
	if err != nil {
		fmt.Println(err)
		return "I can't touch that group dude, do it yourself~"
	}

	delete := true
	for _, member := range g.Members {
		if member.User.ID == UserID {
			continue // Ignore self since it's not updated here yet
		}

		for _, r := range member.Roles {
			if r == "913632866238332978" {
				delete = false
				break
			}
		}
	}

	fmt.Println("Should delete it?", delete)

	if delete {
		err := s.GuildRoleDelete(GuildID, "913632866238332978")
		if err != nil {
			fmt.Println(err)
			return fmt.Sprintf("Unsubscribed from but failed to delete %s~", arg)
		}

		fmt.Println("Unsubscribed and deleted")
		return fmt.Sprintf("Unsubscribed from and deleted %s~", arg)
	}

	fmt.Println("Unsubscribed")
	return fmt.Sprintf("Unsubscribed from %s~", arg)
}


func AddRole(s *discordgo.Session, GuildID string, UserID string, arg string) string {
	_, err := s.State.Guild(GuildID)
	if err != nil {
		fmt.Println(err)
		return "You fucking broke it~"
	}

	member, err := s.GuildMember(GuildID, UserID)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	newRoles := append(member.Roles, arg)

	err = s.GuildMemberEdit(GuildID, UserID, newRoles)
	if err != nil {
		fmt.Println(err)
		return "I can't touch that group dude, do it yourself~"
	}


	return fmt.Sprintf("Created and subscribed to %s", arg)
}
