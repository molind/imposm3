package postgis

import (
	"fmt"
)

type ColumnType interface {
	Name() string
	PrepareInsertSql(i int,
		spec *TableSpec) string
	GeneralizeSql(colSpec *ColumnSpec, spec *GeneralizedTableSpec) string
}

type simpleColumnType struct {
	name string
}

func (t *simpleColumnType) Name() string {
	return t.name
}

func (t *simpleColumnType) PrepareInsertSql(i int, spec *TableSpec) string {
	return fmt.Sprintf("$%d", i)
}

func (t *simpleColumnType) GeneralizeSql(colSpec *ColumnSpec, spec *GeneralizedTableSpec) string {
	return colSpec.Name
}

type geometryType struct {
	name string
}

func (t *geometryType) Name() string {
	return t.name
}

func (t *geometryType) PrepareInsertSql(i int, spec *TableSpec) string {
	return fmt.Sprintf("$%d::Geometry",
		i,
	)
}

func (t *geometryType) GeneralizeSql(colSpec *ColumnSpec, spec *GeneralizedTableSpec) string {
	return fmt.Sprintf(`ST_SimplifyPreserveTopology("%s", %f) as "%s"`,
		colSpec.Name, spec.Tolerance, colSpec.Name,
	)
}

type hstoreType struct {
	name string
}

func (t *hstoreType) Name() string {
	return t.name
}

func (t *hstoreType) PrepareInsertSql(i int, spec *TableSpec) string {
	return fmt.Sprintf("$%d::hstore",
		i,
	)
}

func (t *hstoreType) GeneralizeSql(colSpec *ColumnSpec, spec *GeneralizedTableSpec) string {
	return colSpec.Name
}

var pgTypes map[string]ColumnType

func init() {
	pgTypes = map[string]ColumnType{
		"string":   &simpleColumnType{"VARCHAR"},
		"bool":     &simpleColumnType{"BOOL"},
		"int8":     &simpleColumnType{"SMALLINT"},
		"int32":    &simpleColumnType{"INT"},
		"int64":    &simpleColumnType{"BIGINT"},
		"float32":  &simpleColumnType{"REAL"},
		"geometry": &geometryType{"GEOMETRY"},
		"map":      &hstoreType{"HSTORE"},
	}
}
