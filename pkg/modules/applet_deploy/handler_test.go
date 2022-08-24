package applet_deploy_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/global"
	"github.com/iotexproject/w3bstream/pkg/modules/applet_deploy"
	"github.com/iotexproject/w3bstream/pkg/modules/vm"
	. "github.com/onsi/gomega"
)

func TestLoadHandlers(t *testing.T) {
	var (
		ctx  = global.WithConfContext(context.Background())
		conf = global.ConfFromContext(ctx)
	)
	hdls, err := applet_deploy.LoadHandlers(
		filepath.Join(
			conf.ResourceRoot,
			"88418cfc-54f1-455f-9357-a11826b5b5e0",
			"0.0.6", "Simple", "abis", "Simple.json",
		),
		vm.EventHandler{
			Event:   "add",
			Handler: "",
		},
	)
	NewWithT(t).Expect(err).To(BeNil())
	t.Log(hdls)
}
