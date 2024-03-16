// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"context"
	"reflect"
	"service-api/internal/ent"
	"service-api/internal/entity/models/auth"
)

var registeredConfigs = map[auth.Type]func() auth.Config{}

// RegisterTypeConfig register a config for a provided type
func RegisterTypeConfig(typ auth.Type, exemplar auth.Config) {
	if reflect.TypeOf(exemplar).Kind() == reflect.Ptr {
		// Pointer:
		registeredConfigs[typ] = func() auth.Config {
			return reflect.New(reflect.ValueOf(exemplar).Elem().Type()).Interface().(auth.Config)
		}
		return
	}

	// Not a Pointer
	registeredConfigs[typ] = func() auth.Config {
		return reflect.New(reflect.TypeOf(exemplar)).Elem().Interface().(auth.Config)
	}
}

// SourceSettable configurations can have their authSource set on them
type SourceSettable interface {
	SetAuthSource(*ent.Source)
}

// CreateSource inserts a AuthSource in the DB if not already
// existing with the given name.
//func CreateSource(ctx context.Context, source *ent.Source) error {
//	sourceLimit := &Source{source}
//
//	has, err := models.Client().Source.Query().Where(sourcewhere.NameEQ(source.Name)).Exist(ctx)
//	if err != nil {
//		return err
//	} else if has {
//		return ErrSourceAlreadyExist{source.Name}
//	}
//	// Synchronization is only available with LDAP for now
//	if !sourceLimit.IsLDAP() {
//		sourceLimit.IsSyncEnabled = false
//	}
//
//	has, err = models.Client().Source.
//		CreateBulk()
//
//	_, err = db.GetEngine(ctx).Insert(source)
//	if err != nil {
//		return err
//	}
//
//	if !source.IsActive {
//		return nil
//	}
//
//	if settable, ok := source.Cfg.(SourceSettable); ok {
//		settable.SetAuthSource(source)
//	}
//
//	registerableSource, ok := source.Cfg.(RegisterableSource)
//	if !ok {
//		return nil
//	}
//
//	err = registerableSource.RegisterSource()
//	if err != nil {
//		// remove the AuthSource in case of errors while registering configuration
//		if _, err := db.GetEngine(ctx).ID(source.ID).Delete(new(Source)); err != nil {
//			log.Error("CreateSource: Error while wrapOpenIDConnectInitializeError: %v", err)
//		}
//	}
//	return err
//}
//
//type FindSourcesOptions struct {
//	db.ListOptions
//	IsActive  util.OptionalBool
//	LoginType Type
//}
//
//func (opts FindSourcesOptions) ToConds() builder.Cond {
//	conds := builder.NewCond()
//	if !opts.IsActive.IsNone() {
//		conds = conds.And(builder.Eq{"is_active": opts.IsActive.IsTrue()})
//	}
//	if opts.LoginType != NoType {
//		conds = conds.And(builder.Eq{"`type`": opts.LoginType})
//	}
//	return conds
//}
//
//// IsSSPIEnabled returns true if there is at least one activated login
//// source of type LoginSSPI
//func IsSSPIEnabled(ctx context.Context) bool {
//	exist, err := db.Exist[Source](ctx, FindSourcesOptions{
//		IsActive:  util.OptionalBoolTrue,
//		LoginType: SSPI,
//	}.ToConds())
//	if err != nil {
//		log.Error("IsSSPIEnabled: failed to query active SSPI sources: %v", err)
//		return false
//	}
//	return exist
//}

// GetSourceByID returns login source by given ID.
func GetSourceByID(ctx context.Context, id int64) (*ent.Source, error) {
	source := new(ent.Source)
	if id == 0 {
		source.Cfg = registeredConfigs[auth.NoType]()
		// Set this source to active
		// FIXME: allow disabling of db based password authentication in future
		source.IsActive = true
		return source, nil
	}

	has, err := db.GetEngine(ctx).ID(id).Get(source)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrSourceNotExist{id}
	}
	return source, nil
}

