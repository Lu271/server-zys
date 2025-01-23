// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/Lu271/server-zys/internal/store/model"
)

func newPlayerAchievement(db *gorm.DB) playerAchievement {
	_playerAchievement := playerAchievement{}

	_playerAchievement.playerAchievementDo.UseDB(db)
	_playerAchievement.playerAchievementDo.UseModel(&model.PlayerAchievement{})

	tableName := _playerAchievement.playerAchievementDo.TableName()
	_playerAchievement.ALL = field.NewAsterisk(tableName)
	_playerAchievement.PlayerID = field.NewInt32(tableName, "playerID")
	_playerAchievement.AchievementID = field.NewInt32(tableName, "achievementID")
	_playerAchievement.AchievedAt = field.NewInt64(tableName, "achieved_at")

	_playerAchievement.fillFieldMap()

	return _playerAchievement
}

type playerAchievement struct {
	playerAchievementDo playerAchievementDo

	ALL           field.Asterisk
	PlayerID      field.Int32
	AchievementID field.Int32
	AchievedAt    field.Int64

	fieldMap map[string]field.Expr
}

func (p playerAchievement) Table(newTableName string) *playerAchievement {
	p.playerAchievementDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p playerAchievement) As(alias string) *playerAchievement {
	p.playerAchievementDo.DO = *(p.playerAchievementDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *playerAchievement) updateTableName(table string) *playerAchievement {
	p.ALL = field.NewAsterisk(table)
	p.PlayerID = field.NewInt32(table, "playerID")
	p.AchievementID = field.NewInt32(table, "achievementID")
	p.AchievedAt = field.NewInt64(table, "achieved_at")

	p.fillFieldMap()

	return p
}

func (p *playerAchievement) WithContext(ctx context.Context) *playerAchievementDo {
	return p.playerAchievementDo.WithContext(ctx)
}

func (p playerAchievement) TableName() string { return p.playerAchievementDo.TableName() }

func (p playerAchievement) Alias() string { return p.playerAchievementDo.Alias() }

func (p *playerAchievement) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *playerAchievement) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["playerID"] = p.PlayerID
	p.fieldMap["achievementID"] = p.AchievementID
	p.fieldMap["achieved_at"] = p.AchievedAt
}

func (p playerAchievement) clone(db *gorm.DB) playerAchievement {
	p.playerAchievementDo.ReplaceDB(db)
	return p
}

type playerAchievementDo struct{ gen.DO }

func (p playerAchievementDo) Debug() *playerAchievementDo {
	return p.withDO(p.DO.Debug())
}

func (p playerAchievementDo) WithContext(ctx context.Context) *playerAchievementDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p playerAchievementDo) ReadDB() *playerAchievementDo {
	return p.Clauses(dbresolver.Read)
}

func (p playerAchievementDo) WriteDB() *playerAchievementDo {
	return p.Clauses(dbresolver.Write)
}

func (p playerAchievementDo) Clauses(conds ...clause.Expression) *playerAchievementDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p playerAchievementDo) Returning(value interface{}, columns ...string) *playerAchievementDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p playerAchievementDo) Not(conds ...gen.Condition) *playerAchievementDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p playerAchievementDo) Or(conds ...gen.Condition) *playerAchievementDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p playerAchievementDo) Select(conds ...field.Expr) *playerAchievementDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p playerAchievementDo) Where(conds ...gen.Condition) *playerAchievementDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p playerAchievementDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *playerAchievementDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p playerAchievementDo) Order(conds ...field.Expr) *playerAchievementDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p playerAchievementDo) Distinct(cols ...field.Expr) *playerAchievementDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p playerAchievementDo) Omit(cols ...field.Expr) *playerAchievementDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p playerAchievementDo) Join(table schema.Tabler, on ...field.Expr) *playerAchievementDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p playerAchievementDo) LeftJoin(table schema.Tabler, on ...field.Expr) *playerAchievementDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p playerAchievementDo) RightJoin(table schema.Tabler, on ...field.Expr) *playerAchievementDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p playerAchievementDo) Group(cols ...field.Expr) *playerAchievementDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p playerAchievementDo) Having(conds ...gen.Condition) *playerAchievementDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p playerAchievementDo) Limit(limit int) *playerAchievementDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p playerAchievementDo) Offset(offset int) *playerAchievementDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p playerAchievementDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *playerAchievementDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p playerAchievementDo) Unscoped() *playerAchievementDo {
	return p.withDO(p.DO.Unscoped())
}

func (p playerAchievementDo) Create(values ...*model.PlayerAchievement) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p playerAchievementDo) CreateInBatches(values []*model.PlayerAchievement, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p playerAchievementDo) Save(values ...*model.PlayerAchievement) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p playerAchievementDo) First() (*model.PlayerAchievement, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PlayerAchievement), nil
	}
}

func (p playerAchievementDo) Take() (*model.PlayerAchievement, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PlayerAchievement), nil
	}
}

func (p playerAchievementDo) Last() (*model.PlayerAchievement, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PlayerAchievement), nil
	}
}

func (p playerAchievementDo) Find() ([]*model.PlayerAchievement, error) {
	result, err := p.DO.Find()
	return result.([]*model.PlayerAchievement), err
}

func (p playerAchievementDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PlayerAchievement, err error) {
	buf := make([]*model.PlayerAchievement, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p playerAchievementDo) FindInBatches(result *[]*model.PlayerAchievement, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p playerAchievementDo) Attrs(attrs ...field.AssignExpr) *playerAchievementDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p playerAchievementDo) Assign(attrs ...field.AssignExpr) *playerAchievementDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p playerAchievementDo) Joins(fields ...field.RelationField) *playerAchievementDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p playerAchievementDo) Preload(fields ...field.RelationField) *playerAchievementDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p playerAchievementDo) FirstOrInit() (*model.PlayerAchievement, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PlayerAchievement), nil
	}
}

func (p playerAchievementDo) FirstOrCreate() (*model.PlayerAchievement, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PlayerAchievement), nil
	}
}

func (p playerAchievementDo) FindByPage(offset int, limit int) (result []*model.PlayerAchievement, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p playerAchievementDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p playerAchievementDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p playerAchievementDo) Delete(models ...*model.PlayerAchievement) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *playerAchievementDo) withDO(do gen.Dao) *playerAchievementDo {
	p.DO = *do.(*gen.DO)
	return p
}
