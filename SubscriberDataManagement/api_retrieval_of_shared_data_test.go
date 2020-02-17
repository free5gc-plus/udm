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
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
)

// GetSharedData - retrieve shared data
func TestGetSharedData(t *testing.T) {

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
	go udm_handler.Handle()

	go func() { // fake udr server
		router := gin.Default()

		router.GET("/nudr-dr/v1/subscription-data/shared-data", func(c *gin.Context) {
			sharedDataArray := c.QueryArray("shared-data-ids")
			fmt.Println("==========GetSharedData - retrieve shared data==========")
			fmt.Println("shared-data-ids", sharedDataArray)
			var testsharedData []models.SharedData
			for _, element := range testsharedData {
				element.SharedDataId = "Id1"
			}
			c.JSON(http.StatusOK, testsharedData)
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

	sharedDataIds := []string{"sharedId1", "sharedId2"}
	var getSharedDataParamOpts Nudm_SDM_Client.GetSharedDataParamOpts
	_, resp, err := clientAPI.RetrievalOfSharedDataApi.GetSharedData(context.Background(), sharedDataIds, &getSharedDataParamOpts)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("resp: ", resp)
		// fmt.Println("sharedDatas: ", sharedDatas)
	}
}