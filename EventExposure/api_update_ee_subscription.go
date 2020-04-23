/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package EventExposure

import (
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi/models"
	"free5gc/src/udm/logger"
	"free5gc/src/udm/udm_handler"
	"free5gc/src/udm/udm_handler/udm_message"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateEeSubscription - Patch
func UpdateEeSubscription(c *gin.Context) {

	var eeSubscriptionReq models.PatchItem

	err := c.ShouldBindJSON(&eeSubscriptionReq)
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.EeLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, eeSubscriptionReq)
	req.Params["ueIdentity"] = c.Params.ByName("ueIdentity")
	req.Params["subscriptionID"] = c.Params.ByName("subscriptionId")

	handlerMsg := udm_message.NewHandlerMessage(udm_message.EventUpdateEeSubscription, req)
	udm_handler.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
