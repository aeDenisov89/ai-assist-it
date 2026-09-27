package models

import "time"

// User представляет пользователя системы.
type User struct {
	ID        int       `json:"id"`         // Уникальный идентификатор пользователя
	Email     string    `json:"email"`      // Электронная почта
	Name      string    `json:"name"`       // Имя пользователя
	CreatedAt time.Time `json:"created_at"` // Дата создания
}

// CreateUserRequest описывает тело запроса на создание пользователя.
type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`  // Электронная почта, обязательное поле, должна быть валидной
	Name  string `json:"name" validate:"required,max=100"` // Имя пользователя, обязательное поле, максимум 100 символов
}

// UpdateUserRequest описывает тело запроса на обновление пользователя.
type UpdateUserRequest struct {
	Email *string `json:"email" validate:"omitempty,email"`  // Электронная почта, опциональное поле, должна быть валидной
	Name  *string `json:"name" validate:"omitempty,max=100"` // Имя пользователя, опциональное поле, максимум 100 символов
}

// ErrorResponse структура для описания ошибки.
type ErrorResponse struct {
	Error   string `json:"error"`   // Тип ошибки
	Message string `json:"message"` // Описание ошибки
}

// SuccessResponse универсальная обёртка успешного ответа.
type SuccessResponse struct {
	Success bool        `json:"success"`        // Признак успешного ответа
	Data    interface{} `json:"data,omitempty"` // Данные ответа: объект или список, опционально
}
