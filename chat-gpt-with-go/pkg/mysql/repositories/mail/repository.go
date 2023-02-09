package mail

import (
	"go.tienngay/pkg/mysql/entities"
	"gorm.io/gorm"
	"time"
	"fmt"
)

const (
    YYYYMMDDHIS = "2006-01-02 15:04:05"
    STATUS_WAITING = 1
    STATUS_SUCCESS = 2
    STATUS_ERRORS = 3
    STATUS_INACTIVE = 4
    STATUS_SENDING = 5
)

//Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	GetWaitingMails() (*[]entities.Mail)
	UpdateStatusSending(ID int) (bool)
	UpdateStatusSuccess(mail entities.Mail) (bool)
	UpdateStatusErrors(mail entities.Mail) (bool)
}

type repository struct {
	DB *gorm.DB
}

//NewRepo is the single instance repo that is being created.
func NewRepo(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) GetWaitingMails() (*[]entities.Mail) {
	var mails *[]entities.Mail
	targetTime := time.Now().Add(-360*time.Minute)
	fmt.Println(targetTime.Format(YYYYMMDDHIS))
	r.DB.Where("status = ?", STATUS_WAITING).Where("created_at > ?", targetTime).Where("deleted_at IS NULL").Limit(250).Find(&mails)
	return mails
}

func (r *repository) UpdateStatusSending(ID int) (bool) {
	r.DB.Where("id = ?", ID).Update("status", STATUS_SENDING)
	return true
}

func (r *repository) UpdateStatusSuccess(mail entities.Mail) (bool) {
	r.DB.Model(&mail).Where("id = ?", mail.ID).Updates(map[string]interface{}{"status": STATUS_SUCCESS, "errors": mail.Errors})
	return true
}

func (r *repository) UpdateStatusErrors(mail entities.Mail) (bool) {
	r.DB.Model(&mail).Where("id = ?", mail.ID).Updates(map[string]interface{}{"status": STATUS_ERRORS, "errors": mail.Errors})
	return true
}
