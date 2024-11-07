package repositories

import "EurikaOrmanel/up-charter/internal/models"

func (db DB) FindAdminByPhone(phone string) models.Admin {
	admin := new(models.Admin)
	db.First(admin, "phone = ?", phone)
	return *admin

}

func (db DB) FindAdminByEmail(email string) models.Admin {
	admin := new(models.Admin)
	db.First(admin, "email = ?", email)
	return *admin

}

func (db DB) FindAdminByID(id string) models.Admin {
	admin := new(models.Admin)
	db.First(admin, "id = ?", id)
	return *admin
}

func (db DB) CreateAdmin(admin *models.Admin) error {
	return db.Create(admin).Error
}
