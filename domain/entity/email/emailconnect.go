package email

type EmailConnect struct {
	Login     string
	Password  string
	ServerSmt string
	Port      int
}

func ConnectToEmail() *EmailConnect {
	return &EmailConnect{
		Login:     "<email.login>",
		Password:  "appPassword",
		ServerSmt: "smtp.gmail.com",
		Port:      587,
	}
}
