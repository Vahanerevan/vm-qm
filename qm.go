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

type NullBoolField string

func (f BoolField) ToValue(v bool) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullBoolField) ToNullValue(v *bool) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f BoolField) Equals(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullBoolField) Equals(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f BoolField) GreaterThan(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullBoolField) GreaterThan(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f BoolField) GreaterEqual(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullBoolField) GreaterEqual(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f BoolField) In(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullBoolField) In(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f BoolField) IsNotNull(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullBoolField) IsNotNull(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullBoolField) IsNull(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f BoolField) LessThan(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullBoolField) LessThan(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f BoolField) LessOrEqual(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullBoolField) LessOrEqual(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f BoolField) NotEquals(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullBoolField) NotEquals(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f BoolField) NotIn(v bool) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullBoolField) NotIn(v *bool) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// StringField is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type StringField string

type NullStringField string

func (f StringField) ToValue(v string) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullStringField) ToNullValue(v *string) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f StringField) Equals(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullStringField) Equals(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f StringField) GreaterThan(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullStringField) GreaterThan(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f StringField) GreaterEqual(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullStringField) GreaterEqual(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f StringField) In(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullStringField) In(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f StringField) IsNotNull(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullStringField) IsNotNull(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullStringField) IsNull(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f StringField) LessThan(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullStringField) LessThan(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f StringField) LessOrEqual(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullStringField) LessOrEqual(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f StringField) NotEquals(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullStringField) NotEquals(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f StringField) NotIn(v string) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullStringField) NotIn(v *string) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// IntField is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type IntField string

type NullIntField string

func (f IntField) ToValue(v int) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullIntField) ToNullValue(v *int) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f IntField) Equals(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullIntField) Equals(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f IntField) GreaterThan(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullIntField) GreaterThan(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f IntField) GreaterEqual(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullIntField) GreaterEqual(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f IntField) In(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullIntField) In(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f IntField) IsNotNull(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullIntField) IsNotNull(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullIntField) IsNull(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f IntField) LessThan(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullIntField) LessThan(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f IntField) LessOrEqual(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullIntField) LessOrEqual(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f IntField) NotEquals(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullIntField) NotEquals(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f IntField) NotIn(v int) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullIntField) NotIn(v *int) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Int8Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Int8Field string

type NullInt8Field string

func (f Int8Field) ToValue(v int8) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullInt8Field) ToNullValue(v *int8) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Int8Field) Equals(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullInt8Field) Equals(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Int8Field) GreaterThan(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullInt8Field) GreaterThan(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Int8Field) GreaterEqual(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullInt8Field) GreaterEqual(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Int8Field) In(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullInt8Field) In(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Int8Field) IsNotNull(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt8Field) IsNotNull(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt8Field) IsNull(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Int8Field) LessThan(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullInt8Field) LessThan(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Int8Field) LessOrEqual(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullInt8Field) LessOrEqual(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Int8Field) NotEquals(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullInt8Field) NotEquals(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Int8Field) NotIn(v int8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullInt8Field) NotIn(v *int8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Int16Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Int16Field string

type NullInt16Field string

func (f Int16Field) ToValue(v int16) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullInt16Field) ToNullValue(v *int16) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Int16Field) Equals(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullInt16Field) Equals(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Int16Field) GreaterThan(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullInt16Field) GreaterThan(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Int16Field) GreaterEqual(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullInt16Field) GreaterEqual(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Int16Field) In(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullInt16Field) In(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Int16Field) IsNotNull(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt16Field) IsNotNull(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt16Field) IsNull(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Int16Field) LessThan(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullInt16Field) LessThan(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Int16Field) LessOrEqual(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullInt16Field) LessOrEqual(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Int16Field) NotEquals(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullInt16Field) NotEquals(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Int16Field) NotIn(v int16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullInt16Field) NotIn(v *int16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Int32Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Int32Field string

type NullInt32Field string

func (f Int32Field) ToValue(v int32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullInt32Field) ToNullValue(v *int32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Int32Field) Equals(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullInt32Field) Equals(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Int32Field) GreaterThan(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullInt32Field) GreaterThan(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Int32Field) GreaterEqual(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullInt32Field) GreaterEqual(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Int32Field) In(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullInt32Field) In(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Int32Field) IsNotNull(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt32Field) IsNotNull(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt32Field) IsNull(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Int32Field) LessThan(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullInt32Field) LessThan(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Int32Field) LessOrEqual(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullInt32Field) LessOrEqual(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Int32Field) NotEquals(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullInt32Field) NotEquals(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Int32Field) NotIn(v int32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullInt32Field) NotIn(v *int32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Int64Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Int64Field string

type NullInt64Field string

func (f Int64Field) ToValue(v int64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullInt64Field) ToNullValue(v *int64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Int64Field) Equals(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullInt64Field) Equals(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Int64Field) GreaterThan(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullInt64Field) GreaterThan(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Int64Field) GreaterEqual(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullInt64Field) GreaterEqual(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Int64Field) In(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullInt64Field) In(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Int64Field) IsNotNull(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt64Field) IsNotNull(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullInt64Field) IsNull(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Int64Field) LessThan(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullInt64Field) LessThan(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Int64Field) LessOrEqual(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullInt64Field) LessOrEqual(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Int64Field) NotEquals(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullInt64Field) NotEquals(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Int64Field) NotIn(v int64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullInt64Field) NotIn(v *int64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// UintField is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type UintField string

type NullUintField string

func (f UintField) ToValue(v uint) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullUintField) ToNullValue(v *uint) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f UintField) Equals(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullUintField) Equals(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f UintField) GreaterThan(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullUintField) GreaterThan(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f UintField) GreaterEqual(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullUintField) GreaterEqual(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f UintField) In(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullUintField) In(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f UintField) IsNotNull(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUintField) IsNotNull(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUintField) IsNull(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f UintField) LessThan(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullUintField) LessThan(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f UintField) LessOrEqual(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullUintField) LessOrEqual(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f UintField) NotEquals(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullUintField) NotEquals(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f UintField) NotIn(v uint) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullUintField) NotIn(v *uint) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Uint8Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Uint8Field string

