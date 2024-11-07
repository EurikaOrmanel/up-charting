package schemas

type AdminRegisterInput struct {
	Fullname string `validate:"required|min_len:2" json:"fullname"`
	Email    string `validate:"email" json:"email,omitempty"`
	Phone    string `validate:"ghPhone" json:"phone,omitempty"`
	Password string `validate:"required|min_len:8|max_len:20" json:"password"`
}

type AdminLoginInput struct {
	Email    string `validate:"required|email" json:"email,omitempty"`
	// Phone    string `validate:"required|ghPhone" json:"phone,omitempty"`
	Password string `validate:"required"`
}
