package template

import (
	"fmt"
)

type iOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type otp struct {
	iOtp iOtp
}

func (o *otp)genAndSendOTP(optLength int) error {
	opt := o.iOtp.genRandomOTP(optLength)
	o.iOtp.saveOTPCache(opt)
	message := o.iOtp.getMessage(opt)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOtp.publishMetric()
	return nil
}

type sms struct {
	otp
}

func (s *sms)genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random opt %s\n", randomOTP)
	return randomOTP
}
func (s *sms)saveOTPCache(otp string)  {
	fmt.Printf("SMS: saving otp %s to cache\n", otp)
}
func (s *sms)getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}
func (s *sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
func (s *sms)publishMetric()  {
	fmt.Printf("SMS: publishing metrics\n")
}
type email struct {
	otp
}
func (s *email)genRandomOTP(len int) string {
	randomOTP := "5678"
	fmt.Printf("EMAIL: generating random opt %s\n", randomOTP)
	return randomOTP
}
func (s *email)saveOTPCache(otp string)  {
	fmt.Printf("EMAIL: saving otp %s to cache\n", otp)
}
func (s *email)getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}
func (s *email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending sms: %s\n", message)
	return nil
}
func (s *email)publishMetric()  {
	fmt.Printf("EMAIL: publishing metrics\n")
}

func TestTemplate()  {
	smsOTP := &sms{}
	o := otp{iOtp: smsOTP}
	o.genAndSendOTP(4)
	fmt.Println("====")
	emailOTP := &email{}
	o = otp{iOtp: emailOTP}
	o.genAndSendOTP(4)
}