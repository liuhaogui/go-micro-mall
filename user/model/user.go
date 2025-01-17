package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/liuhaogui/go-micro-mall/common/util/log"
	pb "github.com/liuhaogui/go-micro-mall/user/proto/user"
	config "github.com/micro/go-micro/config"
	"github.com/micro/go-plugins/config/source/consul"
)

var db *gorm.DB

type dbInfo struct {
	Address      string `json:"address"`
	Port         int    `json:"port"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	DbName       string `json:"db_name"`
}

func Init(address string) {
	consulSource := consul.NewSource(consul.WithAddress(address))
	conf := config.NewConfig()

	// Load file source
	err := conf.Load(consulSource)
	if err != nil {
		log.Fatal(err)
	}
	var v dbInfo
	err = conf.Get("micro", "config", "database", "user").Scan(&v)
	if err != nil {
		log.Fatal(err)
	}

	log.Info(v)
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		v.UserName, v.UserPassword, v.Address, v.Port, v.DbName))
	if err != nil {
		log.Fatal("failed to connect database：", err)
	}

	db.AutoMigrate(&pb.User{})
	db.Model(&pb.User{}).AddUniqueIndex("uIndex_phone", "phone")
}

// CreateUser 在数据库中创建一个用户
func CreateUser(user *pb.User) error {
	return db.Create(user).Error
}

// DelUser 删除用户
func DelUser(user *pb.User) error {
	return db.Delete(user).Error
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(user *pb.User) error {
	return db.Model(user).Updates(*user).Error
}

// GetByID 通过id取用户信息
func GetByID(id int64) (pb.User, error) {
	var user pb.User
	err := db.Where("id = ?", id).Find(&user).Error
	return user, err
}

// GetByPhone 通过电话获取用户信息
func GetByTel(phone string) (pb.User, error) {
	var user pb.User
	err := db.Where("phone = ?", phone).Find(&user).Error
	return user, err
}

// GetAllUsers 获取所有用户信息
func GetAllUsers() ([]*pb.User, error) {
	var users []*pb.User
	err := db.Find(&users).Error
	return users, err
}
