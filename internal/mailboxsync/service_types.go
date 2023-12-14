package mailboxsync

import "strconv"

type Sync struct {
	Src       ImapServer `json:"src" validate:"required"`
	Dst       ImapServer `json:"dst" validate:"required"`
	Mailboxes []Mailbox  `json:"mailboxes" validate:"required"`
}

type ImapServer struct {
	Host string `json:"host" validate:"required"`
	Port *int   `json:"port" default:"143"`
}

func (i ImapServer) getHost() string {
	return i.Host
}

func (i ImapServer) getPort() string {
	if i.Port != nil {
		return strconv.Itoa(*i.Port)
	}

	return "143"
}

type Mailbox struct {
	User        string `json:"user" validate:"required"`
	Password    string `json:"password" validate:"required"`
	SrcPassword string `json:"srcPassword"`
	SrcUser     string `json:"srcUser"`
	DstPassword string `json:"dstPassword"`
	DstUser     string `json:"dstUser"`
}

func (m Mailbox) GetSrcUser() string {
	if m.SrcUser != "" {
		return m.SrcUser
	}

	return m.User
}

func (m Mailbox) GetSrcPassword() string {
	if m.SrcPassword != "" {
		return m.SrcPassword
	}

	return m.Password
}

func (m Mailbox) GetDstUser() string {
	if m.DstUser != "" {
		return m.DstUser
	}

	return m.User
}

func (m Mailbox) GetDstPassword() string {
	if m.DstPassword != "" {
		return m.DstPassword
	}

	return m.Password
}
