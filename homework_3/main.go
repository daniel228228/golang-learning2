package main

import (
	"math/rand"
	"strings"
	"time"
)

func main() {
	delim := "\n" + strings.Repeat("=", 30) + "\n"
	names := []string{"James", "Mary", "Robert", "Patricia", "John", "Jennifer", "Michael", "Linda", "David", "Elizabeth", "William", "Barbara", "Richard", "Susan", "Joseph", "Jessica", "Thomas", "Sarah", "Charles", "Karen"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
	rnd, rnd1, rnd2, rnd3 := "", [2]any{rand.Intn(len(names)/2) - 1}, [2]any{rand.Intn(len(names)/4) + len(names)/2}, [2]any{len(names) - 1 - rand.Intn(len(names)/4)}

	output := InitOutputFile("output.log")
	defer output.Close()

	out := NewOut(output)

	server := NewServer(out)

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

	out.Println("Client list:\n")
	server.ShowClientList()
	server.ShowTeamMemberList()
	out.Println(delim)

	server.SetTeamMember(rnd1[1].(string), "administrator")
	server.SetTeamMember(rnd2[1].(string), "moderator")
	server.SetTeamMember(rnd3[1].(string), "moderator")

	out.Println("Some clients given team member rights:\n")
	server.ShowClientList()
	server.ShowTeamMemberList()
	out.Println(delim)

	out.Println("Checking IsTeamMember method:\n")
	out.Printf("%v is %t (expected %t)\n", rnd1[1], server.GetClient(rnd1[1].(string)).IsTeamMember(), true)
	out.Printf("%v is %t (expected %t)\n", rnd, server.GetClient(rnd).IsTeamMember(), false)
	out.Printf("%v is %t (expected %t)\n", rnd3[1], server.GetClient(rnd3[1].(string)).IsTeamMember(), true)
	out.Println(delim)

	server.RevokeTeamMember(rnd1[1].(string))

	out.Println("Revoke admin from client\n")
	server.ShowClientList()
	server.ShowTeamMemberList()
	out.Println(delim)

	server.DeleteClient(rnd1[1].(string))

	out.Println("Deleting random client (ex admin):\n")
	server.ShowClientList()
	server.ShowTeamMemberList()
	out.Println(delim)

	server.DeleteClient(rnd)

	out.Println("Deleting random client (not team member):\n")
	server.ShowClientList()
	server.ShowTeamMemberList()
	out.Println(delim)

	server.DeleteClient(rnd2[1].(string))

	out.Println("Deleting random client (team member):\n")
	server.ShowClientList()
	server.ShowTeamMemberList()
	out.Println(delim)
}
