package qm

import (
	"fmt"
	"strings"
)

type Operand string

const (
	OpEquals        Operand = "="
	OpIN                    = "IN"
	OpNotIN                 = "NOT IN"
	OpGreater               = ">"
	OpGreaterEquals         = ">="
	OpLess                  = "<"
	OpLessEquals            = "<="
	OpNotEquals             = "!="

	OpIsNull = "IS NULL"

	OpIsNotNull = "IS NOT NULL"
)

func Sqlize(col string, op Operand, placeholders ...string) string {
	if len(placeholders) != 0 {
		return fmt.Sprintf("%s %s %s", col, op, strings.Join(placeholders, ""))
	}
	return fmt.Sprintf("%s %s ?", col, op)
}

func SqlizeValueWithoutAlias(col string, op Operand, value interface{}) (string, interface{}) {
	colValues := strings.Split(col, ".")
	v := colValues[len(colValues)-1]
	return fmt.Sprintf("%s %s ?", v, op), value
}

func SqlizeValue(col string, op Operand, value interface{}) (string, interface{}) {
	return fmt.Sprintf("%s %s ?", col, op), value
}

const (
	ASC  string = "ASC"
	DESC        = "DESC"
)

func Relations(rels ...string) string {
	return strings.Join(rels,".")
}

type OrderBy string

func (o OrderBy) DESC() string {
	return fmt.Sprintf("%s %s", string(o), DESC)
}

func (o OrderBy) ASC() string {
	return fmt.Sprintf("%s %s", string(o), ASC)
}

func (o OrderBy) String() string {
	return string(o)
}

type BoolField string

type NullBoolFieldField string

