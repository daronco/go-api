package models

import (
	// "fmt"
	// "strconv"
	// "time"
)

type Meeting struct {
	Id        string `json:"meeting_id"`
	Name      string `json:"name"`
	UpdatedAt string `json:"updated_at"`
}

// func (this *Meeting) Save() (bool, error) {
// 	m := *this
// 	m.UpdatedAt = strconv.FormatInt(time.Now().UnixNano(), 10)
// 	reply, err := bridge.Save(m.Id, m)
// 	*this = m
// 	return reply, err
// }
