package wasmresource

import (
	"context"
	"fmt"
	confid "github.com/machinefi/w3bstream/pkg/depends/conf/id"
	"github.com/machinefi/w3bstream/pkg/models"
	"github.com/machinefi/w3bstream/pkg/types"
	"github.com/pkg/errors"
)

func FetchOrCreateResourceByMd5(ctx context.Context, md5 string) (m *models.WasmResource, err error) {
	d := types.MustDBExecutorFromContext(ctx)
	l := types.MustLoggerFromContext(ctx)
	idg := confid.MustSFIDGeneratorFromContext(ctx)

	_, l = l.Start(ctx, "FetchOrCreateResourceByMd5")
	defer l.End()

	m = &models.WasmResource{WasmResourceInfo: models.WasmResourceInfo{Md5: md5}}

	if err = m.FetchByMd5(d); err != nil {
		l.Error(errors.Wrap(err, fmt.Sprintf("fetch wasm resource by md5 - %s, maybe it doesnt exist.", md5)))
		// TODO check err type
		m.WasmResourceID = idg.MustGenSFID()
		m.WasmResourceInfo.Path = ""
		m.WasmResourceInfo.Md5 = md5
		m.WasmResourceInfo.RefCnt = 0
		if err = m.Create(d); err != nil {
			l.Error(errors.Wrap(err, "create wasm resource db failed"))
			return nil, err
		}
		l.Info("wasm resource created")
	}
	l.Info("get wasm resource from db")
	return m, err
}
