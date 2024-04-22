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

func newClientEvent(db *gorm.DB, opts ...gen.DOOption) clientEvent {
	_clientEvent := clientEvent{}

	_clientEvent.clientEventDo.UseDB(db, opts...)
	_clientEvent.clientEventDo.UseModel(&table.ClientEvent{})

	tableName := _clientEvent.clientEventDo.TableName()
	_clientEvent.ALL = field.NewAsterisk(tableName)
	_clientEvent.ID = field.NewUint32(tableName, "id")
	_clientEvent.ClientID = field.NewUint32(tableName, "client_id")
	_clientEvent.CursorID = field.NewString(tableName, "cursor_id")
	_clientEvent.UID = field.NewString(tableName, "uid")
	_clientEvent.BizID = field.NewUint32(tableName, "biz_id")
	_clientEvent.AppID = field.NewUint32(tableName, "app_id")
	_clientEvent.ClientMode = field.NewString(tableName, "client_mode")
	_clientEvent.OriginalReleaseID = field.NewUint32(tableName, "original_release_id")
	_clientEvent.TargetReleaseID = field.NewUint32(tableName, "target_release_id")
	_clientEvent.StartTime = field.NewTime(tableName, "start_time")
	_clientEvent.EndTime = field.NewTime(tableName, "end_time")
	_clientEvent.TotalSeconds = field.NewFloat64(tableName, "total_seconds")
	_clientEvent.TotalFileSize = field.NewFloat64(tableName, "total_file_size")
	_clientEvent.DownloadFileSize = field.NewFloat64(tableName, "download_file_size")
	_clientEvent.TotalFileNum = field.NewUint32(tableName, "total_file_num")
	_clientEvent.DownloadFileNum = field.NewUint32(tableName, "download_file_num")
	_clientEvent.ReleaseChangeStatus = field.NewString(tableName, "release_change_status")
	_clientEvent.ReleaseChangeFailedReason = field.NewString(tableName, "release_change_failed_reason")
	_clientEvent.SpecificFailedReason = field.NewString(tableName, "specific_failed_reason")
	_clientEvent.FailedDetailReason = field.NewString(tableName, "failed_detail_reason")

	_clientEvent.fillFieldMap()

	return _clientEvent
}

type clientEvent struct {
	clientEventDo clientEventDo

	ALL                       field.Asterisk
	ID                        field.Uint32
	ClientID                  field.Uint32
	CursorID                  field.String
	UID                       field.String
	BizID                     field.Uint32
	AppID                     field.Uint32
	ClientMode                field.String
	OriginalReleaseID         field.Uint32
	TargetReleaseID           field.Uint32
	StartTime                 field.Time
	EndTime                   field.Time
	TotalSeconds              field.Float64
	TotalFileSize             field.Float64
	DownloadFileSize          field.Float64
	TotalFileNum              field.Uint32
	DownloadFileNum           field.Uint32
	ReleaseChangeStatus       field.String
	ReleaseChangeFailedReason field.String
	SpecificFailedReason      field.String
	FailedDetailReason        field.String

	fieldMap map[string]field.Expr
}

