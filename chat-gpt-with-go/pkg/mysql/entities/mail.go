package entities

import "gorm.io/gorm"

type Mail struct {
	gorm.Model
    ID            int     `gorm:"column:id; primaryKey; not null"`
    From          string  `gorm:"column:from; size:50; not null"`
    To            string  `gorm:"column:to; size:50; not null"`
    Subject       string  `gorm:"column:subject; size:50; not null"`
    Message       string  `gorm:"column:message; not null"`
    Errors        string  `gorm:"column:errors"`
    NameFrom      string  `gorm:"column:nameFrom; size:50; not null"`
    Type          uint8   `gorm:"column:type"`
    Status        uint8    `gorm:"column:status"`
    DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}