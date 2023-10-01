package model

type Player struct {
	Session string  `json:"session_id"`
	Name    string  `json:"name"`
	Avatar  string  `json:"avatar"`
	Rounds  []Round `json:"rounds"`
}
