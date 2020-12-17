package avatar

import (
	"bytes"
	"image/png"

	"github.com/laojianzi/mdavatar"
	"github.com/laojianzi/mdavatar/style"

	"github.com/laojianzi/mdclubgo/db"
	"github.com/laojianzi/mdclubgo/internal/database"
)

// UserAvatar impl avatar.Brandable for user avatar
type UserAvatar struct{}

// BrandType user avatar type
func (u *UserAvatar) BrandType() string {
	return UserAvatarBrandType
}

// BrandSize user avatar size
func (u *UserAvatar) BrandSize() map[string][2]int {
	return map[string][2]int{
		"small":  {64, 64},
		"middle": {128, 128},
		"large":  {256, 256},
	}
}

// DefaultBrandUrls default user avatar
func (u *UserAvatar) DefaultBrandUrls() map[string]string {
	return make(map[string]string)
}

// DeleteUserAvatar delete user avatar and reset to default avatar
func DeleteUserAvatar(userID uint) (*database.User, error) {
	var u database.User
	brand := new(UserAvatar)
	if err := db.Instance().Find(&u, userID).Error; err != nil {
		return nil, err
	}

	if u.Avatar != "" {
		if err := DeleteImage(brand, userID, u.Avatar); err != nil {
			return nil, err
		}
	}

	avatar, err := mdavatar.New(u.Username).Builds(style.NewCircle)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer([]byte{})
	if err = png.Encode(buf, avatar); err != nil {
		return nil, err
	}

	filename, err := UploadImage(brand, userID, "image/png", buf)
	if err != nil {
		return nil, err
	}

	u.Avatar = filename
	err = db.Instance().Model(&u).Select("avatar").Updates(u).Error
	return &u, err
}
