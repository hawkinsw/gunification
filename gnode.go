package gunification

type gNode struct {
	self   GType
	parent *gNode
}

// Note -- may be self
func (node *gNode) findCanonical() *gNode {
	if node.parent != nil {
		return node.parent.findCanonical()
	}
	return node
}
