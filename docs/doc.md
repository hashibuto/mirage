<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# mirage

```go
import "github.com/hashibuto/mirage"
```

## Index

- [type Info](<#type-info>)
- [type Reflection](<#type-reflection>)
  - [func Reflect(obj any, tagName string) *Reflection](<#func-reflect>)
  - [func (r *Reflection) FieldByIdx(idx int) (*reflect.StructField, error)](<#func-reflection-fieldbyidx>)
  - [func (r *Reflection) FieldByName(fieldName string) (*reflect.StructField, error)](<#func-reflection-fieldbyname>)
  - [func (r *Reflection) HasFieldName(fieldName string) bool](<#func-reflection-hasfieldname>)
  - [func (r *Reflection) HasTagKey(tagKeyName string) bool](<#func-reflection-hastagkey>)
  - [func (r *Reflection) InfoByIdx(idx int) (*Info, error)](<#func-reflection-infobyidx>)
  - [func (r *Reflection) InfoByName(fieldName string) (*Info, error)](<#func-reflection-infobyname>)
  - [func (r *Reflection) InfoByTagKey(fieldName string) (*Info, error)](<#func-reflection-infobytagkey>)
  - [func (r *Reflection) Io() *ReflectionIo](<#func-reflection-io>)
  - [func (r *Reflection) Keys() []string](<#func-reflection-keys>)
  - [func (r *Reflection) NewIo(obj any) *ReflectionIo](<#func-reflection-newio>)
  - [func (r *Reflection) NumFields() int](<#func-reflection-numfields>)
  - [func (r *Reflection) TagKeys() []string](<#func-reflection-tagkeys>)
- [type ReflectionIo](<#type-reflectionio>)
  - [func (r *ReflectionIo) InstantiateByIdx(idx int) (any, error)](<#func-reflectionio-instantiatebyidx>)
  - [func (r *ReflectionIo) InstantiateByName(name string) (any, error)](<#func-reflectionio-instantiatebyname>)
  - [func (r *ReflectionIo) InstantiateByTagKey(tagKey string) (any, error)](<#func-reflectionio-instantiatebytagkey>)
  - [func (r *ReflectionIo) IsNilPointerByIdx(idx int) (bool, error)](<#func-reflectionio-isnilpointerbyidx>)
  - [func (r *ReflectionIo) IsNilPointerByName(name string) (bool, error)](<#func-reflectionio-isnilpointerbyname>)
  - [func (r *ReflectionIo) IsNilPointerByTagKey(tagKey string) (bool, error)](<#func-reflectionio-isnilpointerbytagkey>)
  - [func (r *ReflectionIo) SetValueByIdx(idx int, value any) error](<#func-reflectionio-setvaluebyidx>)
  - [func (r *ReflectionIo) SetValueByName(name string, value any) error](<#func-reflectionio-setvaluebyname>)
  - [func (r *ReflectionIo) SetValueByTagKey(tagKey string, value any) error](<#func-reflectionio-setvaluebytagkey>)
  - [func (r *ReflectionIo) ValueFromIdx(idx int) (any, error)](<#func-reflectionio-valuefromidx>)
  - [func (r *ReflectionIo) ValueFromName(name string) (any, error)](<#func-reflectionio-valuefromname>)
  - [func (r *ReflectionIo) ValueFromTagKey(tagKey string) (any, error)](<#func-reflectionio-valuefromtagkey>)
- [type StringSet](<#type-stringset>)
  - [func NewStringSet(values []string) StringSet](<#func-newstringset>)
  - [func (ss StringSet) Has(value string) bool](<#func-stringset-has>)


## type [Info](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L9-L14>)

```go
type Info struct {
    Name      string
    TagKey    string
    IsPointer bool
    Kind      reflect.Kind
}
```

## type [Reflection](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L16-L23>)

```go
type Reflection struct {
    // contains filtered or unexported fields
}
```

### func [Reflect](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L26>)

```go
func Reflect(obj any, tagName string) *Reflection
```

Reflect produces a reflected version of "obj", including information about tag key names indicated by "tagName"

### func \(\*Reflection\) [FieldByIdx](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L161>)

```go
func (r *Reflection) FieldByIdx(idx int) (*reflect.StructField, error)
```

FieldByIdx returns a field struct by index

### func \(\*Reflection\) [FieldByName](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L151>)

```go
func (r *Reflection) FieldByName(fieldName string) (*reflect.StructField, error)
```

FieldByName returns a field struct by field name

### func \(\*Reflection\) [HasFieldName](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L180>)

```go
func (r *Reflection) HasFieldName(fieldName string) bool
```

HasFieldName returns true if the field name exists

### func \(\*Reflection\) [HasTagKey](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L174>)

```go
func (r *Reflection) HasTagKey(tagKeyName string) bool
```

HasTagKey returns true if the tag key exists

### func \(\*Reflection\) [InfoByIdx](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L142>)

```go
func (r *Reflection) InfoByIdx(idx int) (*Info, error)
```

### func \(\*Reflection\) [InfoByName](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L125>)

```go
func (r *Reflection) InfoByName(fieldName string) (*Info, error)
```

InfoByName returns the reflect kind for a given field by name

### func \(\*Reflection\) [InfoByTagKey](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L134>)

```go
func (r *Reflection) InfoByTagKey(fieldName string) (*Info, error)
```

InfoByTagKey returns the reflect kind for a given tag key

### func \(\*Reflection\) [Io](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L100>)

```go
func (r *Reflection) Io() *ReflectionIo
```

Io returns a reflection io object for this instance of the reflected object

### func \(\*Reflection\) [Keys](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L80>)

```go
func (r *Reflection) Keys() []string
```

Keys returns an array of key names

### func \(\*Reflection\) [NewIo](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L112>)

```go
func (r *Reflection) NewIo(obj any) *ReflectionIo
```

Io returns a reflection io object for a new instance of the reflected object

### func \(\*Reflection\) [NumFields](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L169>)

```go
func (r *Reflection) NumFields() int
```

NumFields returns the number of fields on the structure

### func \(\*Reflection\) [TagKeys](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L90>)

```go
func (r *Reflection) TagKeys() []string
```

TagKeys returns an array of keys garnered from a given tag name

## type [ReflectionIo](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L185-L188>)

```go
type ReflectionIo struct {
    // contains filtered or unexported fields
}
```

### func \(\*ReflectionIo\) [InstantiateByIdx](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L297>)

```go
func (r *ReflectionIo) InstantiateByIdx(idx int) (any, error)
```

InstantiateByIdx instantiates a new empty value of the field type and returns it

### func \(\*ReflectionIo\) [InstantiateByName](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L309>)

```go
func (r *ReflectionIo) InstantiateByName(name string) (any, error)
```

InstantiateByName instantiates a new empty value of the field type and returns it

### func \(\*ReflectionIo\) [InstantiateByTagKey](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L319>)

```go
func (r *ReflectionIo) InstantiateByTagKey(tagKey string) (any, error)
```

InstantiateByTagKey instantiates a new empty value of the field type and returns it

### func \(\*ReflectionIo\) [IsNilPointerByIdx](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L250>)

```go
func (r *ReflectionIo) IsNilPointerByIdx(idx int) (bool, error)
```

IsNilPointerByIdx returns true if the underlying value is a nil pointer

### func \(\*ReflectionIo\) [IsNilPointerByName](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L263>)

```go
func (r *ReflectionIo) IsNilPointerByName(name string) (bool, error)
```

IsNilPointerByName returns true if the underlying value is a nil pointer

### func \(\*ReflectionIo\) [IsNilPointerByTagKey](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L280>)

```go
func (r *ReflectionIo) IsNilPointerByTagKey(tagKey string) (bool, error)
```

IsNilPointerByTagKey returns true if the underlying value is a nil pointer

### func \(\*ReflectionIo\) [SetValueByIdx](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L219>)

```go
func (r *ReflectionIo) SetValueByIdx(idx int, value any) error
```

SetValueByIdx sets a value on the reflected object using the field index

### func \(\*ReflectionIo\) [SetValueByName](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L228>)

```go
func (r *ReflectionIo) SetValueByName(name string, value any) error
```

SetValueByName sets a value on the reflected object using the field name

### func \(\*ReflectionIo\) [SetValueByTagKey](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L239>)

```go
func (r *ReflectionIo) SetValueByTagKey(tagKey string, value any) error
```

SetValueByTagKey sets a value on the reflected object using the tag key

### func \(\*ReflectionIo\) [ValueFromIdx](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L191>)

```go
func (r *ReflectionIo) ValueFromIdx(idx int) (any, error)
```

ValueFromIdx returns the struct value referenced by the field index

### func \(\*ReflectionIo\) [ValueFromName](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L199>)

```go
func (r *ReflectionIo) ValueFromName(name string) (any, error)
```

ValueFromName returns the struct value referenced by the field name

### func \(\*ReflectionIo\) [ValueFromTagKey](<https://github.com/hashibuto/mirage/blob/master/mirage.go#L209>)

```go
func (r *ReflectionIo) ValueFromTagKey(tagKey string) (any, error)
```

ValueFromTagKey returns the struct value referenced by the tag key

## type [StringSet](<https://github.com/hashibuto/mirage/blob/master/string-set.go#L3>)

```go
type StringSet map[string]struct{}
```

### func [NewStringSet](<https://github.com/hashibuto/mirage/blob/master/string-set.go#L6>)

```go
func NewStringSet(values []string) StringSet
```

NewStringSet returns a new string set

### func \(StringSet\) [Has](<https://github.com/hashibuto/mirage/blob/master/string-set.go#L16>)

```go
func (ss StringSet) Has(value string) bool
```

Has returns true if the value is present in the set



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
