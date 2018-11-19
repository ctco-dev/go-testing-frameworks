package fakes

type service struct {
	rdr reader
}

func NewService(rdr reader) service {
	return service{ rdr: rdr}
}

func (s *service) Serve() string {
	return s.rdr.Read()
}
