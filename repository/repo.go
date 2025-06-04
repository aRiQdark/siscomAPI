package repository

import (
	"gin-gonic-gorm/entity"
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	FindAllMembers() ([]entity.Membersdetail, error)
	FindByID(memberID string) (entity.Membersdetail, error)
	Create(member entity.Membersdetail) (entity.Membersdetail, error)
	FindByUsername(username string) (entity.Membersdetail, error)
	FindByhandphone(handphone string) (entity.Membersdetail, error)
	FindByEmail(email string) (entity.Membersdetail, error)
	Update(member entity.Membersdetail) (entity.Membersdetail, error)
	Oldpassword(password string) (entity.Membersdetail, error)
	UpdatepasswordByEmail(email entity.Membersdetail) (entity.Membersdetail, error)
	FindAllMembersworker() ([]entity.Membersdetail, error)
	UpdateImeibyemail(data entity.Membersdetail) (entity.Membersdetail, error)
	UpdateLongtitudelatitude(data entity.Membersdetail) (entity.Membersdetail, error)
	FindNearbyMembers(lat, lon, radius float64) ([]entity.Membersdetail, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAllMembers() ([]entity.Membersdetail, error) {
	var members []entity.Membersdetail
	err := r.db.Find(&members).Error
	return members, err
}

func (r *repository) FindAllMembersworker() ([]entity.Membersdetail, error) {
	var members []entity.Membersdetail
	err := r.db.Where("Membertype = ?", "2").Find(&members).Error
	return members, err
}

func (r *repository) FindByID(memberID string) (entity.Membersdetail, error) {
	var member entity.Membersdetail
	err := r.db.Where("member_id = ?", memberID).First(&member).Error
	return member, err
}

func (r *repository) Create(member entity.Membersdetail) (entity.Membersdetail, error) {
	log.Printf("Saving member to database: %+v", member)
	err := r.db.Create(&member).Error
	if err != nil {
		log.Printf("Error saving member: %v", err)
	}
	return member, err
}
func (r *repository) FindByUsername(username string) (entity.Membersdetail, error) {
	var member entity.Membersdetail
	err := r.db.Where("username = ?", username).First(&member).Error
	return member, err
}

func (r *repository) Oldpassword(password string) (entity.Membersdetail, error) {
	var member entity.Membersdetail
	err := r.db.Where("password = ?", password).First(&member).Error
	return member, err
}

func (r *repository) FindByhandphone(handphone string) (entity.Membersdetail, error) {
	var member entity.Membersdetail
	err := r.db.Where("handphone = ?", handphone).First(&member).Error
	return member, err
}

func (r *repository) FindByEmail(email string) (entity.Membersdetail, error) {
	var member entity.Membersdetail
	err := r.db.Where("email = ?", email).First(&member).Error
	if err != nil {
		return member, err
	}
	return member, nil
}

func (r *repository) Update(member entity.Membersdetail) (entity.Membersdetail, error) {
	err := r.db.Model(&entity.Membersdetail{}).Where("member_id = ?", member.MemberID).Updates(member).Error
	if err != nil {
		return member, err
	}
	return member, nil
}

func (r *repository) UpdatepasswordByEmail(email entity.Membersdetail) (entity.Membersdetail, error) {
	err := r.db.Model(&entity.Membersdetail{}).Where("email = ?", email.Email).Updates(email.Password).Error
	if err != nil {
		return email, err
	}
	return email, nil
}

func (r *repository) UpdateImeibyemail(data entity.Membersdetail) (entity.Membersdetail, error) {
	err := r.db.
		Model(&entity.Membersdetail{}).
		Where("email = ?", data.Email).
		Update("imei", data.Imei).Error

	if err != nil {
		return entity.Membersdetail{}, err
	}

	return data, nil
}

func (r *repository) UpdateLongtitudelatitude(data entity.Membersdetail) (entity.Membersdetail, error) {
	err := r.db.
		Model(&entity.Membersdetail{}).
		Where("member_id = ?", data.MemberID).
		Update("longtitude", data.Longtitude).Update("lattitude", data.Lattitude).Error

	if err != nil {
		return entity.Membersdetail{}, err
	}

	return data, nil
}

func (r *repository) FindNearbyMembers(lat, lon, radius float64) ([]entity.Membersdetail, error) {
	var members []entity.Membersdetail

	query := `
		SELECT * FROM (
			SELECT *, (
				6371 * acos(
					cos(radians(?)) * cos(radians(lattitude::float)) *
					cos(radians(longtitude::float) - radians(?)) +
					sin(radians(?)) * sin(radians(lattitude::float))
				)
			) AS distance
			FROM membersdetails
			WHERE lattitude IS NOT NULL 
			  AND longtitude IS NOT NULL
		) AS subquery
		WHERE distance <= ?
		ORDER BY distance ASC;
	`

	err := r.db.Raw(query, lat, lon, lat, radius).Scan(&members).Error
	if err != nil {
		return nil, err
	}

	return members, nil
}
