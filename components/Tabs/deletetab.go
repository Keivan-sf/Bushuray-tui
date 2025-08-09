package tabs

func (m *Model) DeleteTab(tid int) {
	if tid == 0 {
		return
	}
	m.Children = append(m.Children[:tid], m.Children[tid+1:]...)
	if m.ActiveTap >= len(m.Children) && m.ActiveTap > 0 {
		m.ActiveTap--
	}
}
