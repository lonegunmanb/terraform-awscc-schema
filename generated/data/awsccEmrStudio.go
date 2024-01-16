package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEmrStudio = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the EMR Studio.",
        "description_kind": "plain",
        "type": "string"
      },
      "auth_mode": {
        "computed": true,
        "description": "Specifies whether the Studio authenticates users using single sign-on (SSO) or IAM. Amazon EMR Studio currently only supports SSO authentication.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_s3_location": {
        "computed": true,
        "description": "The default Amazon S3 location to back up EMR Studio Workspaces and notebook files. A Studio user can select an alternative Amazon S3 location when creating a Workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A detailed description of the Studio.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_key_arn": {
        "computed": true,
        "description": "The AWS KMS key identifier (ARN) used to encrypt AWS EMR Studio workspace and notebook files when backed up to AWS S3.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_security_group_id": {
        "computed": true,
        "description": "The ID of the Amazon EMR Studio Engine security group. The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by VpcId.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "idc_instance_arn": {
        "computed": true,
        "description": "The ARN of the IAM Identity Center instance to create the Studio application.",
        "description_kind": "plain",
        "type": "string"
      },
      "idc_user_assignment": {
        "computed": true,
        "description": "Specifies whether IAM Identity Center user assignment is REQUIRED or OPTIONAL. If the value is set to REQUIRED, users must be explicitly assigned to the Studio application to access the Studio.",
        "description_kind": "plain",
        "type": "string"
      },
      "idp_auth_url": {
        "computed": true,
        "description": "Your identity provider's authentication endpoint. Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "idp_relay_state_parameter_name": {
        "computed": true,
        "description": "The name of relay state parameter for external Identity Provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A descriptive name for the Amazon EMR Studio.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_role": {
        "computed": true,
        "description": "The IAM role that will be assumed by the Amazon EMR Studio. The service role provides a way for Amazon EMR Studio to interoperate with other AWS services.",
        "description_kind": "plain",
        "type": "string"
      },
      "studio_id": {
        "computed": true,
        "description": "The ID of the EMR Studio.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "A list of up to 5 subnet IDs to associate with the Studio. The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to associate with the Studio. Tags are user-defined key-value pairs that consist of a required key string with a maximum of 128 characters, and an optional value string with a maximum of 256 characters.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "trusted_identity_propagation_enabled": {
        "computed": true,
        "description": "A Boolean indicating whether to enable Trusted identity propagation for the Studio. The default value is false.",
        "description_kind": "plain",
        "type": "bool"
      },
      "url": {
        "computed": true,
        "description": "The unique Studio access URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_role": {
        "computed": true,
        "description": "The IAM user role that will be assumed by users and groups logged in to a Studio. The permissions attached to this IAM role can be scoped down for each user or group using session policies.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_security_group_id": {
        "computed": true,
        "description": "The ID of the Amazon EMR Studio Workspace security group. The Workspace security group allows outbound network traffic to resources in the Engine security group, and it must be in the same VPC specified by VpcId.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EMR::Studio",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEmrStudioSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEmrStudio), &result)
	return &result
}
