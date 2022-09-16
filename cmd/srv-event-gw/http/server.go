package http

import (
	"io"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/iotexproject/Bumblebee/conf/log"

	"github.com/iotexproject/w3bstream/pkg/depends/unit"
	me "github.com/iotexproject/w3bstream/pkg/modules/event"
)

const (
	strLenLimit   = 50
	dataSizeLimit = 2 * unit.KiB
)

// Run run http server
func Run(events chan<- me.Event, logger log.Logger) {
	r := gin.Default()

	r.POST("/:project/*applet", func(c *gin.Context) {
		projectID := strings.TrimSpace(c.Param("project"))
		appletID := strings.TrimSpace(c.Param("applet"))
		publisherID := strings.TrimSpace(c.GetHeader("publisher"))
		if !check(projectID, appletID, publisherID) {
			c.Status(http.StatusBadRequest)
			return
		}
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			logger.Error(err)
			c.Status(http.StatusBadRequest)
			return
		}
		if len(data) > dataSizeLimit {
			c.Status(http.StatusBadRequest)
			return
		}

		res := make(chan me.Result)
		events <- &event{
			projectID:   projectID,
			appletID:    appletID,
			publisherID: publisherID,
			data:        data,
			result:      res,
		}
		// TODO timeout
		result := <-res
		s := http.StatusOK
		if !result.Success {
			s = http.StatusInternalServerError
		}
		c.Data(s, "application/octet-stream", result.Data)
	})
	r.Run()
}

func check(project, applet, publisher string) bool {
	if l := utf8.RuneCountInString(project); l <= 0 || l > strLenLimit {
		return false
	}
	if l := utf8.RuneCountInString(publisher); l <= 0 || l > strLenLimit {
		return false
	}
	return utf8.RuneCountInString(applet) <= strLenLimit
}
