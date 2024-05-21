package controllers

import (
	"net/http"
	"privia-staj-backend-todo/mock"
	"privia-staj-backend-todo/models"
	"privia-staj-backend-todo/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetTodos tüm todo listelerini getirir (admin tüm listeleri, kullanıcı sadece kendi listelerini görür)
func GetTodos(c *gin.Context) {
	var todos []models.TodoList
	userType := c.GetString("userType")

	if userType == "admin" {
		mock.DB.Unscoped().Where("deleted_at IS NULL").Find(&todos)
	} else {
		userID := c.GetUint("userID")
		mock.DB.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&todos)
	}

	utils.RespondJSON(c, http.StatusOK, todos)
}

// CreateTodoList yeni bir todo listesi oluşturur
func CreateTodoList(c *gin.Context) {
	var newTodoList models.TodoList
	if err := c.ShouldBindJSON(&newTodoList); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID := c.GetUint("userID")
	newTodoList.UserID = userID
	newTodoList.CreatedAt = time.Now()
	newTodoList.UpdatedAt = time.Now()

	mock.DB.Create(&newTodoList)
	utils.RespondJSON(c, http.StatusCreated, newTodoList)
}

// GetTodo belirli bir ID'ye sahip todo listesini getirir
func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", id).First(&todoList).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "TodoList bulunamadı")
		return
	}

	userID := c.GetUint("userID")
	userType := c.GetString("userType")
	if todoList.UserID != userID && userType != "admin" {
		utils.RespondError(c, http.StatusForbidden, "Bu listeye erişim izniniz yok")
		return
	}

	utils.RespondJSON(c, http.StatusOK, todoList)
}

// UpdateTodoList belirli bir ID'ye sahip todo listesini günceller
func UpdateTodoList(c *gin.Context) {
	var updatedTodoList models.TodoList
	if err := c.ShouldBindJSON(&updatedTodoList); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", id).First(&todoList).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "TodoList bulunamadı")
		return
	}

	// Yetkilendirme kontrolü
	userID := c.GetUint("userID")
	if todoList.UserID != userID {
		utils.RespondError(c, http.StatusForbidden, "Bu listeyi güncelleme izniniz yok")
		return
	}

	todoList.Name = updatedTodoList.Name
	todoList.UpdatedAt = time.Now()
	mock.DB.Save(&todoList)

	utils.RespondJSON(c, http.StatusOK, todoList)
}

// DeleteTodoList belirli bir ID'ye sahip todo listesini siler (soft delete)
func DeleteTodoList(c *gin.Context) {
	id := c.Param("id")
	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", id).First(&todoList).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "TodoList bulunamadı")
		return
	}

	// Yetkilendirme kontrolü
	userID := c.GetUint("userID")
	if todoList.UserID != userID {
		utils.RespondError(c, http.StatusForbidden, "Bu listeyi silme izniniz yok")
		return
	}

	deletedAt := time.Now() // Soft delete
	todoList.DeletedAt = &deletedAt
	mock.DB.Save(&todoList)

	utils.RespondJSON(c, http.StatusOK, "TodoList silindi")
}

// GetItems todo listesindeki tüm öğeleri getirir
func GetItems(c *gin.Context) {
	todoIDStr := c.Param("todoId")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Geçersiz todoID")
		return
	}

	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", todoID).First(&todoList).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "TodoList bulunamadı")
		return
	}

	// Yetkilendirme kontrolü
	userID := c.GetUint("userID")
	userType := c.GetString("userType")
	if todoList.UserID != userID && userType != "admin" {
		utils.RespondError(c, http.StatusForbidden, "Bu listeye erişim izniniz yok")
		return
	}

	var todoItems []models.TodoItem
	mock.DB.Where("todo_list_id = ? AND deleted_at IS NULL", todoID).Find(&todoItems)
	utils.RespondJSON(c, http.StatusOK, todoItems)
}

