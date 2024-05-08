// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCjProduct = "cj_product"

// CjProduct 产品
type CjProduct struct {
	ID        uint32  `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Number    string  `gorm:"column:number;type:varchar(15);not null;uniqueIndex:NUMBER_UNIQUE,priority:1;comment:编号" json:"number"` // 编号
	Shortname *string `gorm:"column:shortname;type:varchar(60)" json:"shortname"`
	Fullname  *string `gorm:"column:fullname;type:text" json:"fullname"`
	IsRemoved uint32  `gorm:"column:is_removed;type:tinyint unsigned;not null" json:"is_removed"`
}

// TableName CjProduct's table name
func (*CjProduct) TableName() string {
	return TableNameCjProduct
}
