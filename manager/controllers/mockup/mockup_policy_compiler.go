// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package mockup

import (
	"log"

	"github.com/mesh-for-data/the-mesh-for-data/manager/controllers/utils"
	pb "github.com/mesh-for-data/the-mesh-for-data/pkg/connectors/protobuf"
	pc "github.com/mesh-for-data/the-mesh-for-data/pkg/policy-compiler/policy-compiler"
)

// MockPolicyCompiler is a mock for PolicyCompiler interface used in tests
type MockPolicyCompiler struct {
	pc.IPolicyCompiler
}

// GetPoliciesDecisions implements the PolicyCompiler interface
func (s *MockPolicyCompiler) GetPoliciesDecisions(in *pb.ApplicationContext) (*pb.PoliciesDecisions, error) {
	log.Printf("Received: ")
	log.Printf("ProcessingGeography: " + in.AppInfo.GetProcessingGeography())
	log.Printf("Secret: " + in.GetCredentialPath())
	log.Printf("Properties:")
	for key, val := range in.AppInfo.GetProperties() {
		log.Printf(key + " : " + val)
	}
	var externalComponents []*pb.ComponentVersion
	externalComponents = append(externalComponents, &pb.ComponentVersion{Id: "PC1", Version: "1.0", Name: "PolicyCompiler"})
	var dataSetWithActions []*pb.DatasetDecision

	for ind, element := range in.GetDatasets() {
		dataset := element.GetDataset()
		log.Printf("Sending DataSet: ")
		log.Printf("   DataSetID: " + dataset.GetDatasetId())
		var enforcementActions []*pb.EnforcementAction
		args := make(map[string]string)

		var operationDecisions []*pb.OperationDecision
		assetID := utils.GetAttribute("asset_id", dataset.GetDatasetId())
		switch assetID {
		case "allow-dataset":
			enforcementActions = append(enforcementActions, &pb.EnforcementAction{
				Name: "Allow",
				Id:   "Allow-ID",
			})
		case "deny-dataset":
			enforcementActions = append(enforcementActions, &pb.EnforcementAction{
				Name: "Deny",
				Id:   "Deny-ID",
			})
		case "allow-theshire":
			if element.GetOperation().Destination == "theshire" {
				enforcementActions = append(enforcementActions, &pb.EnforcementAction{
					Name: "Allow",
					Id:   "Allow-ID",
				})
			} else {
				enforcementActions = append(enforcementActions, &pb.EnforcementAction{
					Name: "Deny",
					Id:   "Deny-ID",
				})
			}
		case "deny-theshire":
			if element.GetOperation().Destination != "theshire" {
				enforcementActions = append(enforcementActions, &pb.EnforcementAction{
					Name: "Allow",
					Id:   "Allow-ID",
				})
			} else {
				enforcementActions = append(enforcementActions, &pb.EnforcementAction{
					Name: "Deny",
					Id:   "Deny-ID",
				})
			}
		default:
			args["column"] = "SSN"
			enforcementActions = append(enforcementActions, &pb.EnforcementAction{
				Name:  "redact",
				Id:    "redact-ID",
				Level: pb.EnforcementAction_COLUMN,
				Args:  args})
		}
		operationDecisions = append(operationDecisions, &pb.OperationDecision{Operation: in.GetDatasets()[0].GetOperation(), EnforcementActions: enforcementActions})
		dataSetWithActions = append(dataSetWithActions, &pb.DatasetDecision{
			Dataset: &pb.DatasetIdentifier{
				DatasetId: in.GetDatasets()[ind].GetDataset().GetDatasetId()},
			Decisions: operationDecisions})
	}

	return &pb.PoliciesDecisions{ComponentVersions: externalComponents,
		DatasetDecisions: dataSetWithActions}, nil
}
