package eth

type MacAddr struct {
	Data []byte
}

func (m *MacAddr) GetAddr() [8]byte {
	for len(m.Data) < 8 {
		m.Data = append(m.Data, 0)
	}

	mac := [8]byte{}
	copy(mac[:], m.Data)

	return mac
}
