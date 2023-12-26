package user

type userRegisterFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func UserRegisterFormatter(user User, token string) userRegisterFormatter {
	var formatter userRegisterFormatter = userRegisterFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}

	return formatter
}
