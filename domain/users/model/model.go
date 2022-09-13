package users

type ReqUser struct {
	ID_Usuario *uint           `json:"id_usuario" gorm:"primaryKey"`
	Nome       *string         `json:"nome"`
	Email      *string         `json:"email"`
	Password   *string         `json:"password"`
}