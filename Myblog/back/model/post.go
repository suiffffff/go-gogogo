package model

import "time"

// Post 文章结构体
type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"` // 摘要
	Content   string    `json:"content"` // 内容 (Markdown)
	CreatedAt time.Time `json:"created_at"`
	Comments  []Comment // 一篇文章有多条评论
}
