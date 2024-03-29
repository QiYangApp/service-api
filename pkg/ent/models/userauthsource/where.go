// Code generated by ent, DO NOT EDIT.

package userauthsource

import (
	"ent/models/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldUserID, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldToken, v))
}

// Channel applies equality check predicate on the "channel" field. It's identical to ChannelEQ.
func Channel(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldChannel, v))
}

// Device applies equality check predicate on the "device" field. It's identical to DeviceEQ.
func Device(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldDevice, v))
}

// DeviceDetail applies equality check predicate on the "device_detail" field. It's identical to DeviceDetailEQ.
func DeviceDetail(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldDeviceDetail, v))
}

// ClientIP applies equality check predicate on the "client_ip" field. It's identical to ClientIPEQ.
func ClientIP(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldClientIP, v))
}

// RemoteIP applies equality check predicate on the "remote_ip" field. It's identical to RemoteIPEQ.
func RemoteIP(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldRemoteIP, v))
}

// Snapshot applies equality check predicate on the "snapshot" field. It's identical to SnapshotEQ.
func Snapshot(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldSnapshot, v))
}

// LoginName applies equality check predicate on the "login_name" field. It's identical to LoginNameEQ.
func LoginName(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldLoginName, v))
}

// LoginSource applies equality check predicate on the "login_source" field. It's identical to LoginSourceEQ.
func LoginSource(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldLoginSource, v))
}

// LoginType applies equality check predicate on the "login_type" field. It's identical to LoginTypeEQ.
func LoginType(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldLoginType, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldUpdateTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldUserID, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContainsFold(FieldToken, v))
}

// ChannelEQ applies the EQ predicate on the "channel" field.
func ChannelEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldChannel, v))
}

// ChannelNEQ applies the NEQ predicate on the "channel" field.
func ChannelNEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldChannel, v))
}

// ChannelIn applies the In predicate on the "channel" field.
func ChannelIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldChannel, vs...))
}

// ChannelNotIn applies the NotIn predicate on the "channel" field.
func ChannelNotIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldChannel, vs...))
}

// ChannelGT applies the GT predicate on the "channel" field.
func ChannelGT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldChannel, v))
}

// ChannelGTE applies the GTE predicate on the "channel" field.
func ChannelGTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldChannel, v))
}

// ChannelLT applies the LT predicate on the "channel" field.
func ChannelLT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldChannel, v))
}

// ChannelLTE applies the LTE predicate on the "channel" field.
func ChannelLTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldChannel, v))
}

// ChannelContains applies the Contains predicate on the "channel" field.
func ChannelContains(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContains(FieldChannel, v))
}

// ChannelHasPrefix applies the HasPrefix predicate on the "channel" field.
func ChannelHasPrefix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasPrefix(FieldChannel, v))
}

// ChannelHasSuffix applies the HasSuffix predicate on the "channel" field.
func ChannelHasSuffix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasSuffix(FieldChannel, v))
}

// ChannelEqualFold applies the EqualFold predicate on the "channel" field.
func ChannelEqualFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEqualFold(FieldChannel, v))
}

// ChannelContainsFold applies the ContainsFold predicate on the "channel" field.
func ChannelContainsFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContainsFold(FieldChannel, v))
}

// DeviceEQ applies the EQ predicate on the "device" field.
func DeviceEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldDevice, v))
}

// DeviceNEQ applies the NEQ predicate on the "device" field.
func DeviceNEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldDevice, v))
}

// DeviceIn applies the In predicate on the "device" field.
func DeviceIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldDevice, vs...))
}

// DeviceNotIn applies the NotIn predicate on the "device" field.
func DeviceNotIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldDevice, vs...))
}

// DeviceGT applies the GT predicate on the "device" field.
func DeviceGT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldDevice, v))
}

// DeviceGTE applies the GTE predicate on the "device" field.
func DeviceGTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldDevice, v))
}

// DeviceLT applies the LT predicate on the "device" field.
func DeviceLT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldDevice, v))
}

// DeviceLTE applies the LTE predicate on the "device" field.
func DeviceLTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldDevice, v))
}

