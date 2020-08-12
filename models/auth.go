package models

type Admin struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Level_ID int    `gorm:"column:level_ID" json:"level_ID"`

}

func CheckAuth(username, password string) (bool, int, int) {
	var admin Admin
	db.Where(Admin{Name: username, Password: password}).First(&admin)
	if admin.ID > 0 {
		return true, admin.ID, admin.Level_ID
	}

	return false, 0, 0
}