func (c clientEvent) Table(newTableName string) *clientEvent {
	c.clientEventDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c clientEvent) As(alias string) *clientEvent {
	c.clientEventDo.DO = *(c.clientEventDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *clientEvent) updateTableName(table string) *clientEvent {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint32(table, "id")
	c.ClientID = field.NewUint32(table, "client_id")
	c.CursorID = field.NewString(table, "cursor_id")
	c.UID = field.NewString(table, "uid")
	c.BizID = field.NewUint32(table, "biz_id")
	c.AppID = field.NewUint32(table, "app_id")
	c.ClientMode = field.NewString(table, "client_mode")
	c.OriginalReleaseID = field.NewUint32(table, "original_release_id")
	c.TargetReleaseID = field.NewUint32(table, "target_release_id")
	c.StartTime = field.NewTime(table, "start_time")
	c.EndTime = field.NewTime(table, "end_time")
	c.TotalSeconds = field.NewFloat64(table, "total_seconds")
	c.TotalFileSize = field.NewFloat64(table, "total_file_size")
	c.DownloadFileSize = field.NewFloat64(table, "download_file_size")
	c.TotalFileNum = field.NewUint32(table, "total_file_num")
	c.DownloadFileNum = field.NewUint32(table, "download_file_num")
	c.ReleaseChangeStatus = field.NewString(table, "release_change_status")
	c.ReleaseChangeFailedReason = field.NewString(table, "release_change_failed_reason")
	c.SpecificFailedReason = field.NewString(table, "specific_failed_reason")
	c.FailedDetailReason = field.NewString(table, "failed_detail_reason")

	c.fillFieldMap()

	return c
}

func (c *clientEvent) WithContext(ctx context.Context) IClientEventDo {
	return c.clientEventDo.WithContext(ctx)
}

func (c clientEvent) TableName() string { return c.clientEventDo.TableName() }

func (c clientEvent) Alias() string { return c.clientEventDo.Alias() }

func (c clientEvent) Columns(cols ...field.Expr) gen.Columns { return c.clientEventDo.Columns(cols...) }

func (c *clientEvent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *clientEvent) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 20)
	c.fieldMap["id"] = c.ID
	c.fieldMap["client_id"] = c.ClientID
	c.fieldMap["cursor_id"] = c.CursorID
	c.fieldMap["uid"] = c.UID
	c.fieldMap["biz_id"] = c.BizID
	c.fieldMap["app_id"] = c.AppID
	c.fieldMap["client_mode"] = c.ClientMode
	c.fieldMap["original_release_id"] = c.OriginalReleaseID
	c.fieldMap["target_release_id"] = c.TargetReleaseID
	c.fieldMap["start_time"] = c.StartTime
	c.fieldMap["end_time"] = c.EndTime
	c.fieldMap["total_seconds"] = c.TotalSeconds
	c.fieldMap["total_file_size"] = c.TotalFileSize
	c.fieldMap["download_file_size"] = c.DownloadFileSize
	c.fieldMap["total_file_num"] = c.TotalFileNum
	c.fieldMap["download_file_num"] = c.DownloadFileNum
	c.fieldMap["release_change_status"] = c.ReleaseChangeStatus
	c.fieldMap["release_change_failed_reason"] = c.ReleaseChangeFailedReason
	c.fieldMap["specific_failed_reason"] = c.SpecificFailedReason
	c.fieldMap["failed_detail_reason"] = c.FailedDetailReason
}

func (c clientEvent) clone(db *gorm.DB) clientEvent {
	c.clientEventDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c clientEvent) replaceDB(db *gorm.DB) clientEvent {
	c.clientEventDo.ReplaceDB(db)
	return c
}

type clientEventDo struct{ gen.DO }

type IClientEventDo interface {
	gen.SubQuery
	Debug() IClientEventDo
	WithContext(ctx context.Context) IClientEventDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IClientEventDo
	WriteDB() IClientEventDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IClientEventDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IClientEventDo
	Not(conds ...gen.Condition) IClientEventDo
	Or(conds ...gen.Condition) IClientEventDo
	Select(conds ...field.Expr) IClientEventDo
	Where(conds ...gen.Condition) IClientEventDo
	Order(conds ...field.Expr) IClientEventDo
	Distinct(cols ...field.Expr) IClientEventDo
	Omit(cols ...field.Expr) IClientEventDo
	Join(table schema.Tabler, on ...field.Expr) IClientEventDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IClientEventDo
	RightJoin(table schema.Tabler, on ...field.Expr) IClientEventDo
	Group(cols ...field.Expr) IClientEventDo
	Having(conds ...gen.Condition) IClientEventDo
	Limit(limit int) IClientEventDo
	Offset(offset int) IClientEventDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IClientEventDo
	Unscoped() IClientEventDo
	Create(values ...*table.ClientEvent) error
	CreateInBatches(values []*table.ClientEvent, batchSize int) error
	Save(values ...*table.ClientEvent) error
	First() (*table.ClientEvent, error)
	Take() (*table.ClientEvent, error)
	Last() (*table.ClientEvent, error)
	Find() ([]*table.ClientEvent, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ClientEvent, err error)
	FindInBatches(result *[]*table.ClientEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.ClientEvent) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IClientEventDo
	Assign(attrs ...field.AssignExpr) IClientEventDo
	Joins(fields ...field.RelationField) IClientEventDo
	Preload(fields ...field.RelationField) IClientEventDo
	FirstOrInit() (*table.ClientEvent, error)
	FirstOrCreate() (*table.ClientEvent, error)
	FindByPage(offset int, limit int) (result []*table.ClientEvent, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IClientEventDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c clientEventDo) Debug() IClientEventDo {
	return c.withDO(c.DO.Debug())
}

func (c clientEventDo) WithContext(ctx context.Context) IClientEventDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c clientEventDo) ReadDB() IClientEventDo {
	return c.Clauses(dbresolver.Read)
}

