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
	"testing"
)

// GetSupi - retrieve multiple data sets
func TestGetSupi(t *testing.T) {

	/*	udm_util.testInitUdmConfig()
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

			go func() { // fake udr server
				router := gin.Default()

				router.GET("/nudr-dr/v1/subscription-data/:ueId/:servingPlmnId/provisioned-data/am-data", func(c *gin.Context) {
					supi := c.Param("supi")
					fmt.Println("==========AM Data==========")
					fmt.Println("supi: ", supi)
					var testAccessAndMobilitySubscriptionData models.AccessAndMobilitySubscriptionData
					c.JSON(http.StatusOK, testAccessAndMobilitySubscriptionData)
				})

				router.GET("/nudr-dr/v1/subscription-data/:ueId/:servingPlmnId/provisioned-data/smf-selection-subscription-data", func(c *gin.Context) {
					supi := c.Param("supi")
					fmt.Println("==========SMF selection subscription data==========")
					fmt.Println("supi: ", supi)
					var testsmfSelectionSubscriptionData models.SmfSelectionSubscriptionData
					c.JSON(http.StatusOK, testsmfSelectionSubscriptionData)
				})

				router.GET("/nudr-dr/v1/subscription-data/:ueId/:servingPlmnId/provisioned-data/trace-data", func(c *gin.Context) {
					supi := c.Param("supi")
					fmt.Println("==========Trace data==========")
					fmt.Println("supi: ", supi)
					var testsmfSelectionSubscriptionData models.SmfSelectionSubscriptionData
					c.JSON(http.StatusOK, testsmfSelectionSubscriptionData)
				})

				router.GET("/nudr-dr/v1/subscription-data/:ueId/:servingPlmnId/provisioned-data/sm-data", func(c *gin.Context) {
					supi := c.Param("supi")
					fmt.Println("==========Session manangeemnt==========")
					fmt.Println("supi: ", supi)
					var testsmfSelectionSubscriptionData models.SmfSelectionSubscriptionData
					c.JSON(http.StatusOK, testsmfSelectionSubscriptionData)
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
			smfSelectionSubscriptionData, resp, err := clientAPI.SMFSelectionSubscriptionDataRetrievalApi.GetSmfSelectData(context.Background(), supi, nil)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("resp: ", resp)
				fmt.Println("smfSelectionSubscriptionData: ", smfSelectionSubscriptionData)
			}
	*/
}
