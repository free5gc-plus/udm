/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package subscriberdatamanagement_test

import (
	"context"
	"fmt"
	Nudm_SDM_Client "free5gc/lib/Nudm_SubscriberDataManagement"
	"free5gc/lib/http2_util"
	"free5gc/lib/openapi/models"
	"free5gc/lib/path_util"
	udm_context "free5gc/src/udm/context"
	"free5gc/src/udm/logger"
	Nudm_SDM_Server "free5gc/src/udm/subscriberdatamanagement"
	"free5gc/src/udm/udm_handler"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Modify - modify the subscription
func TestModify(t *testing.T) {

	go func() { // udm server
		router := gin.Default()
		Nudm_SDM_Server.AddService(router)

		udmLogPath := path_util.Gofree5gcPath("free5gc/udmsslkey.log")
		udmPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.pem")
		udmKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.key")

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

		router.PUT("/nudr-dr/v1/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId", func(c *gin.Context) {
			ueId := c.Param("ueId")
			fmt.Println("\n==========Modify - modify the subscription==========")
			fmt.Println("ueId: ", ueId)

			var sdmSubsMod models.SdmSubsModification
			if err := c.ShouldBindJSON(&sdmSubsMod); err != nil {
				fmt.Println("fake udr server error")
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
			fmt.Println("sdm Subs Modification - ", sdmSubsMod)
			c.JSON(http.StatusOK, gin.H{})
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
	time.Sleep(200 * time.Millisecond)
	cfg := Nudm_SDM_Client.NewConfiguration()
	cfg.SetBasePath("https://localhost:29503")
	clientAPI := Nudm_SDM_Client.NewAPIClient(cfg)

	supi := "SDM1234"
	subscriptionId := "TestSubscriptionId"
	var sdmSubsModification models.SdmSubsModification

	mod, resp, err := clientAPI.SubscriptionModificationApi.Modify(context.TODO(), supi, subscriptionId, sdmSubsModification)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("resp: ", resp)
		fmt.Println("model: ", mod)
	}
}

// ModifyForSharedData - modify the subscription
func TestModifyForSharedData(t *testing.T) {

	go func() { // udm server
		router := gin.Default()
		Nudm_SDM_Server.AddService(router)

		udmLogPath := path_util.Gofree5gcPath("free5gc/udmsslkey.log")
		udmPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.pem")
		udmKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.key")

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

		router.PUT("/nudr-dr/v1/subscription-data/:ueId/context-data/sdm-subscriptions/:subsId", func(c *gin.Context) {
			ueId := c.Param("ueId")
			fmt.Println("==========ModifyForSharedData - modify the subscription==========")
			fmt.Println("ueId: ", ueId)

			var sdmSubsModification models.SdmSubsModification
			if err := c.ShouldBindJSON(&sdmSubsModification); err != nil {
				fmt.Println("fake udr server error")
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
			fmt.Println("sdmSubsModification - ", sdmSubsModification)
			c.JSON(http.StatusOK, gin.H{})
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
	time.Sleep(200 * time.Millisecond)
	cfg := Nudm_SDM_Client.NewConfiguration()
	cfg.SetBasePath("https://localhost:29503")
	clientAPI := Nudm_SDM_Client.NewAPIClient(cfg)

	subscriptionId := "Test_subscriptionId"
	var sdmSubsModification models.SdmSubsModification
	_, resp, err := clientAPI.SubscriptionModificationApi.ModifyForSharedData(context.TODO(), subscriptionId, sdmSubsModification)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("resp: ", resp)
	}

}