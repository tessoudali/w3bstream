package vm

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iotexproject/Bumblebee/conf/log"
	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/global"
	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/pkg/errors"
	w "github.com/wasmerio/wasmer-go/wasmer"
	"gopkg.in/yaml.v2"
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
						fmt.Println(string(data[args[0].I32():]))
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
	broker := global.MqttFromContext(ctx)

	logger.Info("%s subscribe started", topic)
	go func() {
		cli, err := broker.Client(m.AppletID)
		if err != nil {
			panic(err)
		}

		for {
			err = cli.WithTopic(topic).Subscribe(
				func(c mqtt.Client, msg mqtt.Message) {
					// TODO defer log event
					inputs := [2]int{}
					payload := msg.Payload()
					err := json.Unmarshal(payload, &inputs)
					if err != nil {
						logger.Error(err)
						return
					}
					sum, err := m.instance.ExecuteFunction("add", inputs[0], inputs[1])
					if err != nil {
						logger.Error(err)
						return
					}
					logger.Info(
						"topic: %s payload: %s result: %v",
						topic, payload, sum,
					)
				},
			)
		}
	}()
}
