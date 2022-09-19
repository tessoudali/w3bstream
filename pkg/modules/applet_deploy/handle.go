package applet_deploy

import (
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/iotexproject/w3bstream/pkg/modules/vm"
)

func LoadHandlers(path string, events ...vm.EventHandler) ([]models.HandlerInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(f)
	if err != nil {
		return nil, err
	}

	hdls := make([]models.HandlerInfo, 0)

	for _, h := range events {
		if ev, ok := parsed.Methods[h.Event]; ok {
			i := models.HandlerInfo{
				Name:    ev.Name,
				Handler: h.Handler,
			}
			for _, ei := range ev.Inputs {
				i.Params.Inputs = append(i.Params.Inputs, models.HandlerParam{
					Name: ei.Name,
					Type: ei.Type.String(),
				})
			}
			for _, eo := range ev.Outputs {
				i.Params.Outputs = append(i.Params.Outputs, models.HandlerParam{
					Name: eo.Name,
					Type: eo.Type.String(),
				})
			}
			hdls = append(hdls, i)
		}
	}
	return hdls, nil
}
