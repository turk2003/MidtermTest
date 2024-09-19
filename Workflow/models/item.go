package models

import constants "github.com/turk2003/workflow/constant"

type Item struct {
	ID       uint                 `gorm:"primaryKey"`
	Title    string               `json:"title"`
	Amount   int                  `json:"amount"`
	Quantity int                  `json:"quantity"`
	Status   constants.ItemStatus `json:"status"`
	OwnerID  uint                 `json:"owner_id"`
}
