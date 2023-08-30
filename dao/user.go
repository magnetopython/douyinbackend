package dao

type User struct {
	Id            int
	Name          string
	FollowCount   int
	FollowerCount int
	IsFollow      bool
}
