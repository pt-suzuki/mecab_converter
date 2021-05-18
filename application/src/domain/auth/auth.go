package auth

import (
	nullAble "gopkg.in/guregu/null.v4"
)

type MasterSchedule struct {
	Id nullAble.Int
	Name string
	GroupId nullAble.Int
	ProductId nullAble.Int
	Priority nullAble.Int
	TrademarkId nullAble.Int
	Status nullAble.Int
	JsonDetail string
	CreatedUserId nullAble.Int
	UpdatedUserId nullAble.Int
	CreatedAt nullAble.Time
	UpdatedAt nullAble.Time
}
