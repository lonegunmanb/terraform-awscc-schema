package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewiseApplication = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The unique identifier of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time the application was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the application.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_subdomain": {
        "computed": true,
        "description": "The DNS-compliant subdomain label assigned to the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "idc_application_arn": {
        "computed": true,
        "description": "The ARN of the IAM Identity Center application created for the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "idc_instance_arn": {
        "computed": true,
        "description": "The ARN of the IAM Identity Center instance used to create the application.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The time the application was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_name": {
        "description": "The name of the workspace that the application belongs to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Represents an AWS IoT SiteWise application that provides IAM Identity Center based access to a workspace.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotsitewiseApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewiseApplication), &result)
	return &result
}
