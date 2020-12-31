import "time"

type UsableFunction struct {
	ID uint `gorm:"primaryKey;column:id;autoIncrement"`
	Content string `gorm:"column:content"`
	UserId uint
	User User
	CreatedAt time.Time `gorm:"column:created_at"`
}