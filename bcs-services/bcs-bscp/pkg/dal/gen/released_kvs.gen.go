// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/dal/table"
)

func newReleasedKv(db *gorm.DB, opts ...gen.DOOption) releasedKv {
	_releasedKv := releasedKv{}

	_releasedKv.releasedKvDo.UseDB(db, opts...)
	_releasedKv.releasedKvDo.UseModel(&table.ReleasedKv{})

	tableName := _releasedKv.releasedKvDo.TableName()
	_releasedKv.ALL = field.NewAsterisk(tableName)
	_releasedKv.ID = field.NewUint32(tableName, "id")
	_releasedKv.ReleaseID = field.NewUint32(tableName, "release_id")
	_releasedKv.Key = field.NewString(tableName, "key")
	_releasedKv.Memo = field.NewString(tableName, "memo")
	_releasedKv.KvType = field.NewString(tableName, "kv_type")
	_releasedKv.Version = field.NewUint32(tableName, "version")
	_releasedKv.SecretType = field.NewString(tableName, "secret_type")
	_releasedKv.SecretHidden = field.NewBool(tableName, "secret_hidden")
	_releasedKv.BizID = field.NewUint32(tableName, "biz_id")
	_releasedKv.AppID = field.NewUint32(tableName, "app_id")
	_releasedKv.Creator = field.NewString(tableName, "creator")
	_releasedKv.Reviser = field.NewString(tableName, "reviser")
	_releasedKv.CreatedAt = field.NewTime(tableName, "created_at")
	_releasedKv.UpdatedAt = field.NewTime(tableName, "updated_at")
	_releasedKv.Signature = field.NewString(tableName, "signature")
	_releasedKv.ByteSize = field.NewUint64(tableName, "byte_size")
	_releasedKv.Md5 = field.NewString(tableName, "md5")

	_releasedKv.fillFieldMap()

	return _releasedKv
}

type releasedKv struct {
	releasedKvDo releasedKvDo

	ALL          field.Asterisk
	ID           field.Uint32
	ReleaseID    field.Uint32
	Key          field.String
	Memo         field.String
	KvType       field.String
	Version      field.Uint32
	SecretType   field.String
	SecretHidden field.Bool
	BizID        field.Uint32
	AppID        field.Uint32
	Creator      field.String
	Reviser      field.String
	CreatedAt    field.Time
	UpdatedAt    field.Time
	Signature    field.String
	ByteSize     field.Uint64
	Md5          field.String

	fieldMap map[string]field.Expr
}

func (r releasedKv) Table(newTableName string) *releasedKv {
	r.releasedKvDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r releasedKv) As(alias string) *releasedKv {
	r.releasedKvDo.DO = *(r.releasedKvDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *releasedKv) updateTableName(table string) *releasedKv {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewUint32(table, "id")
	r.ReleaseID = field.NewUint32(table, "release_id")
	r.Key = field.NewString(table, "key")
	r.Memo = field.NewString(table, "memo")
	r.KvType = field.NewString(table, "kv_type")
	r.Version = field.NewUint32(table, "version")
	r.SecretType = field.NewString(table, "secret_type")
	r.SecretHidden = field.NewBool(table, "secret_hidden")
	r.BizID = field.NewUint32(table, "biz_id")
	r.AppID = field.NewUint32(table, "app_id")
	r.Creator = field.NewString(table, "creator")
	r.Reviser = field.NewString(table, "reviser")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.Signature = field.NewString(table, "signature")
	r.ByteSize = field.NewUint64(table, "byte_size")
	r.Md5 = field.NewString(table, "md5")

	r.fillFieldMap()

	return r
}

func (r *releasedKv) WithContext(ctx context.Context) IReleasedKvDo {
	return r.releasedKvDo.WithContext(ctx)
}

func (r releasedKv) TableName() string { return r.releasedKvDo.TableName() }

func (r releasedKv) Alias() string { return r.releasedKvDo.Alias() }

func (r releasedKv) Columns(cols ...field.Expr) gen.Columns { return r.releasedKvDo.Columns(cols...) }

func (r *releasedKv) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *releasedKv) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 17)
	r.fieldMap["id"] = r.ID
	r.fieldMap["release_id"] = r.ReleaseID
	r.fieldMap["key"] = r.Key
	r.fieldMap["memo"] = r.Memo
	r.fieldMap["kv_type"] = r.KvType
	r.fieldMap["version"] = r.Version
	r.fieldMap["secret_type"] = r.SecretType
	r.fieldMap["secret_hidden"] = r.SecretHidden
	r.fieldMap["biz_id"] = r.BizID
	r.fieldMap["app_id"] = r.AppID
	r.fieldMap["creator"] = r.Creator
	r.fieldMap["reviser"] = r.Reviser
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["signature"] = r.Signature
	r.fieldMap["byte_size"] = r.ByteSize
	r.fieldMap["md5"] = r.Md5
}

func (r releasedKv) clone(db *gorm.DB) releasedKv {
	r.releasedKvDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r releasedKv) replaceDB(db *gorm.DB) releasedKv {
	r.releasedKvDo.ReplaceDB(db)
	return r
}

type releasedKvDo struct{ gen.DO }

