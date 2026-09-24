package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightCustomization = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name of the namespace-scoped account customization, arn:{Partition}:quicksight:{Region}:{AccountId}:customization/namespace/{Namespace}. Returned by the service and used as the primary identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description": "The AWS account that owns the customization. Supplied by CloudFormation from the invoking account rather than by the template, so it is read-only.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_theme": {
        "computed": true,
        "description": "The ARN of the theme applied by default in the QuickSight console for this namespace. May be an AWS-managed starter theme such as arn:{Partition}:quicksight::aws:theme/MIDNIGHT or a theme owned by this account.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "computed": true,
        "description": "The QuickSight namespace the customization applies to. One customization exists per (account, region, namespace), so this is create-only: changing it addresses a different resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags applied to the customization. QuickSight rejects any key prefixed aws: or quicksight:, so CloudFormation system tags are not propagated.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::QuickSight::Customization",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightCustomizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightCustomization), &result)
	return &result
}
