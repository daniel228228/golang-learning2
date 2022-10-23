package main

import (
	"github.com/google/uuid"
)

type Server struct {
	out            *out
	clientList     map[string]*Client
	teamMemberList map[string]*TeamMember
}

func NewServer(out *out) *Server {
	return &Server{
		out:            out,
		clientList:     map[string]*Client{},
		teamMemberList: make(map[string]*TeamMember),
	}
}

func (s *Server) AddClient(name string, read, write, exec bool) (uuid string) {
	uuid = s.genUUID()
	s.clientList[uuid] = NewClient(uuid, name, read, write, exec)
	return
}

func (s *Server) GetClient(uuid string) *Client {
	if value, ok := s.clientList[uuid]; ok {
		return value
	} else {
		return nil
	}
}

func (s *Server) DeleteClient(uuid string) {
	if value, ok := s.clientList[uuid]; ok {
		if value.IsTeamMember() {
			delete(s.teamMemberList, uuid)
		}
		delete(s.clientList, uuid)
	}
}

func (s *Server) SetTeamMember(uuid string, role string) {
	if value, ok := s.clientList[uuid]; ok {
		teamMember := &TeamMember{Client: value, role: role}
		s.teamMemberList[uuid] = teamMember
		value.read = true
		value.write = true
		value.exec = true
		value.teamMember = teamMember
	}
}

func (s *Server) RevokeTeamMember(uuid string) {
	if value, ok := s.teamMemberList[uuid]; ok {
		delete(s.teamMemberList, uuid)
		value.teamMember = nil
		value.read = true
		value.write = false
		value.exec = false
	}
}

func (s *Server) genUUID() string {
	return uuid.New().String()
}

func (s *Server) ShowClientList() {
	s.out.Println("Client list:")

	for _, v := range s.clientList {
		var checkTeamMember string

		if v.IsTeamMember() {
			checkTeamMember = "[TEAM MEMBER] "
		}

		s.out.Printf("  %s%s: name: %s permissions: %v\n", checkTeamMember, v.uuid, v.name, v.Permissions())
	}
}

func (s *Server) ShowTeamMemberList() {
	s.out.Println("Team member list:")

	if len(s.teamMemberList) == 0 {
		s.out.Println("  (empty)")
	}

	for _, v := range s.teamMemberList {
		s.out.Printf("  %s (%s): name: %s permissions: %v\n", v.uuid, v.role, v.name, v.Permissions())
	}
}
