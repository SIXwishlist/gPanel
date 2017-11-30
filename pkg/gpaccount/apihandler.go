// Package gpaccount handles the logic of the gPanel account server
package gpaccount

import (
	"net/http"
	"strings"

	"github.com/Ennovar/gPanel/pkg/api/ip"
	logapi "github.com/Ennovar/gPanel/pkg/api/log"
	"github.com/Ennovar/gPanel/pkg/api/server"
	"github.com/Ennovar/gPanel/pkg/api/user"
)

func (con *Controller) apiHandler(res http.ResponseWriter, req *http.Request) (bool, bool) {
	path := req.URL.Path[1:]
	if len(path) == 0 {
		path = (con.Directory + "index.html")
	} else {
		path = (con.Directory + path)
	}

	splitUrl := strings.SplitN(path, "api", 2)
	suspectApi := strings.ToLower(splitUrl[len(splitUrl)-1])

	switch suspectApi {
	case "/user/auth":
		return true, user.Auth(res, req, con.APILogger, con.Directory)
	case "/user/register":
		return true, user.Register(res, req, con.APILogger, con.Directory)
	case "/user/logout":
		return true, user.Logout(res, req, con.APILogger, con.Directory)
	case "/server/status":
		return true, server.Status(res, req, con.APILogger, con.Public)
	case "/server/start":
		return true, server.Start(res, req, con.APILogger, con.Public)
	case "/server/shutdown":
		return true, server.Shutdown(res, req, con.APILogger, con.Public)
	case "/server/restart":
		return true, server.Restart(res, req, con.APILogger, con.Public)
	case "/server/maintenance":
		return true, server.Maintenance(res, req, con.APILogger, con.Public)
	case "/log/read":
		return true, logapi.Read(res, req, con.APILogger, con.Directory)
	case "/log/delete":
		return true, logapi.Truncate(res, req, con.APILogger, con.Directory)
	case "/ip/list":
		return true, ip.List(res, req, con.APILogger, con.Directory)
	case "/ip/filter":
		return true, ip.Filter(res, req, con.APILogger, con.Directory)
	case "/ip/unfilter":
		return true, ip.Unfilter(res, req, con.APILogger, con.Directory)
	default:
		return false, false
	}
}
