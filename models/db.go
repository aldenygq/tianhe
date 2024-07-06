package models

import (
	"strings"

	"oncall/middleware"

	"github.com/jinzhu/gorm"

	"errors"
	"fmt"
)

var db = middleware.Sql

const (
	ONCALL_RULE       = "oncall_rule"
	CURRENT_DUTY_INFO = "current_duty_info"
	USERS             = "users"
	USER_TOKEN_EXPIRE = "user_token_expire"
)

type UserTokenExpire struct {
	Ctime    int64  `gorm:"column:ctime;type:int(11)" json:"ctime" description:"创建时间"`
	Id       int64  `gorm:"column:id;PRIMARY_KEY;type:int(10)" json:"id"  description:"用户主键id"`
	Mtime    int64  `gorm:"column:mtime;type:int(11)" json:"mtime" description:"修改时间"`
	EnName   string `gorm:"column:en_name;type:varchar(256)" json:"en_name" description:"用户英文名"`
	ExpireTime int64 `gorm:"column:expire_time;type:int(11)" json:"expire_time" description:"用户token有效时间"`
}
func (e *UserTokenExpire) Create() error{
	tx := middleware.Sql.Begin()

	err := tx.Table(USER_TOKEN_EXPIRE).Create(e).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil

}
func (e *UserTokenExpire) GetByUser() error {
	err := middleware.Sql.Table(USERS).Where("en_name = ?",e.EnName).Take(&e).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	} else if err != nil {
		return err
	}
	return nil 
}
func (e *UserTokenExpire) SetUserTokenExpire() error {
	tx := middleware.Sql.Begin()

	err := tx.Table(USER_TOKEN_EXPIRE).Where("en_name= ?", e.EnName).Update(&e).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
type Users struct {
	Ctime    int64  `gorm:"column:ctime;type:int(11)" json:"ctime" description:"创建时间"`
	Email    string `gorm:"column:email;type:int(11)" json:"email" description:"邮箱"`
	Id       int64  `gorm:"column:id;PRIMARY_KEY;type:int(10)" json:"id"  description:"用户主键id"`
	Mtime    int64  `gorm:"column:mtime;type:int(11)" json:"mtime" description:"修改时间"`
	CnName   string `gorm:"column:cn_name;type:varchar(256)" json:"cname" description:"用户中文名"`
	EnName   string `gorm:"column:en_name;type:varchar(256)" json:"en_name" description:"用户英文名"`
	Password string `gorm:"column:password;type:varchar(256)" json:"password" description:"用户密码"`
	Mobile   string `gorm:"column:mobile;type:varchar(16)" json:"mobile" description:"手机号"`
	Status   int64  `gorm:"column:status;type:int(10)" json:"status" description:"用户状态，1(启用)/2(禁用)/3(删除)"`
}

func (u *Users) Create() error {
	tx := middleware.Sql.Begin()

	err := tx.Table(USERS).Create(u).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
func (u *Users) CheckuserStatus() error {
	_,result,err := u.IsExist()
	if err != nil || result {
		return err 
	}

	return nil
}
/*
func (u *Users) IsDelete() (bool, error) {
	tx, err := u.getRequirement()
	if err != nil {
		return true, err
	}
	err = tx.Where("status = 1").Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return false, gorm.ErrRecordNotFound
	} else if err != nil {
		tx.Rollback()
		return false, err
	}
	err = tx.Commit().Error
	if err != nil {
		return true, err
	}
	return true, nil
}
*/
func (u *Users) CheckUserExistByMobile() (bool,error) {
	err := middleware.Sql.Table(USERS).Where("mobile = ? and status != 3",u.Mobile).Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false,nil
	} else if err != nil {
		return true,errors.New(fmt.Sprintf("check mobile %v failed:%v",u.Mobile,err))
	}
	return true,nil 
}
func (u *Users) CheckUserExistByEnName() (bool,error) {
	err := middleware.Sql.Table(USERS).Where("en_name = ? and status != 3",u.EnName).Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false,nil
	} else if err != nil {
		return true,errors.New(fmt.Sprintf("check user %v failed:%v",u.EnName,err))
	}
	return true,nil
}
func (u *Users) IsExist() (string,bool,error) {
	var (
		err error
		result bool
	) 
	result,err = u.CheckUserExistByMobile()
	if err != nil || result {
		return u.Mobile,result,err 
	}
	result,err = u.CheckUserExistByEnName()
	if err != nil || result {
		return u.EnName,result,err 
	}
	return "",result,nil
}
func (u *Users) getRequirement() (*gorm.DB, error) {
	tx := middleware.Sql.Begin()
	switch {
	case strings.TrimSpace(u.Mobile) != "":
		tx = tx.Table(USERS).Where("mobile  = ?", u.Mobile)
	case strings.TrimSpace(u.EnName) != "":
		tx = tx.Table(USERS).Where("en_name  = ?", u.EnName)
	case strings.TrimSpace(u.Email) != "":
		tx = tx.Table(USERS).Where("email  = ?", u.Email)
	default:
		return nil, errors.New("filed invalid")
	}
	return tx, nil
}

