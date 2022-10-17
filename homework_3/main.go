package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

func main() {
	delim := color.InYellow("\n" + strings.Repeat("=", 20) + "\n")
	names := []string{"James", "Mary", "Robert", "Patricia", "John", "Jennifer", "Michael", "Linda", "David", "Elizabeth", "William", "Barbara", "Richard", "Susan", "Joseph", "Jessica", "Thomas", "Sarah", "Charles", "Karen"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
	rnd, rnd1, rnd2, rnd3 := "", [2]any{rand.Intn(len(names)/2) - 1}, [2]any{rand.Intn(len(names)/4) + len(names)/2}, [2]any{len(names) - 1 - rand.Intn(len(names)/4)}

	server := NewServer()

	for i, v := range names {
		uuid := server.AddClient(v, rand.Intn(2) == 0, rand.Intn(2) == 0, rand.Intn(2) == 0)

		switch i {
		case len(names)/2 - 1:
			rnd = uuid
		case rnd1[0]:
			rnd1[1] = uuid
		case rnd2[0]:
			rnd2[1] = uuid
		case rnd3[0]:
			rnd3[1] = uuid
		}
	}

	fmt.Println(color.InGreen("Client list:"))
	server.ShowClientList()
	server.ShowTeamMemberList()
	fmt.Print(delim)

	server.SetTeamMember(rnd1[1].(string), "administrator")
	server.SetTeamMember(rnd2[1].(string), "moderator")
	server.SetTeamMember(rnd3[1].(string), "moderator")

	fmt.Println(color.InGreen("Some clients given team member rights:"))
	server.ShowClientList()
	server.ShowTeamMemberList()
	fmt.Print(delim)

	fmt.Println(color.InGreen("Checking IsTeamMember method:"))
	fmt.Printf("%v is %t (expected %t)\n", rnd1[1], server.GetClient(rnd1[1].(string)).IsTeamMember(), true)
	fmt.Printf("%v is %t (expected %t)\n", rnd, server.GetClient(rnd).IsTeamMember(), false)
	fmt.Printf("%v is %t (expected %t)\n", rnd3[1], server.GetClient(rnd3[1].(string)).IsTeamMember(), true)
	fmt.Print(delim)

	server.RevokeTeamMember(rnd1[1].(string))

	fmt.Println(color.InGreen("Revoke admin from client"))
	server.ShowClientList()
	server.ShowTeamMemberList()
	fmt.Print(delim)

	server.DeleteClient(rnd1[1].(string))

	fmt.Println(color.InGreen("Deleting random client (ex admin):"))
	server.ShowClientList()
	server.ShowTeamMemberList()
	fmt.Print(delim)

	server.DeleteClient(rnd)

	fmt.Println(color.InGreen("Deleting random client (not team member):"))
	server.ShowClientList()
	server.ShowTeamMemberList()
	fmt.Print(delim)

	server.DeleteClient(rnd2[1].(string))

	fmt.Println(color.InGreen("Deleting random client (team member):"))
	server.ShowClientList()
	server.ShowTeamMemberList()
	fmt.Print(delim)
}
