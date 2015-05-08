// @APIVersion 1.0.0
// @Title BigBlueButton Go API
// @Description Test API in Go for BigBlueButton
// @Contact invalid@bigbluebutton.org
// @TermsOfServiceUrl http://bigbluebutton.org/
// @License ?
// @LicenseUrl ?
package routers

import (
	"github.com/astaxie/beego"
	"github.com/bigbluebutton/go-api/controllers"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/meetings",
			beego.NSInclude(
				&controllers.MeetingController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