func (u *Users) CheckUserDisableByMobile() (bool,error) {
	err := middleware.Sql.Table(USERS).Where("mobile = ? and status = 2",u.Mobile).Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false,nil
	} else if err != nil {
		return true,errors.New(fmt.Sprintf("check mobile %v failed:%v",u.Mobile,err))
	}
	return true,nil 
}
func (u *Users) CheckUserDisableByEnName() (bool,error) {
	err := middleware.Sql.Table(USERS).Where("en_name = ? and status = 2",u.EnName).Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false,nil
	} else if err != nil {
		return true,errors.New(fmt.Sprintf("check user %v failed:%v",u.EnName,err))
	}
	return true,nil
}
func (u *Users) IsDisable() (string,bool, error) {
	var (
		err error
		result bool
	) 
	result,err = u.CheckUserDisableByMobile()
	if err != nil || result {
		return u.Mobile,result,err 
	}
	result,err = u.CheckUserDisableByEnName()
	if err != nil || result {
		return u.EnName,result,err 
	}
	return "",result,nil
}
func (u *Users) List() (int64, []*Users, error) {
	var (
		count int64
		users []*Users = make([]*Users, 0)
	)

	tx := db.Begin().Table(USERS)
	switch {
	case strings.TrimSpace(u.Mobile) != "":
		tx = tx.Where("mobile  = ?", u.Mobile)
	case strings.TrimSpace(u.EnName) != "":
		tx = tx.Where("en_name  = ?", u.EnName)
	case strings.TrimSpace(u.Email) != "":
		tx = tx.Where("email  = ?", u.Email)
	case u.Status > 0 && u.Status != 3:
		tx = tx.Where("status = ?", u.Status)
	case u.CnName != "":
		tx = tx.Where("cn_name like ?", "%"+u.CnName+"%")
	default:
		return count, nil, errors.New("filed invalid")
	}

	err := tx.Find(&users).Error
	if err != nil {
		tx.Rollback()
		return count, nil, err
	}
	err = tx.Count(&count).Error
	if err != nil {
		tx.Rollback()
		return count, nil, err
	}
	err = tx.Commit().Error
	if err != nil {
		return count, nil, err
	}

	return count, users, nil
}
func (u *Users) GetByUname() error {
	err := middleware.Sql.Table(USERS).Where("en_name = ?", u.EnName).Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("user %v not found", u.EnName))
	} else if err != nil {
		return err
	}
	return nil
}
func (u *Users) GetByMobile() error {
	err := middleware.Sql.Table(USERS).Where("mobile = ?", u.Mobile).Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("user %v not found", u.EnName))
	} else if err != nil {
		return err
	}
	return nil
}
func (u *Users) UpdateByEnName() error {
	tx := middleware.Sql.Begin()

	err := tx.Table(USERS).Where("en_name= ?", u.EnName).Update(&u).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
func (u *Users) UpdateByMobile() error {
	tx := middleware.Sql.Begin()

	err := tx.Table(USERS).Where("mobile = ?", u.Mobile).Update(&u).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

type OncallRule struct {
	Id                  int64              `gorm:"column:id;PRIMARY_KEY;type:int(10)" json:"id"  description:"值班规则主键id"`
	CnTitle             string             `gorm:"column:cn_title;type:varchar(256)" json:"cn_title" description:"中文标题"`
	EnTitle             string             `gorm:"column:en_title;type:varchar(256)" json:"en_title" description:"英文标题"`
	OncallCycleType     string             `gorm:"column:oncall_cycle_type;type:varchar(64)" json:"oncall_cycle_type" binding:"required,min=1" description:"值班周期类型,day(天)、custom(自定义)、month(月)，默认周类型，即每轮7天"`
	StartDay            string             `gorm:"column:start_day;type:varchar(64)" json:"start_day" description:"开始日期,日期不得小于当前日期"`
	RotationNum         int64              `gorm:"column:rotation_num;type:int(10)" json:"rotation_num" description:"轮转次数,如为0，则表示持续轮转"`
	PerRotationDays     int64              `gorm:"column:per_rotation_days;type:int(10)" json:"per_rotation_days"  description:"每轮的轮转天数，最小值为1,最大值为30,custom必传"`
	OncallPeopleInfos   []string           `gorm:"column:oncall_people_infos;type:json" json:"oncall_people_infos" description:"值班人员信息"`
	IsSkipWeekend       int64              `gorm:"column:is_skip_weekend;type:int(10)" json:"is_skip_weekend" description:"是否跳过周末值班，0表示跳过，1表示不跳过，默认为0"`
	SubscribeNotifyInfo []*SubscribeNotify `gorm:"column:subscribe_notify_info;type:json" json:"subscribe_notify_info" description:"订阅通知提醒信息"`
	SubscribeGroups     []*SubscribeGroup  `gorm:"column:subscribe_groups;type:json" json:"subscribe_groups" description:"订阅组信息"`
	IsTemporaryOncall   int64              `gorm:"column:is_temporary_oncall;type:int(10)" json:"is_temporary_oncall" description:"是否开启临时值班：0(不开启),1(开启)，默认是0不开启，当临时值班开启后，默认覆盖现有值班规则"`
	TemporaryOncallInfo *TemporaryOncall   `gorm:"column:temporary_oncall_info;type:json" json:"temporary_oncall_info"  description:"临时值班信息"`
	Status              int64              `gorm:"column:status;type:int(10)" json:"status"  description:"是否启用,0表示启用，1表示不启用,默认启用"`
	Creator             string             `gorm:"column:creator;type:varchar(64)" json:"creator"  description:"创建者"`
	CreateTime          string             `gorm:"column:create_time;type:varchar(64)" json:"create_time"  description:"创建时间"`
	Updator             string             `gorm:"column:updator;type:varchar(64)" json:"updator"  description:"最后一次修改人"`
	UpdateTime          string             `gorm:"column:update_time;type:varchar(64)" json:"update_time"  description:"最后一次修改时间"`
}

func (o *OncallRule) Create() error {
	tx := db.Begin()

	err := tx.Table(ONCALL_RULE).Create(&o).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
func (o *OncallRule) Modify() error {
	tx := db.Begin()

	err := tx.Table(ONCALL_RULE).Update(&o).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
func (o *OncallRule) EnabledRule() ([]*OncallRule, error) {
	var rules []*OncallRule = make([]*OncallRule, 0)
	tx := db.Begin()

	err := tx.Table(ONCALL_RULE).Where("status = ?", 1).Find(&rules).Error
	if err != nil {
		return nil, err
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return rules, nil
}
func (o *OncallRule) List() (int64, []*OncallRule, error) {
	var (
		total int64
		rules []*OncallRule = make([]*OncallRule, 0)
	)

	tx := db.Begin()

	err := tx.Table(ONCALL_RULE).Count(&total).Find(&rules).Error
	if err != nil {
		return total, nil, err
	}
	err = tx.Commit().Error
	if err != nil {
		return total, nil, err
	}

	return total, rules, nil
}

func (o *OncallRule) Get() error {
	tx := db.Begin().Table(ONCALL_RULE)
	if o.Id > 0 {
		tx = tx.Where("id = ?", o.Id)
	}
	if o.CnTitle != "" {
		tx = tx.Where("cn_title like ?", "%"+o.CnTitle+"%")
	}
	if o.EnTitle != "" {
		tx = tx.Where("en_title like ?", "%"+o.EnTitle+"%")
	}
	err := tx.Take(&o).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("oncall rule not esixt")
	} else if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

type CurrentDutyInfo struct {
	Id     int64  `gorm:"column:id;PRIMARY_KEY;type:int(10)" json:"id"  description:"当前值班信息主键id"`
	RuleId int64  `gorm:"column:rule_id;type:int(10)" json:"rule_id"  description:"值班规则id"`
	User   string `gorm:"column:user;type:varchar(512)" json:"user"  description:"当前值班人员"`
}

func (c *CurrentDutyInfo) Create() error {
	tx := db.Begin()

	err := tx.Table(CURRENT_DUTY_INFO).Create(&c).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
func (c *CurrentDutyInfo) List() ([]*CurrentDutyInfo, error) {
	var (
		dutys []*CurrentDutyInfo = make([]*CurrentDutyInfo, 0)
		err   error
	)
	tx := db.Begin().Table(CURRENT_DUTY_INFO)
	if c.RuleId > 0 {
		tx = tx.Where("rule_id = ?", c.RuleId)
	}
	if c.User != "" {
		tx = tx.Where("user like ?", "%"+c.User+"%")
	}
	err = tx.Find(&dutys).Error
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return nil, nil
}
