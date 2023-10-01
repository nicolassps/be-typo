package model

type Round struct {
	Number int    `json:"number"`
	Letter string `json:"letter"`
	Words  []Word `json:"words"`
}
