// Code generated by 'goexports debug/pe'. DO NOT EDIT.

// +build go1.13,!go1.14

package stdlib

import (
	"debug/pe"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["debug/pe"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"COFFSymbolSize":                       reflect.ValueOf(constant.MakeFromLiteral("18", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_ARCHITECTURE":   reflect.ValueOf(constant.MakeFromLiteral("7", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_BASERELOC":      reflect.ValueOf(constant.MakeFromLiteral("5", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT":   reflect.ValueOf(constant.MakeFromLiteral("11", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR": reflect.ValueOf(constant.MakeFromLiteral("14", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_DEBUG":          reflect.ValueOf(constant.MakeFromLiteral("6", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT":   reflect.ValueOf(constant.MakeFromLiteral("13", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_EXCEPTION":      reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_EXPORT":         reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_GLOBALPTR":      reflect.ValueOf(constant.MakeFromLiteral("8", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_IAT":            reflect.ValueOf(constant.MakeFromLiteral("12", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_IMPORT":         reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG":    reflect.ValueOf(constant.MakeFromLiteral("10", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_RESOURCE":       reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_SECURITY":       reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"IMAGE_DIRECTORY_ENTRY_TLS":            reflect.ValueOf(constant.MakeFromLiteral("9", token.INT, 0)),
		"IMAGE_FILE_MACHINE_AM33":              reflect.ValueOf(constant.MakeFromLiteral("467", token.INT, 0)),
		"IMAGE_FILE_MACHINE_AMD64":             reflect.ValueOf(constant.MakeFromLiteral("34404", token.INT, 0)),
		"IMAGE_FILE_MACHINE_ARM":               reflect.ValueOf(constant.MakeFromLiteral("448", token.INT, 0)),
		"IMAGE_FILE_MACHINE_ARM64":             reflect.ValueOf(constant.MakeFromLiteral("43620", token.INT, 0)),
		"IMAGE_FILE_MACHINE_ARMNT":             reflect.ValueOf(constant.MakeFromLiteral("452", token.INT, 0)),
		"IMAGE_FILE_MACHINE_EBC":               reflect.ValueOf(constant.MakeFromLiteral("3772", token.INT, 0)),
		"IMAGE_FILE_MACHINE_I386":              reflect.ValueOf(constant.MakeFromLiteral("332", token.INT, 0)),
		"IMAGE_FILE_MACHINE_IA64":              reflect.ValueOf(constant.MakeFromLiteral("512", token.INT, 0)),
		"IMAGE_FILE_MACHINE_M32R":              reflect.ValueOf(constant.MakeFromLiteral("36929", token.INT, 0)),
		"IMAGE_FILE_MACHINE_MIPS16":            reflect.ValueOf(constant.MakeFromLiteral("614", token.INT, 0)),
		"IMAGE_FILE_MACHINE_MIPSFPU":           reflect.ValueOf(constant.MakeFromLiteral("870", token.INT, 0)),
		"IMAGE_FILE_MACHINE_MIPSFPU16":         reflect.ValueOf(constant.MakeFromLiteral("1126", token.INT, 0)),
		"IMAGE_FILE_MACHINE_POWERPC":           reflect.ValueOf(constant.MakeFromLiteral("496", token.INT, 0)),
		"IMAGE_FILE_MACHINE_POWERPCFP":         reflect.ValueOf(constant.MakeFromLiteral("497", token.INT, 0)),
		"IMAGE_FILE_MACHINE_R4000":             reflect.ValueOf(constant.MakeFromLiteral("358", token.INT, 0)),
		"IMAGE_FILE_MACHINE_SH3":               reflect.ValueOf(constant.MakeFromLiteral("418", token.INT, 0)),
		"IMAGE_FILE_MACHINE_SH3DSP":            reflect.ValueOf(constant.MakeFromLiteral("419", token.INT, 0)),
		"IMAGE_FILE_MACHINE_SH4":               reflect.ValueOf(constant.MakeFromLiteral("422", token.INT, 0)),
		"IMAGE_FILE_MACHINE_SH5":               reflect.ValueOf(constant.MakeFromLiteral("424", token.INT, 0)),
		"IMAGE_FILE_MACHINE_THUMB":             reflect.ValueOf(constant.MakeFromLiteral("450", token.INT, 0)),
		"IMAGE_FILE_MACHINE_UNKNOWN":           reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"IMAGE_FILE_MACHINE_WCEMIPSV2":         reflect.ValueOf(constant.MakeFromLiteral("361", token.INT, 0)),
		"NewFile":                              reflect.ValueOf(pe.NewFile),
		"Open":                                 reflect.ValueOf(pe.Open),

		// type definitions
		"COFFSymbol":       reflect.ValueOf((*pe.COFFSymbol)(nil)),
		"DataDirectory":    reflect.ValueOf((*pe.DataDirectory)(nil)),
		"File":             reflect.ValueOf((*pe.File)(nil)),
		"FileHeader":       reflect.ValueOf((*pe.FileHeader)(nil)),
		"FormatError":      reflect.ValueOf((*pe.FormatError)(nil)),
		"ImportDirectory":  reflect.ValueOf((*pe.ImportDirectory)(nil)),
		"OptionalHeader32": reflect.ValueOf((*pe.OptionalHeader32)(nil)),
		"OptionalHeader64": reflect.ValueOf((*pe.OptionalHeader64)(nil)),
		"Reloc":            reflect.ValueOf((*pe.Reloc)(nil)),
		"Section":          reflect.ValueOf((*pe.Section)(nil)),
		"SectionHeader":    reflect.ValueOf((*pe.SectionHeader)(nil)),
		"SectionHeader32":  reflect.ValueOf((*pe.SectionHeader32)(nil)),
		"StringTable":      reflect.ValueOf((*pe.StringTable)(nil)),
		"Symbol":           reflect.ValueOf((*pe.Symbol)(nil)),
	}
}
