package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftEndpointAuthorization = `{
  "block": {
    "attributes": {
      "account": {
        "description": "The target AWS account ID to grant or revoke access for.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "allowed_all_vp_cs": {
        "computed": true,
        "description": "Indicates whether all VPCs in the grantee account are allowed access to the cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "allowed_vp_cs": {
        "computed": true,
        "description": "The VPCs allowed access to the cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "authorize_time": {
        "computed": true,
        "description": "The time (UTC) when the authorization was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_identifier": {
        "description": "The cluster identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_status": {
        "computed": true,
        "description": "The status of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_count": {
        "computed": true,
        "description": "The number of Redshift-managed VPC endpoints created for the authorization.",
        "description_kind": "plain",
        "type": "number"
      },
      "force": {
        "computed": true,
        "description": " Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "grantee": {
        "computed": true,
        "description": "The AWS account ID of the grantee of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "grantor": {
        "computed": true,
        "description": "The AWS account ID of the cluster owner.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the authorization action.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_ids": {
        "computed": true,
        "description": "The virtual private cloud (VPC) identifiers to grant or revoke access to.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Describes an endpoint authorization for authorizing Redshift-managed VPC endpoint access to a cluster across AWS accounts.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRedshiftEndpointAuthorizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftEndpointAuthorization), &result)
	return &result
}
