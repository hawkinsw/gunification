package gunification

type GType interface {
	IsCanonical() bool
	IsConstructor() bool
	Equals(GType) bool
	Constructor() GType
	SubTerms() []GType
	Identifier() string
}

type gNode struct {
	self   GType
	parent *gNode
}

type GUnification struct {
	nodes []*gNode
	state bool
}

// Note -- may be self
func (node *gNode) findCanonical() *gNode {
	if node.parent != nil {
		return node.parent.findCanonical()
	}
	return node
}

func (guni *GUnification) findCanonical(node *gNode) *gNode {
	return node.findCanonical()
}

// Order matters -- if x and y are members of
// different canonical classes, then the x canonization
// will become a member of the y canonization.
func (guni *GUnification) equate(x, y *gNode) {
	xc := guni.findCanonical(x)
	yc := guni.findCanonical(y)

	if xc != yc {
		xc.parent = yc
	}
}

func (guni *GUnification) findOrInsert(x GType) *gNode {
	for _, t := range guni.nodes {
		if t.self.Equals(x) {
			return t.findCanonical()
		}
	}
	// This type is not already in the unification -- add it
	node := &gNode{self: x, parent: nil}
	guni.nodes = append(guni.nodes, node)
	return node
}

func (guni *GUnification) Unify(x, y GType) bool {
	xn := guni.findOrInsert(x)
	yn := guni.findOrInsert(y)

	if xn != yn {
		if !xn.self.IsCanonical() && !yn.self.IsCanonical() { // both are type variables.
			guni.equate(xn, yn)
		} else if !xn.self.IsCanonical() && yn.self.IsCanonical() { // xn is a type variable and yn is a type
			// make the type variable (xn) a "child" of the type (yn).
			guni.equate(xn, yn)
		} else if !yn.self.IsCanonical() && xn.self.IsCanonical() { // yn is a type variable and xn is a type
			// make the type variable (yn) a "child" of the type (xn).
			guni.equate(yn, xn)
		} else if xn.self.IsCanonical() && xn.self.IsCanonical() && // both types ...
			xn.self.Constructor() != nil && yn.self.Constructor() != nil && // have constructors ...
			xn.self.Constructor().Equals(yn.self.Constructor()) { // and their constructors are the same!
			guni.equate(xn, yn)
			if len(xn.self.Constructor().SubTerms()) != len((xn.self.Constructor().SubTerms())) {
				guni.state = false
				return false
			}
			for idx := range xn.self.Constructor().SubTerms() {
				xnt := xn.self.Constructor().SubTerms()[idx]
				ynt := yn.self.Constructor().SubTerms()[idx]
				guni.state = false
				return guni.Unify(xnt, ynt)
			}
		} else {
			return false
		}
	}
	guni.state = true
	return true
}

func (guni *GUnification) Repr() string {
	if !guni.state {
		return "Unification is impossible."
	}

	repr := ""
	for _, n := range guni.nodes {
		canonicalClass := n.findCanonical()
		repr += n.self.Identifier() + ": " + canonicalClass.self.Identifier() + "\n"
	}
	return repr
}

func (guni *GUnification) GetCanonical(t GType) GType {
	if !guni.state {
		panic(!guni.state)
	}
	node := guni.findOrInsert(t)
	return node.findCanonical().self
}
