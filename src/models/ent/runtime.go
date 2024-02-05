// Code generated by ent, DO NOT EDIT.

package ent

import (
	"service-api/src/models/ent/member"
	"service-api/src/models/ent/memberauthorizelog"
	"service-api/src/models/ent/memberrelatedrole"
	"service-api/src/models/ent/memberrole"
	"service-api/src/models/ent/memberrolerelatedpermission"
	"service-api/src/models/ent/permissiongroup"
	"service-api/src/models/ent/permissionrelatedrouter"
	"service-api/src/models/ent/router"
	"service-api/src/models/ent/schema"
	"service-api/src/models/ent/sourcedata"
	"service-api/src/models/ent/wakatime"
	"service-api/src/models/ent/wakatimecategory"
	"service-api/src/models/ent/wakatimedependency"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	memberMixin := schema.Member{}.Mixin()
	memberMixinFields0 := memberMixin[0].Fields()
	_ = memberMixinFields0
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescCreateTime is the schema descriptor for create_time field.
	memberDescCreateTime := memberMixinFields0[0].Descriptor()
	// member.DefaultCreateTime holds the default value on creation for the create_time field.
	member.DefaultCreateTime = memberDescCreateTime.Default.(func() time.Time)
	// memberDescUpdateTime is the schema descriptor for update_time field.
	memberDescUpdateTime := memberMixinFields0[1].Descriptor()
	// member.DefaultUpdateTime holds the default value on creation for the update_time field.
	member.DefaultUpdateTime = memberDescUpdateTime.Default.(func() time.Time)
	// member.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	member.UpdateDefaultUpdateTime = memberDescUpdateTime.UpdateDefault.(func() time.Time)
	// memberDescEmail is the schema descriptor for email field.
	memberDescEmail := memberFields[1].Descriptor()
	// member.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	member.EmailValidator = func() func(string) error {
		validators := memberDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescPasswordSing is the schema descriptor for password_sing field.
	memberDescPasswordSing := memberFields[3].Descriptor()
	// member.PasswordSingValidator is a validator for the "password_sing" field. It is called by the builders before save.
	member.PasswordSingValidator = func() func(string) error {
		validators := memberDescPasswordSing.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password_sing string) error {
			for _, fn := range fns {
				if err := fn(password_sing); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescPassword is the schema descriptor for password field.
	memberDescPassword := memberFields[4].Descriptor()
	// member.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	member.PasswordValidator = func() func(string) error {
		validators := memberDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescMobile is the schema descriptor for mobile field.
	memberDescMobile := memberFields[5].Descriptor()
	// member.MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	member.MobileValidator = memberDescMobile.Validators[0].(func(string) error)
	// memberDescNickname is the schema descriptor for nickname field.
	memberDescNickname := memberFields[6].Descriptor()
	// member.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	member.NicknameValidator = func() func(string) error {
		validators := memberDescNickname.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(nickname string) error {
			for _, fn := range fns {
				if err := fn(nickname); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescState is the schema descriptor for state field.
	memberDescState := memberFields[7].Descriptor()
	// member.StateValidator is a validator for the "state" field. It is called by the builders before save.
	member.StateValidator = func() func(string) error {
		validators := memberDescState.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(state string) error {
			for _, fn := range fns {
				if err := fn(state); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescID is the schema descriptor for id field.
	memberDescID := memberFields[0].Descriptor()
	// member.DefaultID holds the default value on creation for the id field.
	member.DefaultID = memberDescID.Default.(func() uuid.UUID)
	memberauthorizelogMixin := schema.MemberAuthorizeLog{}.Mixin()
	memberauthorizelogMixinFields0 := memberauthorizelogMixin[0].Fields()
	_ = memberauthorizelogMixinFields0
	memberauthorizelogFields := schema.MemberAuthorizeLog{}.Fields()
	_ = memberauthorizelogFields
	// memberauthorizelogDescCreateTime is the schema descriptor for create_time field.
	memberauthorizelogDescCreateTime := memberauthorizelogMixinFields0[0].Descriptor()
	// memberauthorizelog.DefaultCreateTime holds the default value on creation for the create_time field.
	memberauthorizelog.DefaultCreateTime = memberauthorizelogDescCreateTime.Default.(func() time.Time)
	// memberauthorizelogDescUpdateTime is the schema descriptor for update_time field.
	memberauthorizelogDescUpdateTime := memberauthorizelogMixinFields0[1].Descriptor()
	// memberauthorizelog.DefaultUpdateTime holds the default value on creation for the update_time field.
	memberauthorizelog.DefaultUpdateTime = memberauthorizelogDescUpdateTime.Default.(func() time.Time)
	// memberauthorizelog.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	memberauthorizelog.UpdateDefaultUpdateTime = memberauthorizelogDescUpdateTime.UpdateDefault.(func() time.Time)
	// memberauthorizelogDescMemberID is the schema descriptor for member_id field.
	memberauthorizelogDescMemberID := memberauthorizelogFields[1].Descriptor()
	// memberauthorizelog.DefaultMemberID holds the default value on creation for the member_id field.
	memberauthorizelog.DefaultMemberID = memberauthorizelogDescMemberID.Default.(func() uuid.UUID)
	// memberauthorizelogDescToken is the schema descriptor for token field.
	memberauthorizelogDescToken := memberauthorizelogFields[2].Descriptor()
	// memberauthorizelog.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	memberauthorizelog.TokenValidator = func() func(string) error {
		validators := memberauthorizelogDescToken.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(token string) error {
			for _, fn := range fns {
				if err := fn(token); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberauthorizelogDescChannel is the schema descriptor for channel field.
	memberauthorizelogDescChannel := memberauthorizelogFields[3].Descriptor()
	// memberauthorizelog.ChannelValidator is a validator for the "channel" field. It is called by the builders before save.
	memberauthorizelog.ChannelValidator = memberauthorizelogDescChannel.Validators[0].(func(string) error)
	// memberauthorizelogDescDevice is the schema descriptor for device field.
	memberauthorizelogDescDevice := memberauthorizelogFields[4].Descriptor()
	// memberauthorizelog.DeviceValidator is a validator for the "device" field. It is called by the builders before save.
	memberauthorizelog.DeviceValidator = memberauthorizelogDescDevice.Validators[0].(func(string) error)
	// memberauthorizelogDescDeviceDetail is the schema descriptor for device_detail field.
	memberauthorizelogDescDeviceDetail := memberauthorizelogFields[5].Descriptor()
	// memberauthorizelog.DeviceDetailValidator is a validator for the "device_detail" field. It is called by the builders before save.
	memberauthorizelog.DeviceDetailValidator = memberauthorizelogDescDeviceDetail.Validators[0].(func(string) error)
	// memberauthorizelogDescClientIP is the schema descriptor for client_ip field.
	memberauthorizelogDescClientIP := memberauthorizelogFields[6].Descriptor()
	// memberauthorizelog.ClientIPValidator is a validator for the "client_ip" field. It is called by the builders before save.
	memberauthorizelog.ClientIPValidator = func() func(string) error {
		validators := memberauthorizelogDescClientIP.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(client_ip string) error {
			for _, fn := range fns {
				if err := fn(client_ip); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberauthorizelogDescRemoteIP is the schema descriptor for remote_ip field.
	memberauthorizelogDescRemoteIP := memberauthorizelogFields[7].Descriptor()
	// memberauthorizelog.RemoteIPValidator is a validator for the "remote_ip" field. It is called by the builders before save.
	memberauthorizelog.RemoteIPValidator = func() func(string) error {
		validators := memberauthorizelogDescRemoteIP.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(remote_ip string) error {
			for _, fn := range fns {
				if err := fn(remote_ip); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberauthorizelogDescSnapshot is the schema descriptor for snapshot field.
	memberauthorizelogDescSnapshot := memberauthorizelogFields[8].Descriptor()
	// memberauthorizelog.SnapshotValidator is a validator for the "snapshot" field. It is called by the builders before save.
	memberauthorizelog.SnapshotValidator = memberauthorizelogDescSnapshot.Validators[0].(func(string) error)
	// memberauthorizelogDescID is the schema descriptor for id field.
	memberauthorizelogDescID := memberauthorizelogFields[0].Descriptor()
	// memberauthorizelog.DefaultID holds the default value on creation for the id field.
	memberauthorizelog.DefaultID = memberauthorizelogDescID.Default.(func() uuid.UUID)
	memberrelatedroleMixin := schema.MemberRelatedRole{}.Mixin()
	memberrelatedroleMixinFields0 := memberrelatedroleMixin[0].Fields()
	_ = memberrelatedroleMixinFields0
	memberrelatedroleFields := schema.MemberRelatedRole{}.Fields()
	_ = memberrelatedroleFields
	// memberrelatedroleDescCreateTime is the schema descriptor for create_time field.
	memberrelatedroleDescCreateTime := memberrelatedroleMixinFields0[0].Descriptor()
	// memberrelatedrole.DefaultCreateTime holds the default value on creation for the create_time field.
	memberrelatedrole.DefaultCreateTime = memberrelatedroleDescCreateTime.Default.(func() time.Time)
	// memberrelatedroleDescUpdateTime is the schema descriptor for update_time field.
	memberrelatedroleDescUpdateTime := memberrelatedroleMixinFields0[1].Descriptor()
	// memberrelatedrole.DefaultUpdateTime holds the default value on creation for the update_time field.
	memberrelatedrole.DefaultUpdateTime = memberrelatedroleDescUpdateTime.Default.(func() time.Time)
	// memberrelatedrole.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	memberrelatedrole.UpdateDefaultUpdateTime = memberrelatedroleDescUpdateTime.UpdateDefault.(func() time.Time)
	// memberrelatedroleDescID is the schema descriptor for id field.
	memberrelatedroleDescID := memberrelatedroleFields[0].Descriptor()
	// memberrelatedrole.DefaultID holds the default value on creation for the id field.
	memberrelatedrole.DefaultID = memberrelatedroleDescID.Default.(func() uuid.UUID)
	memberroleMixin := schema.MemberRole{}.Mixin()
	memberroleMixinFields0 := memberroleMixin[0].Fields()
	_ = memberroleMixinFields0
	memberroleFields := schema.MemberRole{}.Fields()
	_ = memberroleFields
	// memberroleDescCreateTime is the schema descriptor for create_time field.
	memberroleDescCreateTime := memberroleMixinFields0[0].Descriptor()
	// memberrole.DefaultCreateTime holds the default value on creation for the create_time field.
	memberrole.DefaultCreateTime = memberroleDescCreateTime.Default.(func() time.Time)
	// memberroleDescUpdateTime is the schema descriptor for update_time field.
	memberroleDescUpdateTime := memberroleMixinFields0[1].Descriptor()
	// memberrole.DefaultUpdateTime holds the default value on creation for the update_time field.
	memberrole.DefaultUpdateTime = memberroleDescUpdateTime.Default.(func() time.Time)
	// memberrole.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	memberrole.UpdateDefaultUpdateTime = memberroleDescUpdateTime.UpdateDefault.(func() time.Time)
	// memberroleDescRoleName is the schema descriptor for role_name field.
	memberroleDescRoleName := memberroleFields[1].Descriptor()
	// memberrole.RoleNameValidator is a validator for the "role_name" field. It is called by the builders before save.
	memberrole.RoleNameValidator = func() func(string) error {
		validators := memberroleDescRoleName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(role_name string) error {
			for _, fn := range fns {
				if err := fn(role_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberroleDescState is the schema descriptor for state field.
	memberroleDescState := memberroleFields[2].Descriptor()
	// memberrole.DefaultState holds the default value on creation for the state field.
	memberrole.DefaultState = memberroleDescState.Default.(string)
	// memberrole.StateValidator is a validator for the "state" field. It is called by the builders before save.
	memberrole.StateValidator = func() func(string) error {
		validators := memberroleDescState.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(state string) error {
			for _, fn := range fns {
				if err := fn(state); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberroleDescID is the schema descriptor for id field.
	memberroleDescID := memberroleFields[0].Descriptor()
	// memberrole.DefaultID holds the default value on creation for the id field.
	memberrole.DefaultID = memberroleDescID.Default.(func() uuid.UUID)
	memberrolerelatedpermissionMixin := schema.MemberRoleRelatedPermission{}.Mixin()
	memberrolerelatedpermissionMixinFields0 := memberrolerelatedpermissionMixin[0].Fields()
	_ = memberrolerelatedpermissionMixinFields0
	memberrolerelatedpermissionFields := schema.MemberRoleRelatedPermission{}.Fields()
	_ = memberrolerelatedpermissionFields
	// memberrolerelatedpermissionDescCreateTime is the schema descriptor for create_time field.
	memberrolerelatedpermissionDescCreateTime := memberrolerelatedpermissionMixinFields0[0].Descriptor()
	// memberrolerelatedpermission.DefaultCreateTime holds the default value on creation for the create_time field.
	memberrolerelatedpermission.DefaultCreateTime = memberrolerelatedpermissionDescCreateTime.Default.(func() time.Time)
	// memberrolerelatedpermissionDescUpdateTime is the schema descriptor for update_time field.
	memberrolerelatedpermissionDescUpdateTime := memberrolerelatedpermissionMixinFields0[1].Descriptor()
	// memberrolerelatedpermission.DefaultUpdateTime holds the default value on creation for the update_time field.
	memberrolerelatedpermission.DefaultUpdateTime = memberrolerelatedpermissionDescUpdateTime.Default.(func() time.Time)
	// memberrolerelatedpermission.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	memberrolerelatedpermission.UpdateDefaultUpdateTime = memberrolerelatedpermissionDescUpdateTime.UpdateDefault.(func() time.Time)
	// memberrolerelatedpermissionDescID is the schema descriptor for id field.
	memberrolerelatedpermissionDescID := memberrolerelatedpermissionFields[0].Descriptor()
	// memberrolerelatedpermission.DefaultID holds the default value on creation for the id field.
	memberrolerelatedpermission.DefaultID = memberrolerelatedpermissionDescID.Default.(func() uuid.UUID)
	permissiongroupMixin := schema.PermissionGroup{}.Mixin()
	permissiongroupMixinFields0 := permissiongroupMixin[0].Fields()
	_ = permissiongroupMixinFields0
	permissiongroupFields := schema.PermissionGroup{}.Fields()
	_ = permissiongroupFields
	// permissiongroupDescCreateTime is the schema descriptor for create_time field.
	permissiongroupDescCreateTime := permissiongroupMixinFields0[0].Descriptor()
	// permissiongroup.DefaultCreateTime holds the default value on creation for the create_time field.
	permissiongroup.DefaultCreateTime = permissiongroupDescCreateTime.Default.(func() time.Time)
	// permissiongroupDescUpdateTime is the schema descriptor for update_time field.
	permissiongroupDescUpdateTime := permissiongroupMixinFields0[1].Descriptor()
	// permissiongroup.DefaultUpdateTime holds the default value on creation for the update_time field.
	permissiongroup.DefaultUpdateTime = permissiongroupDescUpdateTime.Default.(func() time.Time)
	// permissiongroup.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	permissiongroup.UpdateDefaultUpdateTime = permissiongroupDescUpdateTime.UpdateDefault.(func() time.Time)
	// permissiongroupDescPermissionName is the schema descriptor for permission_name field.
	permissiongroupDescPermissionName := permissiongroupFields[1].Descriptor()
	// permissiongroup.DefaultPermissionName holds the default value on creation for the permission_name field.
	permissiongroup.DefaultPermissionName = permissiongroupDescPermissionName.Default.(string)
	// permissiongroup.PermissionNameValidator is a validator for the "permission_name" field. It is called by the builders before save.
	permissiongroup.PermissionNameValidator = func() func(string) error {
		validators := permissiongroupDescPermissionName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(permission_name string) error {
			for _, fn := range fns {
				if err := fn(permission_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// permissiongroupDescIoc is the schema descriptor for ioc field.
	permissiongroupDescIoc := permissiongroupFields[2].Descriptor()
	// permissiongroup.DefaultIoc holds the default value on creation for the ioc field.
	permissiongroup.DefaultIoc = permissiongroupDescIoc.Default.(string)
	// permissiongroup.IocValidator is a validator for the "ioc" field. It is called by the builders before save.
	permissiongroup.IocValidator = func() func(string) error {
		validators := permissiongroupDescIoc.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(ioc string) error {
			for _, fn := range fns {
				if err := fn(ioc); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// permissiongroupDescSort is the schema descriptor for sort field.
	permissiongroupDescSort := permissiongroupFields[3].Descriptor()
	// permissiongroup.DefaultSort holds the default value on creation for the sort field.
	permissiongroup.DefaultSort = permissiongroupDescSort.Default.(int32)
	// permissiongroupDescLeft is the schema descriptor for left field.
	permissiongroupDescLeft := permissiongroupFields[4].Descriptor()
	// permissiongroup.DefaultLeft holds the default value on creation for the left field.
	permissiongroup.DefaultLeft = permissiongroupDescLeft.Default.(int32)
	// permissiongroupDescRight is the schema descriptor for right field.
	permissiongroupDescRight := permissiongroupFields[5].Descriptor()
	// permissiongroup.DefaultRight holds the default value on creation for the right field.
	permissiongroup.DefaultRight = permissiongroupDescRight.Default.(int32)
	// permissiongroupDescState is the schema descriptor for state field.
	permissiongroupDescState := permissiongroupFields[6].Descriptor()
	// permissiongroup.DefaultState holds the default value on creation for the state field.
	permissiongroup.DefaultState = permissiongroupDescState.Default.(string)
	// permissiongroup.StateValidator is a validator for the "state" field. It is called by the builders before save.
	permissiongroup.StateValidator = permissiongroupDescState.Validators[0].(func(string) error)
	// permissiongroupDescID is the schema descriptor for id field.
	permissiongroupDescID := permissiongroupFields[0].Descriptor()
	// permissiongroup.DefaultID holds the default value on creation for the id field.
	permissiongroup.DefaultID = permissiongroupDescID.Default.(func() uuid.UUID)
	permissionrelatedrouterMixin := schema.PermissionRelatedRouter{}.Mixin()
	permissionrelatedrouterMixinFields0 := permissionrelatedrouterMixin[0].Fields()
	_ = permissionrelatedrouterMixinFields0
	permissionrelatedrouterFields := schema.PermissionRelatedRouter{}.Fields()
	_ = permissionrelatedrouterFields
	// permissionrelatedrouterDescCreateTime is the schema descriptor for create_time field.
	permissionrelatedrouterDescCreateTime := permissionrelatedrouterMixinFields0[0].Descriptor()
	// permissionrelatedrouter.DefaultCreateTime holds the default value on creation for the create_time field.
	permissionrelatedrouter.DefaultCreateTime = permissionrelatedrouterDescCreateTime.Default.(func() time.Time)
	// permissionrelatedrouterDescUpdateTime is the schema descriptor for update_time field.
	permissionrelatedrouterDescUpdateTime := permissionrelatedrouterMixinFields0[1].Descriptor()
	// permissionrelatedrouter.DefaultUpdateTime holds the default value on creation for the update_time field.
	permissionrelatedrouter.DefaultUpdateTime = permissionrelatedrouterDescUpdateTime.Default.(func() time.Time)
	// permissionrelatedrouter.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	permissionrelatedrouter.UpdateDefaultUpdateTime = permissionrelatedrouterDescUpdateTime.UpdateDefault.(func() time.Time)
	// permissionrelatedrouterDescID is the schema descriptor for id field.
	permissionrelatedrouterDescID := permissionrelatedrouterFields[0].Descriptor()
	// permissionrelatedrouter.DefaultID holds the default value on creation for the id field.
	permissionrelatedrouter.DefaultID = permissionrelatedrouterDescID.Default.(func() uuid.UUID)
	routerMixin := schema.Router{}.Mixin()
	routerMixinFields0 := routerMixin[0].Fields()
	_ = routerMixinFields0
	routerFields := schema.Router{}.Fields()
	_ = routerFields
	// routerDescCreateTime is the schema descriptor for create_time field.
	routerDescCreateTime := routerMixinFields0[0].Descriptor()
	// router.DefaultCreateTime holds the default value on creation for the create_time field.
	router.DefaultCreateTime = routerDescCreateTime.Default.(func() time.Time)
	// routerDescUpdateTime is the schema descriptor for update_time field.
	routerDescUpdateTime := routerMixinFields0[1].Descriptor()
	// router.DefaultUpdateTime holds the default value on creation for the update_time field.
	router.DefaultUpdateTime = routerDescUpdateTime.Default.(func() time.Time)
	// router.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	router.UpdateDefaultUpdateTime = routerDescUpdateTime.UpdateDefault.(func() time.Time)
	// routerDescRouteName is the schema descriptor for route_name field.
	routerDescRouteName := routerFields[1].Descriptor()
	// router.DefaultRouteName holds the default value on creation for the route_name field.
	router.DefaultRouteName = routerDescRouteName.Default.(string)
	// router.RouteNameValidator is a validator for the "route_name" field. It is called by the builders before save.
	router.RouteNameValidator = func() func(string) error {
		validators := routerDescRouteName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(route_name string) error {
			for _, fn := range fns {
				if err := fn(route_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// routerDescRoute is the schema descriptor for route field.
	routerDescRoute := routerFields[2].Descriptor()
	// router.DefaultRoute holds the default value on creation for the route field.
	router.DefaultRoute = routerDescRoute.Default.(string)
	// router.RouteValidator is a validator for the "route" field. It is called by the builders before save.
	router.RouteValidator = func() func(string) error {
		validators := routerDescRoute.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(route string) error {
			for _, fn := range fns {
				if err := fn(route); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// routerDescDescription is the schema descriptor for description field.
	routerDescDescription := routerFields[3].Descriptor()
	// router.DefaultDescription holds the default value on creation for the description field.
	router.DefaultDescription = routerDescDescription.Default.(string)
	// router.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	router.DescriptionValidator = func() func(string) error {
		validators := routerDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// routerDescState is the schema descriptor for state field.
	routerDescState := routerFields[4].Descriptor()
	// router.DefaultState holds the default value on creation for the state field.
	router.DefaultState = routerDescState.Default.(string)
	// router.StateValidator is a validator for the "state" field. It is called by the builders before save.
	router.StateValidator = func() func(string) error {
		validators := routerDescState.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(state string) error {
			for _, fn := range fns {
				if err := fn(state); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// routerDescID is the schema descriptor for id field.
	routerDescID := routerFields[0].Descriptor()
	// router.DefaultID holds the default value on creation for the id field.
	router.DefaultID = routerDescID.Default.(func() uuid.UUID)
	sourcedataMixin := schema.SourceData{}.Mixin()
	sourcedataMixinFields0 := sourcedataMixin[0].Fields()
	_ = sourcedataMixinFields0
	sourcedataFields := schema.SourceData{}.Fields()
	_ = sourcedataFields
	// sourcedataDescCreateTime is the schema descriptor for create_time field.
	sourcedataDescCreateTime := sourcedataMixinFields0[0].Descriptor()
	// sourcedata.DefaultCreateTime holds the default value on creation for the create_time field.
	sourcedata.DefaultCreateTime = sourcedataDescCreateTime.Default.(func() time.Time)
	// sourcedataDescUpdateTime is the schema descriptor for update_time field.
	sourcedataDescUpdateTime := sourcedataMixinFields0[1].Descriptor()
	// sourcedata.DefaultUpdateTime holds the default value on creation for the update_time field.
	sourcedata.DefaultUpdateTime = sourcedataDescUpdateTime.Default.(func() time.Time)
	// sourcedata.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	sourcedata.UpdateDefaultUpdateTime = sourcedataDescUpdateTime.UpdateDefault.(func() time.Time)
	// sourcedataDescType is the schema descriptor for type field.
	sourcedataDescType := sourcedataFields[2].Descriptor()
	// sourcedata.DefaultType holds the default value on creation for the type field.
	sourcedata.DefaultType = sourcedataDescType.Default.(string)
	// sourcedata.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	sourcedata.TypeValidator = func() func(string) error {
		validators := sourcedataDescType.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_type string) error {
			for _, fn := range fns {
				if err := fn(_type); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// sourcedataDescSubType is the schema descriptor for sub_type field.
	sourcedataDescSubType := sourcedataFields[3].Descriptor()
	// sourcedata.DefaultSubType holds the default value on creation for the sub_type field.
	sourcedata.DefaultSubType = sourcedataDescSubType.Default.(string)
	// sourcedata.SubTypeValidator is a validator for the "sub_type" field. It is called by the builders before save.
	sourcedata.SubTypeValidator = sourcedataDescSubType.Validators[0].(func(string) error)
	// sourcedataDescInfo is the schema descriptor for info field.
	sourcedataDescInfo := sourcedataFields[4].Descriptor()
	// sourcedata.DefaultInfo holds the default value on creation for the info field.
	sourcedata.DefaultInfo = sourcedataDescInfo.Default.(string)
	// sourcedataDescSnapshot is the schema descriptor for snapshot field.
	sourcedataDescSnapshot := sourcedataFields[5].Descriptor()
	// sourcedata.DefaultSnapshot holds the default value on creation for the snapshot field.
	sourcedata.DefaultSnapshot = sourcedataDescSnapshot.Default.(string)
	// sourcedataDescID is the schema descriptor for id field.
	sourcedataDescID := sourcedataFields[0].Descriptor()
	// sourcedata.DefaultID holds the default value on creation for the id field.
	sourcedata.DefaultID = sourcedataDescID.Default.(func() uuid.UUID)
	wakatimeMixin := schema.Wakatime{}.Mixin()
	wakatimeMixinFields0 := wakatimeMixin[0].Fields()
	_ = wakatimeMixinFields0
	wakatimeFields := schema.Wakatime{}.Fields()
	_ = wakatimeFields
	// wakatimeDescCreateTime is the schema descriptor for create_time field.
	wakatimeDescCreateTime := wakatimeMixinFields0[0].Descriptor()
	// wakatime.DefaultCreateTime holds the default value on creation for the create_time field.
	wakatime.DefaultCreateTime = wakatimeDescCreateTime.Default.(func() time.Time)
	// wakatimeDescUpdateTime is the schema descriptor for update_time field.
	wakatimeDescUpdateTime := wakatimeMixinFields0[1].Descriptor()
	// wakatime.DefaultUpdateTime holds the default value on creation for the update_time field.
	wakatime.DefaultUpdateTime = wakatimeDescUpdateTime.Default.(func() time.Time)
	// wakatime.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	wakatime.UpdateDefaultUpdateTime = wakatimeDescUpdateTime.UpdateDefault.(func() time.Time)
	// wakatimeDescKey is the schema descriptor for key field.
	wakatimeDescKey := wakatimeFields[2].Descriptor()
	// wakatime.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	wakatime.KeyValidator = wakatimeDescKey.Validators[0].(func(string) error)
	// wakatimeDescAPI is the schema descriptor for api field.
	wakatimeDescAPI := wakatimeFields[3].Descriptor()
	// wakatime.APIValidator is a validator for the "api" field. It is called by the builders before save.
	wakatime.APIValidator = wakatimeDescAPI.Validators[0].(func(string) error)
	// wakatimeDescState is the schema descriptor for state field.
	wakatimeDescState := wakatimeFields[4].Descriptor()
	// wakatime.StateValidator is a validator for the "state" field. It is called by the builders before save.
	wakatime.StateValidator = wakatimeDescState.Validators[0].(func(string) error)
	// wakatimeDescID is the schema descriptor for id field.
	wakatimeDescID := wakatimeFields[0].Descriptor()
	// wakatime.DefaultID holds the default value on creation for the id field.
	wakatime.DefaultID = wakatimeDescID.Default.(func() uuid.UUID)
	wakatimecategoryMixin := schema.WakatimeCategory{}.Mixin()
	wakatimecategoryMixinFields0 := wakatimecategoryMixin[0].Fields()
	_ = wakatimecategoryMixinFields0
	wakatimecategoryFields := schema.WakatimeCategory{}.Fields()
	_ = wakatimecategoryFields
	// wakatimecategoryDescCreateTime is the schema descriptor for create_time field.
	wakatimecategoryDescCreateTime := wakatimecategoryMixinFields0[0].Descriptor()
	// wakatimecategory.DefaultCreateTime holds the default value on creation for the create_time field.
	wakatimecategory.DefaultCreateTime = wakatimecategoryDescCreateTime.Default.(func() time.Time)
	// wakatimecategoryDescUpdateTime is the schema descriptor for update_time field.
	wakatimecategoryDescUpdateTime := wakatimecategoryMixinFields0[1].Descriptor()
	// wakatimecategory.DefaultUpdateTime holds the default value on creation for the update_time field.
	wakatimecategory.DefaultUpdateTime = wakatimecategoryDescUpdateTime.Default.(func() time.Time)
	// wakatimecategory.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	wakatimecategory.UpdateDefaultUpdateTime = wakatimecategoryDescUpdateTime.UpdateDefault.(func() time.Time)
	// wakatimecategoryDescName is the schema descriptor for name field.
	wakatimecategoryDescName := wakatimecategoryFields[3].Descriptor()
	// wakatimecategory.DefaultName holds the default value on creation for the name field.
	wakatimecategory.DefaultName = wakatimecategoryDescName.Default.(string)
	// wakatimecategory.NameValidator is a validator for the "name" field. It is called by the builders before save.
	wakatimecategory.NameValidator = wakatimecategoryDescName.Validators[0].(func(string) error)
	// wakatimecategoryDescTotalSeconds is the schema descriptor for total_seconds field.
	wakatimecategoryDescTotalSeconds := wakatimecategoryFields[4].Descriptor()
	// wakatimecategory.DefaultTotalSeconds holds the default value on creation for the total_seconds field.
	wakatimecategory.DefaultTotalSeconds = wakatimecategoryDescTotalSeconds.Default.(int64)
	// wakatimecategoryDescID is the schema descriptor for id field.
	wakatimecategoryDescID := wakatimecategoryFields[0].Descriptor()
	// wakatimecategory.DefaultID holds the default value on creation for the id field.
	wakatimecategory.DefaultID = wakatimecategoryDescID.Default.(func() uuid.UUID)
	wakatimedependencyMixin := schema.WakatimeDependency{}.Mixin()
	wakatimedependencyMixinFields0 := wakatimedependencyMixin[0].Fields()
	_ = wakatimedependencyMixinFields0
	wakatimedependencyFields := schema.WakatimeDependency{}.Fields()
	_ = wakatimedependencyFields
	// wakatimedependencyDescCreateTime is the schema descriptor for create_time field.
	wakatimedependencyDescCreateTime := wakatimedependencyMixinFields0[0].Descriptor()
	// wakatimedependency.DefaultCreateTime holds the default value on creation for the create_time field.
	wakatimedependency.DefaultCreateTime = wakatimedependencyDescCreateTime.Default.(func() time.Time)
	// wakatimedependencyDescUpdateTime is the schema descriptor for update_time field.
	wakatimedependencyDescUpdateTime := wakatimedependencyMixinFields0[1].Descriptor()
	// wakatimedependency.DefaultUpdateTime holds the default value on creation for the update_time field.
	wakatimedependency.DefaultUpdateTime = wakatimedependencyDescUpdateTime.Default.(func() time.Time)
	// wakatimedependency.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	wakatimedependency.UpdateDefaultUpdateTime = wakatimedependencyDescUpdateTime.UpdateDefault.(func() time.Time)
	// wakatimedependencyDescName is the schema descriptor for name field.
	wakatimedependencyDescName := wakatimedependencyFields[3].Descriptor()
	// wakatimedependency.DefaultName holds the default value on creation for the name field.
	wakatimedependency.DefaultName = wakatimedependencyDescName.Default.(string)
	// wakatimedependency.NameValidator is a validator for the "name" field. It is called by the builders before save.
	wakatimedependency.NameValidator = wakatimedependencyDescName.Validators[0].(func(string) error)
	// wakatimedependencyDescTotalSeconds is the schema descriptor for total_seconds field.
	wakatimedependencyDescTotalSeconds := wakatimedependencyFields[4].Descriptor()
	// wakatimedependency.DefaultTotalSeconds holds the default value on creation for the total_seconds field.
	wakatimedependency.DefaultTotalSeconds = wakatimedependencyDescTotalSeconds.Default.(int64)
	// wakatimedependencyDescID is the schema descriptor for id field.
	wakatimedependencyDescID := wakatimedependencyFields[0].Descriptor()
	// wakatimedependency.DefaultID holds the default value on creation for the id field.
	wakatimedependency.DefaultID = wakatimedependencyDescID.Default.(func() uuid.UUID)
}
