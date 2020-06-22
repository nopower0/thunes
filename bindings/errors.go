package bindings

type Error struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

var (
	ServerError = &Error{
		Code:    "S0001",
		Message: "Internal Server Error",
	}
	TokenError = &Error{
		Code:    "T0001",
		Message: "Invalid Token",
	}
	UserNotLoginError = &Error{
		Code:    "U0001",
		Message: "User not login",
	}
	InvalidUsernameOrPasswordError = &Error{
		Code:    "U0002",
		Message: "Invalid username or password",
	}
	UserAlreadyLoginError = &Error{
		Code:    "U0003",
		Message: "User already login",
	}
)

func NewParamError(msg string) *Error {
	return &Error{
		Code:    "P0001",
		Message: msg,
	}
}
