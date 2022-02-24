// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"strconv"
	"strings"
	"testing"

	apiv1alpha1 "fybrik.io/fybrik/manager/apis/app/v1alpha1"
	"fybrik.io/fybrik/pkg/test"
	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/flight"
	"github.com/apache/arrow/go/arrow/memory"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/apache/arrow/go/arrow/csv"
	"github.com/aws/aws-sdk-go/service/s3"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestS3NotebookWrite(t *testing.T) {

	gomega.RegisterFailHandler(Fail)

	g := gomega.NewGomegaWithT(t)
	defer GinkgoRecover()

	// Copy data.csv file to S3
	// S3 is assumed to be exposed on localhost at port 9090
	region := "theshire"
	endpoint := "http://localhost:9090"
	//bucket := "bucket1"
	//key1 := "data.csv"
	//filename := "../../../samples/kubeflow/data.csv"
	s3credentials := credentials.NewStaticCredentials("ak", "sk", "")

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials:      s3credentials,
		Endpoint:         &endpoint,
		Region:           &region,
		S3ForcePathStyle: aws.Bool(true),
	}))
	svc := s3.New(sess)
	/*
		_, err := svc.GetObject(&s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &key1,
		})
		if err != nil { // Could not retrieve object. Assume it does not exist
			uploader := s3manager.NewUploader(sess)

			f, ferr := os.Open(filename)
			g.Expect(ferr).To(gomega.BeNil(), "Opening local test data file")

			// Upload the file to S3.
			result, err := uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(key1),
				Body:   f,
			})
			g.Expect(err).To(gomega.BeNil(), "S3 upload")
			if result != nil {
				log.Println(fmt.Sprintf("file uploaded to, %s\n", result.Location))
			}
		} else {
			//g.Expect(object).ToNot(gomega.BeNil())
			log.Println("Object already exists in S3!")
		}
	*/
	err := apiv1alpha1.AddToScheme(scheme.Scheme)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	k8sClient, err := client.New(ctrl.GetConfigOrDie(), client.Options{Scheme: scheme.Scheme})
	g.Expect(err).To(gomega.BeNil())

	// Create Kubernetes objects for test
	// - namespace (in setup before)
	// - asset (in setup before)
	// - asset secret (in setup before)
	// - arrow flight module (in setup before)
	// - rego policy (in setup before)

	// Push read module and application

	// Module installed by setup script directly from remote arrow-flight-module repository
	// Installing application
	err = apiv1alpha1.AddToScheme(scheme.Scheme)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	application := &apiv1alpha1.FybrikApplication{}
	g.Expect(readObjectFromFile("../../testdata/notebook/fybrikapplication.yaml", application)).ToNot(gomega.HaveOccurred())
	applicationKey := client.ObjectKeyFromObject(application)

	// Create FybrikApplication and FybrikModule
	By("Expecting application creation to succeed")
	g.Expect(k8sClient.Create(context.Background(), application)).Should(gomega.Succeed())

	// Ensure getting cleaned up after tests finish
	/*defer func() {
		application := &apiv1alpha1.FybrikApplication{ObjectMeta: metav1.ObjectMeta{Namespace: applicationKey.Namespace, Name: applicationKey.Name}}
		_ = k8sClient.Get(context.Background(), applicationKey, application)
		_ = k8sClient.Delete(context.Background(), application)
	}()*/

	By("Expecting application to be created")
	g.Eventually(func() error {
		return k8sClient.Get(context.Background(), applicationKey, application)
	}, timeout, interval).Should(gomega.Succeed())
	By("Expecting plotter to be constructed")
	g.Eventually(func() *apiv1alpha1.ResourceReference {
		_ = k8sClient.Get(context.Background(), applicationKey, application)
		return application.Status.Generated
	}, timeout, interval).ShouldNot(gomega.BeNil())

	// The plotter has to be created
	plotter := &apiv1alpha1.Plotter{}
	plotterObjectKey := client.ObjectKey{Namespace: application.Status.Generated.Namespace, Name: application.Status.Generated.Name}
	By("Expecting plotter to be fetchable")
	g.Eventually(func() error {
		return k8sClient.Get(context.Background(), plotterObjectKey, plotter)
	}, timeout, interval).Should(gomega.Succeed())

	By("Expecting application to be ready")
	g.Eventually(func() bool {
		err := k8sClient.Get(context.Background(), applicationKey, application)
		if err != nil {
			return false
		}
		return application.Status.Ready
	}, timeout, interval).Should(gomega.Equal(true))

	modulesNamespace := plotter.Spec.ModulesNamespace
	fmt.Printf("data access module namespace notebook test: %s\n", modulesNamespace)

	// Forward port of arrow flight service to local port
	connection := application.Status.AssetStates["fybrik-notebook-sample/data-csv"].Endpoint.AdditionalProperties.Items["fybrik-arrow-flight"].(map[string]interface{})
	hostname := fmt.Sprintf("%v", connection["hostname"])
	port := fmt.Sprintf("%v", connection["port"])
	svcName := strings.Replace(hostname, "."+modulesNamespace, "", 1)

	By("Starting kubectl port-forward for arrow-flight")
	portNum, err := strconv.Atoi(port)
	g.Expect(err).To(gomega.BeNil())
	listenPort, err := test.RunPortForward(modulesNamespace, svcName, portNum)
	g.Expect(err).To(gomega.BeNil())

	// Reading data via arrow flight
	opts := make([]grpc.DialOption, 0)
	opts = append(opts, grpc.WithInsecure())
	flightClient, err := flight.NewFlightClient(net.JoinHostPort("localhost", listenPort), nil, opts...)
	g.Expect(err).To(gomega.BeNil(), "Connect to arrow-flight service")
	defer flightClient.Close()

	request := ArrowRequest{
		Asset: "fybrik-notebook-sample/data-csv",
	}

	marshal, err := json.Marshal(request)
	g.Expect(err).To(gomega.BeNil())

	info, err := flightClient.GetFlightInfo(context.Background(), &flight.FlightDescriptor{
		Type: flight.FlightDescriptor_CMD,
		Cmd:  marshal,
	})
	g.Expect(err).To(gomega.BeNil())
	/*schema, err := flightClient.GetSchema(context.Background(), &flight.FlightDescriptor{
		Type: flight.FlightDescriptor_CMD,
		Cmd:  marshal,
	})
	g.Expect(err).To(gomega.BeNil())*/
	stream, err := flightClient.DoGet(context.Background(), info.Endpoint[0].Ticket)
	g.Expect(err).To(gomega.BeNil())

	reader, err := flight.NewRecordReader(stream)
	g.Expect(err).To(gomega.BeNil())
	newBuff := new(bytes.Buffer)
	By("Revit")

	schema1, err := flight.DeserializeSchema(info.Schema, memory.DefaultAllocator)
	g.Expect(err).To(gomega.BeNil())
	w := csv.NewWriter(newBuff, schema1, csv.WithComma(';'))

	By("Revit77")
	for reader.Next() {
		record := reader.Record()
		defer record.Release()
		By("Revit1")

		g.Expect(record.ColumnName(0)).To(gomega.Equal("step"))
		g.Expect(record.ColumnName(1)).To(gomega.Equal("type"))
		g.Expect(record.ColumnName(3)).To(gomega.Equal("nameOrig"))
		column := record.Column(3) // Check out the third 4th column that should be nameOrig and redacted

		// Check that data of nameOrig column is the correct size and all records are redacted
		dt := column.Data().DataType()
		g.Expect(column.Data().Len()).To(gomega.Equal(100))
		g.Expect(dt.ID()).To(gomega.Equal(arrow.STRING))
		g.Expect(dt.Name()).To(gomega.Equal((&arrow.StringType{}).Name()))
		data := array.NewStringData(column.Data())
		for i := 0; i < data.Len(); i++ {
			g.Expect(data.Value(i)).To(gomega.Equal("XXXXX"))
		}
		By("Revit2")
		err := w.Write(record)
		g.Expect(err).To(gomega.BeNil())
		By("Revit22")
	}
	By("Test1")
	err = w.Flush()
	g.Expect(err).To(gomega.BeNil())

	err = w.Error()
	g.Expect(err).To(gomega.BeNil())

	uploader := s3manager.NewUploader(sess)

	newBucket := "bucket2"
	key2 := "masked_data.csv"
	By("Test2")
	// Upload the masked data to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(newBucket),
		Key:    aws.String(key2),
		Body:   newBuff,
	})
	g.Expect(err).To(gomega.BeNil(), "S3 upload")
	if result != nil {
		log.Println(fmt.Sprintf("file uploaded to, %s\n", result.Location))
	}
	By("Test3")
	// Create a downloader with the s3 client and default options
	downloader := s3manager.NewDownloaderWithClient(svc)

	writeCsv := &aws.WriteAtBuffer{}
	_, err = downloader.Download(writeCsv, &s3.GetObjectInput{
		Bucket: aws.String(newBucket),
		Key:    aws.String(key2),
	})
	g.Expect(err).To(gomega.BeNil(), "S3 download")
	g.Expect(bytes.Compare(writeCsv.Bytes(), newBuff.Bytes()) == 0)
	By("Test finished")
}
