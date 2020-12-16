package core

// Put all View Method here

func (m *MinorVersion) View() interface{} {
	b := m.Bytes
	return int(U2(b))
}

func (af *AccessFlag) View() interface{} {
	b := af.Bytes
	f := U2(b)
	return getFlag(int(f))
}

func (m *Magic) View() interface{} {
	return m.Hex
}

func (m *MajorVersion) View() interface{} {
	b := m.Bytes
	return int(U2(b))
}
