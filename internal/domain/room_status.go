package model

type RoomStatus string

const (
	Created  RoomStatus = "CREATED"
	Starting RoomStatus = "STARTING"
	Started  RoomStatus = "STARTED"
	Voting   RoomStatus = "VOTING"
	Finished RoomStatus = "FINISHED"
)