//// UpdateSource updates a Source record in DB.
//func UpdateSource(ctx context.Context, source *Source) error {
//	var originalSource *Source
//	if source.IsOAuth2() {
//		// keep track of the original values so we can restore in case of errors while registering OAuth2 providers
//		var err error
//		if originalSource, err = GetSourceByID(ctx, source.ID); err != nil {
//			return err
//		}
//	}
//
//	has, err := db.GetEngine(ctx).Where("name=? AND id!=?", source.Name, source.ID).Exist(new(Source))
//	if err != nil {
//		return err
//	} else if has {
//		return ErrSourceAlreadyExist{source.Name}
//	}
//
//	_, err = db.GetEngine(ctx).ID(source.ID).AllCols().Update(source)
//	if err != nil {
//		return err
//	}
//
//	if !source.IsActive {
//		return nil
//	}
//
//	if settable, ok := source.Cfg.(SourceSettable); ok {
//		settable.SetAuthSource(source)
//	}
//
//	registerableSource, ok := source.Cfg.(RegisterableSource)
//	if !ok {
//		return nil
//	}
//
//	err = registerableSource.RegisterSource()
//	if err != nil {
//		// restore original values since we cannot update the provider it self
//		if _, err := db.GetEngine(ctx).ID(source.ID).AllCols().Update(originalSource); err != nil {
//			log.Error("UpdateSource: Error while wrapOpenIDConnectInitializeError: %v", err)
//		}
//	}
//	return err
//}
//
//// ErrSourceNotExist represents a "SourceNotExist" kind of error.
//type ErrSourceNotExist struct {
//	ID int64
//}
//
//// IsErrSourceNotExist checks if an error is a ErrSourceNotExist.
//func IsErrSourceNotExist(err error) bool {
//	_, ok := err.(ErrSourceNotExist)
//	return ok
//}
//
//func (err ErrSourceNotExist) Error() string {
//	return fmt.Sprintf("login source does not exist [id: %d]", err.ID)
//}
//
//// Unwrap unwraps this as a ErrNotExist err
//func (err ErrSourceNotExist) Unwrap() error {
//	return util.ErrNotExist
//}
//
//// ErrSourceAlreadyExist represents a "SourceAlreadyExist" kind of error.
//type ErrSourceAlreadyExist struct {
//	Name string
//}
//
//// IsErrSourceAlreadyExist checks if an error is a ErrSourceAlreadyExist.
//func IsErrSourceAlreadyExist(err error) bool {
//	_, ok := err.(ErrSourceAlreadyExist)
//	return ok
//}
//
//func (err ErrSourceAlreadyExist) Error() string {
//	return fmt.Sprintf("login source already exists [name: %s]", err.Name)
//}
//
//// Unwrap unwraps this as a ErrExist err
//func (err ErrSourceAlreadyExist) Unwrap() error {
//	return util.ErrAlreadyExist
//}
//
//// ErrSourceInUse represents a "SourceInUse" kind of error.
//type ErrSourceInUse struct {
//	ID int64
//}
//
//// IsErrSourceInUse checks if an error is a ErrSourceInUse.
//func IsErrSourceInUse(err error) bool {
//	_, ok := err.(ErrSourceInUse)
//	return ok
//}
//
//func (err ErrSourceInUse) Error() string {
//	return fmt.Sprintf("login source is still used by some users [id: %d]", err.ID)
//}
//
//type Source struct {
//	*ent.Source
//}
//
//// BeforeSet is invoked from XORM before setting the value of a field of this object.
//func (source *Source) BeforeSet(colName string, val xorm.Cell) {
//	if colName == "type" {
//		typ := Type(db.Cell2Int64(val))
//		constructor, ok := registeredConfigs[typ]
//		if !ok {
//			return
//		}
//		source.Cfg = constructor()
//		if settable, ok := source.Cfg.(SourceSettable); ok {
//			settable.SetAuthSource(source)
//		}
//	}
//}
//
//// TypeName return name of this login source type.
//func (source *Source) TypeName() string {
//	return Names[source.Type]
//}
//
//// IsLDAP returns true of this source is of the LDAP type.
//func (source *Source) IsLDAP() bool {
//	return source.Source.Type == LDAP
//}
//
//// IsDLDAP returns true of this source is of the DLDAP type.
//func (source *Source) IsDLDAP() bool {
//	return source.Type == DLDAP
//}
//
//// IsSMTP returns true of this source is of the SMTP type.
//func (source *Source) IsSMTP() bool {
//	return source.Type == SMTP
//}
//
//// IsPAM returns true of this source is of the PAM type.
//func (source *Source) IsPAM() bool {
//	return source.Type == PAM
//}
//
//// IsOAuth2 returns true of this source is of the OAuth2 type.
//func (source *Source) IsOAuth2() bool {
//	return source.Type == OAuth2
//}
//
//// IsSSPI returns true of this source is of the SSPI type.
//func (source *Source) IsSSPI() bool {
//	return source.Type == SSPI
//}
//
//// HasTLS returns true of this source supports TLS.
//func (source *Source) HasTLS() bool {
//	hasTLSer, ok := source.Cfg.(HasTLSer)
//	return ok && hasTLSer.HasTLS()
//}
//
//// UseTLS returns true of this source is configured to use TLS.
//func (source *Source) UseTLS() bool {
//	useTLSer, ok := source.Cfg.(UseTLSer)
//	return ok && useTLSer.UseTLS()
//}
//
//// SkipVerify returns true if this source is configured to skip SSL
//// verification.
//func (source *Source) SkipVerify() bool {
//	skipVerifiable, ok := source.Cfg.(SkipVerifiable)
//	return ok && skipVerifiable.IsSkipVerify()
//}
