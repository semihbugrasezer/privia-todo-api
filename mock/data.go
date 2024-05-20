package mock

import (
	"time"

	"github.com/semihbugrasezer/privia-todo-api/models"
	"gorm.io/driver/sqlite" // SQLite veritabanı sürücüsü
	"gorm.io/gorm"
)

// Mock veritabanı bağlantısı
var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("Veritabanı bağlantısı başarısız!")
	}
	DB = db
	DB.AutoMigrate(&models.TodoList{}, &models.TodoItem{}) // Tabloları otomatik oluştur
	LoadMockData()
}

var Users = map[string]models.User{
	"user1": {ID: 1, Username: "user1", Password: "password", Type: "user"},
	"admin": {ID: 2, Username: "admin", Password: "password", Type: "admin"},
}

func LoadMockData() {

	DB.Create(&[]models.TodoList{
		{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: "Alışveriş Listesi", CompletionRate: 0.5, UserID: 1},
		{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: "İş Listesi", CompletionRate: 0.2, UserID: 1},
		{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: "Tatil Planı", CompletionRate: 0.0, UserID: 2},
	})

	DB.Create(&[]models.TodoItem{
		{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), TodoListID: 1, Content: "Süt", Completed: true},
		{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now(), TodoListID: 1, Content: "Yumurta", Completed: false},
		{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now(), TodoListID: 2, Content: "Rapor hazırla", Completed: false},
		{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now(), TodoListID: 2, Content: "Toplantı", Completed: true},
		{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now(), TodoListID: 3, Content: "Uçak bileti al", Completed: false},
	})
}
