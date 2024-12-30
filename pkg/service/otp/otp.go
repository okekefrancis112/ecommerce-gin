package otp

// type OtpAuth interface {
// 	SentOtp(phoneNumber string) (string, error)
// 	VerifyOtp(phoneNumber string, code string) (valid bool, err error)
// }

// OtpAuth interface for OTP operations
type OtpAuth interface {
	SendOtp(email string) (string, error)
	VerifyOtp(email string, code string) (bool, error)
}
