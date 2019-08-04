package user

type Repository interface {
	Create(user *User) (id string, err error)
	Read(email string) (*User, error)
	Update(user *User) error
	Delete(email string) error
	Exist(email string) (bool, error)
}