type IReleasedKvDo interface {
	gen.SubQuery
	Debug() IReleasedKvDo
	WithContext(ctx context.Context) IReleasedKvDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReleasedKvDo
	WriteDB() IReleasedKvDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReleasedKvDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReleasedKvDo
	Not(conds ...gen.Condition) IReleasedKvDo
	Or(conds ...gen.Condition) IReleasedKvDo
	Select(conds ...field.Expr) IReleasedKvDo
	Where(conds ...gen.Condition) IReleasedKvDo
	Order(conds ...field.Expr) IReleasedKvDo
	Distinct(cols ...field.Expr) IReleasedKvDo
	Omit(cols ...field.Expr) IReleasedKvDo
	Join(table schema.Tabler, on ...field.Expr) IReleasedKvDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReleasedKvDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReleasedKvDo
	Group(cols ...field.Expr) IReleasedKvDo
	Having(conds ...gen.Condition) IReleasedKvDo
	Limit(limit int) IReleasedKvDo
	Offset(offset int) IReleasedKvDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReleasedKvDo
	Unscoped() IReleasedKvDo
	Create(values ...*table.ReleasedKv) error
	CreateInBatches(values []*table.ReleasedKv, batchSize int) error
	Save(values ...*table.ReleasedKv) error
	First() (*table.ReleasedKv, error)
	Take() (*table.ReleasedKv, error)
	Last() (*table.ReleasedKv, error)
	Find() ([]*table.ReleasedKv, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ReleasedKv, err error)
	FindInBatches(result *[]*table.ReleasedKv, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.ReleasedKv) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReleasedKvDo
	Assign(attrs ...field.AssignExpr) IReleasedKvDo
	Joins(fields ...field.RelationField) IReleasedKvDo
	Preload(fields ...field.RelationField) IReleasedKvDo
	FirstOrInit() (*table.ReleasedKv, error)
	FirstOrCreate() (*table.ReleasedKv, error)
	FindByPage(offset int, limit int) (result []*table.ReleasedKv, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReleasedKvDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r releasedKvDo) Debug() IReleasedKvDo {
	return r.withDO(r.DO.Debug())
}

func (r releasedKvDo) WithContext(ctx context.Context) IReleasedKvDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r releasedKvDo) ReadDB() IReleasedKvDo {
	return r.Clauses(dbresolver.Read)
}

func (r releasedKvDo) WriteDB() IReleasedKvDo {
	return r.Clauses(dbresolver.Write)
}

func (r releasedKvDo) Session(config *gorm.Session) IReleasedKvDo {
	return r.withDO(r.DO.Session(config))
}

func (r releasedKvDo) Clauses(conds ...clause.Expression) IReleasedKvDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r releasedKvDo) Returning(value interface{}, columns ...string) IReleasedKvDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r releasedKvDo) Not(conds ...gen.Condition) IReleasedKvDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r releasedKvDo) Or(conds ...gen.Condition) IReleasedKvDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r releasedKvDo) Select(conds ...field.Expr) IReleasedKvDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r releasedKvDo) Where(conds ...gen.Condition) IReleasedKvDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r releasedKvDo) Order(conds ...field.Expr) IReleasedKvDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r releasedKvDo) Distinct(cols ...field.Expr) IReleasedKvDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r releasedKvDo) Omit(cols ...field.Expr) IReleasedKvDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r releasedKvDo) Join(table schema.Tabler, on ...field.Expr) IReleasedKvDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r releasedKvDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReleasedKvDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r releasedKvDo) RightJoin(table schema.Tabler, on ...field.Expr) IReleasedKvDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r releasedKvDo) Group(cols ...field.Expr) IReleasedKvDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r releasedKvDo) Having(conds ...gen.Condition) IReleasedKvDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r releasedKvDo) Limit(limit int) IReleasedKvDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r releasedKvDo) Offset(offset int) IReleasedKvDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r releasedKvDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReleasedKvDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r releasedKvDo) Unscoped() IReleasedKvDo {
	return r.withDO(r.DO.Unscoped())
}

func (r releasedKvDo) Create(values ...*table.ReleasedKv) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r releasedKvDo) CreateInBatches(values []*table.ReleasedKv, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r releasedKvDo) Save(values ...*table.ReleasedKv) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r releasedKvDo) First() (*table.ReleasedKv, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedKv), nil
	}
}

func (r releasedKvDo) Take() (*table.ReleasedKv, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedKv), nil
	}
}

func (r releasedKvDo) Last() (*table.ReleasedKv, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedKv), nil
	}
}

func (r releasedKvDo) Find() ([]*table.ReleasedKv, error) {
	result, err := r.DO.Find()
	return result.([]*table.ReleasedKv), err
}

func (r releasedKvDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ReleasedKv, err error) {
	buf := make([]*table.ReleasedKv, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r releasedKvDo) FindInBatches(result *[]*table.ReleasedKv, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r releasedKvDo) Attrs(attrs ...field.AssignExpr) IReleasedKvDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r releasedKvDo) Assign(attrs ...field.AssignExpr) IReleasedKvDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r releasedKvDo) Joins(fields ...field.RelationField) IReleasedKvDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r releasedKvDo) Preload(fields ...field.RelationField) IReleasedKvDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r releasedKvDo) FirstOrInit() (*table.ReleasedKv, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedKv), nil
	}
}

func (r releasedKvDo) FirstOrCreate() (*table.ReleasedKv, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.ReleasedKv), nil
	}
}

func (r releasedKvDo) FindByPage(offset int, limit int) (result []*table.ReleasedKv, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r releasedKvDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r releasedKvDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r releasedKvDo) Delete(models ...*table.ReleasedKv) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *releasedKvDo) withDO(do gen.Dao) *releasedKvDo {
	r.DO = *do.(*gen.DO)
	return r
}
