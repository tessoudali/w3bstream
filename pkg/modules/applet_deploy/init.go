package applet_deploy

import (
	"context"

	"github.com/iotexproject/Bumblebee/conf/log"
	"github.com/iotexproject/Bumblebee/kit/sqlx/builder"

	"github.com/iotexproject/w3bstream/cmd/srv-applet-mgr/global"
	"github.com/iotexproject/w3bstream/pkg/models"
	"github.com/iotexproject/w3bstream/pkg/modules/vm"
)

func StartAppletVMs(ctx context.Context) error {
	d := global.DBExecutorFromContext(ctx)
	l := log.Std()
	ma := &models.Applet{}
	md := &models.AppletDeploy{}
	mh := &models.Handler{}

	monitors := make([]struct {
		Location      string               `db:"f_location"`
		AppletID      string               `db:"f_applet_id"`
		AppletName    string               `db:"f_applet_name"`
		Version       string               `db:"f_version"`
		HandlerName   string               `db:"f_handler_name"`
		HandlerParams models.HandlerParams `db:"f_handler_params"`
	}, 0)
	err := d.QueryAndScan(
		builder.Select(builder.MultiWith(
			",",
			builder.Alias(ma.ColAppletID(), `f_applet_id`),
			builder.Alias(ma.ColName(), `f_applet_name`),
			builder.Alias(md.ColVersion(), `f_version`),
			builder.Alias(mh.ColName(), `f_handler_name`),
			builder.Alias(mh.ColParams(), `f_handler_params`),
			builder.Alias(md.ColLocation(), `f_location`),
		)).
			From(
				d.T(ma),
				builder.Join(d.T(md)).On(ma.ColAppletID().Eq(md.ColAppletID())),
				builder.Join(d.T(mh)).On(mh.ColAppletID().Eq(mh.ColAppletID())),
			),
		&monitors,
	)
	if err != nil {
		return err
	}

	m := make(map[string]map[string]*vm.Monitor)

	for _, v := range monitors {
		c, err := vm.Load(v.Location)
		if err != nil {
			l.Error(err)
			return err
		}

		if m[v.AppletID] == nil {
			m[v.AppletID] = make(map[string]*vm.Monitor)
		}
		monitor, ok := m[v.AppletID][v.Version]
		if !ok {
			monitor = vm.NewMonitorContext(c, v.AppletID, v.AppletName, v.Version)
			m[v.AppletID][v.Version] = monitor
		}
		monitor.Handlers[v.HandlerName] = models.HandlerInfo{
			Name:   v.HandlerName,
			Params: v.HandlerParams,
		}
	}

	for name := range m {
		for version := range m[name] {
			vm.Start(ctx, m[name][version])
		}
	}
	return nil
}
