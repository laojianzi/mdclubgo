package avatar

import (
	"fmt"
)

// UserCover impl avatar.Brandable for user cover
type UserCover struct{}

// BrandType user cover type
func (u *UserCover) BrandType() string {
	return UserCoverBrandType
}

// BrandSize user cover size
func (u *UserCover) BrandSize() map[string][2]int {
	return map[string][2]int{
		"small":  {600, 336},
		"middle": {1050, 588},
		"large":  {1450, 812},
	}
}

// DefaultBrandUrls default user cover
func (u *UserCover) DefaultBrandUrls() map[string]string {
	urls := make(map[string]string)
	urls["original"] = "/default/user_cover"
	for size := range u.BrandSize() {
		urls[size] = fmt.Sprintf("/default/user_cover_%s", size)
	}

	return urls
}
