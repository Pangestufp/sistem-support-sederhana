package dto

type RegisterRequest struct {
	Name                 string `json:"name"`
	Username             string `json:"username"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirm"`
	SuperiorID           *int   `json:"superior_id"`
	DepartmentID         *int   `json:"department_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	DepartmentID *int   `json:"department_id"`
	SuperiorID   *int   `json:"superior_id"`
	Status       int    `json:"status"`
}

type UpdatePasswordRequest struct {
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	NewPasswordConfirm string `json:"new_password_confirm"`
}

type UserResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	DepartmentID *int   `json:"department_id"`
	SuperiorID   *int   `json:"superior_id"`
	Status       int    `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
