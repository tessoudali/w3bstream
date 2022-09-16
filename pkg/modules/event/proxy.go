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
			var data []byte
			result, ok := e.(EventResult)
			defer func() {
				if ok {
					result.ResultChan() <- Result{success, data}
				}
			}()

			if !p.f.filter(e) {
				return
			}
			res, err := p.d.dispatch(e)
			if err != nil {
				p.logger.Error(err)
				return
			}
			success = true
			data = res
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