func (c clientEventDo) WriteDB() IClientEventDo {
	return c.Clauses(dbresolver.Write)
}

func (c clientEventDo) Session(config *gorm.Session) IClientEventDo {
	return c.withDO(c.DO.Session(config))
}

func (c clientEventDo) Clauses(conds ...clause.Expression) IClientEventDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c clientEventDo) Returning(value interface{}, columns ...string) IClientEventDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c clientEventDo) Not(conds ...gen.Condition) IClientEventDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c clientEventDo) Or(conds ...gen.Condition) IClientEventDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c clientEventDo) Select(conds ...field.Expr) IClientEventDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c clientEventDo) Where(conds ...gen.Condition) IClientEventDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c clientEventDo) Order(conds ...field.Expr) IClientEventDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c clientEventDo) Distinct(cols ...field.Expr) IClientEventDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c clientEventDo) Omit(cols ...field.Expr) IClientEventDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c clientEventDo) Join(table schema.Tabler, on ...field.Expr) IClientEventDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c clientEventDo) LeftJoin(table schema.Tabler, on ...field.Expr) IClientEventDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c clientEventDo) RightJoin(table schema.Tabler, on ...field.Expr) IClientEventDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c clientEventDo) Group(cols ...field.Expr) IClientEventDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c clientEventDo) Having(conds ...gen.Condition) IClientEventDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c clientEventDo) Limit(limit int) IClientEventDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c clientEventDo) Offset(offset int) IClientEventDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c clientEventDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IClientEventDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c clientEventDo) Unscoped() IClientEventDo {
	return c.withDO(c.DO.Unscoped())
}

func (c clientEventDo) Create(values ...*table.ClientEvent) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c clientEventDo) CreateInBatches(values []*table.ClientEvent, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c clientEventDo) Save(values ...*table.ClientEvent) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c clientEventDo) First() (*table.ClientEvent, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.ClientEvent), nil
	}
}

func (c clientEventDo) Take() (*table.ClientEvent, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.ClientEvent), nil
	}
}

func (c clientEventDo) Last() (*table.ClientEvent, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.ClientEvent), nil
	}
}

func (c clientEventDo) Find() ([]*table.ClientEvent, error) {
	result, err := c.DO.Find()
	return result.([]*table.ClientEvent), err
}

func (c clientEventDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ClientEvent, err error) {
	buf := make([]*table.ClientEvent, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c clientEventDo) FindInBatches(result *[]*table.ClientEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c clientEventDo) Attrs(attrs ...field.AssignExpr) IClientEventDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c clientEventDo) Assign(attrs ...field.AssignExpr) IClientEventDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c clientEventDo) Joins(fields ...field.RelationField) IClientEventDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c clientEventDo) Preload(fields ...field.RelationField) IClientEventDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c clientEventDo) FirstOrInit() (*table.ClientEvent, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.ClientEvent), nil
	}
}

func (c clientEventDo) FirstOrCreate() (*table.ClientEvent, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.ClientEvent), nil
	}
}

func (c clientEventDo) FindByPage(offset int, limit int) (result []*table.ClientEvent, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c clientEventDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c clientEventDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c clientEventDo) Delete(models ...*table.ClientEvent) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *clientEventDo) withDO(do gen.Dao) *clientEventDo {
	c.DO = *do.(*gen.DO)
	return c
}
