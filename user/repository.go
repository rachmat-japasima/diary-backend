package user

import (
	"database/sql"
	"time"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository{
	return &repository{db:db}
}

func (r *repository) Save(user User) (User, error) {
	// err := r.db.Create(&user).Error
	_, err := r.db.Exec("INSERT INTO users (name, gender, email, password_hash, file_avatar, role, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?)", user.Name, user.Gender, user.Email, user.PasswordHash, user.FileAvatar, user.Role, time.Now(), time.Now())

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error){
	var user User
	// err := r.db.Where("email = ?", email).Find(&user).Error
	row := r.db.QueryRow("SELECT * FROM users WHERE email = ?", email)

	err := row.Scan(&user.ID, &user.Name, &user.Gender ,&user.Email ,&user.PasswordHash, &user.FileAvatar, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (User, error){
	var user User
	// err := r.db.Where("id = ?", ID).Find(&user).Error
	row := r.db.QueryRow("SELECT * FROM users WHERE id = ?", ID)

	err := row.Scan(&user.ID, &user.Name, &user.Gender ,&user.Email ,&user.PasswordHash, &user.FileAvatar, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error){
	// err := r.db.Save(&user).Error
	_, err := r.db.Exec("update users set name = ?, gender = ?, email = ?, password_hash = ?, file_avatar = ?, role = ?, updated_at = ? where id = ?", user.Name, user.Gender, user.Email, user.PasswordHash, user.FileAvatar, user.Role, time.Now(), user.ID)


	if err != nil {
		return user, err
	}

	return user, nil
}