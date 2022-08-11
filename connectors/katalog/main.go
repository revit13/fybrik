// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"emperror.dev/errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
	kconfig "sigs.k8s.io/controller-runtime/pkg/client/config"

	"fybrik.io/fybrik/connectors/katalog/pkg/apis/katalog/v1alpha1"
	"fybrik.io/fybrik/connectors/katalog/pkg/connector"
	"fybrik.io/fybrik/pkg/environment"
	fybrikTLS "fybrik.io/fybrik/pkg/tls"
)

const (
	envServicePort = "SERVICE_PORT"
)

var (
	gitCommit string
	gitTag    string
)

// RootCmd defines the root cli command
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "katalog",
		Short: "Kubernetes based data catalog for Fybrik",
	}
	cmd.AddCommand(RunCmd())
	return cmd
}

// RunCmd defines the command for running the connector
func RunCmd() *cobra.Command {
	ip := ""
	portStr, err := environment.MustGetEnv(envServicePort)
	if err != nil {
		log.Err(err).Msg(envServicePort + " env var is not defined")
		return nil
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Err(err).Msg(fmt.Sprintf("error in converting %s = [%s] to integer", envServicePort, portStr))
		return nil
	}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the connector",
		RunE: func(cmd *cobra.Command, args []string) error {
			gin.SetMode(gin.ReleaseMode)

			scheme := runtime.NewScheme()
			err := v1alpha1.AddToScheme(scheme)
			if err != nil {
				return errors.Wrap(err, "unable to add katalog v1alpha1 to schema")
			}
			err = corev1.AddToScheme(scheme)
			if err != nil {
				return errors.Wrap(err, "unable to add corev1 to schema")
			}

			client, err := kclient.New(kconfig.GetConfigOrDie(), kclient.Options{Scheme: scheme})
			if err != nil {
				return errors.Wrap(err, "failed to create a Kubernetes client")
			}

			handler := connector.NewHandler(client)
			handler.Log.Info().Msg("based on: gitTag=" + gitTag + ", latest gitCommit=" + gitCommit)
			router := connector.NewRouter(handler)
			router.Use(gin.Logger())
			bindAddress := fmt.Sprintf("%s:%d", ip, port)

			if environment.IsUsingTLS() {
				tlsConfig, err := fybrikTLS.GetServerConfig(&handler.Log)
				if err != nil {
					return errors.Wrap(err, "failed to get tls config")
				}
				server := http.Server{Addr: bindAddress, Handler: router, TLSConfig: tlsConfig}
				return server.ListenAndServeTLS("", "")
			}

			handler.Log.Info().Msg(fybrikTLS.TLSDisabledMsg)
			return router.Run(bindAddress)
		},
	}
	cmd.Flags().StringVar(&ip, "ip", ip, "IP address")
	cmd.Flags().IntVar(&port, "port", port, "Listening port")
	return cmd
}

func main() {
	// Run the cli
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
