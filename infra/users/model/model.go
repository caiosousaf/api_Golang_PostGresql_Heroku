package users

type ReqUser struct {
	Nome         *string    `json:"nome"`
	Email        *string    `json:"email"`
	Password     string    `json:"password"`
}
