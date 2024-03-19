package user

import "service-api/internal/modules/optional"

// CreateUserOverwriteOptions are an optional options who overwrite system defaults on user creation
type CreateUserOverwriteOptions struct {
	KeepEmailPrivate             optional.Option[bool]
	Visibility                   *VisibleType
	AllowCreateOrganization      optional.Option[bool]
	EmailNotificationsPreference *string
	MaxRepoCreation              *int
	Theme                        *string
	IsRestricted                 optional.Option[bool]
	IsActive                     optional.Option[bool]
}
