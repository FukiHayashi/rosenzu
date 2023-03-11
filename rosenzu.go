package rosenzuapi

import (
	"context"
	"log"
	"rosenzu/database/cast"
	"rosenzu/database/repository"
	rosenzu "rosenzu/gen/rosenzu"
)

// rosenzu service example implementation.
// The example methods log the requests and return zero values.
type rosenzusrvc struct {
	logger *log.Logger
}

// NewRosenzu returns the rosenzu service implementation.
func NewRosenzu(logger *log.Logger) rosenzu.Service {
	return &rosenzusrvc{logger}
}

// Find implements find.
func (s *rosenzusrvc) Find(ctx context.Context, p *rosenzu.FindPayload) (res *rosenzu.Line, err error) {
	var line = repository.FindLine(p.Name)
	res = cast.CastedLine(line)
	s.logger.Print("rosenzu.find")
	return
}
