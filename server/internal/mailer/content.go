package mailer

import (
	"an-overengineered-app/internal/logger"
	"context"
	"strings"
)

func GetSignupContent(ctx context.Context, otp string) string {
	logger.PrintInfo(ctx, "Generating signup mail content", nil)

	content := `<p>Hi,</p>
		<p>Thank you so much for signin up to <em>an overengineered app</em>.
		Please use the following otp to verify your email. OTP will be expired after <strong>two minutes</strong></p>
		<p
		style="text-align: center; margin: 32px 0"><span style="text-align: center;
		color: #0000ff; font-size: 32px; letter-spacing: 3px; border:1px solid black; padding: 8px">
		<strong>{OTP_PLACEHOLDER}</strong></span>
		</p>
		<p style="text-align: left;">
		<span style="text-align: center; color: #ff0000; font-size: 12px;">
		***If it's not you, you can safely ignore this message.***
		</span></p>`

	content = strings.Replace(content, "{OTP_PLACEHOLDER}", otp, -1)

	logger.PrintInfo(ctx, "Constructed signup mail content", map[string]string{"mailContent": content})

	return content
}
