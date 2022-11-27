package gunification

import (
	"testing"
)

type variable struct {
	identifier string
}

func (v *variable) IsCanonical() bool {
	return false
}
func (v *variable) IsConstructor() bool {
	return false
}
func (v *variable) Identifier() string {
	return v.identifier
}
func (v *variable) Equals(o GType) bool {
	return v.identifier == o.Identifier()
}
func (v *variable) SubTerms() []GType {
	return nil
}
func (v *variable) Constructor() GType {
	return nil
}

type variableType struct {
	identifier string
}

func (v *variableType) IsCanonical() bool {
	return true
}
func (v *variableType) IsConstructor() bool {
	return false
}
func (v *variableType) Identifier() string {
	return v.identifier
}
func (v *variableType) Equals(o GType) bool {
	return v.identifier == o.Identifier()
}
func (v *variableType) SubTerms() []GType {
	return nil
}
func (v *variableType) Constructor() GType {
	return nil
}

var a_variable = &variable{identifier: "A"}
var b_variable = &variable{identifier: "B"}
var c_variable = &variable{identifier: "C"}
var d_variable = &variable{identifier: "D"}
var x_variable = &variable{identifier: "X"}
var y_variable = &variable{identifier: "Y"}
var z_variable = &variable{identifier: "Z"}

var int_type = &variableType{identifier: "int"}
var double_type = &variableType{identifier: "double"}

func TestSimpleUnify(t *testing.T) {

	/*
		int X{};
		...
		Y := X
	*/
	gunification := GUnification{}

	if !gunification.Unify(x_variable, int_type) {
		t.Fatalf("Oops: Could not unify X variable with an int type.")
	}
	if !gunification.Unify(x_variable, y_variable) {
		t.Fatalf("Oops: Could not unify X variable with an int type.")
	}
	t.Log(gunification.Repr())
}

func TestSimpleNoUnify(t *testing.T) {

	/*
		int X{};
		double Z{};
		...
		Y := X
		Z = Y
	*/
	gunification := GUnification{}

	if !gunification.Unify(x_variable, int_type) {
		t.Fatalf("Oops: Could not unify X variable with an int type.")
	}
	if !gunification.Unify(z_variable, double_type) {
		t.Fatalf("Oops: Could not unify Z variable with a double type.")
	}
	if !gunification.Unify(x_variable, y_variable) {
		t.Fatalf("Oops: Could not unify X variable with Y variable.")
	}
	if gunification.Unify(x_variable, z_variable) {
		t.Log(gunification.Repr())
		t.Fatalf("Oops: Unified a double-typed variable with an int-typed variable.")
	}
	t.Log(gunification.Repr())
}

func TestTransitiveUnify(t *testing.T) {
	/*
		int A{};
		int C{};
		...
		B = A
		C = B
		D := B
	*/
	gunification := GUnification{}

	if !gunification.Unify(a_variable, int_type) {
		t.Fatalf("Oops: Could not unify A variable with an int type.")
	}
	if !gunification.Unify(c_variable, int_type) {
		t.Fatalf("Oops: Could not unify C variable with an int type.")
	}
	if !gunification.Unify(b_variable, a_variable) {
		t.Fatalf("Oops: Could not unify B variable with A variable.")
	}
	if !gunification.Unify(c_variable, b_variable) {
		t.Fatalf("Oops: Could not unify B variable with A variable.")
	}
	if !gunification.Unify(d_variable, b_variable) {
		t.Fatalf("Oops: Unified a double-typed variable with an int-typed variable.")
	}
	if !gunification.GetCanonical(d_variable).Equals(int_type) {
		t.Fatalf("Oops: D variable's type was incorrectly derived as %v instead of %v", gunification.GetCanonical(d_variable).Identifier(), int_type.Identifier())
	}
	t.Log(gunification.Repr())
}
