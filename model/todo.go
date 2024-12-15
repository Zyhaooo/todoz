package model

type Todo struct {
	Id             uint64 // Todo id
	Title          string // Todo 标题
	Description    string // Todo 详细说明
	Level          int    // 等级
	CreatedTime    string // 创建时间
	ExpirationTime string // 超时时间
	Finish         bool   // 是否已经完成
}
