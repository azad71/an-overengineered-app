package mailer

import "strings"

func GetSignupContent(otp string) string {

	content := `<p>Hi,</p>
	<p>Thank you so much for signin up to <em>an overengineered social app</em>. 
		Please use following otp to verify your email.</p>
	<p 
		style="text-align: center; margin: 32px 0"><span style="text-align: center; 
		color: #0000ff; font-size: 32px; letter-spacing: 3px; border:1px solid black; padding: 8px">
		<strong>{OTP_PLACEHOLDER}</strong></span>
	</p>
	<p style="text-align: left;">
	<span style="text-align: center; color: #ff0000; font-size: 12px;">
		***If it's not you, you can safely ignore this message.***
	</span></p>`

	return strings.Replace(content, "{OTP_PLACEHOLDER}", otp, -1)
}
