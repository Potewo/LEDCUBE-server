package server

type Server interface {
	Generate(frame int) [8][8]byte
}

func Send(s Server, l int) {
	for i := 0; i < l; i++ {
		matrix := s.Generate(l)
		_ = matrix
	}
}
