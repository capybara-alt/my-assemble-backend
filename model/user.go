package model

import (
	"time"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
)

type User struct {
	ID        string `gorm:"primaryKey;index:user_idx"`
	Name      string
	Icon      string
	CreatedAt time.Time
}

func (u *User) ToPB() *myassemblyv1.User {
	return &myassemblyv1.User{
		Id:        u.ID,
		Name:      u.Name,
		Icon:      u.Icon,
		CreatedAt: u.CreatedAt.UnixMilli(),
	}
}

func (u *User) FromPB(user *myassemblyv1.User) *User {
	return &User{
		ID:        user.Id,
		Name:      user.Name,
		Icon:      user.Icon,
		CreatedAt: time.UnixMilli(user.CreatedAt),
	}
}
