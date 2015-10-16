package main
type settings struct {
	// users configured
	users map[*user]bool
	// require client to log in
	requireUser bool

	// ===== source settings =====
	// what.cd settings
	whatCD	*whatCDSetings

	// === downloader settings ===
	// transmission settings
	transmissionBT	*transmissionBTSettings

	// === automation settings ===
	// move files after download
	moveAfterDownload	bool
	// use github.com/go-fsnotify/fsnotify
	useNotify		bool
	// if not, how often to poll
	pollInterval	int
	// root of music library
	moveRoot		string
	// match files with 
}

type user struct {
	name		string
	password	string
}

type whatCDSetings struct {}

type transmissionBTSettings struct {}