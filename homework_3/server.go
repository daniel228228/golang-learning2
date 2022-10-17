package main

import (
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/google/uuid"
)

type Server struct {
	clientList     map[string]/* TODO: complete me */
	teamMemberList map[string]/* TODO: complete me */
}

func NewServer() *Server {
	return &Server{
		clientList:     /* TODO: complete me */
		teamMemberList: /* TODO: complete me */
	}
}

func (s *Server) AddClient(name string, read, write, exec bool) (uuid string) {
	// TODO: implement me
}

// TODO: implement GetClient method

func (s *Server) DeleteClient(uuid string) {
	// TODO: implement me
}

func (s *Server) SetTeamMember(uuid string, role string) {
	// TODO: implement me
}

// TODO: implement RevokeTeamMember method

func (s *Server) genUUID() string {
	return uuid.New().String()
}

func (s *Server) ShowClientList() {
	fmt.Println("Client list:")

	for _, v := range s.clientList {
		var checkTeamMember string

		if v.IsTeamMember() {
			checkTeamMember = color.InRed("[TEAM MEMBER] ")
		}

		fmt.Printf("  %s%s: name: %s permissions: %v\n", checkTeamMember, v.uuid, v.Permissions())
	}
}

func (s *Server) ShowTeamMemberList() {
	fmt.Println("Team member list:")

	if len(s.teamMemberList) == 0 {
		fmt.Println("  (empty)")
	}

	for _, v := range s.teamMemberList {
		fmt.Printf("  %s (%s): name: %s permissions: %v\n", v.uuid, v.role, v.name, v.Permissions())
	}
}
