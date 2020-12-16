package core

// Put all View Method here
// Make all obj in class is view able
func (m *MinorVersion) View() interface{} {
	b := m.Bytes
	return int(U2(b))
}

func (m *MinorVersion) String() string {
	b := m.Bytes
	return string(U2(b))
}

func (af *AccessFlag) View() interface{} {
	b := af.Bytes
	f := U2(b)
	return getFlag(f)
}

func (af *AccessFlag) String() string {
	b := af.Bytes
	f := U2(b)
	return getFlag(f)
}

func (m *Magic) View() interface{} {
	return m.Hex
}

func (m *MajorVersion) View() interface{} {
	b := m.Bytes
	return int(U2(b))
}

func (tc *ThisClass) View() interface{} {
	b := tc.Bytes
	return int(U2(b))
}
