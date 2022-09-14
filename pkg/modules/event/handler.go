package event

import "github.com/iotexproject/Bumblebee/conf/log"

type Handler struct {
	events <-chan Event
	logger log.Logger
}

func (h *Handler) Run() {
	for e := range h.events {
		if err := h.dispatch(e); err != nil {
			h.logger.Error(err)
		}
	}
}

func (h *Handler) dispatch(e Event) error {
	return nil
}
