package user

type UserFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

type UserRegisterFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:     user.ID,
		Name:   user.Name,
		Gender: user.Gender,
		Email:  user.Email,
		Token:  token,
	}

	return formatter
}

func FormatRegisterUser(user User) UserRegisterFormatter {
	formatter := UserRegisterFormatter{
		ID:     user.ID,
		Name:   user.Name,
		Gender: user.Gender,
		Email:  user.Email,
	}

	return formatter
}