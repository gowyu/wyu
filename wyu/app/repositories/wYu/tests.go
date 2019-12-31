package wYu

type Tests struct {
	Id   int    `xorm:"not null pk autoincr comment('AutoIdentity') INT(11)" json:"id,omitempty"`
	Name string `xorm:"default '' VARCHAR(50)" json:"name,omitempty"`
}
