package models

type SessionSummary struct {
	TotalSessions int `json:"total_sessions"`
	TotalDistance int `json:"total_distance"`
	TotalTime     int `json:"total_time"`
}

type SwimmerSummary struct {
	Swimmer       Swimmer `json:"swimmer"`
	TotalSessions int     `json:"total_sessions"`
	TotalDistance int     `json:"total_distance"`
	TotalTime     int     `json:"total_time"`
}
