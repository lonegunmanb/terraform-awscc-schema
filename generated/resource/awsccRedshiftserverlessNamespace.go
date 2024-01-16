package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftserverlessNamespace = `{
  "block": {
    "attributes": {
      "admin_user_password": {
        "computed": true,
        "description": "The password associated with the admin user for the namespace that is being created. Password must be at least 8 characters in length, should be any printable ASCII character. Must contain at least one lowercase letter, one uppercase letter and one decimal digit.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "admin_username": {
        "computed": true,
        "description": "The user name associated with the admin user for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_name": {
        "computed": true,
        "description": "The database name associated for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_iam_role_arn": {
        "computed": true,
        "description": "The default IAM role ARN for the namespace that is being created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "final_snapshot_name": {
        "computed": true,
        "description": "The name of the namespace the source snapshot was created from. Please specify the name if needed before deleting namespace",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "final_snapshot_retention_period": {
        "computed": true,
        "description": "The number of days to retain automated snapshot in the destination region after they are copied from the source region. If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3,653.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "iam_roles": {
        "computed": true,
        "description": "A list of AWS Identity and Access Management (IAM) roles that can be used by the namespace to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. The Default role limit for each request is 10.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the namespace.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_exports": {
        "computed": true,
        "description": "The collection of log types to be exported provided by the customer. Should only be one of the three supported log types: userlog, useractivitylog and connectionlog",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "namespace": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "admin_username": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "creation_date": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "db_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "default_iam_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "iam_roles": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "log_exports": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "namespace_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "namespace_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "namespace_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "namespace_name": {
        "description": "A unique identifier for the namespace. You use this identifier to refer to the namespace for any subsequent namespace operations such as deleting or modifying. All alphabetical characters must be lower case. Namespace name should be unique for all namespaces within an AWS account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The list of tags for the namespace.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::RedshiftServerless::Namespace Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRedshiftserverlessNamespaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftserverlessNamespace), &result)
	return &result
}
