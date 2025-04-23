package eth

type Mac struct {
	Data []byte
}

func (m *Mac) GetAddr() [8]byte {
	for len(m.Data) < 8 {
		m.Data = append(m.Data, 0)
	}

	mac := [8]byte{}
	copy(mac[:], m.Data)

	return mac
}
