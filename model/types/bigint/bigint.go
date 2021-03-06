package bigint

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"
)

type _bi = big.Int
type BigInt struct {
	_bi
}

func (bi *BigInt) BI() _bi {
	return bi._bi
}

var ZERO BigInt

func init() {
	if err := ZERO.CreateFromString("0", 10); err != nil {
		panic(err)
	}
}

func (bi BigInt) Value() (driver.Value, error) {
	return []byte(bi.String()), nil
}
func (bi *BigInt) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return bi.scanBytes(src)
	case string:
		return bi.scanBytes([]byte(src))
	case nil:
		bi = nil
		return nil
	default:
		return fmt.Errorf("pq: cannot convert %T to BigInt", src)
	}
}

func (bi *BigInt) scanBytes(src []byte) error {
	return bi.CreateFromString(string(src), 10)
}

func (bi *BigInt) CreateFromString(s string, base int) error {
	_, ok := bi.SetString(s, base)
	if !ok {
		return errors.New("create bigint from string failed: " + s)
	}
	return nil
}

//@todo xml protobuf ...
func (bi *BigInt) MarshalJSON() ([]byte, error) {
	return []byte(bi.String()), nil
}

func (bi *BigInt) SetUint64(i uint64) *BigInt {
	bi._bi.SetUint64(i)
	return bi
}
func (bi *BigInt) SetInt64(i int64) *BigInt {
	bi._bi.SetInt64(i)
	return bi
}

func (bi *BigInt) Add(a BigInt, b BigInt) {
	bi._bi.Add(&a._bi, &b._bi)
}
func (bi *BigInt) Sub(a BigInt, b BigInt) {
	bi._bi.Sub(&a._bi, &b._bi)
}
func (bi *BigInt) Mul(a BigInt, b BigInt) {
	bi._bi.Mul(&a._bi, &b._bi)
}
func (bi *BigInt) Div(a BigInt, b BigInt) {
	bi._bi.Quo(&a._bi, &b._bi)
}
func (bi *BigInt) Abs(a BigInt) {
	bi._bi.Abs(&a._bi)
}
func (bi *BigInt) Cmp(a BigInt) int {
	return bi._bi.Cmp(&a._bi)
}

//
//func main(){
//	a := BigInt{}
//	a.SetString("10", 10)
//	b := BigInt{}
//	b.SetString("11", 10)
//	c := BigInt{}
//	c.Add(&a.BI, &b.BI)
//}
