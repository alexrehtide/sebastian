package model

type Platform string

const (
	Google Platform = "google"
	Yandex Platform = "yandex"
	Twitch Platform = "twitch"
)

type RemoteAccount struct {
	ID          uint
	AccountID   uint
	RemoteID    string
	RemoteEmail string
	Platform    Platform
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
