package entity

import "time"

const TableNameAlbum = "public.Album"

type Album struct {
	Id          int       `gorm:"column:Id" json:"id"`
	Name        string    `gorm:"column:Name" json:"name"`
	FileId      int       `gorm:"column:FileId" json:"fileId"`
	Description string    `gorm:"column:Description" json:"description"`
	CreateBy    string    `gorm:"column:CreateBy" json:"createBy"`
	CreateDate  time.Time `gorm:"column:CreateDate" json:"createDate"`
	UpdateBy    string    `gorm:"column:UpdateBy" json:"updateBy"`
	UpdateDate  time.Time `gorm:"column:UpdateDate" json:"updateDate"`
	IsDelete    bool      `gorm:"column:IsDelete" json:"isDelete"`
	Songs       []Song    `gorm:"foreignKey:AlbumId;references:Id" json:"songs"`
}

func (*Album) TableName() string {
	return TableNameAlbum
}
