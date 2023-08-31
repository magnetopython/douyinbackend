package dao

type User struct {
	Id            int    `gorm:"column:user_id" json:"id"`
	Name          string `gorm:"column:name" json:"name" `
	FollowCount   int    `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount int    `gorm:"column:follower_count" json:"follower_count"`
	IsFollow      bool   `gorm:"column:is_follow" json:"is_follow"`
	Password      string `gorm:"column:password"`
}

func (User) TableName() string {
	return "user"
}
