package models

type Student struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    int32  `json:"gender"`
	Status    bool   `json:"status"`
}
