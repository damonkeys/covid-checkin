package languagehelper

import (
	"testing"
)

func TestEmptyHeader(t *testing.T) {
	emptyHeader := ""
	expectEN := Retrieve(emptyHeader)
	if expectEN != "en" {
		t.Errorf("Expected en, but got %s", expectEN)
	}
}

func TestRetrieveLangSimpleWithEN(t *testing.T) {
	simpleHeader := "en-GB;q=0.9"
	expectEN := Retrieve(simpleHeader)
	if expectEN != "en" {
		t.Errorf("Expected en, but got %s", expectEN)
	}
}

func TestRetrieveLangSimpleWithDE(t *testing.T) {
	simpleHeader := "de-GB;q=0.9"
	expectDE := Retrieve(simpleHeader)
	if expectDE != "de" {
		t.Errorf("Expected DE, but got %s", expectDE)
	}
}

func TestRetriveDefaultToEN(t *testing.T) {
	simpleHeader := "xx-GB;q=0.9"
	expectEN := Retrieve(simpleHeader)
	if expectEN != "en" {
		t.Errorf("Expected en, but got %s", expectEN)
	}
}

func TestComplexWithDElast(t *testing.T) {
	complex := "en-GB,en-US;q=0.9,en;q=0.8,de;q=0.74"
	expectEN := Retrieve(complex)
	if expectEN != "en" {
		t.Errorf("Expected en, but got %s", expectEN)
	}
}

func TestComplexWithDEfirst(t *testing.T) {
	complex := "de;q=0.95,en-GB,en-US;q=0.9,en;q=0.8"
	expectDE := Retrieve(complex)
	if expectDE != "de" {
		t.Errorf("Expected de, but got %s", expectDE)
	}
}

func TestComplexWithDElaterButBeforeEn(t *testing.T) {
	complex := "xx-GB,yy-US;q=0.9,de;q=0.85,en;q=0.8"
	expectDE := Retrieve(complex)
	if expectDE != "de" {
		t.Errorf("Expected de, but got %s", expectDE)
	}
}

func TestComplexOnlyDEAndOthers(t *testing.T) {
	complex := "xx-GB,yy-US;q=0.9,de;q=0.85,zz;q=0.8"
	expectDE := Retrieve(complex)
	if expectDE != "de" {
		t.Errorf("Expected de, but got %s", expectDE)
	}
}

func TestComplexOnlyENAndOthers(t *testing.T) {
	complex := "xx-GB,yy-US;q=0.9,en;q=0.85,zz;q=0.8"
	expectEN := Retrieve(complex)
	if expectEN != "en" {
		t.Errorf("Expected en, but got %s", expectEN)
	}
}