type NullUint8Field string

func (f Uint8Field) ToValue(v uint8) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullUint8Field) ToNullValue(v *uint8) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Uint8Field) Equals(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullUint8Field) Equals(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Uint8Field) GreaterThan(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullUint8Field) GreaterThan(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Uint8Field) GreaterEqual(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullUint8Field) GreaterEqual(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Uint8Field) In(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullUint8Field) In(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Uint8Field) IsNotNull(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint8Field) IsNotNull(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint8Field) IsNull(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Uint8Field) LessThan(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullUint8Field) LessThan(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Uint8Field) LessOrEqual(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullUint8Field) LessOrEqual(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Uint8Field) NotEquals(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullUint8Field) NotEquals(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Uint8Field) NotIn(v uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullUint8Field) NotIn(v *uint8) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Uint16Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Uint16Field string

type NullUint16Field string

func (f Uint16Field) ToValue(v uint16) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullUint16Field) ToNullValue(v *uint16) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Uint16Field) Equals(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullUint16Field) Equals(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Uint16Field) GreaterThan(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullUint16Field) GreaterThan(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Uint16Field) GreaterEqual(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullUint16Field) GreaterEqual(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Uint16Field) In(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullUint16Field) In(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Uint16Field) IsNotNull(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint16Field) IsNotNull(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint16Field) IsNull(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Uint16Field) LessThan(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullUint16Field) LessThan(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Uint16Field) LessOrEqual(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullUint16Field) LessOrEqual(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Uint16Field) NotEquals(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullUint16Field) NotEquals(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Uint16Field) NotIn(v uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullUint16Field) NotIn(v *uint16) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Uint32Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Uint32Field string

type NullUint32Field string

func (f Uint32Field) ToValue(v uint32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullUint32Field) ToNullValue(v *uint32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Uint32Field) Equals(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullUint32Field) Equals(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Uint32Field) GreaterThan(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullUint32Field) GreaterThan(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Uint32Field) GreaterEqual(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullUint32Field) GreaterEqual(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Uint32Field) In(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullUint32Field) In(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Uint32Field) IsNotNull(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint32Field) IsNotNull(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint32Field) IsNull(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Uint32Field) LessThan(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullUint32Field) LessThan(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Uint32Field) LessOrEqual(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullUint32Field) LessOrEqual(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Uint32Field) NotEquals(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullUint32Field) NotEquals(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Uint32Field) NotIn(v uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullUint32Field) NotIn(v *uint32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Uint64Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Uint64Field string

type NullUint64Field string

func (f Uint64Field) ToValue(v uint64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullUint64Field) ToNullValue(v *uint64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Uint64Field) Equals(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullUint64Field) Equals(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Uint64Field) GreaterThan(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullUint64Field) GreaterThan(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Uint64Field) GreaterEqual(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullUint64Field) GreaterEqual(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Uint64Field) In(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullUint64Field) In(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Uint64Field) IsNotNull(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint64Field) IsNotNull(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullUint64Field) IsNull(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Uint64Field) LessThan(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullUint64Field) LessThan(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Uint64Field) LessOrEqual(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullUint64Field) LessOrEqual(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Uint64Field) NotEquals(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullUint64Field) NotEquals(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Uint64Field) NotIn(v uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullUint64Field) NotIn(v *uint64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// ByteField is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type ByteField string

type NullByteField string

func (f ByteField) ToValue(v byte) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullByteField) ToNullValue(v *byte) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f ByteField) Equals(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullByteField) Equals(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f ByteField) GreaterThan(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullByteField) GreaterThan(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f ByteField) GreaterEqual(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullByteField) GreaterEqual(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f ByteField) In(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullByteField) In(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f ByteField) IsNotNull(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullByteField) IsNotNull(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullByteField) IsNull(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f ByteField) LessThan(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullByteField) LessThan(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f ByteField) LessOrEqual(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullByteField) LessOrEqual(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f ByteField) NotEquals(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullByteField) NotEquals(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f ByteField) NotIn(v byte) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullByteField) NotIn(v *byte) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// RuneField is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type RuneField string

