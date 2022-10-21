package vm

import (
	"context"
	"os"

	"github.com/iotexproject/w3bstream/pkg/modules/vm/common"
	"github.com/iotexproject/w3bstream/pkg/modules/vm/wasmtime"
	"github.com/iotexproject/w3bstream/pkg/types"
	"github.com/iotexproject/w3bstream/pkg/types/wasm"
)

func NewInstance(ctx context.Context, path string, opts ...common.InstanceOptionSetter) (types.SFID, error) {
	l := types.MustLoggerFromContext(ctx)

	_, l = l.Start(ctx, "NewInstanceWithID")
	defer l.End()

	code, err := os.ReadFile(path)
	if err != nil {
		l.Error(err)
		return 0, err
	}
	i, err := newInstanceByCode(ctx, code, opts...)
	if err != nil {
		l.Error(err)
		return 0, err
	}
	return AddInstance(ctx, i), nil
}

func NewInstanceWithID(ctx context.Context, path string, by types.SFID, opts ...common.InstanceOptionSetter) error {
	l := types.MustLoggerFromContext(ctx)

	_, l = l.Start(ctx, "NewInstanceWithID")
	defer l.End()

	code, err := os.ReadFile(path)
	if err != nil {
		l.Error(err)
		return err
	}
	i, err := newInstanceByCode(ctx, code, opts...)
	if err != nil {
		l.Error(err)
		return err
	}

	AddInstanceByID(ctx, by, i)
	return nil
}

func newInstanceByCode(ctx context.Context, code []byte, opts ...common.InstanceOptionSetter) (wasm.Instance, error) {
	return wasmtime.NewInstanceByCode(ctx, code, opts...)
}
