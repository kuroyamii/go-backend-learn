package entity

import "time"

type ValorantMember struct {
	Id          uint64    `db:"id"`
	MemberId    uint64    `db:"member_id"`
	Role        string    `db:"role"`
	InGameName  string    `db:"ingame_name"`
	ImagePath   string    `db:"image"`
	Nationality string    `db:"nationality"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type ValorantMembers []ValorantMember
