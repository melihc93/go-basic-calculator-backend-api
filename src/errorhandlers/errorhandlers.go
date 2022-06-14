package errorhandlers

func UsernameNotExist() string {
	errorMessage := "Username not exist"

	return errorMessage
}

func PasswordNotExist() string {
	errorMessage := "Password not exist"

	return errorMessage
}

func InvalidUsernameOrPassword() string {
	errorMessage := "Invalid username or password"

	return errorMessage
}

func QueryParamOutOfRange() string {
	errorMessage := "query param should not exceed 5 or less than 1"

	return errorMessage
}
