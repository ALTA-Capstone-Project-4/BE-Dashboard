package delivery

import "warehouse/features/user"

type UserResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Phone           string `json:"phone,omitempty"`
	Email           string `json:"email,omitempty"`
	Address         string `json:"address,omitempty"`
	Photo           string `json:"photo,omitempty"`
	Role            string `json:"role,omitempty"`
	FileKTP         string `json:"file_ktp,omitempty"`
	GudangName      string `json:"gudangname,omitempty"`
	GudangLocation  string `json:"gudanglocation,omitempty"`
	GudangPhoto     string `json:"gudangphoto,omitempty"`
	GudangLatitude  string `json:"gudanglatitude,omitempty"`
	GudangLongitude string `json:"gudanglongitude,omitempty"`
}

func fromCore(data user.Core) UserResponse {
	return UserResponse{
		ID:              uint(data.ID),
		Name:            data.Name,
		Phone:           data.Phone,
		Email:           data.Email,
		Address:         data.Address,
		Photo:           data.Photo,
		Role:            data.Role,
		FileKTP:         data.FileKTP,
		GudangName:      data.GudangName,
		GudangLocation:  data.GudangLocation,
		GudangPhoto:     data.GudangPhoto,
		GudangLatitude:  data.GudangLatitude,
		GudangLongitude: data.GudangLongitude,
	}
}
