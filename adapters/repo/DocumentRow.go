package repo

type DocumentRow struct {
	Id          int    `gorm:"primaryKey; autoIncrement"`
	Name        string `gorm:"not ull"`
	Description string
}
