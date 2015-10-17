package main
import (
	"net/http"
	"encoding/json"
)

type Settings struct {
	// users configured
	Users map[string]User
	// require client to log in
	RequireUser bool

	// ===== source settings =====
	// what.cd settings
	WhatCD	*WhatCDSetings

	// === downloader settings ===
	// transmission settings
	TransmissionBT	*TransmissionBTSettings

	// === automation settings ===
	// move files after download
	MoveAfterDownload	bool
	// use github.com/go-fsnotify/fsnotify
	UseNotify		bool
	// if not, how often to poll
	PollInterval	int
	// root of music library
	MoveRoot		string
	// match files with
}

type User struct {
	Password	string
}

type WhatCDSetings struct {}

type TransmissionBTSettings struct {}

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	users := make(map[string]User)
	users["Peter"] = User{
		Password: "p@$$word",
	}
	settings := Settings{
		Users:				users,
		RequireUser: 		false,
		WhatCD:				&WhatCDSetings{},
		TransmissionBT:		&TransmissionBTSettings{},
		MoveAfterDownload:	false,
		UseNotify:			false,
		PollInterval:		300,
		MoveRoot:			"/tmp",
	}
	b, err := json.Marshal(settings)
	if err != nil {
		panic(err)
	}
    w.Write(b)
}
