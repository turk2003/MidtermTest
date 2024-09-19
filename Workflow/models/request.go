package models

import constants "github.com/turk2003/workflow/constant"

type RequestItem struct {
	Title    string
	Amount   float64
	Quantity uint
}
type RequestFindItem struct {
	Statuses []constants.ItemStatus `form:"status[]"`
}

// internal/model/request.go

type RequestUpdateItem struct {
	Status constants.ItemStatus
}
