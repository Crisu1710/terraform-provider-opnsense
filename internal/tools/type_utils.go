package tools

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"log"
	"strconv"
)

// Ints

func Int64ToString(i int64) string {
	return fmt.Sprintf("%d", i)
}

func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return i
	}
	return -1
}

func StringToInt64Null(s string) types.Int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return types.Int64Value(i)
	}
	return types.Int64Null()
}

func Int64ToStringNegative(i int64) string {
	s := fmt.Sprintf("%d", i)
	if i == -1 {
		s = ""
	}
	return s
}

// Floats

func Float64ToString(i float64) string {
	return fmt.Sprintf("%f", i)
}

func Float64ToStringNegative(i float64) string {
	s := fmt.Sprintf("%f", i)
	if i == -1 {
		s = ""
	}
	return s
}

func StringToFloat64(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return i
	}
	return -1
}

// Bools

func BoolToString(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

func StringToBool(s string) bool {
	return s == "1"
}

// Strings

func StringOrNull(s string) types.String {
	if s != "" {
		return types.StringValue(s)
	} else {
		return types.StringNull()
	}
}

// Sets

func EmptySetValue() types.Set {
	sv, _ := types.SetValue(types.StringType, []attr.Value{})
	return sv
}

func AppendSetValue(s []string) types.Set {
	set := make([]attr.Value, len(s))
	for k, v := range s {
		set[k] = types.StringValue(v)
	}
	sv, _ := types.SetValue(types.StringType, set)
	return sv
}

func SetToString(set types.Set) []string {
	var list []string
	set.ElementsAs(context.Background(), &list, false)
	log.Println(list)
	return list
}

func StringToSet(s []string) types.Set {
	var list []attr.Value
	for _, i := range s {
		list = append(list, basetypes.NewStringValue(i))
	}
	typeList, _ := types.SetValue(types.StringType, list)
	log.Println(typeList)
	return typeList
}
