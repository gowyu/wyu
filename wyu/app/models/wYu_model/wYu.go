package wYu_model

import (
	"wyu/app/repositories/wYu"
)

const db string = "wYu" 	// Database Name

type Tests wYu.Tests 		// Table: tests
type TestTest wYu.TestTest 	// Table: test_test

type TestToTest struct {
	Tests `xorm:"extends"`
	TestTest `xorm:"extends"`
}


