package halflife

import "minmax.uk/game-of-life/pkg/engine"

var _ engine.Engine = (*Universe)(nil)

func (m *Universe) Name() string {
	return "HalflifeEngine"
}

func (m *Universe) Iterate() {
	m.root = m.root.Expand()
	m.root = m.root.Iterate()
	m.size >>= 1
}

func (m *Universe) String() string {
	return m.BoardString()
}
