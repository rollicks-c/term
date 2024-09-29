package datetime

import "time"

func (p Parser) getNow() time.Time {
	if p.now == nil {
		return time.Now()
	}
	return *p.now
}
