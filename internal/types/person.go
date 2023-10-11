package types

type Person struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Income float64 `json:"income"`
	User   User    `json:"user"`
}
