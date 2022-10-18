package main

type Client struct {
  /* TODO: complete me */
  uuid              string
  name              string
  read, write, exec bool
}

func NewClient(uuid, name string, read, write, exec bool) *Client {
  return &Client{
    uuid:  uuid,
    name:  name,
    read:  read,
    write: write,
    exec:  exec,
  }
}

func (c *Client) Permissions() []bool {
  return []bool{c.read, c.write, c.exec}
}

func (c *Client) IsTeamMember() bool {
  return c.teamMember != /* TODO: complete me */
}
