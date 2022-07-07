package retry

import (
	"time"

	"github.com/go-courier/envconf"
	"github.com/sirupsen/logrus"
)

type Retry struct {
	Repeats  int
	Interval envconf.Duration
}

func (r *Retry) SetDefaults() {
	if r.Repeats == 0 {
		r.Repeats = 3
	}
	if r.Interval == 0 {
		r.Interval = envconf.Duration(10 * time.Second)
	}
}

func (r Retry) Do(exec func() error) (err error) {
	if r.Repeats <= 0 {
		return exec()
	}
	for i := 0; i < r.Repeats; i++ {
		if err = exec(); err != nil {
			logrus.Warningf("retry in seconds [%d]", r.Interval)
			time.Sleep(time.Duration(r.Interval))
			continue
		}
		break
	}
	return
}

func Do(repeats int, interval time.Duration, exec func() error) error {
	r := &Retry{
		Repeats:  repeats,
		Interval: envconf.Duration(interval),
	}
	return r.Do(exec)
}
