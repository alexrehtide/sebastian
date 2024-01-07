package model

type RequestReceived struct {
	IP        string
	SessionID uint
	AccountID uint
	Path      string
	Body      string
}
