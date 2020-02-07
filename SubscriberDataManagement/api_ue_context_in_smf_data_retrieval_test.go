/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SubscriberDataManagement_test

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	Nudm_SDM_Client "gofree5gc/lib/Nudm_SubscriberDataManagement"
	"gofree5gc/lib/http2_util"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/lib/path_util"
	Nudm_SDM_Server "gofree5gc/src/udm/SubscriberDataManagement"
	"gofree5gc/src/udm/logger"
	"gofree5gc/src/udm/udm_context"
	"gofree5gc/src/udm/udm_handler"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// GetUeContextInSmfData - retrieve a UE's UE Context In SMF Data
func TestGetUeContextInSmfData(t *testing.T) {

	go udm_handler.Handle()
	go func() { // udm server
		router := gin.Default()
		Nudm_SDM_Server.AddService(router)

		udmLogPath := path_util.Gofree5gcPath("gofree5gc/udmsslkey.log")
		udmPemPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udm.pem")
		udmKeyPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udm.key")
		server, err := http2_util.NewServer(":29503", udmLogPath, router)
		if err == nil && server != nil {
			logger.InitLog.Infoln(server.ListenAndServeTLS(udmPemPath, udmKeyPath))
			assert.True(t, err == nil)
		}
	}()

	udm_context.TestInit()

	go func() { // fake udr server
		router := gin.Default()

		router.GET("/nudr-dr/v1/subscription-data/:ueId/context-data/smf-registrations", func(c *gin.Context) {
			supi := c.Param("supi")
			fmt.Println("==========GetUeContextInSmfData - retrieve a UE's UE Context In SMF Data==========")
			fmt.Println("supi: ", supi)

			var testueContextInSmfData models.UeContextInSmfData
			// testueContextInSmfData.EmergencyInfo.PgwFqdn = "TEst_00"
			c.JSON(http.StatusOK, testueContextInSmfData)
		})

		udrLogPath := path_util.Gofree5gcPath("gofree5gc/udrsslkey.log")
		udrPemPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udr.pem")
		udrKeyPath := path_util.Gofree5gcPath("gofree5gc/support/TLS/udr.key")

		server, err := http2_util.NewServer(":29504", udrLogPath, router)
		if err == nil && server != nil {
			logger.InitLog.Infoln(server.ListenAndServeTLS(udrPemPath, udrKeyPath))
			assert.True(t, err == nil)
		}
	}()

	udm_context.Init()
	cfg := Nudm_SDM_Client.NewConfiguration()
	cfg.SetBasePath("https://localhost:29503")
	clientAPI := Nudm_SDM_Client.NewAPIClient(cfg)

	supi := "SDM1234"
	var getUeContextInSmfDataParamOpts Nudm_SDM_Client.GetUeContextInSmfDataParamOpts
	getUeContextInSmfDataParamOpts.SupportedFeatures = optional.NewString("supportedFeatures")
	_, resp, err := clientAPI.UEContextInSMFDataRetrievalApi.GetUeContextInSmfData(context.Background(), supi, &getUeContextInSmfDataParamOpts)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("resp: ", resp)
	}
}
