package main

import (
	"encoding/binary"
	"fmt"
	"hash"
	"hash/crc32"
	"time"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type HashMap struct {
	data HashMaper
}

type DataList struct {
	hasher hash.Hash
	data   []*ListValue
}

type ListValue struct {
	key   string
	value interface{}
	next  *ListValue
}

type DataSlice struct {
	hasher hash.Hash
	data   [][]SliceValue
}

type SliceValue struct {
	key   string
	value interface{}
}

func GetHasher() hash.Hash {
	return crc32.NewIEEE()
}

func CalculateHash(h hash.Hash, input string) uint32 {
	h.Reset()
	h.Write([]byte(input))
	return binary.BigEndian.Uint32(h.Sum(nil))
}

func GetCellNumber(cap, hash uint32) int {
	return int(hash % cap)
}

func (hm HashMap) Set(key string, value interface{}) {
	hm.data.Set(key, value)
}

func (hm HashMap) Get(key string) (interface{}, bool) {
	return hm.data.Get(key)
}

func (list DataList) Set(key string, value interface{}) {
	keysCapacity := uint32(cap(list.data))
	if keysCapacity == 0 {
		fmt.Println("hash map capacity = 0")
		return
	}

	hash := CalculateHash(list.hasher, key)
	targetCell := int(hash % keysCapacity)

	if list.data[targetCell] == nil {
		list.data[targetCell] = &ListValue{
			key:   key,
			value: value,
			next:  nil,
		}

		return
	}

	checkData := list.data[targetCell]
	for {
		if checkData.key == key {
			checkData.value = value
			return
		}

		if checkData.next == nil {
			checkData.next = &ListValue{
				key:   key,
				value: value,
				next:  nil,
			}
			return
		}

		checkData = checkData.next
	}
}

func (list DataList) Get(key string) (interface{}, bool) {
	keysCapacity := uint32(cap(list.data))
	if keysCapacity == 0 {
		fmt.Println("hash map capacity = 0")
		return nil, false
	}

	hash := CalculateHash(list.hasher, key)
	targetCell := int(hash % keysCapacity)

	checkData := list.data[targetCell]
	if checkData == nil {
		return nil, false
	}

	for {
		if checkData.key == key {
			return checkData.value, true
		}

		if checkData.next == nil {
			return nil, false
		}

		checkData = checkData.next
	}
}

func (slice DataSlice) Set(key string, value interface{}) {
	keysCapacity := uint32(cap(slice.data))
	if keysCapacity == 0 {
		fmt.Println("hash map capacity = 0")
		return
	}

	hash := CalculateHash(slice.hasher, key)
	targetCell := int(hash % keysCapacity)

	if len(slice.data[targetCell]) == 0 {
		newValue := SliceValue{
			key:   key,
			value: value,
		}
		slice.data[targetCell] = append(slice.data[targetCell], newValue)
	}

	var lastValue int
	for lastValue = 0; lastValue < len(slice.data[targetCell]); lastValue++ {
		value := slice.data[targetCell][lastValue]

		if value.key == key {
			value.value = value
			return
		}
	}

	newValue := SliceValue{
		key:   key,
		value: value,
	}
	slice.data[targetCell] = append(slice.data[targetCell], newValue)
}

func (slice DataSlice) Get(key string) (interface{}, bool) {
	keysCapacity := uint32(cap(slice.data))
	if keysCapacity == 0 {
		fmt.Println("hash map capacity = 0")
		return nil, false
	}

	hash := CalculateHash(slice.hasher, key)
	targetCell := int(hash % keysCapacity)

	checkData := slice.data[targetCell]
	for _, data := range checkData {
		if data.key == key {
			return data.value, true
		}
	}

	return nil, false
}

func NewHashMapSlice(count int, options ...func(*HashMap)) HashMaper {
	data := make([][]SliceValue, count)
	hasher := GetHasher()
	hashData := DataSlice{
		hasher: hasher,
		data:   data,
	}
	hashMap := HashMap{
		data: hashData,
	}

	for _, option := range options {
		option(&hashMap)
	}

	return hashMap
}
func NewHashMapList(count int, options ...func(*HashMap)) HashMaper {
	data := make([]*ListValue, count)
	hasher := GetHasher()
	hashData := DataList{
		hasher: hasher,
		data:   data,
	}
	hashMap := HashMap{
		data: hashData,
	}

	for _, option := range options {
		option(&hashMap)
	}

	return hashMap
}
func MeassureTime(f func()) time.Duration {
	start := time.Now()
	f()
	since := time.Since(start)
	return since
}
func main() {
	time := MeassureTime(TestSlice16)
	fmt.Println(time)
	time = MeassureTime(TestSlice1000)
	fmt.Println(time)
	time = MeassureTime(TestSlice100000)
	fmt.Println(time)
	time = MeassureTime(TestList16)
	fmt.Println(time)
	time = MeassureTime(TestList1000)
	fmt.Println(time)
	time = MeassureTime(TestList100000)
	fmt.Println(time)
}
func TestList16() {
	m := NewHashMapList(16)
	for i := 0; i < 16; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 16; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}
func TestList1000() {
	m := NewHashMapList(1000)
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 1000; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}

func TestList100000() {
	m := NewHashMapList(100000)
	for i := 0; i < 100000; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 100000; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}
func TestSlice16() {
	m := NewHashMapSlice(16)
	for i := 0; i < 16; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 16; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}
func TestSlice1000() {
	m := NewHashMapSlice(1000)
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 1000; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}

func TestSlice100000() {
	m := NewHashMapSlice(100000)
	for i := 0; i < 100000; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 100000; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}
