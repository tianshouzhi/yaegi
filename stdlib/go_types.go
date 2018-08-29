package stdlib

// Generated by 'goexports go/types'. Do not edit!

import (
	"go/types"
	"reflect"
)

func init() {
	Value["go/types"] = map[string]reflect.Value{
		"AssertableTo":            reflect.ValueOf(types.AssertableTo),
		"AssignableTo":            reflect.ValueOf(types.AssignableTo),
		"Bool":                    reflect.ValueOf(types.Bool),
		"Byte":                    reflect.ValueOf(types.Byte),
		"Comparable":              reflect.ValueOf(types.Comparable),
		"Complex128":              reflect.ValueOf(types.Complex128),
		"Complex64":               reflect.ValueOf(types.Complex64),
		"ConvertibleTo":           reflect.ValueOf(types.ConvertibleTo),
		"DefPredeclaredTestFuncs": reflect.ValueOf(types.DefPredeclaredTestFuncs),
		"Default":                 reflect.ValueOf(types.Default),
		"Eval":                    reflect.ValueOf(types.Eval),
		"ExprString":              reflect.ValueOf(types.ExprString),
		"FieldVal":                reflect.ValueOf(types.FieldVal),
		"Float32":                 reflect.ValueOf(types.Float32),
		"Float64":                 reflect.ValueOf(types.Float64),
		"Id":                      reflect.ValueOf(types.Id),
		"Identical":               reflect.ValueOf(types.Identical),
		"IdenticalIgnoreTags":     reflect.ValueOf(types.IdenticalIgnoreTags),
		"Implements":              reflect.ValueOf(types.Implements),
		"Int":                     reflect.ValueOf(types.Int),
		"Int16":                   reflect.ValueOf(types.Int16),
		"Int32":                   reflect.ValueOf(types.Int32),
		"Int64":                   reflect.ValueOf(types.Int64),
		"Int8":                    reflect.ValueOf(types.Int8),
		"Invalid":                 reflect.ValueOf(types.Invalid),
		"IsBoolean":               reflect.ValueOf(types.IsBoolean),
		"IsComplex":               reflect.ValueOf(types.IsComplex),
		"IsConstType":             reflect.ValueOf(types.IsConstType),
		"IsFloat":                 reflect.ValueOf(types.IsFloat),
		"IsInteger":               reflect.ValueOf(types.IsInteger),
		"IsInterface":             reflect.ValueOf(types.IsInterface),
		"IsNumeric":               reflect.ValueOf(types.IsNumeric),
		"IsOrdered":               reflect.ValueOf(types.IsOrdered),
		"IsString":                reflect.ValueOf(types.IsString),
		"IsUnsigned":              reflect.ValueOf(types.IsUnsigned),
		"IsUntyped":               reflect.ValueOf(types.IsUntyped),
		"LookupFieldOrMethod":     reflect.ValueOf(types.LookupFieldOrMethod),
		"MethodExpr":              reflect.ValueOf(types.MethodExpr),
		"MethodVal":               reflect.ValueOf(types.MethodVal),
		"MissingMethod":           reflect.ValueOf(types.MissingMethod),
		"NewArray":                reflect.ValueOf(types.NewArray),
		"NewChan":                 reflect.ValueOf(types.NewChan),
		"NewChecker":              reflect.ValueOf(types.NewChecker),
		"NewConst":                reflect.ValueOf(types.NewConst),
		"NewField":                reflect.ValueOf(types.NewField),
		"NewFunc":                 reflect.ValueOf(types.NewFunc),
		"NewInterface":            reflect.ValueOf(types.NewInterface),
		"NewInterfaceType":        reflect.ValueOf(types.NewInterfaceType),
		"NewLabel":                reflect.ValueOf(types.NewLabel),
		"NewMap":                  reflect.ValueOf(types.NewMap),
		"NewMethodSet":            reflect.ValueOf(types.NewMethodSet),
		"NewNamed":                reflect.ValueOf(types.NewNamed),
		"NewPackage":              reflect.ValueOf(types.NewPackage),
		"NewParam":                reflect.ValueOf(types.NewParam),
		"NewPkgName":              reflect.ValueOf(types.NewPkgName),
		"NewPointer":              reflect.ValueOf(types.NewPointer),
		"NewScope":                reflect.ValueOf(types.NewScope),
		"NewSignature":            reflect.ValueOf(types.NewSignature),
		"NewSlice":                reflect.ValueOf(types.NewSlice),
		"NewStruct":               reflect.ValueOf(types.NewStruct),
		"NewTuple":                reflect.ValueOf(types.NewTuple),
		"NewTypeName":             reflect.ValueOf(types.NewTypeName),
		"NewVar":                  reflect.ValueOf(types.NewVar),
		"ObjectString":            reflect.ValueOf(types.ObjectString),
		"RecvOnly":                reflect.ValueOf(types.RecvOnly),
		"RelativeTo":              reflect.ValueOf(types.RelativeTo),
		"Rune":                    reflect.ValueOf(types.Rune),
		"SelectionString":         reflect.ValueOf(types.SelectionString),
		"SendOnly":                reflect.ValueOf(types.SendOnly),
		"SendRecv":                reflect.ValueOf(types.SendRecv),
		"SizesFor":                reflect.ValueOf(types.SizesFor),
		"String":                  reflect.ValueOf(types.String),
		"Typ":                     reflect.ValueOf(types.Typ),
		"TypeString":              reflect.ValueOf(types.TypeString),
		"Uint":                    reflect.ValueOf(types.Uint),
		"Uint16":                  reflect.ValueOf(types.Uint16),
		"Uint32":                  reflect.ValueOf(types.Uint32),
		"Uint64":                  reflect.ValueOf(types.Uint64),
		"Uint8":                   reflect.ValueOf(types.Uint8),
		"Uintptr":                 reflect.ValueOf(types.Uintptr),
		"Universe":                reflect.ValueOf(types.Universe),
		"Unsafe":                  reflect.ValueOf(types.Unsafe),
		"UnsafePointer":           reflect.ValueOf(types.UnsafePointer),
		"UntypedBool":             reflect.ValueOf(types.UntypedBool),
		"UntypedComplex":          reflect.ValueOf(types.UntypedComplex),
		"UntypedFloat":            reflect.ValueOf(types.UntypedFloat),
		"UntypedInt":              reflect.ValueOf(types.UntypedInt),
		"UntypedNil":              reflect.ValueOf(types.UntypedNil),
		"UntypedRune":             reflect.ValueOf(types.UntypedRune),
		"UntypedString":           reflect.ValueOf(types.UntypedString),
		"WriteExpr":               reflect.ValueOf(types.WriteExpr),
		"WriteSignature":          reflect.ValueOf(types.WriteSignature),
		"WriteType":               reflect.ValueOf(types.WriteType),
	}
	Type["go/types"] = map[string]reflect.Type{
		"Array":         reflect.TypeOf((*types.Array)(nil)).Elem(),
		"Basic":         reflect.TypeOf((*types.Basic)(nil)).Elem(),
		"BasicInfo":     reflect.TypeOf((*types.BasicInfo)(nil)).Elem(),
		"BasicKind":     reflect.TypeOf((*types.BasicKind)(nil)).Elem(),
		"Builtin":       reflect.TypeOf((*types.Builtin)(nil)).Elem(),
		"Chan":          reflect.TypeOf((*types.Chan)(nil)).Elem(),
		"ChanDir":       reflect.TypeOf((*types.ChanDir)(nil)).Elem(),
		"Checker":       reflect.TypeOf((*types.Checker)(nil)).Elem(),
		"Config":        reflect.TypeOf((*types.Config)(nil)).Elem(),
		"Const":         reflect.TypeOf((*types.Const)(nil)).Elem(),
		"Error":         reflect.TypeOf((*types.Error)(nil)).Elem(),
		"Func":          reflect.TypeOf((*types.Func)(nil)).Elem(),
		"ImportMode":    reflect.TypeOf((*types.ImportMode)(nil)).Elem(),
		"Importer":      reflect.TypeOf((*types.Importer)(nil)).Elem(),
		"ImporterFrom":  reflect.TypeOf((*types.ImporterFrom)(nil)).Elem(),
		"Info":          reflect.TypeOf((*types.Info)(nil)).Elem(),
		"Initializer":   reflect.TypeOf((*types.Initializer)(nil)).Elem(),
		"Interface":     reflect.TypeOf((*types.Interface)(nil)).Elem(),
		"Label":         reflect.TypeOf((*types.Label)(nil)).Elem(),
		"Map":           reflect.TypeOf((*types.Map)(nil)).Elem(),
		"MethodSet":     reflect.TypeOf((*types.MethodSet)(nil)).Elem(),
		"Named":         reflect.TypeOf((*types.Named)(nil)).Elem(),
		"Nil":           reflect.TypeOf((*types.Nil)(nil)).Elem(),
		"Object":        reflect.TypeOf((*types.Object)(nil)).Elem(),
		"Package":       reflect.TypeOf((*types.Package)(nil)).Elem(),
		"PkgName":       reflect.TypeOf((*types.PkgName)(nil)).Elem(),
		"Pointer":       reflect.TypeOf((*types.Pointer)(nil)).Elem(),
		"Qualifier":     reflect.TypeOf((*types.Qualifier)(nil)).Elem(),
		"Scope":         reflect.TypeOf((*types.Scope)(nil)).Elem(),
		"Selection":     reflect.TypeOf((*types.Selection)(nil)).Elem(),
		"SelectionKind": reflect.TypeOf((*types.SelectionKind)(nil)).Elem(),
		"Signature":     reflect.TypeOf((*types.Signature)(nil)).Elem(),
		"Sizes":         reflect.TypeOf((*types.Sizes)(nil)).Elem(),
		"Slice":         reflect.TypeOf((*types.Slice)(nil)).Elem(),
		"StdSizes":      reflect.TypeOf((*types.StdSizes)(nil)).Elem(),
		"Struct":        reflect.TypeOf((*types.Struct)(nil)).Elem(),
		"Tuple":         reflect.TypeOf((*types.Tuple)(nil)).Elem(),
		"Type":          reflect.TypeOf((*types.Type)(nil)).Elem(),
		"TypeAndValue":  reflect.TypeOf((*types.TypeAndValue)(nil)).Elem(),
		"TypeName":      reflect.TypeOf((*types.TypeName)(nil)).Elem(),
		"Var":           reflect.TypeOf((*types.Var)(nil)).Elem(),
	}
}
