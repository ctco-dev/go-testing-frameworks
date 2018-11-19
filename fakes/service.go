package fakes

//Service is an example of a component with injected dependency
type Service struct {
	rdr Reader
}

//NewService provides a new service
func NewService(rdr Reader) Service {
	return Service{rdr: rdr}
}

//Serve method gets string from underlying Reader and adds underscore char
func (s *Service) Serve() string {
	return s.rdr.Read() + "_"
}
