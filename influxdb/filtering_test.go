package influxdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfluxDBFiltering(t *testing.T) {

	tests := []struct {
		rest     string
		influxdb string
		err      error
	}{
		{
			"not field1 == 'value1' or field2 == 'value2' and field3 != 'value3'",
			"\"field1\" != 'value1' OR (\"field2\" = 'value2' AND (\"field3\" != 'value3'))",
			nil,
		},
		{
			"field1 == 22",
			"\"field1\" = 22.000000",
			nil,
		},
		// {
		// 	"not field1 == 22",
		// 	"NOT(entities.field1 = ?)",
		// 	[]interface{}{22.0},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"field1 > 22",
		// 	"(entities.field1 > ?)",
		// 	[]interface{}{22.0},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"not field1 > 22",
		// 	"NOT(entities.field1 > ?)",
		// 	[]interface{}{22.0},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"field1 >= 22",
		// 	"(entities.field1 >= ?)",
		// 	[]interface{}{22.0},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"not field1 >= 22",
		// 	"NOT(entities.field1 >= ?)",
		// 	[]interface{}{22.0},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"field1 < 22",
		// 	"(entities.field1 < ?)",
		// 	[]interface{}{22.0},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"not field1 < 22",
		// 	"NOT(entities.field1 < ?)",
		// 	[]interface{}{22.0},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"field1 <= 22",
		// 	"(entities.field1 <= ?)",
		// 	[]interface{}{22.0},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"not field1 <= 22",
		// 	"NOT(entities.field1 <= ?)",
		// 	[]interface{}{22.0},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"field_string > 'str'",
		// 	"(entities.field_string > ?)",
		// 	[]interface{}{"str"},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"field_string >= 'str'",
		// 	"(entities.field_string >= ?)",
		// 	[]interface{}{"str"},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"field_string < 'str'",
		// 	"(entities.field_string < ?)",
		// 	[]interface{}{"str"},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"field_string <= 'str'",
		// 	"(entities.field_string <= ?)",
		// 	[]interface{}{"str"},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"field1 == null",
		// 	"(entities.field1 IS NULL)",
		// 	nil,
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"field1 != null",
		// 	"NOT(entities.field1 IS NULL)",
		// 	nil,
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"field1 != null",
		// 	"NOT(entities.field1 IS NULL)",
		// 	nil,
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"nested_entity.nested_field1 == 11 and nested_entity.nested_field2 == 22",
		// 	"((nested_entity.nested_field1 = ?) AND (nested_entity.nested_field2 = ?))",
		// 	[]interface{}{11.0, 22.0},
		// 	map[string]struct{}{"NestedEntity": {}},
		// 	nil,
		// },
		// {
		// 	"field1 === null",
		// 	"",
		// 	nil,
		// 	nil,
		// 	&query.UnexpectedSymbolError{},
		// },
		// {
		// 	"id == 'id' and ref == 'ref'",
		// 	"((entities.id = ?) AND (entities.ref = ?))",
		// 	[]interface{}{"convertedid", "convertedref"},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"id := 'ID'",
		// 	"(lower(entities.id) = lower(?))",
		// 	[]interface{}{"convertedid"},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"not(id := 'sOmeId')",
		// 	"NOT(lower(entities.id) = lower(?))",
		// 	[]interface{}{"convertedid"},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"id in ['sOmeId', 'egegeg']",
		// 	"(entities.id  IN (?, ?))",
		// 	[]interface{}{"convertedid", "convertedid"},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"not(id in ['sOmeId', 'egegeg'])",
		// 	"(entities.id NOT IN (?, ?))",
		// 	[]interface{}{"convertedid", "convertedid"},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"id in [1, 2]",
		// 	"(entities.id  IN (?, ?))",
		// 	[]interface{}{1.0, 2.0},
		// 	nil,
		// 	nil,
		// },
		// {
		// 	"not(id in [1, 2])",
		// 	"(entities.id NOT IN (?, ?))",
		// 	[]interface{}{1.0, 2.0},
		// 	nil,
		// 	nil,
		// },
	}

	for _, test := range tests {
		influxdb, err := FilterStringToInflux(test.rest)
		assert.Equal(t, test.influxdb, influxdb)
		assert.IsType(t, test.err, err)
	}
}
