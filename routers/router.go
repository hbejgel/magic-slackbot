// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/hbejgel/magic-slackbot/controllers"
)

func init() {
	beego.Get("/meta/healthcheck", func(ctx *context.Context) {
		ctx.ResponseWriter.WriteHeader(200)
	})
	beego.Include(&controllers.MagicController{})
}
