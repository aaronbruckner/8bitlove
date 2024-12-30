package config

const APP_DATA_CACHE = ".localAppDataCache"

type Event struct {
	Key     string
	Title   string
	Date    string
	Summary string
}

type AppData struct {
	Events []Event
}
