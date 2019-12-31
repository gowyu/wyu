package wYu

type TestTest struct {
	Id      int    `xorm:"not null pk autoincr INT(11)" json:"id,omitempty"`
	TestId  int    `xorm:"index INT(11)" json:"test_id,omitempty"`
	Content string `xorm:"VARCHAR(50)" json:"content,omitempty"`
}