// CreateTodoItem todo listesine yeni bir öğe ekler
func CreateTodoItem(c *gin.Context) {
	var newTodoItem models.TodoItem
	if err := c.ShouldBindJSON(&newTodoItem); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	todoIDStr := c.Param("todoId")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Geçersiz todoID")
		return
	}
	// Liste var mı kontrolü
	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", todoID).First(&todoList).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "TodoList bulunamadı")
		return
	}
	// Yetkilendirme kontrolü
	userID := c.GetUint("userID")
	if todoList.UserID != userID {
		utils.RespondError(c, http.StatusForbidden, "Bu listeye erişim izniniz yok")
		return
	}

	newTodoItem.TodoListID = uint(todoID)
	newTodoItem.CreatedAt = time.Now()
	newTodoItem.UpdatedAt = time.Now()

	mock.DB.Create(&newTodoItem)
	utils.RespondJSON(c, http.StatusCreated, newTodoItem)
}

// GetItem todo listesindeki belirli bir öğeyi getirir
func GetItem(c *gin.Context) {
	todoIDStr := c.Param("todoId")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Geçersiz todoID")
		return
	}

	itemIDStr := c.Param("itemId")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Geçersiz öğe ID'si")
		return
	}

	// Todo listesini bul ve yetkilendirme kontrolü yap
	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", todoID).First(&todoList).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "TodoList bulunamadı")
		return
	}

	userID := c.GetUint("userID")
	userType := c.GetString("userType")
	if todoList.UserID != userID && userType != "admin" {
		utils.RespondError(c, http.StatusForbidden, "Bu listeye erişim izniniz yok")
		return
	}

	var todoItem models.TodoItem
	if err := mock.DB.Where("id = ? AND todo_list_id = ? AND deleted_at IS NULL", itemID, todoID).First(&todoItem).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "TodoItem bulunamadı")
		return
	}

	utils.RespondJSON(c, http.StatusOK, todoItem)
}

// UpdateTodoItem belirli bir ID'ye sahip todo öğesini günceller
// UpdateTodoItem belirli bir ID'ye sahip todo öğesini günceller
func UpdateTodoItem(c *gin.Context) {
	todoIDStr := c.Param("todoId")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Geçersiz todoID")
		return
	}

	itemIDStr := c.Param("itemId")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Geçersiz öğe ID'si")
		return
	}

	var updatedTodoItem models.TodoItem
	if err := c.ShouldBindJSON(&updatedTodoItem); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var todoItem models.TodoItem
	if err := mock.DB.Where("id = ? AND todo_list_id = ? AND deleted_at IS NULL", itemID, todoID).First(&todoItem).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "TodoItem bulunamadı")
		return
	}

	// Yetkilendirme kontrolü
	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", todoID).First(&todoList).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "TodoList bulunamadı")
		return
	}
	userID := c.GetUint("userID")
	if todoList.UserID != userID {
		utils.RespondError(c, http.StatusForbidden, "Bu listeyi güncelleme izniniz yok")
		return
	}

	todoItem.Title = updatedTodoItem.Title
	todoItem.Description = updatedTodoItem.Description
	todoItem.UpdatedAt = time.Now()
	if err := mock.DB.Save(&todoItem).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "TodoItem güncellenirken bir hata oluştu: "+err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, todoItem)
}

// DeleteTodoItem belirli bir ID'ye sahip todo öğesini siler (soft delete)
func DeleteTodoItem(c *gin.Context) {
	todoIDStr := c.Param("todoId")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Geçersiz todoID")
		return
	}

	itemIDStr := c.Param("itemId")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Geçersiz öğe ID'si")
		return
	}

	var todoItem models.TodoItem
	if err := mock.DB.Where("id = ? AND todo_list_id = ? AND deleted_at IS NULL", itemID, todoID).First(&todoItem).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "TodoItem bulunamadı")
		return
	}

	// Yetkilendirme kontrolü
	var todoList models.TodoList
	if err := mock.DB.Where("id = ? AND deleted_at IS NULL", todoID).First(&todoList).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "TodoList bulunamadı")
		return
	}
	userID := c.GetUint("userID")
	if todoList.UserID != userID {
		utils.RespondError(c, http.StatusForbidden, "Bu öğeyi silme izniniz yok")
		return
	}

	deletedAt := time.Now() // Soft delete
	todoItem.DeletedAt = &deletedAt
	mock.DB.Save(&todoItem)

	utils.RespondJSON(c, http.StatusOK, "TodoItem silindi")
}
