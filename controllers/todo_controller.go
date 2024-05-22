package controllers

import (
	"strconv"
	"time"

	"privia-staj-backend-todo/mock"
	"privia-staj-backend-todo/models"
	"privia-staj-backend-todo/utils"

	"github.com/gin-gonic/gin"
)

// GetTodos retrieves all todo lists (admin sees all lists, user sees only their own lists)
func GetTodos(c *gin.Context) {
	var todos []models.TodoList
	userType := c.GetString("userType")

	if userType == "admin" {
		mock.DB.Unscoped().Where("deleted_at IS NULL").Find(&todos)
	} else {
		userID := c.GetUint("userID")
		mock.DB.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&todos)
	}

	utils.OK(c, "Todos retrieved successfully", todos)
}

func CreateTodoList(c *gin.Context) {
	var newTodoList models.TodoList
	if err := c.ShouldBindJSON(&newTodoList); err != nil {
		utils.BadRequest(c, "Invalid input", err)
		return
	}

	userID := c.GetUint("userID")
	newTodoList.UserID = userID
	newTodoList.CreatedAt = time.Now()
	newTodoList.UpdatedAt = time.Now()

	mock.DB.Create(&newTodoList)
	utils.Created(c, "TodoList created successfully", newTodoList)
}

// GetTodo retrieves a specific todo list by ID
func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", id).First(&todoList).Error; err != nil {
		utils.NotFound(c, "TodoList not found", err)
		return
	}

	userID := c.GetUint("userID")
	userType := c.GetString("userType")
	if todoList.UserID != userID && userType != "admin" {
		utils.Forbidden(c, "You do not have access to this list", nil)
		return
	}

	utils.OK(c, "TodoList retrieved successfully", todoList)
}

// UpdateTodoList updates a specific todo list by ID
func UpdateTodoList(c *gin.Context) {
	id := c.Param("id")
	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", id).First(&todoList).Error; err != nil {
		utils.NotFound(c, "TodoList not found", err)
		return
	}

	userID := c.GetUint("userID")
	if todoList.UserID != userID {
		utils.Forbidden(c, "You do not have permission to update this list", nil)
		return
	}

	var updatedTodoList models.TodoList
	if err := c.ShouldBindJSON(&updatedTodoList); err != nil {
		utils.BadRequest(c, "Invalid input", err)
		return
	}

	updatedTodoList.ID = todoList.ID
	updatedTodoList.UserID = todoList.UserID
	updatedTodoList.CreatedAt = todoList.CreatedAt
	updatedTodoList.UpdatedAt = time.Now()

	mock.DB.Save(&updatedTodoList)
	utils.OK(c, "TodoList updated successfully", updatedTodoList)
}

// DeleteTodoList soft deletes a specific todo list by ID
func DeleteTodoList(c *gin.Context) {
	id := c.Param("id")
	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", id).First(&todoList).Error; err != nil {
		utils.NotFound(c, "TodoList not found", err)
		return
	}

	userID := c.GetUint("userID")
	if todoList.UserID != userID {
		utils.Forbidden(c, "You do not have permission to delete this list", nil)
		return
	}

	deletedAt := time.Now()
	todoList.DeletedAt = &deletedAt
	mock.DB.Save(&todoList)

	utils.OK(c, "TodoList deleted successfully", nil)
}

// GetItems retrieves all items in a specific todo list
func GetItems(c *gin.Context) {
	todoIDStr := c.Param("id")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		utils.BadRequest(c, "Invalid todo ID", err)
		return
	}

	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", todoID).First(&todoList).Error; err != nil {
		utils.NotFound(c, "TodoList not found", err)
		return
	}

	userID := c.GetUint("userID")
	if todoList.UserID != userID {
		utils.Forbidden(c, "You do not have access to this list", nil)
		return
	}

	var todoItems []models.TodoItem
	mock.DB.Where("todo_list_id = ? AND deleted_at IS NULL", todoID).Find(&todoItems)
	utils.OK(c, "TodoItems retrieved successfully", todoItems)
}

// CreateTodoItem creates a new item in a specific todo list
func CreateTodoItem(c *gin.Context) {
	todoIDStr := c.Param("id")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		utils.BadRequest(c, "Invalid todo ID", err)
		return
	}

	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", todoID).First(&todoList).Error; err != nil {
		utils.NotFound(c, "TodoList not found", err)
		return
	}

	userID := c.GetUint("userID")
	if todoList.UserID != userID {
		utils.Forbidden(c, "You do not have access to this list", nil)
		return
	}

	var newTodoItem models.TodoItem
	if err := c.ShouldBindJSON(&newTodoItem); err != nil {
		utils.BadRequest(c, "Invalid input", err)
		return
	}

	newTodoItem.TodoListID = uint(todoID)
	newTodoItem.CreatedAt = time.Now()
	newTodoItem.UpdatedAt = time.Now()

	mock.DB.Create(&newTodoItem)
	utils.Created(c, "TodoItem created successfully", newTodoItem)
}

// GetItem retrieves a specific item in a todo list
func UpdateTodoItem(c *gin.Context) {
	todoID := c.Param("todoId") // Parametre adı ':id' yerine ':todoId' olarak değiştirildi
	itemID := c.Param("itemId") // Parametre adı ':item_id' yerine ':itemId' olarak değiştirildi

	var todoItem models.TodoItem
	if err := mock.DB.Where("id = ? AND todo_list_id = ? AND deleted_at IS NULL", itemID, todoID).First(&todoItem).Error; err != nil {
		utils.NotFound(c, "TodoItem not found", err)
		return
	}

	userID := c.GetUint("userID")
	if todoItem.UserID != userID {
		utils.Forbidden(c, "You do not have permission to update this item", nil)
		return
	}

	var updatedTodoItem models.TodoItem
	if err := c.ShouldBindJSON(&updatedTodoItem); err != nil {
		utils.BadRequest(c, "Invalid input", err)
		return
	}

	updatedTodoItem.ID = todoItem.ID
	updatedTodoItem.TodoListID = todoItem.TodoListID
	updatedTodoItem.CreatedAt = todoItem.CreatedAt
	updatedTodoItem.UpdatedAt = time.Now()

	mock.DB.Save(&updatedTodoItem)

	utils.OK(c, "TodoItem updated successfully", updatedTodoItem)
}

// DeleteTodoItem soft deletes a specific item in a todo list
func DeleteTodoItem(c *gin.Context) {
	todoID := c.Param("todoId") // Parametre adı ':id' yerine ':todoId' olarak değiştirildi
	itemID := c.Param("itemId") // Parametre adı ':item_id' yerine ':itemId' olarak değiştirildi

	var todoItem models.TodoItem
	if err := mock.DB.Where("id = ? AND todo_list_id = ? AND deleted_at IS NULL", itemID, todoID).First(&todoItem).Error; err != nil {
		utils.NotFound(c, "TodoItem not found", err)
		return
	}

	userID := c.GetUint("userID")
	if todoItem.UserID != userID {
		utils.Forbidden(c, "You do not have permission to delete this item", nil)
		return
	}

	deletedAt := time.Now()
	todoItem.DeletedAt = &deletedAt
	mock.DB.Save(&todoItem)

	utils.OK(c, "TodoItem deleted successfully", nil)
}
