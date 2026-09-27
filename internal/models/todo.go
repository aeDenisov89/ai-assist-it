package models

import "time"

// Todo представляет данные о списке дел.
type Todo struct {
	ID          int       `json:"id"`          // Уникальный идентификатор задачи
	UserID      int       `json:"user_id"`     // Уникальный идентификатор пользователя
	Title       string    `json:"title"`       // Название задачи
	Description string    `json:"description"` // Подробное описание задачи
	Done        bool      `json:"done"`        // Статус выполнения задачи
	CreatedAt   time.Time `json:"created_at"`  // Время создания задачи
	UpdatedAt   time.Time `json:"updated_at"`  // Время последнего обновления задачи
}

// CreateTodoRequest описывает тело запроса на создание задачи.
type CreateTodoRequest struct {
	UserID      int    `json:"user_id" validate:"required"`       // ID пользователя, обязательное поле
	Title       string `json:"title" validate:"required,max=200"` // Название задачи, обязательное, максимум 200 символов
	Description string `json:"description" validate:"max=1000"`   // Подробное описание задачи, максимально 1000 символов
	Done        bool   `json:"done"`                              // Статус выполнения задачи
}

// UpdateTodoRequest описывает тело запроса на обновление задачи.
type UpdateTodoRequest struct {
	Title       *string `json:"title" validate:"omitempty,max=200"`        // Новое название задачи, опционально, максимум 200 символов
	Description *string `json:"description" validate:"omitempty,max=1000"` // Новое описание задачи, опционально, максимум 1000 символов
	Done        *bool   `json:"done"`                                      // Новый статус задачи, опционально
}
