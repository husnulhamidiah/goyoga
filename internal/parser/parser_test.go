package parser

import "testing"

func TestParseExportedVariable(t *testing.T) {
	api, err := Parse("YG_EXPORT extern const YGValue YGValueAuto;")
	if err != nil {
		t.Fatal(err)
	}
	if len(api.Variables) != 1 {
		t.Fatalf("variables = %d, want 1", len(api.Variables))
	}
	if api.Variables[0].Name != "YGValueAuto" {
		t.Fatalf("name = %q", api.Variables[0].Name)
	}
}

func TestParseYGEnumDecl(t *testing.T) {
	api, err := Parse("YG_ENUM_DECL(YGAlign, YGAlignAuto, YGAlignCenter)")
	if err != nil {
		t.Fatal(err)
	}
	if len(api.Enums) != 1 {
		t.Fatalf("enums = %d, want 1", len(api.Enums))
	}
	if len(api.Functions) != 1 || api.Functions[0].Name != "YGAlignToString" {
		t.Fatalf("functions = %#v", api.Functions)
	}
}
