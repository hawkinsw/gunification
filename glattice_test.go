package gunification

import (
	"log"
	"reflect"
	"testing"
)

var sta_paper_page40b = func(a GThing, b GThing) bool {
	if a == "f" {
		return b == "f" || b == "d" || b == "e" || b == "b" || b == "c" || b == "a"
	} else if a == "d" {
		return b == "d" || b == "b" || b == "c" || b == "a"
	} else if a == "e" {
		return b == "e" || b == "b" || b == "c" || b == "a"
	} else if a == "b" {
		return b == "b" || b == "a"
	} else if a == "c" {
		return b == "c" || b == "a"
	} else if a == "a" {
		return b == "a"
	}
	return false
}
var sta_paper_page40b_nodes = []GThing{"a", "b", "c", "d", "e", "f"}
var sta_paper_page40b_po = newGLattice(sta_paper_page40b, sta_paper_page40b_nodes)

func TestUpperBoundSimplePage40b(t *testing.T) {
	result := sta_paper_page40b_po.UpperBound([]GThing{"f"})
	expected := []GThing{"a", "b", "c", "d", "e", "f"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40b_po.UpperBound([]GThing{"d"})
	expected = []GThing{"a", "b", "c", "d"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40b_po.UpperBound([]GThing{"e"})
	expected = []GThing{"a", "b", "c", "e"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40b_po.UpperBound([]GThing{"b"})
	expected = []GThing{"a", "b"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40b_po.UpperBound([]GThing{"c"})
	expected = []GThing{"a", "c"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40b_po.UpperBound([]GThing{"a"})
	expected = []GThing{"a"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
}
func TestLowerBoundSimplePage40b(t *testing.T) {
	result := sta_paper_page40b_po.LowerBound([]GThing{"f"})
	expected := []GThing{"f"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40b_po.LowerBound([]GThing{"d"})
	expected = []GThing{"d", "f"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40b_po.LowerBound([]GThing{"e"})
	expected = []GThing{"e", "f"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40b_po.LowerBound([]GThing{"b"})
	expected = []GThing{"b", "d", "e", "f"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40b_po.LowerBound([]GThing{"c"})
	expected = []GThing{"c", "d", "e", "f"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40b_po.LowerBound([]GThing{"a"})
	expected = []GThing{"a", "b", "c", "d", "e", "f"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40b) failure: Got %v, expected %v!", result, expected)
	}
}

func TestLatticeConstructorFalse(t *testing.T) {
	if NewGLattice(sta_paper_page40b, sta_paper_page40b_nodes) != nil {
		log.Fatalf("TestLatticeConstructorFalse (page 40b) expected to fail to construct non-lattice")
	}
}

var sta_paper_page40a = func(a GThing, b GThing) bool {
	if a == "a" {
		return true
	} else if a == "b" {
		return b == "b" || b == "d" || b == "e"
	} else if a == "c" {
		return b == "c" || b == "d" || b == "e"
	} else if a == "d" {
		return b == "d" || b == "e"
	} else if a == "e" {
		return b == "e"
	}
	return false
}
var sta_paper_page40a_nodes = []GThing{"a", "b", "c", "d", "e"}
var sta_paper_page40a_po = newGLattice(sta_paper_page40a, sta_paper_page40a_nodes)

func TestUpperBoundSimplePage40a(t *testing.T) {
	result := sta_paper_page40a_po.UpperBound([]GThing{"a"})
	expected := []GThing{"a", "b", "c", "d", "e"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.UpperBound([]GThing{"b"})
	expected = []GThing{"b", "d", "e"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.UpperBound([]GThing{"c"})
	expected = []GThing{"c", "d", "e"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.UpperBound([]GThing{"d"})
	expected = []GThing{"d", "e"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.UpperBound([]GThing{"e"})
	expected = []GThing{"e"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
}

func TestLeastUpperBoundSimplePage40a(t *testing.T) {
	result := sta_paper_page40a_po.LeastUpperBound([]GThing{"a"})
	expected := "a"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LeastUpperBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.LeastUpperBound([]GThing{"b"})
	expected = "b"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LeastUpperBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.LeastUpperBound([]GThing{"c"})
	expected = "c"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LeastUpperBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.LeastUpperBound([]GThing{"d"})
	expected = "d"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LeastUpperBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.LeastUpperBound([]GThing{"e"})
	expected = "e"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LeastUpperBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
}

func TestUpperBoundExtendedPage40a(t *testing.T) {
	result := sta_paper_page40a_po.UpperBound([]GThing{"a", "e"})
	expected := []GThing{"e"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundExtended (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.UpperBound([]GThing{"b", "c"})
	expected = []GThing{"d", "e"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("UpperBoundExtended (page 40a) failure: Got %v, expected %v!", result, expected)
	}
}
func TestLeastUpperBoundExtendedPage40a(t *testing.T) {
	result := sta_paper_page40a_po.LeastUpperBound([]GThing{"a", "e"})
	expected := "e"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LeastUpperBoundExtended (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.LeastUpperBound([]GThing{"b", "c"})
	expected = "d"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LeastUpperBoundExtended (page 40a) failure: Got %v, expected %v!", result, expected)
	}
}

func TestLowerBoundSimplePage40a(t *testing.T) {
	result := sta_paper_page40a_po.LowerBound([]GThing{"e"})
	expected := []GThing{"a", "b", "c", "d", "e"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.LowerBound([]GThing{"d"})
	expected = []GThing{"a", "b", "c", "d"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.LowerBound([]GThing{"b"})
	expected = []GThing{"a", "b"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.LowerBound([]GThing{"c"})
	expected = []GThing{"a", "c"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.LowerBound([]GThing{"a"})
	expected = []GThing{"a"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
}

func TestGreatestLowerBoundSimplePage40a(t *testing.T) {
	result := sta_paper_page40a_po.GreatestLowerBound([]GThing{"e"})
	expected := "e"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.GreatestLowerBound([]GThing{"d"})
	expected = "d"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.GreatestLowerBound([]GThing{"b"})
	expected = "b"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.GreatestLowerBound([]GThing{"c"})
	expected = "c"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.GreatestLowerBound([]GThing{"a"})
	expected = "a"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundSimple (page 40a) failure: Got %v, expected %v!", result, expected)
	}
}

func TestLowerBoundExtendedPage40a(t *testing.T) {
	result := sta_paper_page40a_po.LowerBound([]GThing{"b", "c"})
	expected := []GThing{"a"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundExtended (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.LowerBound([]GThing{"d"})
	expected = []GThing{"a", "b", "c", "d"}
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LowerBoundExtended (page 40a) failure: Got %v, expected %v!", result, expected)
	}
}
func TestGreatestLowerBoundExtendedPage40a(t *testing.T) {
	result := sta_paper_page40a_po.GreatestLowerBound([]GThing{"a", "e"})
	expected := "a"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LeastLowerBoundExtended (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.GreatestLowerBound([]GThing{"b", "c"})
	expected = "a"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LeastLowerBoundExtended (page 40a) failure: Got %v, expected %v!", result, expected)
	}
	result = sta_paper_page40a_po.GreatestLowerBound([]GThing{"b", "d"})
	expected = "b"
	if !reflect.DeepEqual(result, expected) {
		log.Fatalf("LeastLowerBoundExtended (page 40a) failure: Got %v, expected %v!", result, expected)
	}
}
func TestLatticeConstructorTrue(t *testing.T) {
	if NewGLattice(sta_paper_page40a, sta_paper_page40a_nodes) == nil {
		log.Fatalf("TestLatticeConstructorTrue (page 40a) expected to fail to construct non-lattice")
	}
}
