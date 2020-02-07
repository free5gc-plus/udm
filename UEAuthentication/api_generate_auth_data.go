/*
 * NudmUEAU
 *
 * UDM UE Authentication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package UEAuthentication

import (
	"github.com/gin-gonic/gin"
	"gofree5gc/lib/http_wrapper"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/src/udm/logger"
	"gofree5gc/src/udm/udm_handler"
	"gofree5gc/src/udm/udm_handler/udm_message"
	// "fmt"
)

// GenerateAuthData - Generate authentication data for the UE
func GenerateAuthData(c *gin.Context) {
	var authInfoReq models.AuthenticationInfoRequest
	err := c.ShouldBindJSON(&authInfoReq)
	if err != nil {
		logger.UeauLog.Errorln(err)
	}

	req := http_wrapper.NewRequest(c.Request, authInfoReq)
	req.Params["supiOrSuci"] = c.Param("supiOrSuci")

	handlerMsg := udm_message.NewHandlerMessage(udm_message.EventGenerateAuthData, req)
	udm_handler.SendMessage(handlerMsg)
	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
	return
}
