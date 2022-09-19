package vm

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iotexproject/Bumblebee/conf/log"
	"github.com/pkg/errors"
	w "github.com/wasmerio/wasmer-go/wasmer"
	"gopkg.in/yaml.v2"

	"github.com/iotexproject/w3bstream/pkg/types"

	"github.com/iotexproject/w3bstream/pkg/models"
)

type VM struct {
	*Config
	*Wasm
}

func Load(root string) (*VM, error) {
	var wasm *Wasm

	configFilename := filepath.Join(root, "applet.yaml")
	_, err := os.Stat(configFilename)
	if err != nil {
		return nil, errors.Wrap(ErrInvalidConfigPath, err.Error())
	}
	f, err := os.Open(configFilename)
	if err != nil {
		return nil, err
	}
	dec := yaml.NewDecoder(f)
	cfg := &Config{}
	if err = dec.Decode(cfg); err != nil {
		return nil, err
	}

	imports := []WasmImport{
		{
			Namespace: "env",
			Functions: []WasmImportFunc{
				{
					Name: "log",
					InputTypes: []w.ValueKind{
						w.I32,
					},
					OutputTypes: []w.ValueKind{},
					NativeFunc: func(args []w.Value) ([]w.Value, error) {
						data, e := wasm.GetMemory("memory")
						if e != nil {
							return nil, e
						}
						offset := args[0].I32()
						fmt.Println(string(data[offset:]))
						return []w.Value{}, nil
					},
				}, {
					Name: "abort",
					InputTypes: []w.ValueKind{
						w.I32,
						w.I32,
						w.I32,
						w.I32,
					},
					OutputTypes: []w.ValueKind{},
					NativeFunc: func(args []w.Value) ([]w.Value, error) {
						return []w.Value{}, nil
					},
				}, {
					Name:        "run",
					InputTypes:  []w.ValueKind{w.I32},
					OutputTypes: []w.ValueKind{},
					NativeFunc: func(args []w.Value) ([]w.Value, error) {
						return []w.Value{}, nil
					},
				}, {
					Name:        "alloc",
					InputTypes:  []w.ValueKind{w.I32},
					OutputTypes: []w.ValueKind{w.I32},
					NativeFunc: func(args []w.Value) ([]w.Value, error) {
						data, e := wasm.GetMemory("memory")
						if e != nil {
							return nil, err
						}
						return []w.Value{
							w.NewValue(data[args[0].I32():], w.AnyRef),
						}, nil
					},
				},
			},
		},
	}

	wasmFilename := filepath.Join(root, cfg.DataSources[0].File)
	code, err := ioutil.ReadFile(wasmFilename)
	if err != nil {
		return nil, err
	}

	wasm, err = NewWasm(code, imports)
	if err != nil {
		return nil, err
	}

	return &VM{cfg, wasm}, nil
}

var (
	ErrInvalidConfigPath = errors.New("invalid config path")
)

type Monitor struct {
	instance   *VM
	AppletID   string `db:"f_applet_id"`
	AppletName string `db:"f_applet_name"`
	Version    string `db:"f_version"`
	Handlers   map[string]models.HandlerInfo
}

func NewMonitorContext(c *VM, appletID, appletName, version string, hdls ...models.HandlerInfo) *Monitor {
	m := &Monitor{
		instance:   c,
		AppletID:   appletID,
		AppletName: appletName,
		Version:    version,
		Handlers:   make(map[string]models.HandlerInfo),
	}
	for i := range hdls {
		m.Handlers[hdls[i].Name] = hdls[i]
	}
	return m
}

func Start(ctx context.Context, m *Monitor) {
	topic := fmt.Sprintf("%s@%s", m.AppletName, m.Version)
	logger := log.Std()
	broker := types.MustMqttBrokerFromContext(ctx)

	logger.Info("%s subscribe started", topic)
	go func() {
		cli, err := broker.Client(m.AppletID)
		if err != nil {
			panic(err)
		}

		for {
			err = cli.WithTopic(topic).Subscribe(
				func(c mqtt.Client, msg mqtt.Message) {
					// TODO: defer log event
					payload := msg.Payload() // route data
					// instance = dispatch(route)
					// instance.exec(data)
					// TODO get wasm addr(size)
					// TODO pl -> addr
					sum, err := m.instance.ExecuteFunction("run", payload)
					if err != nil {
						logger.Error(err)
						return
					}
					// msg seq id
					// result.Method DB output
					logger.Info(
						"topic: %s payload: %s result: %v",
						topic, payload, sum,
					)
				},
			)
		}
	}()
}
