// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/menggggggg/go-web-template/internal/app/model"
)

func newOjStandardIo(db *gorm.DB, opts ...gen.DOOption) ojStandardIo {
	_ojStandardIo := ojStandardIo{}

	_ojStandardIo.ojStandardIoDo.UseDB(db, opts...)
	_ojStandardIo.ojStandardIoDo.UseModel(&model.OjStandardIo{})

	tableName := _ojStandardIo.ojStandardIoDo.TableName()
	_ojStandardIo.ALL = field.NewAsterisk(tableName)
	_ojStandardIo.ID = field.NewInt64(tableName, "id")
	_ojStandardIo.Input = field.NewString(tableName, "input")
	_ojStandardIo.Output = field.NewString(tableName, "output")
	_ojStandardIo.ProblemID = field.NewInt64(tableName, "problem_id")
	_ojStandardIo.CreateTime = field.NewTime(tableName, "create_time")
	_ojStandardIo.UpdateTime = field.NewTime(tableName, "update_time")

	_ojStandardIo.fillFieldMap()

	return _ojStandardIo
}

type ojStandardIo struct {
	ojStandardIoDo

	ALL        field.Asterisk
	ID         field.Int64
	Input      field.String // 输入
	Output     field.String // 输出
	ProblemID  field.Int64  // 所属题目
	CreateTime field.Time
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (o ojStandardIo) Table(newTableName string) *ojStandardIo {
	o.ojStandardIoDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o ojStandardIo) As(alias string) *ojStandardIo {
	o.ojStandardIoDo.DO = *(o.ojStandardIoDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *ojStandardIo) updateTableName(table string) *ojStandardIo {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.Input = field.NewString(table, "input")
	o.Output = field.NewString(table, "output")
	o.ProblemID = field.NewInt64(table, "problem_id")
	o.CreateTime = field.NewTime(table, "create_time")
	o.UpdateTime = field.NewTime(table, "update_time")

	o.fillFieldMap()

	return o
}

func (o *ojStandardIo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *ojStandardIo) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["id"] = o.ID
	o.fieldMap["input"] = o.Input
	o.fieldMap["output"] = o.Output
	o.fieldMap["problem_id"] = o.ProblemID
	o.fieldMap["create_time"] = o.CreateTime
	o.fieldMap["update_time"] = o.UpdateTime
}

func (o ojStandardIo) clone(db *gorm.DB) ojStandardIo {
	o.ojStandardIoDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o ojStandardIo) replaceDB(db *gorm.DB) ojStandardIo {
	o.ojStandardIoDo.ReplaceDB(db)
	return o
}

type ojStandardIoDo struct{ gen.DO }

type IOjStandardIoDo interface {
	gen.SubQuery
	Debug() IOjStandardIoDo
	WithContext(ctx context.Context) IOjStandardIoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOjStandardIoDo
	WriteDB() IOjStandardIoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOjStandardIoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOjStandardIoDo
	Not(conds ...gen.Condition) IOjStandardIoDo
	Or(conds ...gen.Condition) IOjStandardIoDo
	Select(conds ...field.Expr) IOjStandardIoDo
	Where(conds ...gen.Condition) IOjStandardIoDo
	Order(conds ...field.Expr) IOjStandardIoDo
	Distinct(cols ...field.Expr) IOjStandardIoDo
	Omit(cols ...field.Expr) IOjStandardIoDo
	Join(table schema.Tabler, on ...field.Expr) IOjStandardIoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOjStandardIoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOjStandardIoDo
	Group(cols ...field.Expr) IOjStandardIoDo
	Having(conds ...gen.Condition) IOjStandardIoDo
	Limit(limit int) IOjStandardIoDo
	Offset(offset int) IOjStandardIoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOjStandardIoDo
	Unscoped() IOjStandardIoDo
	Create(values ...*model.OjStandardIo) error
	CreateInBatches(values []*model.OjStandardIo, batchSize int) error
	Save(values ...*model.OjStandardIo) error
	First() (*model.OjStandardIo, error)
	Take() (*model.OjStandardIo, error)
	Last() (*model.OjStandardIo, error)
	Find() ([]*model.OjStandardIo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OjStandardIo, err error)
	FindInBatches(result *[]*model.OjStandardIo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OjStandardIo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOjStandardIoDo
	Assign(attrs ...field.AssignExpr) IOjStandardIoDo
	Joins(fields ...field.RelationField) IOjStandardIoDo
	Preload(fields ...field.RelationField) IOjStandardIoDo
	FirstOrInit() (*model.OjStandardIo, error)
	FirstOrCreate() (*model.OjStandardIo, error)
	FindByPage(offset int, limit int) (result []*model.OjStandardIo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOjStandardIoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o ojStandardIoDo) Debug() IOjStandardIoDo {
	return o.withDO(o.DO.Debug())
}

func (o ojStandardIoDo) WithContext(ctx context.Context) IOjStandardIoDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o ojStandardIoDo) ReadDB() IOjStandardIoDo {
	return o.Clauses(dbresolver.Read)
}

func (o ojStandardIoDo) WriteDB() IOjStandardIoDo {
	return o.Clauses(dbresolver.Write)
}

func (o ojStandardIoDo) Session(config *gorm.Session) IOjStandardIoDo {
	return o.withDO(o.DO.Session(config))
}

func (o ojStandardIoDo) Clauses(conds ...clause.Expression) IOjStandardIoDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o ojStandardIoDo) Returning(value interface{}, columns ...string) IOjStandardIoDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o ojStandardIoDo) Not(conds ...gen.Condition) IOjStandardIoDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o ojStandardIoDo) Or(conds ...gen.Condition) IOjStandardIoDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o ojStandardIoDo) Select(conds ...field.Expr) IOjStandardIoDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o ojStandardIoDo) Where(conds ...gen.Condition) IOjStandardIoDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o ojStandardIoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IOjStandardIoDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o ojStandardIoDo) Order(conds ...field.Expr) IOjStandardIoDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o ojStandardIoDo) Distinct(cols ...field.Expr) IOjStandardIoDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o ojStandardIoDo) Omit(cols ...field.Expr) IOjStandardIoDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o ojStandardIoDo) Join(table schema.Tabler, on ...field.Expr) IOjStandardIoDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o ojStandardIoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOjStandardIoDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o ojStandardIoDo) RightJoin(table schema.Tabler, on ...field.Expr) IOjStandardIoDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o ojStandardIoDo) Group(cols ...field.Expr) IOjStandardIoDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o ojStandardIoDo) Having(conds ...gen.Condition) IOjStandardIoDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o ojStandardIoDo) Limit(limit int) IOjStandardIoDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o ojStandardIoDo) Offset(offset int) IOjStandardIoDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o ojStandardIoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOjStandardIoDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o ojStandardIoDo) Unscoped() IOjStandardIoDo {
	return o.withDO(o.DO.Unscoped())
}

