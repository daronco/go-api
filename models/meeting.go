package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	MeetingList map[string]*Meeting
)

func init() {
	MeetingList = make(map[string]*Meeting)
	m := Meeting{"meeting1", "My Meeting"}
	// u := Meeting{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	MeetingList["meeting1"] = &m
}

type Meeting struct {
	Id   string
	Name string
}

func Add(m Meeting) string {
	m.Id = "meeting-" + strconv.FormatInt(time.Now().UnixNano(), 10)
	MeetingList[m.Id] = &m
	return m.Id
}

func Get(meetingId string) (m *Meeting, err error) {
	if m, ok := MeetingList[meetingId]; ok {
		return m, nil
	}
	return nil, errors.New("Meeting does not exist")
}

func GetAll() map[string]*Meeting {
	return MeetingList
}

// func UpdateUser(uid string, uu *User) (a *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		if uu.Username != "" {
// 			u.Username = uu.Username
// 		}
// 		if uu.Password != "" {
// 			u.Password = uu.Password
// 		}
// 		if uu.Profile.Age != 0 {
// 			u.Profile.Age = uu.Profile.Age
// 		}
// 		if uu.Profile.Address != "" {
// 			u.Profile.Address = uu.Profile.Address
// 		}
// 		if uu.Profile.Gender != "" {
// 			u.Profile.Gender = uu.Profile.Gender
// 		}
// 		if uu.Profile.Email != "" {
// 			u.Profile.Email = uu.Profile.Email
// 		}
// 		return u, nil
// 	}
// 	return nil, errors.New("User Not Exist")
// }

// func Login(username, password string) bool {
// 	for _, u := range UserList {
// 		if u.Username == username && u.Password == password {
// 			return true
// 		}
// 	}
// 	return false
// }

// func DeleteUser(uid string) {
// 	delete(UserList, uid)
// }
