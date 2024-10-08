package channeld

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"unsafe"

	"github.com/channeldorg/channeld/pkg/common"
	"github.com/stretchr/testify/assert"
)

//go:linkname typelinks reflect.typelinks
func typelinks() (sections []unsafe.Pointer, offset [][]int32)

//go:linkname add reflect.add
func add(p unsafe.Pointer, x uintptr, whySafe string) unsafe.Pointer

func TestReflectionTypes(t *testing.T) {

	sections, offsets := typelinks()
	for i, base := range sections {
		for _, offset := range offsets[i] {
			typeAddr := add(base, uintptr(offset), "")
			typ := reflect.TypeOf(*(*interface{})(unsafe.Pointer(&typeAddr)))
			fmt.Println(typ)
		}
	}
}

func TestGetNextId(t *testing.T) {
	m := make(map[uint32]interface{})
	var index uint32 = 1
	var ok bool

	index, _ = GetNextId(&m, index, 1, 3)
	assert.EqualValues(t, 1, index)
	m[index] = "aaa"

	index, _ = GetNextId(&m, index, 1, 3)
	assert.EqualValues(t, 2, index)
	m[index] = "bbb"

	index, _ = GetNextId(&m, index, 1, 3)
	assert.EqualValues(t, 3, index)
	m[index] = "ccc"

	_, ok = GetNextId(&m, index, 1, 3)
	assert.False(t, ok)
}

func BenchmarkGetNextId(b *testing.B) {
	m := make(map[uint32]interface{})
	var index uint32 = 1
	var ok bool

	for i := 0; i < b.N; i++ {
		index, ok = GetNextId(&m, index, 1, 65535)
		if ok {
			m[index] = i
		} else {
			break
		}
	}
}

func TestGetNextIdSync(t *testing.T) {
	m := sync.Map{}
	var index common.ChannelId = 1
	var ok bool

	index, _ = GetNextIdSync(&m, index, 1, 3)
	assert.EqualValues(t, 1, index)
	m.Store(index, "aaa")

	index, _ = GetNextIdSync(&m, index, 1, 3)
	assert.EqualValues(t, 2, index)
	m.Store(index, "bbb")

	index, _ = GetNextIdSync(&m, index, 1, 3)
	assert.EqualValues(t, 3, index)
	m.Store(index, "ccc")

	_, ok = GetNextIdSync(&m, index, 1, 3)
	assert.False(t, ok)

}

// 3x slower as of BenchmarkGetNextId
func BenchmarkGetNextIdSync(b *testing.B) {
	m := sync.Map{}
	var index common.ChannelId = 1
	var ok bool

	for i := 0; i < b.N; i++ {
		index, ok = GetNextIdSync(&m, index, 1, 65535)
		if ok {
			m.Store(index, i)
		} else {
			break
		}
	}
}
