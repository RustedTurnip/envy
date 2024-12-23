// package envy provides an interface similar to that of the flag package to
// assist in enviuronment variable handling. The functions of this package make
// it easier to pull in environment variables as different types in a succinc
// manner.
package envy

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
)

type caster[T any] func(string) (T, error)

// StringVar will assign the value of the env variable name to the memory
// address of addr. If there is no env variable set, then the defaultValue will
// be assigned instead.
//
// If there is an env variable matching name, but it is not set (i.e. it has an
// empty string value) then the empty value will be used instead of the
// defaultValue.
func StringVar(addr *string, name, defaultValue string) {
	setVar(addr, name, defaultValue, castString)
}

// IntVar will attempt to convert the value of the env variable name to an int,
// and then assign the converted value to the address of addr. If there is no
// env variable set, then the defaultValue will be assigned instead.
//
// If the value of the env variable name cannot be converted to an int, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func IntVar(addr *int, name string, defaultValue int) {
	setVar(addr, name, defaultValue, strconv.Atoi)
}

// Int64Var will attempt to convert the value of the env variable name to an
// int64, and then assign the converted value to the address of addr. If there
// is no env variable set, then the defaultValue will be assigned instead.
//
// If the value of the env variable name cannot be converted to an int64, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func Int64Var(addr *int64, name string, defaultValue int64) {
	setVar(addr, name, defaultValue, castInt64)
}

// UintVar will attempt to convert the value of the env variable name to a uint,
// and then assign the converted value to the address of addr. If there is no
// env variable set, then the defaultValue will be assigned instead.
//
// If the value of the env variable name cannot be converted to an uint, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func UintVar(addr *uint, name string, defaultValue uint) {
	setVar(addr, name, defaultValue, castUint)
}

// Uint64Var will attempt to convert the value of the env variable name to a
// uint64, and then assign the converted value to the address of addr. If there
// is no env variable set, then the defaultValue will be assigned instead.
//
// If the value of the env variable name cannot be converted to an uint64, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func Uint64Var(addr *uint64, name string, defaultValue uint64) {
	setVar(addr, name, defaultValue, castUint64)
}

// Float64Var will attempt to convert the value of the env variable name to a
// float64, and then assign the converted value to the address of addr. If there
// is no env variable set, then the defaultValue will be assigned instead.
//
// If the value of the env variable name cannot be converted to a float64, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func Float64Var(addr *float64, name string, defaultValue float64) {
	setVar(addr, name, defaultValue, castFloat64)
}

// BoolVar will attempt to convert the value of the env variable name to a bool
// and then assign the converted value to the address of addr. If there is no
// env variable set, then the defaultValue will be assigned instead.
//
// If the value of the env variable name cannot be converted to a bool, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func BoolVar(addr *bool, name string, defaultValue bool) {
	setVar(addr, name, defaultValue, strconv.ParseBool)
}

// DurationVar will attempt to convert the value of the env variable name to a
// time.Duration and then assign the converted value to the address of addr. If
// there is no env variable set, then the defaultValue will be assigned instead.
//
// If the value of the env variable name cannot be converted to a time.Duration,
// then this function will panic. This includes if the env variable is set as
// empty ("").
func DurationVar(addr *time.Duration, name string, defaultValue time.Duration) {
	setVar(addr, name, defaultValue, castDuration)
}

func setVar[T any](addr *T, name string, defaultValue T, cast caster[T]) {
	v, ok := os.LookupEnv(name)
	if !ok {
		*addr = defaultValue

		return
	}

	cv, err := cast(v)
	if err != nil {
		failCast[T](name, err)
	}

	*addr = cv
}

// String returns the value of the env variable name to the memory address of
// addr. If there is no env variable set, then the defaultValue will be returned
// instead.
//
// If there is an env variable matching name, but it is not set (i.e. it has an
// empty string value) then the empty value will be returned instead of the
// defaultValue.
func String(name, defaultValue string) string {
	return value(name, defaultValue, castString)
}

// Int will attempt to convert the value of the env variable name to an int and,
// if successful, will return that value. If the env variable name is not set,
// the defaultValue will be returned instead.
//
// If the value of the env variable name cannot be converted to an int, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func Int(name string, defaultValue int) int {
	return value(name, defaultValue, strconv.Atoi)
}

// Int64 will attempt to convert the value of the env variable name to an int64
// and, if successful, will return that value. If the env variable name is not
// set, the defaultValue will be returned instead.
//
// If the value of the env variable name cannot be converted to an int64, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func Int64(name string, defaultValue int64) int64 {
	return value(name, defaultValue, castInt64)
}

// Uint will attempt to convert the value of the env variable name to a uint
// and, if successful, will return that value. If the env variable name is not
// set, the defaultValue will be returned instead.
//
// If the value of the env variable name cannot be converted to an uint, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func Uint(name string, defaultValue uint) uint {
	return value(name, defaultValue, castUint)
}

// Uint64 will attempt to convert the value of the env variable name to a uint64
// and, if successful, will return that value. If the env variable name is not
// set, the defaultValue will be returned instead.
//
// If the value of the env variable name cannot be converted to an uint64, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func Uint64(name string, defaultValue uint64) uint64 {
	return value(name, defaultValue, castUint64)
}

// Float64 will attempt to convert the value of the env variable name to a
// float64 and, if successful, will return that value. If the env variable name
// is not set, the defaultValue will be returned instead.
//
// If the value of the env variable name cannot be converted to a float64, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func Float64(name string, defaultValue float64) float64 {
	return value(name, defaultValue, castFloat64)
}

// Bool will attempt to convert the value of the env variable name to a bool
// and, if successful, will return that value. If the env variable name is not
// set, the defaultValue will be returned instead.
//
// If the value of the env variable name cannot be converted to a bool, then
// this function will panic. This includes if the env variable is set as empty
// ("").
func Bool(name string, defaultValue bool) bool {
	return value(name, defaultValue, strconv.ParseBool)
}

// Duration will attempt to convert the value of the env variable name to a
// time.Duration and, if successful, will return that value. If the env variable
// name is not set, the defaultValue will be returned instead.
//
// If the value of the env variable name cannot be converted to a time.Duration,
// then this function will panic. This includes if the env variable is set as
// empty ("").
func Duration(name string, defaultValue time.Duration) time.Duration {
	return value(name, defaultValue, castDuration)
}

func value[T any](name string, value T, convert caster[T]) T {
	s, ok := os.LookupEnv(name)
	if !ok {
		return value()
	}

	v, err := convert(s)
	if err != nil {
		failCast[T](name, err)
	}

	return v
}

func castString(s string) (string, error) {
	return s, nil
}

func castInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func castUint(s string) (uint, error) {
	i, err := castUint64(s)
	if err != nil {
		return 0, err
	}

	return uint(i), nil
}

func castUint64(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func castFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

func castDuration(s string) (time.Duration, error) {
	i, err := castInt64(s)
	if err != nil {
		return 0, err
	}

	return time.Duration(i), nil
}

func failCast[T any](variable string, err error) {
	var zero T

	panic(fmt.Errorf(
		"failed to parse %s as %s: %w",
		variable,
		reflect.TypeOf(zero).String(),
		err,
	))
}
