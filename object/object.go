package object

import "fmt"

// Type is the representation of the type
// of an object
type Type string

// Object constants
const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

// Object provides an interface that is fulfilled by structs that
// wrap values in the object system
type Object interface {
	Type() Type
	Inspect() string
}

// Integer is the representation of an integer literal in the object
// system
type Integer struct {
	Value int64
}

// Type returns the type of the object
func (i *Integer) Type() Type { return INTEGER_OBJ }

// Inspect returns the value of the integer literal as a string
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Boolean is the representation of a boolean literal in the object
// system
type Boolean struct {
	Value bool
}

// Type returns the type of the object
func (b *Boolean) Type() Type { return BOOLEAN_OBJ }

// Inspect returns the value of the boolean literal as a string
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Null is the respresentation of null (absence of value) in the object
// system
type Null struct{}

// Type returns the type of the object
func (n *Null) Type() Type { return NULL_OBJ }

// Inspect returns the string "null"
func (n *Null) Inspect() string { return "null" }
