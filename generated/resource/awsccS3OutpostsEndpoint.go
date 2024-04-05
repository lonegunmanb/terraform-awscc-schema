package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3OutpostsEndpoint = `{
  "block": {
    "attributes": {
      "access_type": {
        "computed": true,
        "description": "The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "cidr_block": {
        "computed": true,
        "description": "The VPC CIDR committed by this endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The time the endpoint was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_owned_ipv_4_pool": {
        "computed": true,
        "description": "The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_id": {
        "computed": true,
        "description": "The ID of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "failed_reason": {
        "computed": true,
        "description": "The failure reason, if any, for a create or delete endpoint operation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "error_code": {
              "computed": true,
              "description": "The failure code, if any, for a create or delete endpoint operation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "message": {
              "computed": true,
              "description": "Additional error details describing the endpoint failure and recommended action.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_interfaces": {
        "computed": true,
        "description": "The network interfaces of the endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "network_interface_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "outpost_id": {
        "description": "The id of the customer outpost on which the bucket resides.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_group_id": {
        "description": "The ID of the security group to use with the endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "description": "The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type Definition for AWS::S3Outposts::Endpoint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3OutpostsEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3OutpostsEndpoint), &result)
	return &result
}
