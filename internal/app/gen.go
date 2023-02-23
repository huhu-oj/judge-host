package app
import (
	// "fmt"
	// "gorm.io/driver/mysql"
	"gorm.io/gen"
	// "gorm.io/gorm"

)
// const MysqlConfig = "root:birthdayis1123@(43.139.47.68:3306)/eladmin?charset=utf8mb4&parseTime=True&loc=Local"

func InitGen() {
	db := InitGormDB()
	g := gen.NewGenerator(gen.Config{
		OutPath:          "./internal/app/dao",
		Mode:             gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:    true,
		FieldCoverable:   false,
		FieldSignable:    false,
		FieldWithTypeTag: true,
	})
	g.UseDB(db)
	dataMap := map[string]func(detailType string) (dataType string){
		"tinyint":   func(detailType string) (dataType string) { return "int64" },
		"smallint":  func(detailType string) (dataType string) { return "int64" },
		"mediumint": func(detailType string) (dataType string) { return "int64" },
		"bigint":    func(detailType string) (dataType string) { return "int64" },
		"int":       func(detailType string) (dataType string) { return "int64" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)
	//jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	//	toStringField := `balance, `
	//	if strings.Contains(toStringField, columnName) {
	//		return columnName + ",string"
	//	}
	//	return columnName
	//})
	//autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;autoUpdateTime")
	//autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")
	//softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")
	//fieldOpts := []gen.ModelOpt{jsonField, autoCreateTimeField, autoUpdateTimeField, softDeleteField}
	osi := g.GenerateModel("oj_standard_io")
	//allModel := g.GenerateAllTable(fieldOpts...)
	g.ApplyBasic(osi)
	//g.ApplyBasic(allModel...)

	g.Execute()
}
