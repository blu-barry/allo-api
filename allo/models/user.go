package models

import {
	"allo/models/item"
}

type User struct {
	userId     string
	userEmail  string
	userName   string
	firstName  string
	lastName   string
	items      []Item
	itemIds    []string
	sessionIds []string
}