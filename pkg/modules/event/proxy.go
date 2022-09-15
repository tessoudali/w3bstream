package event

import "github.com/iotexproject/Bumblebee/conf/log"

type Proxifier struct {
	d      *dispatcher
	f      *filtration
	logger log.Logger
}

func (p *Proxifier) Proxy(events <-chan Event) {
	for e := range events {
		func() {
			success := false
			result, ok := e.(EventResult)
			defer func() {
				if ok {
					result.ResultChan() <- success
				}
			}()

			if !p.f.filter(e) {
				return
			}
			if err := p.d.dispatch(e); err != nil {
				p.logger.Error(err)
				return
			}
			success = true
		}()
	}
}

func NewProxifier(logger log.Logger) *Proxifier {
	return &Proxifier{
		d:      &dispatcher{},
		f:      &filtration{},
		logger: logger,
	}
}
