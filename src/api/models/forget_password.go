package models

// Password is struct use for forgot password api
type Password struct {
	Email    string `json:"email"`
	Otp      string `json:"opt"`
	Password string `json:"password"`
}

// func GenerateOpt() int {
// 	Opt := 12345
// 	return Opt
// }

// func VerifyOpt() bool {
// 	opt := GenerateOpt()
// 	if opt == 12345 {
// 		fmt.Println(opt)
// 		return true
// 	}
// 	return false
// }
