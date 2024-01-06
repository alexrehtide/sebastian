package model

type Platform string

const (
	Google Platform = "google"
	Yandex Platform = "yandex"
	Twitch Platform = "twitch"
)

type RemoteAccount struct {
	ID          uint     `db:"id"`
	AccountID   uint     `db:"account_id"`
	RemoteID    string   `db:"remote_id"`
	RemoteEmail string   `db:"remote_email"`
	Platform    Platform `db:"platform"`
}
type CreateRemoteAccountOptions struct {
	AccountID   uint
	RemoteID    string
	RemoteEmail string
	Platform    Platform
}

type ReadRemoteAccountOptions struct {
	ID          uint
	AccountID   uint
	RemoteID    string
	RemoteEmail string
	Platform    Platform
}

type UpdateRemoteAccountOptions struct {
	AccountID uint
}
