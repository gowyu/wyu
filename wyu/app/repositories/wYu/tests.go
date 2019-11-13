package wYu

type Tests struct {
	Id   int    `xorm:"not null pk autoincr comment('AutoIdentity') INT(11)"`
	Name string `xorm:"default '' VARCHAR(50)"`
}
