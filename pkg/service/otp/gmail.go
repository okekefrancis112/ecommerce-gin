package otp

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/okekefrancis112/ecommerce-gin/pkg/config"

	"gopkg.in/gomail.v2"
)

// // OtpAuth interface for OTP operations
// type OtpAuth interface {
// 	SendOtp(email string) (string, error)
// 	VerifyOtp(email string, code string) (bool, error)
// }

type gmailOtp struct {
	smtpHost    string
	smtpPort    int
	email       string
	password    string
	otpStore    map[string]string
	expiration  map[string]time.Time
	storeMutex  sync.RWMutex
	otpLifetime time.Duration
}

// NewOtpAuth initializes the Gmail OTP service
func NewOtpAuth(cfg config.Config) OtpAuth {
	// func NewOtpAuth(smtpHost string, smtpPort int, email, password string, otpLifetime time.Duration) OtpAuth {
	return &gmailOtp{
		smtpHost:    cfg.smtpHost,
		smtpPort:    cfg.smtpPort,
		email:       cfg.email,
		password:    cfg.password,
		otpStore:    make(map[string]string),
		expiration:  make(map[string]time.Time),
		otpLifetime: 5 * time.Minute,
	}
}

// GenerateOtp generates a random 6-digit OTP
func generateOtp() (string, error) {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", fmt.Errorf("failed to generate OTP: %w", err)
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}

// SendOtp sends an OTP to the specified email
func (g *gmailOtp) SendOtp(email string) (string, error) {
	otp, err := generateOtp()
	if err != nil {
		return "", err
	}

	// Send OTP via Gmail
	mail := gomail.NewMessage()
	mail.SetHeader("From", g.email)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Your OTP Code")
	mail.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))

	dialer := gomail.NewDialer(g.smtpHost, g.smtpPort, g.email, g.password)

	if err := dialer.DialAndSend(mail); err != nil {
		return "", fmt.Errorf("failed to send OTP email: %w", err)
	}

	// Store OTP with expiration time
	g.storeMutex.Lock()
	defer g.storeMutex.Unlock()
	g.otpStore[email] = otp
	g.expiration[email] = time.Now().Add(g.otpLifetime)

	log.Printf("OTP sent to %s", email)
	return otp, nil
}

// VerifyOtp verifies the OTP for the given email
func (g *gmailOtp) VerifyOtp(email string, code string) (bool, error) {
	g.storeMutex.RLock()
	defer g.storeMutex.RUnlock()

	otp, exists := g.otpStore[email]
	if !exists {
		return false, fmt.Errorf("no OTP found for email: %s", email)
	}

	if time.Now().After(g.expiration[email]) {
		delete(g.otpStore, email) // Clean up expired OTP
		delete(g.expiration, email)
		return false, fmt.Errorf("OTP expired for email: %s", email)
	}

	if otp != code {
		return false, fmt.Errorf("invalid OTP for email: %s", email)
	}

	// Clean up after successful verification
	delete(g.otpStore, email)
	delete(g.expiration, email)
	log.Printf("OTP verified successfully for %s", email)
	return true, nil
}
