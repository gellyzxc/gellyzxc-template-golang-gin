package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Post struct {
	bun.BaseModel `bun:"table:posts"`

	ID        int64     `json:"id" bun:",pk,autoincrement"`
	Name      string    `json:"name" bun:"name,unique" binding:"required"`
	Text      string    `json:"text" bun:"text" binding:"required"`
	CreatedAt time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UserID    int64     `json:"user_id" bun:"user_id"`
	User      *User     `json:"user" bun:"rel:belongs-to,join:user_id=id"`
}
