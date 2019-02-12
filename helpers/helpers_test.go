package helpers

import "testing"

func TestInSliceString(t *testing.T) {
	needle := "fuck"
	slice := []interface{}{"shit", "fuck1", "fuck"}
	if !InSlice(needle, slice) {
		t.Error(`InSlice("fuck", []interface{}{"shit", "fuck1", "fuck"}) = true`)
	}
}
func TestInSliceInt(t *testing.T) {
	needle := 123
	slice := []interface{}{123, 456, "fuck"}
	if !InSlice(needle, slice) {
		t.Error(`InSlice(123, []interface{}{123, 456, "fuck"}) = true`)
	}
}
func TestNotInSlice(t *testing.T) {
	needle := 1234
	slice := []interface{}{123, 456, "fuck"}
	if InSlice(needle, slice) {
		t.Error(`InSlice(1234, []interface{}{123, 456, "fuck"}) = false`)
	}
}
