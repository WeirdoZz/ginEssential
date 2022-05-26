package dto

import "ginEssential/model"

type UserDto struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}
}
