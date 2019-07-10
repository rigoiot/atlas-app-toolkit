package influxdb

import (
	"fmt"

	"github.com/rigoiot/atlas-app-toolkit/query"
)

// FilterStringToInflux is a shortcut to parse a filter string using default FilteringParser implementation
// and call FilteringToInflux on the returned filtering expression.
func FilterStringToInflux(filter string) (string, error) {
	f, err := query.ParseFiltering(filter)
	if err != nil {
		return "", err
	}
	return FilteringToInflux(f)
}

// FilteringToInflux returns InfluxDB Plain SQL representation of the filtering expression.
func FilteringToInflux(m *query.Filtering) (string, error) {
	if m == nil || m.Root == nil {
		return "", nil
	}
	switch r := m.Root.(type) {
	case *query.Filtering_Operator:
		return LogicalOperatorToInflux(r.Operator)
	case *query.Filtering_StringCondition:
		return StringConditionToInflux(r.StringCondition)
	case *query.Filtering_NumberCondition:
		return NumberConditionToInflux(r.NumberCondition)
	default:
		return "", fmt.Errorf("%T type is not supported in Filtering", r)
	}
}

// LogicalOperatorToInflux returns InfluxDB Plain SQL representation of the logical operator.
func LogicalOperatorToInflux(lop *query.LogicalOperator) (string, error) {
	var lres string
	var err error
	switch l := lop.Left.(type) {
	case *query.LogicalOperator_LeftOperator:
		lres, err = LogicalOperatorToInflux(l.LeftOperator)
	case *query.LogicalOperator_LeftStringCondition:
		lres, err = StringConditionToInflux(l.LeftStringCondition)
	case *query.LogicalOperator_LeftNumberCondition:
		lres, err = NumberConditionToInflux(l.LeftNumberCondition)
	default:
		return "", fmt.Errorf("%T type is not supported in Filtering", l)
	}
	if err != nil {
		return "", err
	}

	var rres string
	switch r := lop.Right.(type) {
	case *query.LogicalOperator_RightOperator:
		rres, err = LogicalOperatorToInflux(r.RightOperator)
	case *query.LogicalOperator_RightStringCondition:
		rres, err = StringConditionToInflux(r.RightStringCondition)
	case *query.LogicalOperator_RightNumberCondition:
		rres, err = NumberConditionToInflux(r.RightNumberCondition)
	default:
		return "", fmt.Errorf("%T type is not supported in Filtering", r)
	}
	if err != nil {
		return "", err
	}

	var o string
	switch lop.Type {
	case query.LogicalOperator_AND:
		o = "AND"
	case query.LogicalOperator_OR:
		o = "OR"
	}

	return fmt.Sprintf("%s %s (%s)", lres, o, rres), nil
}

// StringConditionToInflux returns InfluxDB Plain SQL representation of the string condition.
func StringConditionToInflux(c *query.StringCondition) (string, error) {
	if len(c.FieldPath) < 1 {
		return "", fmt.Errorf("Field path longer than 1 is not supported")
	}

	var o string
	switch c.Type {
	case query.StringCondition_EQ:
		if c.IsNegative {
			o = "!="
		} else {
			o = "="
		}
	case query.StringCondition_GT:
		o = ">"
	case query.StringCondition_GE:
		o = ">="
	case query.StringCondition_LT:
		o = "<"
	case query.StringCondition_LE:
		o = "<="
	}

	return fmt.Sprintf("\"%s\" %s '%s'", c.FieldPath[0], o, c.Value), nil
}

// NumberConditionToInflux returns InfluxDB Plain SQL representation of the number condition.
func NumberConditionToInflux(c *query.NumberCondition) (string, error) {
	if len(c.FieldPath) < 1 {
		return "", fmt.Errorf("Field path longer than 1 is not supported")
	}

	var o string
	switch c.Type {
	case query.NumberCondition_EQ:
		if c.IsNegative {
			o = "!="
		} else {
			o = "="
		}
	case query.NumberCondition_GT:
		o = ">"
	case query.NumberCondition_GE:
		o = ">="
	case query.NumberCondition_LT:
		o = "<"
	case query.NumberCondition_LE:
		o = "<="
	}

	return fmt.Sprintf("\"%s\" %s %f", c.FieldPath[0], o, c.Value), nil
}
