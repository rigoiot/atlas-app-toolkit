package gorm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	ptypes "github.com/gogo/protobuf/types"
)

type childTest struct {
	FieldOne   int
	FieldTwo   string
	FieldThree *int
	FieldFour  []int
}

type topTest struct {
	FieldA childTest
	FieldB *childTest
}

func wrapInt(x int) *int {
	return &x
}

func TestMergeWithMask(t *testing.T) {
	source := &topTest{
		FieldA: childTest{FieldOne: 22, FieldTwo: "catch", FieldThree: wrapInt(2), FieldFour: []int{1, 2, 3}},
		FieldB: &childTest{FieldOne: 3, FieldTwo: "string", FieldThree: wrapInt(1), FieldFour: []int{3, 2, 1}},
	}
	dest := &topTest{}
	err := MergeWithMask(source, dest, &ptypes.FieldMask{Paths: []string{"FieldB.FieldOne", "FieldA.FieldTwo", "FieldA.FieldThree", "FieldB.FieldFour"}})
	assert.Equal(t, &topTest{
		FieldA: childTest{FieldTwo: "catch", FieldThree: wrapInt(2)},
		FieldB: &childTest{FieldOne: 3, FieldFour: []int{3, 2, 1}},
	}, dest)
	assert.Nil(t, err)

	err = MergeWithMask(source, dest, &ptypes.FieldMask{Paths: []string{"FieldB.FieldDNE", "FieldA.FieldTwo"}})
	assert.Equal(t, errors.New("Field path \"FieldB.FieldDNE\" doesn't exist in type *gorm.topTest"), err)

	err = MergeWithMask(nil, dest, &ptypes.FieldMask{Paths: []string{"FieldB.FieldDNE"}})
	assert.Equal(t, errors.New("Source object is nil"), err)

	for _, fm := range []*ptypes.FieldMask{nil, {}} {
		err = MergeWithMask(nil, nil, fm)
		assert.Nil(t, err)
		err = MergeWithMask(nil, dest, fm)
		assert.Nil(t, err)
		err = MergeWithMask(source, nil, fm)
		assert.Nil(t, err)
		err = MergeWithMask(source, dest.FieldA, fm)
		assert.Nil(t, err)
	}
	err = MergeWithMask(source, nil, &ptypes.FieldMask{Paths: []string{"FieldB"}})
	assert.Equal(t, errors.New("Destination object is nil"), err)
	err = MergeWithMask(source, dest.FieldA, &ptypes.FieldMask{Paths: []string{"FieldB"}})
	assert.Equal(t, errors.New("Types of source and destination objects do not match"), err)
	dest = &topTest{}
	err = MergeWithMask(source, dest, &ptypes.FieldMask{Paths: []string{"FieldA.FieldTwo", "FieldA.FieldFour.Anything"}})
	assert.Equal(t, &topTest{
		FieldA: childTest{FieldTwo: "catch"},
	}, dest)
	assert.Nil(t, err)
}
