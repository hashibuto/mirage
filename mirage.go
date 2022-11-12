package mirage

import (
	"fmt"
	"reflect"
	"strings"
)

type Kind struct {
	IsPointer bool
	Kind      reflect.Kind
}

var Invalid = &Kind{
	IsPointer: false,
	Kind:      reflect.Invalid,
}

type Reflection struct {
	idxByName   map[string]int
	idxByTagKey map[string]int
	fieldByIdx  map[int]reflect.StructField
	kindByIdx   map[int]*Kind
	elem        reflect.Value
}

// Reflect produces a reflected version of "obj", including information about tag key names indicated by "tagName"
func Reflect(obj any, tagName string) *Reflection {
	idxByName := map[string]int{}
	idxByTagKey := map[string]int{}
	fieldByIdx := map[int]reflect.StructField{}
	kindByIdx := map[int]*Kind{}

	value := reflect.ValueOf(obj)
	elem := value.Elem()
	numFields := elem.NumField()
	elemType := elem.Type()
	for idx := 0; idx < numFields; idx++ {
		field := elemType.Field(idx)
		if tagName == "" {
			idxByTagKey[field.Name] = idx
		} else {
			tag := field.Tag.Get(tagName)
			if tag == "" {
				idxByTagKey[field.Name] = idx
			} else {
				idxByTagKey[strings.Split(tag, ",")[0]] = idx
			}
		}
		idxByName[field.Name] = idx
		fieldByIdx[idx] = field

		kind := field.Type.Kind()
		isPointer := false
		if kind == reflect.Pointer {
			// Indirect to get the underlying kind
			kind = field.Type.Elem().Kind()
			isPointer = true
		}
		kindByIdx[idx] = &Kind{
			IsPointer: isPointer,
			Kind:      kind,
		}
	}

	return &Reflection{
		idxByName:   idxByName,
		idxByTagKey: idxByTagKey,
		fieldByIdx:  fieldByIdx,
		kindByIdx:   kindByIdx,
		elem:        elem,
	}
}

// Keys returns an array of key names
func (r *Reflection) Keys() []string {
	keys := make([]string, len(r.fieldByIdx))
	for idx, field := range r.fieldByIdx {
		keys[idx] = field.Name
	}

	return keys
}

// TagKeys returns an array of keys garnered from a given tag name
func (r *Reflection) TagKeys() []string {
	tagKeys := make([]string, len(r.fieldByIdx))
	for tagKey, idx := range r.idxByTagKey {
		tagKeys[idx] = tagKey
	}

	return tagKeys
}

// Io returns a reflection io object for this instance of the reflected object
func (r *Reflection) Io() *ReflectionIo {
	valueByIdx := map[int]reflect.Value{}
	for idx := 0; idx < len(r.fieldByIdx); idx++ {
		valueByIdx[idx] = r.elem.Field(idx)
	}
	return &ReflectionIo{
		reflection: r,
		valueByIdx: valueByIdx,
	}
}

// Io returns a reflection io object for a new instance of the reflected object
func (r *Reflection) NewIo(obj any) *ReflectionIo {
	valueByIdx := map[int]reflect.Value{}
	elem := reflect.ValueOf(obj).Elem()
	for idx := 0; idx < len(r.fieldByIdx); idx++ {
		valueByIdx[idx] = elem.Field(idx)
	}
	return &ReflectionIo{
		reflection: r,
		valueByIdx: valueByIdx,
	}
}

// KindByName returns the reflect kind for a given field by name
func (r *Reflection) KindByName(fieldName string) (*Kind, error) {
	idx, ok := r.idxByName[fieldName]
	if !ok {
		return Invalid, fmt.Errorf("Unknown field name \"%s\"", fieldName)
	}
	return r.kindByIdx[idx], nil
}

// KindByTagKey returns the reflect kind for a given tag key
func (r *Reflection) KindByTagKey(fieldName string) (*Kind, error) {
	idx, ok := r.idxByTagKey[fieldName]
	if !ok {
		return Invalid, fmt.Errorf("Unknown tag key name \"%s\"", fieldName)
	}
	return r.kindByIdx[idx], nil
}

type ReflectionIo struct {
	reflection *Reflection
	valueByIdx map[int]reflect.Value
}

// ValueFromName returns the struct value referenced by the field name
func (r *ReflectionIo) ValueFromName(name string) (any, error) {
	idx, ok := r.reflection.idxByName[name]
	if !ok {
		return nil, fmt.Errorf("Unable to locate field name %s in object", name)
	}

	return r.valueByIdx[idx].Interface(), nil
}

// ValueFromTagKey returns the struct value referenced by the tag key
func (r *ReflectionIo) ValueFromTagKey(tagKey string) (any, error) {
	idx, ok := r.reflection.idxByTagKey[tagKey]
	if !ok {
		return nil, fmt.Errorf("Unable to locate tag key %s in object", tagKey)
	}

	return r.valueByIdx[idx].Interface(), nil
}

// SetValueByName sets a value on the reflected object using the field name
func (r *ReflectionIo) SetValueByName(name string, value any) error {
	idx, ok := r.reflection.idxByName[name]
	if !ok {
		return fmt.Errorf("Unable to locate field name %s in object", name)
	}

	r.valueByIdx[idx].Set(reflect.ValueOf(value))
	return nil
}

// SetValueByTagKey sets a value on the reflected object using the tag key
func (r *ReflectionIo) SetValueByTagKey(tagKey string, value any) error {
	idx, ok := r.reflection.idxByTagKey[tagKey]
	if !ok {
		return fmt.Errorf("Unable to locate tag key %s in object", tagKey)
	}

	r.valueByIdx[idx].Set(reflect.ValueOf(value))
	return nil
}
