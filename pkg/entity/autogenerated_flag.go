package entity

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set FlagQuerySet

// FlagQuerySet is an queryset type for Flag
type FlagQuerySet struct {
	db *gorm.DB
}

// NewFlagQuerySet constructs new FlagQuerySet
func NewFlagQuerySet(db *gorm.DB) FlagQuerySet {
	return FlagQuerySet{
		db: db.Model(&Flag{}),
	}
}

func (qs FlagQuerySet) w(db *gorm.DB) FlagQuerySet {
	return NewFlagQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) All(ret *[]Flag) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Flag) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtEq(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtGt(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtGte(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtLt(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtLte(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtNe(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// CreatedByEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedByEq(createdBy string) FlagQuerySet {
	return qs.w(qs.db.Where("created_by = ?", createdBy))
}

// CreatedByIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedByIn(createdBy string, createdByRest ...string) FlagQuerySet {
	iArgs := []interface{}{createdBy}
	for _, arg := range createdByRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("created_by IN (?)", iArgs))
}

// CreatedByNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedByNe(createdBy string) FlagQuerySet {
	return qs.w(qs.db.Where("created_by != ?", createdBy))
}

// CreatedByNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedByNotIn(createdBy string, createdByRest ...string) FlagQuerySet {
	iArgs := []interface{}{createdBy}
	for _, arg := range createdByRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("created_by NOT IN (?)", iArgs))
}

// DataRecordsEnabledEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DataRecordsEnabledEq(dataRecordsEnabled bool) FlagQuerySet {
	return qs.w(qs.db.Where("data_records_enabled = ?", dataRecordsEnabled))
}

// DataRecordsEnabledIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DataRecordsEnabledIn(dataRecordsEnabled bool, dataRecordsEnabledRest ...bool) FlagQuerySet {
	iArgs := []interface{}{dataRecordsEnabled}
	for _, arg := range dataRecordsEnabledRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("data_records_enabled IN (?)", iArgs))
}

// DataRecordsEnabledNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DataRecordsEnabledNe(dataRecordsEnabled bool) FlagQuerySet {
	return qs.w(qs.db.Where("data_records_enabled != ?", dataRecordsEnabled))
}

// DataRecordsEnabledNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DataRecordsEnabledNotIn(dataRecordsEnabled bool, dataRecordsEnabledRest ...bool) FlagQuerySet {
	iArgs := []interface{}{dataRecordsEnabled}
	for _, arg := range dataRecordsEnabledRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("data_records_enabled NOT IN (?)", iArgs))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Flag) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) Delete() error {
	return qs.db.Delete(Flag{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtEq(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtGt(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtGte(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtIsNotNull() FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtIsNull() FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtLt(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtLte(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtNe(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DescriptionEq(description string) FlagQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DescriptionIn(description string, descriptionRest ...string) FlagQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description IN (?)", iArgs))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DescriptionNe(description string) FlagQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DescriptionNotIn(description string, descriptionRest ...string) FlagQuerySet {
	iArgs := []interface{}{description}
	for _, arg := range descriptionRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", iArgs))
}

// EnabledEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EnabledEq(enabled bool) FlagQuerySet {
	return qs.w(qs.db.Where("enabled = ?", enabled))
}

// EnabledIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EnabledIn(enabled bool, enabledRest ...bool) FlagQuerySet {
	iArgs := []interface{}{enabled}
	for _, arg := range enabledRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("enabled IN (?)", iArgs))
}

// EnabledNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EnabledNe(enabled bool) FlagQuerySet {
	return qs.w(qs.db.Where("enabled != ?", enabled))
}

// EnabledNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EnabledNotIn(enabled bool, enabledRest ...bool) FlagQuerySet {
	iArgs := []interface{}{enabled}
	for _, arg := range enabledRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("enabled NOT IN (?)", iArgs))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) GetUpdater() FlagUpdater {
	return NewFlagUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDEq(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDGt(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDGte(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDIn(ID uint, IDRest ...uint) FlagQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDLt(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDLte(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDNe(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDNotIn(ID uint, IDRest ...uint) FlagQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) Limit(limit int) FlagQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs FlagQuerySet) One(ret *Flag) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderAscByCreatedAt() FlagQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderAscByDeletedAt() FlagQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderAscByID() FlagQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderAscByUpdatedAt() FlagQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderDescByCreatedAt() FlagQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderDescByDeletedAt() FlagQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderDescByID() FlagQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderDescByUpdatedAt() FlagQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// PreloadFlagEvaluation is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) PreloadFlagEvaluation() FlagQuerySet {
	return qs.w(qs.db.Preload("FlagEvaluation"))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetCreatedAt(createdAt time.Time) FlagUpdater {
	u.fields[string(FlagDBSchema.CreatedAt)] = createdAt
	return u
}

// SetCreatedBy is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetCreatedBy(createdBy string) FlagUpdater {
	u.fields[string(FlagDBSchema.CreatedBy)] = createdBy
	return u
}

// SetDataRecordsEnabled is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetDataRecordsEnabled(dataRecordsEnabled bool) FlagUpdater {
	u.fields[string(FlagDBSchema.DataRecordsEnabled)] = dataRecordsEnabled
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetDescription(description string) FlagUpdater {
	u.fields[string(FlagDBSchema.Description)] = description
	return u
}

// SetEnabled is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetEnabled(enabled bool) FlagUpdater {
	u.fields[string(FlagDBSchema.Enabled)] = enabled
	return u
}

// SetFlagEvaluation is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetFlagEvaluation(flagEvaluation FlagEvaluation) FlagUpdater {
	u.fields[string(FlagDBSchema.FlagEvaluation)] = flagEvaluation
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetID(ID uint) FlagUpdater {
	u.fields[string(FlagDBSchema.ID)] = ID
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetUpdatedAt(updatedAt time.Time) FlagUpdater {
	u.fields[string(FlagDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUpdatedBy is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetUpdatedBy(updatedBy string) FlagUpdater {
	u.fields[string(FlagDBSchema.UpdatedBy)] = updatedBy
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u FlagUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u FlagUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtEq(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtGt(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtGte(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtLt(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtLte(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtNe(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// UpdatedByEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedByEq(updatedBy string) FlagQuerySet {
	return qs.w(qs.db.Where("updated_by = ?", updatedBy))
}

// UpdatedByIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedByIn(updatedBy string, updatedByRest ...string) FlagQuerySet {
	iArgs := []interface{}{updatedBy}
	for _, arg := range updatedByRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("updated_by IN (?)", iArgs))
}

// UpdatedByNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedByNe(updatedBy string) FlagQuerySet {
	return qs.w(qs.db.Where("updated_by != ?", updatedBy))
}

// UpdatedByNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedByNotIn(updatedBy string, updatedByRest ...string) FlagQuerySet {
	iArgs := []interface{}{updatedBy}
	for _, arg := range updatedByRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("updated_by NOT IN (?)", iArgs))
}

// ===== END of query set FlagQuerySet

// ===== BEGIN of Flag modifiers

type flagDBSchemaField string

// FlagDBSchema stores db field names of Flag
var FlagDBSchema = struct {
	ID                 flagDBSchemaField
	CreatedAt          flagDBSchemaField
	UpdatedAt          flagDBSchemaField
	DeletedAt          flagDBSchemaField
	Description        flagDBSchemaField
	CreatedBy          flagDBSchemaField
	UpdatedBy          flagDBSchemaField
	Enabled            flagDBSchemaField
	Segments           flagDBSchemaField
	Variants           flagDBSchemaField
	DataRecordsEnabled flagDBSchemaField
	FlagEvaluation     flagDBSchemaField
}{

	ID:                 flagDBSchemaField("id"),
	CreatedAt:          flagDBSchemaField("created_at"),
	UpdatedAt:          flagDBSchemaField("updated_at"),
	DeletedAt:          flagDBSchemaField("deleted_at"),
	Description:        flagDBSchemaField("description"),
	CreatedBy:          flagDBSchemaField("created_by"),
	UpdatedBy:          flagDBSchemaField("updated_by"),
	Enabled:            flagDBSchemaField("enabled"),
	Segments:           flagDBSchemaField("segments"),
	Variants:           flagDBSchemaField("variants"),
	DataRecordsEnabled: flagDBSchemaField("data_records_enabled"),
	FlagEvaluation:     flagDBSchemaField("flag_evaluation"),
}

// Update updates Flag fields by primary key
func (o *Flag) Update(db *gorm.DB, fields ...flagDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":                   o.ID,
		"created_at":           o.CreatedAt,
		"updated_at":           o.UpdatedAt,
		"deleted_at":           o.DeletedAt,
		"description":          o.Description,
		"created_by":           o.CreatedBy,
		"updated_by":           o.UpdatedBy,
		"enabled":              o.Enabled,
		"segments":             o.Segments,
		"variants":             o.Variants,
		"data_records_enabled": o.DataRecordsEnabled,
		"flag_evaluation":      o.FlagEvaluation,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := string(f)
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Flag %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// FlagUpdater is an Flag updates manager
type FlagUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewFlagUpdater creates new Flag updater
func NewFlagUpdater(db *gorm.DB) FlagUpdater {
	return FlagUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Flag{}),
	}
}

// ===== END of Flag modifiers

// ===== END of all query sets
