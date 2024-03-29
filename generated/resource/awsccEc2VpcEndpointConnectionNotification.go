package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcEndpointConnectionNotification = `{
  "block": {
    "attributes": {
      "connection_events": {
        "description": "The endpoint events for which to receive notifications.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "connection_notification_arn": {
        "description": "The ARN of the SNS topic for the notifications.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_id": {
        "computed": true,
        "description": "The ID of the endpoint service.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_endpoint_connection_notification_id": {
        "computed": true,
        "description": "VPC Endpoint Connection ID generated by service",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_endpoint_id": {
        "computed": true,
        "description": "The ID of the endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::VPCEndpointConnectionNotification",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcEndpointConnectionNotificationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcEndpointConnectionNotification), &result)
	return &result
}
