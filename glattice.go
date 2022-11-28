package gunification

type GLattice struct {
	relation GRelation
	nodes    []GThing
}

func calculateLeastUpperBound(upperBounds []GThing, relation GRelation) GThing {
	for _, candidate := range upperBounds {
		isCandidateLub := true
		for _, compare := range upperBounds {
			if !relation(candidate, compare) {
				// Then candidate is not less than compare so it cannot be the LUB -- move on!
				isCandidateLub = false
				break
			}
		}
		if isCandidateLub {
			return candidate
		}
	}
	return nil
}

func calculateGreatestLowerBound(lowerBounds []GThing, relation GRelation) GThing {
	for _, candidate := range lowerBounds {
		isCandidateLub := true
		for _, compare := range lowerBounds {
			if !relation(compare, candidate) {
				// Then candidate is not less than compare so it cannot be the LUB -- move on!
				isCandidateLub = false
				break
			}
		}
		if isCandidateLub {
			return candidate
		}
	}
	return nil
}

func NewGLattice(relation GRelation, nodes []GThing) *GLattice {
	lattice := &GLattice{relation: relation, nodes: nodes}
	if lattice.isLattice() {
		return lattice
	}
	return nil
}

// A special internal ctor for testing purposes.
func newGLattice(relation GRelation, nodes []GThing) *GLattice {
	return &GLattice{relation: relation, nodes: nodes}
}

func (lattice *GLattice) isLattice() bool {
	for _, pair := range lattice.nodePairs() {
		ub := lattice.UpperBound(pair)
		if calculateLeastUpperBound(ub, lattice.relation) == nil {
			return false
		}
		lb := lattice.LowerBound(pair)
		if calculateGreatestLowerBound(lb, lattice.relation) == nil {
			return false
		}
	}
	return true
}

func (lattice *GLattice) nodePairs() (result [][]GThing) {
	result = make([][]GThing, 0)
	for o := range lattice.nodes {
		for i := range lattice.nodes {
			n := make([]GThing, 2)
			n[0] = lattice.nodes[o]
			n[1] = lattice.nodes[i]
			result = append(result, n)
		}
	}
	return
}

func (lattice *GLattice) LeastUpperBound(nodes []GThing) GThing {
	ub := lattice.UpperBound(nodes)
	return calculateLeastUpperBound(ub, lattice.relation)
}

func (lattice *GLattice) UpperBound(nodes []GThing) []GThing {
	result := make([]GThing, 0)
	for _, candidate := range lattice.nodes {
		failure := false
		for _, n := range nodes {
			if !lattice.relation(n, candidate) {
				failure = true
				break
			}
		}
		if !failure {
			result = append(result, candidate)
		}
	}
	return result
}

func (lattice *GLattice) GreatestLowerBound(nodes []GThing) GThing {
	lb := lattice.LowerBound(nodes)
	return calculateGreatestLowerBound(lb, lattice.relation)
}
func (lattice *GLattice) LowerBound(nodes []GThing) []GThing {
	result := make([]GThing, 0)
	for _, candidate := range lattice.nodes {
		failure := false
		for _, n := range nodes {
			if !lattice.relation(candidate, n) {
				failure = true
				break
			}
		}
		if !failure {
			result = append(result, candidate)
		}
	}
	return result
}
