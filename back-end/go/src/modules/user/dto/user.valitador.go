package dto

type CreateUserDto struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
