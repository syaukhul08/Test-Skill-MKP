package models

import "gorm.io/gorm"

type Terminal struct {
	gorm.Model
	TerminalId       int    `json:"terminal_id" gorm:"primary_key"`
	Name     string `json:"name"`
	Location string `json:"location"`
}