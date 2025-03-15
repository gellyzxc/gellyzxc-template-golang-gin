package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int       `json:"id" bun:",pk,autoincrement"`
	Name      string    `json:"name" bun:"name,notnull"  binding:"required"`
	Email     string    `json:"email" bun:"email,unique,notnull"  binding:"required"`
	Password  string    `json:"password" bun:"password,notnull" binding:"required"`
	CreatedAt time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.Password = u.Password + "_hashed"
	}
	return nil
}
