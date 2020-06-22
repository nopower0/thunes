package models

import (
	"time"
)

type AbstractTimeModel struct {
	AddTime    time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
