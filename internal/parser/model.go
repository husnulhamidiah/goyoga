package parser

type API struct {
	Functions []Function
	Variables []Variable
	Enums     []Enum
	Structs   []Struct
	Typedefs  []Typedef
}

type Type struct {
	Name     string
	Const    bool
	Struct   bool
	Pointers int
}

type Field struct {
	Name string
	Type Type
}

type Param struct {
	Name string
	Type Type
}

type Function struct {
	Name   string
	Return Type
	Params []Param
}

type Variable struct {
	Name string
	Type Type
}

type Enum struct {
	Name   string
	Values []EnumValue
}

type EnumValue struct {
	Name  string
	Value string
}

type Struct struct {
	Name   string
	Fields []Field
}

type Typedef struct {
	Name            string
	Type            Type
	FunctionPointer *Function
}
