package models

type SystemUser struct {
	Model

	Account  string `json:"account"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

func SelectOneByAccountAndPassword(account string, password string) (result SystemUser) {
	db.Model(&result).Where(map[string]any{
		"account":  account,
		"password": password,
	}).First(&result)

	return
}
