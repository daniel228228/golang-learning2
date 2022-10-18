package main

import (
  "github.com/google/uuid"
)

type Server struct {
  out            *out
  clientList     map[string]/* TODO: complete me */
  teamMemberList map[string]/* TODO: complete me */
}

func NewServer(out *out) *Server {
  return &Server{
    out:            out,
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