func (f BoolField) ToValue(v bool) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullBoolFieldField) ToNullValue(v *bool) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f BoolField) Equals(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullBoolFieldField) Equals(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f BoolField) GreaterThan(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullBoolFieldField) GreaterThan(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f BoolField) GreaterEqual(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullBoolFieldField) GreaterEqual(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f BoolField) In(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullBoolFieldField) In(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f BoolField) IsNotNull(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullBoolFieldField) IsNotNull(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullBoolFieldField) IsNull(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f BoolField) LessThan(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullBoolFieldField) LessThan(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f BoolField) LessOrEqual(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullBoolFieldField) LessOrEqual(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f BoolField) NotEquals(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullBoolFieldField) NotEquals(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f BoolField) NotIn(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullBoolFieldField) NotIn(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// StringField is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type StringField string

type NullStringFieldField string

func (f StringField) ToValue(v string) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullStringFieldField) ToNullValue(v *string) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f StringField) Equals(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullStringFieldField) Equals(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f StringField) GreaterThan(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullStringFieldField) GreaterThan(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f StringField) GreaterEqual(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullStringFieldField) GreaterEqual(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f StringField) In(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullStringFieldField) In(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f StringField) IsNotNull(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullStringFieldField) IsNotNull(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullStringFieldField) IsNull(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f StringField) LessThan(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullStringFieldField) LessThan(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f StringField) LessOrEqual(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullStringFieldField) LessOrEqual(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f StringField) NotEquals(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullStringFieldField) NotEquals(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f StringField) NotIn(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullStringFieldField) NotIn(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// IntField is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type IntField string

type NullIntFieldField string

func (f IntField) ToValue(v int) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullIntFieldField) ToNullValue(v *int) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f IntField) Equals(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullIntFieldField) Equals(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f IntField) GreaterThan(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullIntFieldField) GreaterThan(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f IntField) GreaterEqual(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullIntFieldField) GreaterEqual(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f IntField) In(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullIntFieldField) In(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f IntField) IsNotNull(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullIntFieldField) IsNotNull(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullIntFieldField) IsNull(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f IntField) LessThan(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullIntFieldField) LessThan(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f IntField) LessOrEqual(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullIntFieldField) LessOrEqual(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f IntField) NotEquals(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullIntFieldField) NotEquals(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f IntField) NotIn(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullIntFieldField) NotIn(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Int8Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Int8Field string

type NullInt8FieldField string

func (f Int8Field) ToValue(v int8) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullInt8FieldField) ToNullValue(v *int8) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Int8Field) Equals(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullInt8FieldField) Equals(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Int8Field) GreaterThan(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullInt8FieldField) GreaterThan(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Int8Field) GreaterEqual(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullInt8FieldField) GreaterEqual(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Int8Field) In(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullInt8FieldField) In(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Int8Field) IsNotNull(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt8FieldField) IsNotNull(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt8FieldField) IsNull(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Int8Field) LessThan(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullInt8FieldField) LessThan(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Int8Field) LessOrEqual(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullInt8FieldField) LessOrEqual(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Int8Field) NotEquals(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullInt8FieldField) NotEquals(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Int8Field) NotIn(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullInt8FieldField) NotIn(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Int16Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Int16Field string

type NullInt16FieldField string

func (f Int16Field) ToValue(v int16) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullInt16FieldField) ToNullValue(v *int16) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Int16Field) Equals(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullInt16FieldField) Equals(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Int16Field) GreaterThan(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullInt16FieldField) GreaterThan(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Int16Field) GreaterEqual(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullInt16FieldField) GreaterEqual(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Int16Field) In(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullInt16FieldField) In(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Int16Field) IsNotNull(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt16FieldField) IsNotNull(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt16FieldField) IsNull(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Int16Field) LessThan(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullInt16FieldField) LessThan(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Int16Field) LessOrEqual(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullInt16FieldField) LessOrEqual(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Int16Field) NotEquals(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullInt16FieldField) NotEquals(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Int16Field) NotIn(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullInt16FieldField) NotIn(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Int32Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Int32Field string

type NullInt32FieldField string

func (f Int32Field) ToValue(v int32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullInt32FieldField) ToNullValue(v *int32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Int32Field) Equals(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullInt32FieldField) Equals(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Int32Field) GreaterThan(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullInt32FieldField) GreaterThan(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Int32Field) GreaterEqual(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullInt32FieldField) GreaterEqual(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Int32Field) In(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullInt32FieldField) In(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Int32Field) IsNotNull(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt32FieldField) IsNotNull(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt32FieldField) IsNull(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Int32Field) LessThan(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullInt32FieldField) LessThan(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Int32Field) LessOrEqual(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullInt32FieldField) LessOrEqual(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Int32Field) NotEquals(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullInt32FieldField) NotEquals(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Int32Field) NotIn(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullInt32FieldField) NotIn(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Int64Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Int64Field string

type NullInt64FieldField string

func (f Int64Field) ToValue(v int64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullInt64FieldField) ToNullValue(v *int64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Int64Field) Equals(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullInt64FieldField) Equals(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Int64Field) GreaterThan(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullInt64FieldField) GreaterThan(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Int64Field) GreaterEqual(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullInt64FieldField) GreaterEqual(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Int64Field) In(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullInt64FieldField) In(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Int64Field) IsNotNull(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt64FieldField) IsNotNull(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt64FieldField) IsNull(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Int64Field) LessThan(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullInt64FieldField) LessThan(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Int64Field) LessOrEqual(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullInt64FieldField) LessOrEqual(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Int64Field) NotEquals(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullInt64FieldField) NotEquals(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Int64Field) NotIn(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullInt64FieldField) NotIn(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// UintField is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type UintField string

type NullUintFieldField string

func (f UintField) ToValue(v uint) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullUintFieldField) ToNullValue(v *uint) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f UintField) Equals(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullUintFieldField) Equals(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f UintField) GreaterThan(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullUintFieldField) GreaterThan(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f UintField) GreaterEqual(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullUintFieldField) GreaterEqual(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f UintField) In(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullUintFieldField) In(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f UintField) IsNotNull(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUintFieldField) IsNotNull(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUintFieldField) IsNull(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f UintField) LessThan(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullUintFieldField) LessThan(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f UintField) LessOrEqual(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullUintFieldField) LessOrEqual(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f UintField) NotEquals(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullUintFieldField) NotEquals(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f UintField) NotIn(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullUintFieldField) NotIn(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Uint8Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Uint8Field string

type NullUint8FieldField string

func (f Uint8Field) ToValue(v uint8) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullUint8FieldField) ToNullValue(v *uint8) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Uint8Field) Equals(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullUint8FieldField) Equals(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Uint8Field) GreaterThan(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullUint8FieldField) GreaterThan(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Uint8Field) GreaterEqual(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullUint8FieldField) GreaterEqual(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Uint8Field) In(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullUint8FieldField) In(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Uint8Field) IsNotNull(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint8FieldField) IsNotNull(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint8FieldField) IsNull(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Uint8Field) LessThan(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullUint8FieldField) LessThan(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Uint8Field) LessOrEqual(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullUint8FieldField) LessOrEqual(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Uint8Field) NotEquals(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullUint8FieldField) NotEquals(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Uint8Field) NotIn(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullUint8FieldField) NotIn(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Uint16Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Uint16Field string

type NullUint16FieldField string

func (f Uint16Field) ToValue(v uint16) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullUint16FieldField) ToNullValue(v *uint16) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Uint16Field) Equals(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullUint16FieldField) Equals(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Uint16Field) GreaterThan(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullUint16FieldField) GreaterThan(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Uint16Field) GreaterEqual(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullUint16FieldField) GreaterEqual(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Uint16Field) In(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullUint16FieldField) In(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Uint16Field) IsNotNull(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint16FieldField) IsNotNull(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint16FieldField) IsNull(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Uint16Field) LessThan(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullUint16FieldField) LessThan(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Uint16Field) LessOrEqual(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullUint16FieldField) LessOrEqual(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Uint16Field) NotEquals(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullUint16FieldField) NotEquals(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Uint16Field) NotIn(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullUint16FieldField) NotIn(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Uint32Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Uint32Field string

type NullUint32FieldField string

func (f Uint32Field) ToValue(v uint32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullUint32FieldField) ToNullValue(v *uint32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Uint32Field) Equals(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullUint32FieldField) Equals(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Uint32Field) GreaterThan(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullUint32FieldField) GreaterThan(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Uint32Field) GreaterEqual(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullUint32FieldField) GreaterEqual(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Uint32Field) In(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullUint32FieldField) In(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Uint32Field) IsNotNull(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint32FieldField) IsNotNull(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint32FieldField) IsNull(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Uint32Field) LessThan(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullUint32FieldField) LessThan(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Uint32Field) LessOrEqual(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullUint32FieldField) LessOrEqual(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Uint32Field) NotEquals(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullUint32FieldField) NotEquals(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Uint32Field) NotIn(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullUint32FieldField) NotIn(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Uint64Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Uint64Field string

type NullUint64FieldField string

func (f Uint64Field) ToValue(v uint64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullUint64FieldField) ToNullValue(v *uint64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Uint64Field) Equals(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullUint64FieldField) Equals(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Uint64Field) GreaterThan(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullUint64FieldField) GreaterThan(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Uint64Field) GreaterEqual(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullUint64FieldField) GreaterEqual(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Uint64Field) In(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullUint64FieldField) In(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Uint64Field) IsNotNull(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint64FieldField) IsNotNull(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint64FieldField) IsNull(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Uint64Field) LessThan(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullUint64FieldField) LessThan(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Uint64Field) LessOrEqual(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullUint64FieldField) LessOrEqual(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Uint64Field) NotEquals(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullUint64FieldField) NotEquals(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Uint64Field) NotIn(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullUint64FieldField) NotIn(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// ByteField is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type ByteField string

type NullByteFieldField string

func (f ByteField) ToValue(v byte) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullByteFieldField) ToNullValue(v *byte) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f ByteField) Equals(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullByteFieldField) Equals(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f ByteField) GreaterThan(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullByteFieldField) GreaterThan(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f ByteField) GreaterEqual(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullByteFieldField) GreaterEqual(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f ByteField) In(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullByteFieldField) In(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f ByteField) IsNotNull(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullByteFieldField) IsNotNull(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullByteFieldField) IsNull(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f ByteField) LessThan(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullByteFieldField) LessThan(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f ByteField) LessOrEqual(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullByteFieldField) LessOrEqual(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f ByteField) NotEquals(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullByteFieldField) NotEquals(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f ByteField) NotIn(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullByteFieldField) NotIn(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// RuneField is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type RuneField string

type NullRuneFieldField string

func (f RuneField) ToValue(v rune) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullRuneFieldField) ToNullValue(v *rune) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f RuneField) Equals(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullRuneFieldField) Equals(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f RuneField) GreaterThan(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullRuneFieldField) GreaterThan(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f RuneField) GreaterEqual(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullRuneFieldField) GreaterEqual(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f RuneField) In(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullRuneFieldField) In(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f RuneField) IsNotNull(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullRuneFieldField) IsNotNull(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullRuneFieldField) IsNull(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f RuneField) LessThan(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullRuneFieldField) LessThan(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f RuneField) LessOrEqual(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullRuneFieldField) LessOrEqual(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f RuneField) NotEquals(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullRuneFieldField) NotEquals(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f RuneField) NotIn(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullRuneFieldField) NotIn(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Float32Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Float32Field string

type NullFloat32FieldField string

func (f Float32Field) ToValue(v float32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullFloat32FieldField) ToNullValue(v *float32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Float32Field) Equals(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullFloat32FieldField) Equals(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Float32Field) GreaterThan(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullFloat32FieldField) GreaterThan(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Float32Field) GreaterEqual(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullFloat32FieldField) GreaterEqual(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Float32Field) In(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullFloat32FieldField) In(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Float32Field) IsNotNull(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullFloat32FieldField) IsNotNull(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullFloat32FieldField) IsNull(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Float32Field) LessThan(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullFloat32FieldField) LessThan(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Float32Field) LessOrEqual(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullFloat32FieldField) LessOrEqual(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Float32Field) NotEquals(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullFloat32FieldField) NotEquals(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Float32Field) NotIn(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullFloat32FieldField) NotIn(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Float64Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Float64Field string

type NullFloat64FieldField string

func (f Float64Field) ToValue(v float64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullFloat64FieldField) ToNullValue(v *float64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Float64Field) Equals(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullFloat64FieldField) Equals(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Float64Field) GreaterThan(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullFloat64FieldField) GreaterThan(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Float64Field) GreaterEqual(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullFloat64FieldField) GreaterEqual(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Float64Field) In(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullFloat64FieldField) In(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Float64Field) IsNotNull(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullFloat64FieldField) IsNotNull(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullFloat64FieldField) IsNull(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Float64Field) LessThan(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullFloat64FieldField) LessThan(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Float64Field) LessOrEqual(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullFloat64FieldField) LessOrEqual(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Float64Field) NotEquals(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullFloat64FieldField) NotEquals(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Float64Field) NotIn(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullFloat64FieldField) NotIn(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Complex64Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Complex64Field string

type NullComplex64FieldField string

func (f Complex64Field) ToValue(v complex64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullComplex64FieldField) ToNullValue(v *complex64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Complex64Field) Equals(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullComplex64FieldField) Equals(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Complex64Field) GreaterThan(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullComplex64FieldField) GreaterThan(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Complex64Field) GreaterEqual(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullComplex64FieldField) GreaterEqual(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Complex64Field) In(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullComplex64FieldField) In(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Complex64Field) IsNotNull(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullComplex64FieldField) IsNotNull(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullComplex64FieldField) IsNull(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Complex64Field) LessThan(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullComplex64FieldField) LessThan(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Complex64Field) LessOrEqual(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullComplex64FieldField) LessOrEqual(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Complex64Field) NotEquals(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullComplex64FieldField) NotEquals(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Complex64Field) NotIn(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullComplex64FieldField) NotIn(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Complex128Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Complex128Field string

type NullComplex128FieldField string

func (f Complex128Field) ToValue(v complex128) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullComplex128FieldField) ToNullValue(v *complex128) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Complex128Field) Equals(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullComplex128FieldField) Equals(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Complex128Field) GreaterThan(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullComplex128FieldField) GreaterThan(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Complex128Field) GreaterEqual(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullComplex128FieldField) GreaterEqual(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Complex128Field) In(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullComplex128FieldField) In(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Complex128Field) IsNotNull(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullComplex128FieldField) IsNotNull(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullComplex128FieldField) IsNull(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Complex128Field) LessThan(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullComplex128FieldField) LessThan(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Complex128Field) LessOrEqual(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullComplex128FieldField) LessOrEqual(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Complex128Field) NotEquals(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullComplex128FieldField) NotEquals(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Complex128Field) NotIn(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullComplex128FieldField) NotIn(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
