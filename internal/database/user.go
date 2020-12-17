package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/laojianzi/mdclubgo/internal/web"
)

// User db table with user
type User struct {
	ID                     uint   `gorm:"COLUMN:user_id;TYPE:INT(11) UNSIGNED;NOT NULL;autoIncrement;primaryKey;->;COMMENT:用户ID" json:"user_id"`
	Username               string `gorm:"TYPE:VARCHAR(20);NOT NULL;INDEX:user_name;COMMENT:用户名" json:"username"`
	Email                  string `gorm:"TYPE:VARCHAR(320) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL;INDEX:email;COMMENT:邮箱" json:"email"`
	Avatar                 string `gorm:"TYPE:VARCHAR(50);COMMENT:头像token" json:"-"`
	Cover                  string `gorm:"TYPE:VARCHAR(50);COMMENT:封面图片token" json:"-"`
	Password               string `gorm:"TYPE:VARCHAR(255);NOT NULL;COMMENT:密码" json:"-"`
	CreateIP               string `gorm:"TYPE:VARCHAR(80);COMMENT:注册IP" json:"create_ip"`
	CreateLocation         string `gorm:"TYPE:VARCHAR(100);COMMENT:注册地址" json:"create_location"`
	LastLoginTime          uint64 `gorm:"TYPE:INT(10) UNSIGNED;DEFAULT:0;NOT NULL;COMMENT:最后登录时间" json:"last_login_time"`
	LastLoginIP            string `gorm:"TYPE:VARCHAR(80);COMMENT:最后登陆IP" json:"last_login_ip"`
	LastLoginLocation      string `gorm:"TYPE:VARCHAR(100);COMMENT:最后登录地址" json:"last_login_location"`
	FollowerCount          uint   `gorm:"TYPE:INT(11) UNSIGNED;NOT NULL;DEFAULT:0;INDEX:follower_count;COMMENT:关注我的人数" json:"follower_count"`
	FolloweeCount          uint   `gorm:"TYPE:INT(11) UNSIGNED;NOT NULL;DEFAULT:0;COMMENT:我关注的人数" json:"followee_count"`
	FollowingArticleCount  uint   `gorm:"TYPE:INT(11) UNSIGNED;NOT NULL;DEFAULT:0;COMMENT:我关注的文章数" json:"following_article_count"`
	FollowingQuestionCount uint   `gorm:"TYPE:INT(11) UNSIGNED;NOT NULL;DEFAULT:0;COMMENT:我关注的问题数" json:"following_question_count"`
	FollowingTopicCount    uint   `gorm:"TYPE:INT(11) UNSIGNED;NOT NULL;DEFAULT:0;COMMENT:我关注的话题数" json:"following_topic_count"`
	ArticleCount           uint   `gorm:"TYPE:INT(11) UNSIGNED;NOT NULL;DEFAULT:0;COMMENT:我发表的文章数量" json:"article_count"`
	QuestionCount          uint   `gorm:"TYPE:INT(11) UNSIGNED;NOT NULL;DEFAULT:0;COMMENT:我发表的问题数量" json:"question_count"`
	AnswerCount            uint   `gorm:"TYPE:INT(11) UNSIGNED;NOT NULL;DEFAULT:0;COMMENT:我发表的回答数量" json:"answer_count"`
	NotificationUnread     uint   `gorm:"TYPE:INT(11) UNSIGNED;NOT NULL;DEFAULT:0;COMMENT:未读通知数" json:"notification_unread"`
	InboxUnread            uint   `gorm:"TYPE:INT(11) UNSIGNED;NOT NULL;DEFAULT:0;COMMENT:未读私信数" json:"inbox_unread"`
	Headline               string `gorm:"TYPE:VARCHAR(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;COMMENT:一句话介绍" json:"headline"`
	Bio                    string `gorm:"TYPE:VARCHAR(160);COMMENT:个人简介" json:"bio"`
	Blog                   string `gorm:"TYPE:VARCHAR(255);COMMENT:个人主页" json:"blog"`
	Company                string `gorm:"TYPE:VARCHAR(255);COMMENT:公司名称" json:"company"`
	Location               string `gorm:"TYPE:VARCHAR(255);COMMENT:地址" json:"location"`
	CreateTime             uint64 `gorm:"autoCreateTime;TYPE:INT(10) UNSIGNED;INDEX:create_time;DEFAULT:0;NOT NULL;COMMENT:注册时间" json:"create_time"`
	UpdateTime             uint64 `gorm:"autoUpdateTime;TYPE:INT(10) UNSIGNED;DEFAULT:0;NOT NULL;COMMENT:更新时间" json:"update_time"`
	DisableTime            uint64 `gorm:"TYPE:INT(10) UNSIGNED;DEFAULT:0;NOT NULL;COMMENT:禁用时间" json:"disable_time"`
}

// TableName return user table name
func (User) TableName() string {
	return WithTablePrefix("user")
}

// BeforeCreate before create action for user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = passwordHash(u.Password)
	return err
}

// Register add a user save to database
func Register(db *gorm.DB, ctx echo.Context, username, email, password string) (*User, error) {
	reqTimeSec, err := strconv.ParseInt(ctx.Request().Header.Get("REQUEST_TIME"), 10, 64)
	reqTime := time.Now()
	if err == nil {
		reqTime = time.Unix(reqTimeSec, 0)
	}

	ip := web.IP(ctx)
	location := web.Location(ctx, ip)

	u := &User{
		Username:          username,
		Email:             email,
		Password:          password,
		CreateIP:          ip,
		CreateLocation:    location,
		LastLoginTime:     uint64(reqTime.Unix()),
		LastLoginIP:       ip,
		LastLoginLocation: location,
	}

	err = db.Save(u).Error
	return u, err
}

// use bcrypt encode password
func passwordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("password hash: %w", err)
	}

	return string(hash), nil
}
