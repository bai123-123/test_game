package model

import "gorm.io/gorm"

type Map struct {
	gorm.Model
	Rules string
}

type Rule struct {
	Snake  map[int]int `json:"snake"`
	Ladder map[int]int `json:"ladder"`
}

func CreateMap(rules string) Map {
	var mmap Map
	mmap.Rules = rules
	db.Create(&mmap)
	return mmap
}
