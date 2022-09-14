package http

import (
	"io"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/iotexproject/Bumblebee/conf/log"

	me "github.com/iotexproject/w3bstream/pkg/modules/event"
)

const strLenLimit = 50

// Run run http server
func Run(events chan<- me.Event) {
	r := gin.Default()

	r.POST("/:project/*applet", func(c *gin.Context) {
		logger := log.Std()
		project := strings.TrimSpace(c.Param("project"))
		applet := strings.TrimSpace(c.Param("applet"))
		publisher := strings.TrimSpace(c.GetHeader("publisher"))
		if !check(project, applet, publisher) {
			c.Status(http.StatusBadRequest)
			return
		}
		// TODO data size check
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			logger.Error(err)
			c.Status(http.StatusBadRequest)
			return
		}

		events <- &event{
			project:   project,
			applet:    applet,
			publisher: publisher,
			data:      data,
		}
		c.Status(http.StatusOK)
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
