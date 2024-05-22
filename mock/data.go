package mock

import (
	"time"

	"privia-staj-backend-todo/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Database connection failed!")
	}
	DB = db
	DB.Migrator().DropTable(&models.TodoList{}, &models.TodoItem{}, &models.User{})
	DB.AutoMigrate(&models.TodoList{}, &models.TodoItem{}, &models.User{})
	LoadMockData()
}

var Users = map[string]models.User{
	"user":    {ID: 1, Username: "user", Password: "user", UserType: "user"},
	"admin":   {ID: 2, Username: "admin", Password: "admin", UserType: "admin"},
	"newuser": {ID: 3, Username: "newuser", Password: "password", UserType: "normal"},
}

func LoadMockData() {
	for _, user := range Users {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
		DB.Create(&user)
	}

	DB.Create(&[]models.TodoList{
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: "Shopping List", CompletionRate: 0.5, UserID: 1},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: "Work List", CompletionRate: 0.2, UserID: 1},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: "Vacation Plan", CompletionRate: 0.0, UserID: 2},
	})

	DB.Create(&[]models.TodoItem{
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), TodoListID: 1, Content: "Milk", Completed: true},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), TodoListID: 1, Content: "Eggs", Completed: false},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), TodoListID: 2, Content: "Prepare report", Completed: false},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), TodoListID: 2, Content: "Meeting", Completed: true},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), TodoListID: 3, Content: "Buy plane ticket", Completed: false},
	})
}
