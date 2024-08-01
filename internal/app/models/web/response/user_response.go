package response

type UserRegisterResponse struct {
	Name         string `json:"name"`
	Biography    string `json:"biography"`
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
}
