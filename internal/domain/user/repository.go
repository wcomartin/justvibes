package user

type Repository interface {
    Save(user *User) error
    FindByID(id UserID) (*User, error)
    FindAll() ([]*User, error)
    FindByEmail(email string) (*User, error)
}