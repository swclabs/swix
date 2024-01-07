package mailers

var Env *MailerEnv = &MailerEnv{}

func Config(email, appPassword string) {
	Env.Email = email
	Env.AppPassword = appPassword
}
