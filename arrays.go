package arrays

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// Check this value is in the array/slice.
func In[T any](v T, array []T) bool {
	if n := len(array); array != nil && n != 0 {
		for i := 0; i < n; i++ {
			if reflect.DeepEqual(v, array[i]) {

				return true
			}
		}
	}
	return false
}

// Find the index of the value in the array/slice.
// If not exist return -1 else return index.
func Index[T any](v T, array []T) int {
	if n := len(array); array != nil && n != 0 {
		for i := 0; i < n; i++ {
			if reflect.DeepEqual(v, array[i]) {
				return i
			}
		}
	}
	return -1
}

// Returns the value of a single column in the input array/slice.
func Column[T, V any](array []T, k any) any {
	n := len(array)
	values := make([]V, n)
	if array != nil && n != 0 {
		switch reflect.TypeOf(array).Elem().Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < n; i++ {
				values[i] = reflect.ValueOf(array[i]).Index(int(reflect.ValueOf(k).Int())).Interface().(V)
			}
			break
		case reflect.Map:
			for i := 0; i < n; i++ {
				values[i] = reflect.ValueOf(array[i]).MapIndex(reflect.ValueOf(k)).Interface().(V)
			}
			break
		case reflect.Struct:
			for i := 0; i < n; i++ {
				values[i] = reflect.ValueOf(array[i]).FieldByName(reflect.ValueOf(k).String()).Interface().(V)
			}
			break
		}
	}
	return values
}

// Count the number of occurrences of all values
func CountValues[T any](array []T) map[any]int {
	var container = make(map[any]int)
	if n := len(array); array != nil && n != 0 {
		for _, v := range array {
			if _, ok := container[v]; !ok {
				container[v] = 1
			} else {
				container[v] += 1
			}
		}
	}
	return container
}

// Slice difference
func Diff[T any](array1, array2 []T) (result []T) {
	result = make([]T, 0)
	if array1 != nil && array2 != nil {
		for k, v := range CountValues(append(array1, array2...)) {
			if v == 1 {
				result = append(result, k.(T))
			}
		}
	}
	return result
}

// Slice intersection
func Intersection[T any](array1, array2 []T) (result []T) {
	result = make([]T, 0)
	if array1 != nil && array2 != nil {
		for k, v := range CountValues(append(array1, array2...)) {
			if v > 1 {
				result = append(result, k.(T))
			}
		}
	}
	return result
}

// slice/array remove duplicate values
func Distinct[T any](array []T) (result []T) {
	if n := len(array); array != nil && n != 0 {
		container := make(map[any]any)
		for i := 0; i < n; i++ {
			container[array[i]] = nil
		}

		result = make([]T, len(container))
		var i = 0
		for k, _ := range container {
			result[i] = k.(T)
			i++
		}
	}
	return result
}

// according to the specified key remove value
func Remove[T any](array []T, k int) []T {
	switch {
	case k == 0:
		array = array[1:]
		break
	case k > 0 && k < len(array)-1:
		array = append(array[0:k], array[k+1:]...)
		break
	case k == len(array):
		array = array[0:k]
		break
	default:
		panic(fmt.Sprintf("index out of range [%d]", k))
	}
	return array
}

// according to the specified key insert value
func Insert[T any](array []T, v T, k int) []T {
	switch {
	case k == 0:
		array = append([]T{v}, array...)
		break
	case k > 0 && k < len(array):
		array = append(append(array[0:k], v), array[k:]...)
		break
	case k == len(array):
		array = append(array, v)
		break
	default:
		panic(fmt.Sprintf("index out of range [%d]", k))
	}
	return array
}

// shuffle slice/array
func Shuffle[T any](array []T) []T {
	if n := len(array); array != nil && n != 0 {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < n; i++ {
			j := rand.Intn(n)
			array[i], array[j] = array[j], array[i]
		}
	}
	return array
}

// type number
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// sum number total
func Sum[T Number](array []T) T {
	var total T
	n := len(array)
	for i := 0; i < n; i++ {
		total += array[i]
	}
	return total
}

// sum string number total return float64
func SumStrNumber(array []string) (total float64) {
	n := len(array)
	for i := 0; i < n; i++ {
		float, _ := strconv.ParseFloat(array[i], 64)
		total += float
	}
	return
}