// DeviceContains applies the Contains predicate on the "device" field.
func DeviceContains(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContains(FieldDevice, v))
}

// DeviceHasPrefix applies the HasPrefix predicate on the "device" field.
func DeviceHasPrefix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasPrefix(FieldDevice, v))
}

// DeviceHasSuffix applies the HasSuffix predicate on the "device" field.
func DeviceHasSuffix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasSuffix(FieldDevice, v))
}

// DeviceEqualFold applies the EqualFold predicate on the "device" field.
func DeviceEqualFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEqualFold(FieldDevice, v))
}

// DeviceContainsFold applies the ContainsFold predicate on the "device" field.
func DeviceContainsFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContainsFold(FieldDevice, v))
}

// DeviceDetailEQ applies the EQ predicate on the "device_detail" field.
func DeviceDetailEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldDeviceDetail, v))
}

// DeviceDetailNEQ applies the NEQ predicate on the "device_detail" field.
func DeviceDetailNEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldDeviceDetail, v))
}

// DeviceDetailIn applies the In predicate on the "device_detail" field.
func DeviceDetailIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldDeviceDetail, vs...))
}

// DeviceDetailNotIn applies the NotIn predicate on the "device_detail" field.
func DeviceDetailNotIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldDeviceDetail, vs...))
}

// DeviceDetailGT applies the GT predicate on the "device_detail" field.
func DeviceDetailGT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldDeviceDetail, v))
}

// DeviceDetailGTE applies the GTE predicate on the "device_detail" field.
func DeviceDetailGTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldDeviceDetail, v))
}

// DeviceDetailLT applies the LT predicate on the "device_detail" field.
func DeviceDetailLT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldDeviceDetail, v))
}

// DeviceDetailLTE applies the LTE predicate on the "device_detail" field.
func DeviceDetailLTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldDeviceDetail, v))
}

// DeviceDetailContains applies the Contains predicate on the "device_detail" field.
func DeviceDetailContains(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContains(FieldDeviceDetail, v))
}

// DeviceDetailHasPrefix applies the HasPrefix predicate on the "device_detail" field.
func DeviceDetailHasPrefix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasPrefix(FieldDeviceDetail, v))
}

// DeviceDetailHasSuffix applies the HasSuffix predicate on the "device_detail" field.
func DeviceDetailHasSuffix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasSuffix(FieldDeviceDetail, v))
}

// DeviceDetailEqualFold applies the EqualFold predicate on the "device_detail" field.
func DeviceDetailEqualFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEqualFold(FieldDeviceDetail, v))
}

// DeviceDetailContainsFold applies the ContainsFold predicate on the "device_detail" field.
func DeviceDetailContainsFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContainsFold(FieldDeviceDetail, v))
}

// ClientIPEQ applies the EQ predicate on the "client_ip" field.
func ClientIPEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldClientIP, v))
}

// ClientIPNEQ applies the NEQ predicate on the "client_ip" field.
func ClientIPNEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldClientIP, v))
}

// ClientIPIn applies the In predicate on the "client_ip" field.
func ClientIPIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldClientIP, vs...))
}

// ClientIPNotIn applies the NotIn predicate on the "client_ip" field.
func ClientIPNotIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldClientIP, vs...))
}

// ClientIPGT applies the GT predicate on the "client_ip" field.
func ClientIPGT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldClientIP, v))
}

// ClientIPGTE applies the GTE predicate on the "client_ip" field.
func ClientIPGTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldClientIP, v))
}

// ClientIPLT applies the LT predicate on the "client_ip" field.
func ClientIPLT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldClientIP, v))
}

// ClientIPLTE applies the LTE predicate on the "client_ip" field.
func ClientIPLTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldClientIP, v))
}

// ClientIPContains applies the Contains predicate on the "client_ip" field.
func ClientIPContains(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContains(FieldClientIP, v))
}

// ClientIPHasPrefix applies the HasPrefix predicate on the "client_ip" field.
func ClientIPHasPrefix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasPrefix(FieldClientIP, v))
}

// ClientIPHasSuffix applies the HasSuffix predicate on the "client_ip" field.
func ClientIPHasSuffix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasSuffix(FieldClientIP, v))
}

