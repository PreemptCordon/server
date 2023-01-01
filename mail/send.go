package mail

import "net/smtp"

func send(recipient string, message MarkdownEmail) {
	auth := smtp.PlainAuth(
		"",
		ServerConfig.MailUser,
		ServerConfig.SMTPpassword,
		ServerConfig.SMTPhost)
	smtp.SendMail(
		ServerConfig.SMTPhost+":"+ServerConfig.SMTPport,
		auth,
		ServerConfig.MailUser,
		recipient,
		render(message),
	)
}
