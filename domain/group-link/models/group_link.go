package models

type GroupLink struct {
	ID        uint   `json:"group_id" gorm:"column:group_id"`
	GroupName string `json:"group_name" gorm:"column:group_name"`
}

type APIGroupLink struct {
	ID        uint   `json:"group_id" gorm:"column:group_id"`
	GroupName string `json:"group_name" gorm:"column:group_name"`
}

func (GroupLink) TableName() string {
	return "tb_groups"
}

func (APIGroupLink) TableName() string {
	return "tb_groups"
}
