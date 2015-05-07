// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/bigbluebutton/api-labs/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSNamespace("/meetings",
				beego.NSInclude(
					&controllers.MeetingController{},
				),
			),
			// beego.NSNamespace("/recordings",
			// 	beego.NSInclude(
			// 		&controllers.RecordingController{},
			// 	),
			// ),
		),
	)
	beego.AddNamespace(ns)
}