func (o ojStandardIoDo) Create(values ...*model.OjStandardIo) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o ojStandardIoDo) CreateInBatches(values []*model.OjStandardIo, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o ojStandardIoDo) Save(values ...*model.OjStandardIo) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o ojStandardIoDo) First() (*model.OjStandardIo, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OjStandardIo), nil
	}
}

func (o ojStandardIoDo) Take() (*model.OjStandardIo, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OjStandardIo), nil
	}
}

func (o ojStandardIoDo) Last() (*model.OjStandardIo, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OjStandardIo), nil
	}
}

func (o ojStandardIoDo) Find() ([]*model.OjStandardIo, error) {
	result, err := o.DO.Find()
	return result.([]*model.OjStandardIo), err
}

func (o ojStandardIoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OjStandardIo, err error) {
	buf := make([]*model.OjStandardIo, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o ojStandardIoDo) FindInBatches(result *[]*model.OjStandardIo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o ojStandardIoDo) Attrs(attrs ...field.AssignExpr) IOjStandardIoDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o ojStandardIoDo) Assign(attrs ...field.AssignExpr) IOjStandardIoDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o ojStandardIoDo) Joins(fields ...field.RelationField) IOjStandardIoDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o ojStandardIoDo) Preload(fields ...field.RelationField) IOjStandardIoDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o ojStandardIoDo) FirstOrInit() (*model.OjStandardIo, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OjStandardIo), nil
	}
}

func (o ojStandardIoDo) FirstOrCreate() (*model.OjStandardIo, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OjStandardIo), nil
	}
}

func (o ojStandardIoDo) FindByPage(offset int, limit int) (result []*model.OjStandardIo, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o ojStandardIoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o ojStandardIoDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o ojStandardIoDo) Delete(models ...*model.OjStandardIo) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *ojStandardIoDo) withDO(do gen.Dao) *ojStandardIoDo {
	o.DO = *do.(*gen.DO)
	return o
}
