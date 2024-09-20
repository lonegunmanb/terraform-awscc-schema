package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediapackagePackagingGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the PackagingGroup.",
        "description_kind": "plain",
        "type": "string"
      },
      "authorization": {
        "computed": true,
        "description": "CDN Authorization",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cdn_identifier_secret": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that is used for CDN authorization.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secrets_role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "domain_name": {
        "computed": true,
        "description": "The fully qualified domain name for Assets in the PackagingGroup.",
        "description_kind": "plain",
        "type": "string"
      },
      "egress_access_logs": {
        "computed": true,
        "description": "The configuration parameters for egress access logging.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_group_name": {
              "computed": true,
              "description": "Sets a custom AWS CloudWatch log group name for egress logs. If a log group name isn't specified, the default name is used: /aws/MediaPackage/VodEgressAccessLogs.",
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
      "packaging_group_id": {
        "description": "The ID of the PackagingGroup.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::MediaPackage::PackagingGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediapackagePackagingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackagePackagingGroup), &result)
	return &result
}
