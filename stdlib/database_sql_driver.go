package stdlib

// Generated by 'goexports database/sql/driver'. Do not edit!

import (
	"database/sql/driver"
	"reflect"
)

func init() {
	Value["database/sql/driver"] = map[string]reflect.Value{
		"Bool":                      reflect.ValueOf(driver.Bool),
		"DefaultParameterConverter": reflect.ValueOf(driver.DefaultParameterConverter),
		"ErrBadConn":                reflect.ValueOf(driver.ErrBadConn),
		"ErrRemoveArgument":         reflect.ValueOf(driver.ErrRemoveArgument),
		"ErrSkip":                   reflect.ValueOf(driver.ErrSkip),
		"Int32":                     reflect.ValueOf(driver.Int32),
		"IsScanValue":               reflect.ValueOf(driver.IsScanValue),
		"IsValue":                   reflect.ValueOf(driver.IsValue),
		"ResultNoRows":              reflect.ValueOf(driver.ResultNoRows),
		"String":                    reflect.ValueOf(driver.String),
	}
	Type["database/sql/driver"] = map[string]reflect.Type{
		"ColumnConverter":                reflect.TypeOf((*driver.ColumnConverter)(nil)).Elem(),
		"Conn":                           reflect.TypeOf((*driver.Conn)(nil)).Elem(),
		"ConnBeginTx":                    reflect.TypeOf((*driver.ConnBeginTx)(nil)).Elem(),
		"ConnPrepareContext":             reflect.TypeOf((*driver.ConnPrepareContext)(nil)).Elem(),
		"Connector":                      reflect.TypeOf((*driver.Connector)(nil)).Elem(),
		"Driver":                         reflect.TypeOf((*driver.Driver)(nil)).Elem(),
		"DriverContext":                  reflect.TypeOf((*driver.DriverContext)(nil)).Elem(),
		"Execer":                         reflect.TypeOf((*driver.Execer)(nil)).Elem(),
		"ExecerContext":                  reflect.TypeOf((*driver.ExecerContext)(nil)).Elem(),
		"IsolationLevel":                 reflect.TypeOf((*driver.IsolationLevel)(nil)).Elem(),
		"NamedValue":                     reflect.TypeOf((*driver.NamedValue)(nil)).Elem(),
		"NamedValueChecker":              reflect.TypeOf((*driver.NamedValueChecker)(nil)).Elem(),
		"NotNull":                        reflect.TypeOf((*driver.NotNull)(nil)).Elem(),
		"Null":                           reflect.TypeOf((*driver.Null)(nil)).Elem(),
		"Pinger":                         reflect.TypeOf((*driver.Pinger)(nil)).Elem(),
		"Queryer":                        reflect.TypeOf((*driver.Queryer)(nil)).Elem(),
		"QueryerContext":                 reflect.TypeOf((*driver.QueryerContext)(nil)).Elem(),
		"Result":                         reflect.TypeOf((*driver.Result)(nil)).Elem(),
		"Rows":                           reflect.TypeOf((*driver.Rows)(nil)).Elem(),
		"RowsAffected":                   reflect.TypeOf((*driver.RowsAffected)(nil)).Elem(),
		"RowsColumnTypeDatabaseTypeName": reflect.TypeOf((*driver.RowsColumnTypeDatabaseTypeName)(nil)).Elem(),
		"RowsColumnTypeLength":           reflect.TypeOf((*driver.RowsColumnTypeLength)(nil)).Elem(),
		"RowsColumnTypeNullable":         reflect.TypeOf((*driver.RowsColumnTypeNullable)(nil)).Elem(),
		"RowsColumnTypePrecisionScale":   reflect.TypeOf((*driver.RowsColumnTypePrecisionScale)(nil)).Elem(),
		"RowsColumnTypeScanType":         reflect.TypeOf((*driver.RowsColumnTypeScanType)(nil)).Elem(),
		"RowsNextResultSet":              reflect.TypeOf((*driver.RowsNextResultSet)(nil)).Elem(),
		"SessionResetter":                reflect.TypeOf((*driver.SessionResetter)(nil)).Elem(),
		"Stmt":                           reflect.TypeOf((*driver.Stmt)(nil)).Elem(),
		"StmtExecContext":                reflect.TypeOf((*driver.StmtExecContext)(nil)).Elem(),
		"StmtQueryContext":               reflect.TypeOf((*driver.StmtQueryContext)(nil)).Elem(),
		"Tx":                             reflect.TypeOf((*driver.Tx)(nil)).Elem(),
		"TxOptions":                      reflect.TypeOf((*driver.TxOptions)(nil)).Elem(),
		"Value":                          reflect.TypeOf((*driver.Value)(nil)).Elem(),
		"ValueConverter":                 reflect.TypeOf((*driver.ValueConverter)(nil)).Elem(),
		"Valuer":                         reflect.TypeOf((*driver.Valuer)(nil)).Elem(),
	}
}
