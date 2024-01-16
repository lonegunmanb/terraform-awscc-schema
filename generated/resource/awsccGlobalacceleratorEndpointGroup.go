package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlobalacceleratorEndpointGroup = `{
  "block": {
    "attributes": {
      "endpoint_configurations": {
        "computed": true,
        "description": "The list of endpoint objects.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attachment_arn": {
              "computed": true,
              "description": "Attachment ARN that provides access control to the cross account endpoint. Not required for resources hosted in the same account as the endpoint group.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_ip_preservation_enabled": {
              "computed": true,
              "description": "true if client ip should be preserved",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "endpoint_id": {
              "description": "Id of the endpoint. For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weight": {
              "computed": true,
              "description": "The weight for the endpoint.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "endpoint_group_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the endpoint group",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_group_region": {
        "description": "The name of the AWS Region where the endpoint group is located",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "health_check_interval_seconds": {
        "computed": true,
        "description": "The time in seconds between each health check for an endpoint. Must be a value of 10 or 30",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_path": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_port": {
        "computed": true,
        "description": "The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_protocol": {
        "computed": true,
        "description": "The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "listener_arn": {
        "description": "The Amazon Resource Name (ARN) of the listener",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port_overrides": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint_port": {
              "description": "A network port number",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "listener_port": {
              "description": "A network port number",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "threshold_count": {
        "computed": true,
        "description": "The number of consecutive health checks required to set the state of the endpoint to unhealthy.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "traffic_dial_percentage": {
        "computed": true,
        "description": "The percentage of traffic to sent to an AWS Region",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::GlobalAccelerator::EndpointGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlobalacceleratorEndpointGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlobalacceleratorEndpointGroup), &result)
	return &result
}
