package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNeptunegraphPrivateGraphEndpoint = `{
  "block": {
    "attributes": {
      "graph_identifier": {
        "computed": true,
        "description": "The auto-generated Graph Id assigned by the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "private_graph_endpoint_identifier": {
        "computed": true,
        "description": "PrivateGraphEndpoint resource identifier generated by concatenating the associated GraphIdentifier and VpcId with an underscore separator.\n\n For example, if GraphIdentifier is ` + "`" + `g-12a3bcdef4` + "`" + ` and VpcId is ` + "`" + `vpc-0a12bc34567de8f90` + "`" + `, the generated PrivateGraphEndpointIdentifier will be ` + "`" + `g-12a3bcdef4_vpc-0a12bc34567de8f90` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description": "The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_ids": {
        "computed": true,
        "description": "The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_endpoint_id": {
        "computed": true,
        "description": "VPC endpoint that provides a private connection between the Graph and specified VPC.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NeptuneGraph::PrivateGraphEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNeptunegraphPrivateGraphEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNeptunegraphPrivateGraphEndpoint), &result)
	return &result
}
