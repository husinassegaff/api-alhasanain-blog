package response

import "api-alhasanain-blog/structs"

func CreateUserResponse(user structs.User) structs.UserResponse {
	var userResponse structs.UserResponse

	userResponse.ID = user.ID
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.Role = user.Role
	userResponse.Token = user.Token

	return userResponse
}
