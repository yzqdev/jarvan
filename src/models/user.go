package models

type User struct {
	Model

	Id        int64  `json:"id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var Users []User

func (user User) Insert() (id int64, err error) {
	if err = db.Create(&user).Error; err != nil {
		return -1, err
	}

	return user.Id, nil
}

func (user *User) Users() (users []User, err error) {
	if err = db.Find(&Users).Error; err != nil {
		return
	}

	return
}

func (user *User) Destroy(id int64) (Result User, err error) {
	if err = db.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = db.Delete(&user).Error; err != nil {
		return
	}

	Result = *user
	return
}