type NullRuneField string

func (f RuneField) ToValue(v rune) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullRuneField) ToNullValue(v *rune) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f RuneField) Equals(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullRuneField) Equals(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f RuneField) GreaterThan(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullRuneField) GreaterThan(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f RuneField) GreaterEqual(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullRuneField) GreaterEqual(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f RuneField) In(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullRuneField) In(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f RuneField) IsNotNull(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullRuneField) IsNotNull(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullRuneField) IsNull(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f RuneField) LessThan(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullRuneField) LessThan(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f RuneField) LessOrEqual(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullRuneField) LessOrEqual(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f RuneField) NotEquals(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullRuneField) NotEquals(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f RuneField) NotIn(v rune) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullRuneField) NotIn(v *rune) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Float32Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Float32Field string

type NullFloat32Field string

func (f Float32Field) ToValue(v float32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullFloat32Field) ToNullValue(v *float32) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Float32Field) Equals(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullFloat32Field) Equals(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Float32Field) GreaterThan(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullFloat32Field) GreaterThan(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Float32Field) GreaterEqual(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullFloat32Field) GreaterEqual(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Float32Field) In(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullFloat32Field) In(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Float32Field) IsNotNull(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullFloat32Field) IsNotNull(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullFloat32Field) IsNull(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Float32Field) LessThan(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullFloat32Field) LessThan(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Float32Field) LessOrEqual(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullFloat32Field) LessOrEqual(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Float32Field) NotEquals(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullFloat32Field) NotEquals(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Float32Field) NotIn(v float32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullFloat32Field) NotIn(v *float32) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Float64Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Float64Field string

type NullFloat64Field string

func (f Float64Field) ToValue(v float64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullFloat64Field) ToNullValue(v *float64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Float64Field) Equals(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullFloat64Field) Equals(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Float64Field) GreaterThan(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullFloat64Field) GreaterThan(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Float64Field) GreaterEqual(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullFloat64Field) GreaterEqual(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Float64Field) In(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullFloat64Field) In(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Float64Field) IsNotNull(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullFloat64Field) IsNotNull(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullFloat64Field) IsNull(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Float64Field) LessThan(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullFloat64Field) LessThan(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Float64Field) LessOrEqual(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullFloat64Field) LessOrEqual(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Float64Field) NotEquals(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullFloat64Field) NotEquals(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Float64Field) NotIn(v float64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullFloat64Field) NotIn(v *float64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Complex64Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Complex64Field string

type NullComplex64Field string

func (f Complex64Field) ToValue(v complex64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullComplex64Field) ToNullValue(v *complex64) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Complex64Field) Equals(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullComplex64Field) Equals(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Complex64Field) GreaterThan(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullComplex64Field) GreaterThan(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Complex64Field) GreaterEqual(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullComplex64Field) GreaterEqual(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Complex64Field) In(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullComplex64Field) In(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Complex64Field) IsNotNull(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullComplex64Field) IsNotNull(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullComplex64Field) IsNull(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Complex64Field) LessThan(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullComplex64Field) LessThan(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Complex64Field) LessOrEqual(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullComplex64Field) LessOrEqual(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Complex64Field) NotEquals(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullComplex64Field) NotEquals(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Complex64Field) NotIn(v complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullComplex64Field) NotIn(v *complex64) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}

// Complex128Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.

type Complex128Field string

type NullComplex128Field string

func (f Complex128Field) ToValue(v complex128) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f NullComplex128Field) ToNullValue(v *complex128) (string, interface{}) {
	return SqlizeValueWithoutAlias(string(f), OpEquals, v)
}

func (f Complex128Field) Equals(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}
func (f NullComplex128Field) Equals(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpEquals, v)
}

func (f Complex128Field) GreaterThan(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}
func (f NullComplex128Field) GreaterThan(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpGreater, v)
}

func (f Complex128Field) GreaterEqual(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}
func (f NullComplex128Field) GreaterEqual(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpGreaterEquals, v)
}

func (f Complex128Field) In(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}
func (f NullComplex128Field) In(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpIN, v)
}

func (f Complex128Field) IsNotNull(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullComplex128Field) IsNotNull(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNotNull, v)
}
func (f NullComplex128Field) IsNull(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpIsNull, v)
}

func (f Complex128Field) LessThan(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}
func (f NullComplex128Field) LessThan(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpLess, v)
}

func (f Complex128Field) LessOrEqual(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}
func (f NullComplex128Field) LessOrEqual(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpLessEquals, v)
}

func (f Complex128Field) NotEquals(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}
func (f NullComplex128Field) NotEquals(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpNotEquals, v)
}

func (f Complex128Field) NotIn(v complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
func (f NullComplex128Field) NotIn(v *complex128) (string, interface{}) {
	return SqlizeValue(string(f), OpNotIN, v)
}
