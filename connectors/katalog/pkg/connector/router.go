// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package connector

import (
	"github.com/gin-gonic/gin"
)

// NewRouter returns a new router.
func NewRouter(handler *Handler) *gin.Engine {
	router := gin.Default()
	/*	loadedCertServer, err := tls.LoadX509KeyPair("/server-cert.pem", "/server-key.pem")
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		config := &tls.Config{
			Certificates: []tls.Certificate{loadedCertServer},
			ClientAuth:   tls.NoClientCert,
		}*/

	//err := http.ListenAndServeTLS(tmp, "/server-cert.pem", "/server-key.pem", router)
	/*server := http.Server{Addr: tmp, Handler: router, TLSConfig: config}
	glog.Fatal(server.ListenAndServeTLS("", ""))*/
	/*if err != nil {
		fmt.Println(err.Error())
		return nil
	}*/
	router.POST("/getAssetInfo", handler.getAssetInfo)
	router.POST("/createAsset", handler.createAsset)
	router.DELETE("/deleteAsset", handler.deleteAsset)
	router.PATCH("/updateAsset", handler.updateAsset)
	return router
}
