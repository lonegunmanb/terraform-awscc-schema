package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationFsxWindows = `{
  "block": {
    "attributes": {
      "cmk_secret_config": {
        "computed": true,
        "description": "Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "Specifies the ARN for the customer-managed AWS KMS key used to encrypt the secret specified for SecretArn. DataSync provides this key to AWS Secrets Manager.",
              "description_kind": "plain",
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description": "Specifies the ARN for an AWS Secrets Manager secret, managed by DataSync.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "custom_secret_config": {
        "computed": true,
        "description": "Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "secret_access_role_arn": {
              "computed": true,
              "description": "Specifies the ARN for the AWS Identity and Access Management role that DataSync uses to access the secret specified for SecretArn.",
              "description_kind": "plain",
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description": "Specifies the ARN for a customer created AWS Secrets Manager secret.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain": {
        "computed": true,
        "description": "The name of the Windows domain that the FSx for Windows server belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "fsx_filesystem_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the FSx for Windows file system.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the FSx for Windows location that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "managed_secret_config": {
        "computed": true,
        "description": "Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location. DataSync uses the default AWS-managed KMS key to encrypt this secret in AWS Secrets Manager.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "secret_arn": {
              "computed": true,
              "description": "Specifies the ARN for an AWS Secrets Manager secret.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "password": {
        "computed": true,
        "description": "The password of the user who has the permissions to access files and folders in the FSx for Windows file system.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_arns": {
        "computed": true,
        "description": "The ARNs of the security groups that are to use to configure the FSx for Windows file system.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "subdirectory": {
        "computed": true,
        "description": "A subdirectory in the location's path.",
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
              "description": "The key for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "user": {
        "computed": true,
        "description": "The user who has the permissions to access files and folders in the FSx for Windows file system.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataSync::LocationFSxWindows",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatasyncLocationFsxWindowsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationFsxWindows), &result)
	return &result
}
