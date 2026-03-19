package entity

import "time"

const TableNameSong = "public.Song"

type Song struct {
	Id         int       `gorm:"column:Id" json:"id"`
	AlbumId    int       `gorm:"column:AlbumId" json:"albumId"`
	Name       string    `gorm:"column:Name" json:"name"`
	CreateBy   string    `gorm:"column:CreateBy" json:"createBy"`
	CreateDate time.Time `gorm:"column:CreateDate" json:"createDate"`
	UpdateBy   string    `gorm:"column:UpdateBy" json:"updateBy"`
	UpdateDate time.Time `gorm:"column:UpdateDate" json:"updateDate"`
	IsDelete   bool      `gorm:"column:IsDelete" json:"isDelete"`
	Album      Album     `gorm:"foreignKey:AlbumId;references:Id" json:"album"`
}

func (*Song) TableName() string {
	return TableNameSong
}
