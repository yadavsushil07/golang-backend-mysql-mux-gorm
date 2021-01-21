package security

import "math/rand"

// import (
// 	"github.com/xlzd/gotp"
// )

const letterBytes = "0123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// func main() {
// 	fmt.Println("Random secret:", gotp.RandomSecret(16))
// 	defaultTOTPUsage()
// 	defaultHOTPUsage()
// }

// func GetOTP() string {
// 	otp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")

// 	return otp.Now()

// }

// func VerifyOTP(otp string) bool {
// 	return otp.Verify(otp, 1524485781)
// }


