/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eventexposure_test

import (
	"context"
	"fmt"
	"free5gc/lib/http2_util"
	Nudm_EE_Client "free5gc/lib/openapi/Nudm_EventExposure"
	"free5gc/lib/openapi/models"
	"free5gc/lib/path_util"
	udm_context "free5gc/src/udm/context"
	Nudm_EE_Server "free5gc/src/udm/eventexposure"
	"free5gc/src/udm/handler"
	"free5gc/src/udm/logger"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// CreateEeSubscription - Subscribe
func TestCreateEeSubscription(t *testing.T) {
	udm_context.TestInit()
	go handler.Handle()
	go func() { // udm server
		router := gin.Default()
		Nudm_EE_Server.AddService(router)

		udmLogPath := path_util.Gofree5gcPath("free5gc/udmsslkey.log")
		udmPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.pem")
		udmKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.key")
		server, err := http2_util.NewServer(":29503", udmLogPath, router)
		if err == nil && server != nil {
			logger.InitLog.Infoln(server.ListenAndServeTLS(udmPemPath, udmKeyPath))
			assert.True(t, err == nil)
		}
	}()

	go func() { // fake udr server
		router := gin.Default()

		router.POST("/nudr-dr/v1/subscription-data/:ueId/context-data/ee-subscriptions", func(c *gin.Context) {
			ueIdentity := c.Param("ueIdentity")
			fmt.Println("==========CreateEeSubscription - Subscribe==========")
			fmt.Println("ueIdentity: ", ueIdentity)
			var testCreatedEeSubscription models.CreatedEeSubscription
			testCreatedEeSubscription.NumberOfUes = 2
			c.JSON(http.StatusCreated, testCreatedEeSubscription)
		})

		udrLogPath := path_util.Gofree5gcPath("free5gc/udrsslkey.log")
		udrPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udr.pem")
		udrKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udr.key")

		server, err := http2_util.NewServer(":29504", udrLogPath, router)
		if err == nil && server != nil {
			logger.InitLog.Infoln(server.ListenAndServeTLS(udrPemPath, udrKeyPath))
			assert.True(t, err == nil)
		}
	}()

	udm_context.Init()
	cfg := Nudm_EE_Client.NewConfiguration()
	cfg.SetBasePath("https://localhost:29503")
	clientAPI := Nudm_EE_Client.NewAPIClient(cfg)

	var eeSubscription models.EeSubscription
	eeSubscription.SupportedFeatures = "Test001"
	ueIdentity := "SDM1234"
	createdEeSubscription, resp, err := clientAPI.CreateEESubscriptionApi.CreateEeSubscription(context.Background(), ueIdentity, eeSubscription)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("resp: ", resp)
		fmt.Println("createdEeSubscription: ", createdEeSubscription)
	}
}
