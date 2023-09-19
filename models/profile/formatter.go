package profile

import (
	"eoffice-backend/models/employee"
	"eoffice-backend/models/user"
)

type ProfileFormatter struct {
	User     user.UserFormatter         `json:"user"`
	Employee employee.EmployeeFormatter `json:"employees"`
}

func FormatProfile(profile Profile) ProfileFormatter {
	formatter := ProfileFormatter{
		User:     user.FormatUser(profile.User),
		Employee: employee.FormatEmployee(profile.Employee),
	}

	return formatter
}

func FormatProfiles(profiles []Profile) []ProfileFormatter {
	profilesFormatter := []ProfileFormatter{}

	for _, profile := range profiles {
		profileFormatter := FormatProfile(profile)
		profilesFormatter = append(profilesFormatter, profileFormatter)
	}

	return profilesFormatter
}