// ClientIPEqualFold applies the EqualFold predicate on the "client_ip" field.
func ClientIPEqualFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEqualFold(FieldClientIP, v))
}

// ClientIPContainsFold applies the ContainsFold predicate on the "client_ip" field.
func ClientIPContainsFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContainsFold(FieldClientIP, v))
}

// RemoteIPEQ applies the EQ predicate on the "remote_ip" field.
func RemoteIPEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldRemoteIP, v))
}

// RemoteIPNEQ applies the NEQ predicate on the "remote_ip" field.
func RemoteIPNEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldRemoteIP, v))
}

// RemoteIPIn applies the In predicate on the "remote_ip" field.
func RemoteIPIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldRemoteIP, vs...))
}

// RemoteIPNotIn applies the NotIn predicate on the "remote_ip" field.
func RemoteIPNotIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldRemoteIP, vs...))
}

// RemoteIPGT applies the GT predicate on the "remote_ip" field.
func RemoteIPGT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldRemoteIP, v))
}

// RemoteIPGTE applies the GTE predicate on the "remote_ip" field.
func RemoteIPGTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldRemoteIP, v))
}

// RemoteIPLT applies the LT predicate on the "remote_ip" field.
func RemoteIPLT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldRemoteIP, v))
}

// RemoteIPLTE applies the LTE predicate on the "remote_ip" field.
func RemoteIPLTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldRemoteIP, v))
}

// RemoteIPContains applies the Contains predicate on the "remote_ip" field.
func RemoteIPContains(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContains(FieldRemoteIP, v))
}

// RemoteIPHasPrefix applies the HasPrefix predicate on the "remote_ip" field.
func RemoteIPHasPrefix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasPrefix(FieldRemoteIP, v))
}

// RemoteIPHasSuffix applies the HasSuffix predicate on the "remote_ip" field.
func RemoteIPHasSuffix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasSuffix(FieldRemoteIP, v))
}

// RemoteIPEqualFold applies the EqualFold predicate on the "remote_ip" field.
func RemoteIPEqualFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEqualFold(FieldRemoteIP, v))
}

// RemoteIPContainsFold applies the ContainsFold predicate on the "remote_ip" field.
func RemoteIPContainsFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContainsFold(FieldRemoteIP, v))
}

// SnapshotEQ applies the EQ predicate on the "snapshot" field.
func SnapshotEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldSnapshot, v))
}

// SnapshotNEQ applies the NEQ predicate on the "snapshot" field.
func SnapshotNEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldSnapshot, v))
}

// SnapshotIn applies the In predicate on the "snapshot" field.
func SnapshotIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldSnapshot, vs...))
}

// SnapshotNotIn applies the NotIn predicate on the "snapshot" field.
func SnapshotNotIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldSnapshot, vs...))
}

// SnapshotGT applies the GT predicate on the "snapshot" field.
func SnapshotGT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldSnapshot, v))
}

// SnapshotGTE applies the GTE predicate on the "snapshot" field.
func SnapshotGTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldSnapshot, v))
}

// SnapshotLT applies the LT predicate on the "snapshot" field.
func SnapshotLT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldSnapshot, v))
}

// SnapshotLTE applies the LTE predicate on the "snapshot" field.
func SnapshotLTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldSnapshot, v))
}

// SnapshotContains applies the Contains predicate on the "snapshot" field.
func SnapshotContains(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContains(FieldSnapshot, v))
}

// SnapshotHasPrefix applies the HasPrefix predicate on the "snapshot" field.
func SnapshotHasPrefix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasPrefix(FieldSnapshot, v))
}

// SnapshotHasSuffix applies the HasSuffix predicate on the "snapshot" field.
func SnapshotHasSuffix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasSuffix(FieldSnapshot, v))
}

// SnapshotEqualFold applies the EqualFold predicate on the "snapshot" field.
func SnapshotEqualFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEqualFold(FieldSnapshot, v))
}

// SnapshotContainsFold applies the ContainsFold predicate on the "snapshot" field.
func SnapshotContainsFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContainsFold(FieldSnapshot, v))
}

