package main

func (s *server) routes() {
	s.router.Get("/", s.handleIndex())
}
