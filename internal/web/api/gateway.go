package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/internal/core"
)

func gatewayProperties(ctx *gin.Context) {
	device := core.Devices.Load(ctx.GetString("id"))
	if device == nil {
		replyFail(ctx, "找不到设备变量")
		return
	}
	replyOk(ctx, device.Properties)
}