// LoginNameEQ applies the EQ predicate on the "login_name" field.
func LoginNameEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldLoginName, v))
}

// LoginNameNEQ applies the NEQ predicate on the "login_name" field.
func LoginNameNEQ(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldLoginName, v))
}

// LoginNameIn applies the In predicate on the "login_name" field.
func LoginNameIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldLoginName, vs...))
}

// LoginNameNotIn applies the NotIn predicate on the "login_name" field.
func LoginNameNotIn(vs ...string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldLoginName, vs...))
}

// LoginNameGT applies the GT predicate on the "login_name" field.
func LoginNameGT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldLoginName, v))
}

// LoginNameGTE applies the GTE predicate on the "login_name" field.
func LoginNameGTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldLoginName, v))
}

// LoginNameLT applies the LT predicate on the "login_name" field.
func LoginNameLT(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldLoginName, v))
}

// LoginNameLTE applies the LTE predicate on the "login_name" field.
func LoginNameLTE(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldLoginName, v))
}

// LoginNameContains applies the Contains predicate on the "login_name" field.
func LoginNameContains(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContains(FieldLoginName, v))
}

// LoginNameHasPrefix applies the HasPrefix predicate on the "login_name" field.
func LoginNameHasPrefix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasPrefix(FieldLoginName, v))
}

// LoginNameHasSuffix applies the HasSuffix predicate on the "login_name" field.
func LoginNameHasSuffix(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldHasSuffix(FieldLoginName, v))
}

// LoginNameEqualFold applies the EqualFold predicate on the "login_name" field.
func LoginNameEqualFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEqualFold(FieldLoginName, v))
}

// LoginNameContainsFold applies the ContainsFold predicate on the "login_name" field.
func LoginNameContainsFold(v string) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldContainsFold(FieldLoginName, v))
}

// LoginSourceEQ applies the EQ predicate on the "login_source" field.
func LoginSourceEQ(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldLoginSource, v))
}

// LoginSourceNEQ applies the NEQ predicate on the "login_source" field.
func LoginSourceNEQ(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldLoginSource, v))
}

// LoginSourceIn applies the In predicate on the "login_source" field.
func LoginSourceIn(vs ...int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldLoginSource, vs...))
}

// LoginSourceNotIn applies the NotIn predicate on the "login_source" field.
func LoginSourceNotIn(vs ...int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldLoginSource, vs...))
}

// LoginSourceGT applies the GT predicate on the "login_source" field.
func LoginSourceGT(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldLoginSource, v))
}

// LoginSourceGTE applies the GTE predicate on the "login_source" field.
func LoginSourceGTE(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldLoginSource, v))
}

// LoginSourceLT applies the LT predicate on the "login_source" field.
func LoginSourceLT(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldLoginSource, v))
}

// LoginSourceLTE applies the LTE predicate on the "login_source" field.
func LoginSourceLTE(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldLoginSource, v))
}

// LoginTypeEQ applies the EQ predicate on the "login_type" field.
func LoginTypeEQ(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldLoginType, v))
}

// LoginTypeNEQ applies the NEQ predicate on the "login_type" field.
func LoginTypeNEQ(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldLoginType, v))
}

// LoginTypeIn applies the In predicate on the "login_type" field.
func LoginTypeIn(vs ...int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldLoginType, vs...))
}

// LoginTypeNotIn applies the NotIn predicate on the "login_type" field.
func LoginTypeNotIn(vs ...int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldLoginType, vs...))
}

// LoginTypeGT applies the GT predicate on the "login_type" field.
func LoginTypeGT(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldLoginType, v))
}

// LoginTypeGTE applies the GTE predicate on the "login_type" field.
func LoginTypeGTE(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldLoginType, v))
}

// LoginTypeLT applies the LT predicate on the "login_type" field.
func LoginTypeLT(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldLoginType, v))
}

// LoginTypeLTE applies the LTE predicate on the "login_type" field.
func LoginTypeLTE(v int) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldLoginType, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.FieldLTE(FieldUpdateTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserAuthSource) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserAuthSource) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserAuthSource) predicate.UserAuthSource {
	return predicate.UserAuthSource(sql.NotPredicates(p))
}
