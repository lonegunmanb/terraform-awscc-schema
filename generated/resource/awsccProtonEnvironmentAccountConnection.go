package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccProtonEnvironmentAccountConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the environment account connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "codebuild_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM service role in the environment account. AWS Proton uses this role to provision infrastructure resources using CodeBuild-based provisioning in the associated environment account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "component_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM service role that AWS Proton uses when provisioning directly defined components in the associated environment account. It determines the scope of infrastructure that a component can provision in the account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_account_connection_id": {
        "computed": true,
        "description": "The ID of the environment account connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_account_id": {
        "computed": true,
        "description": "The environment account that's connected to the environment account connection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_name": {
        "computed": true,
        "description": "The name of the AWS Proton environment that's created in the associated management account.",
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
      "management_account_id": {
        "computed": true,
        "description": "The ID of the management account that accepts or rejects the environment account connection. You create an manage the AWS Proton environment in this account. If the management account accepts the environment account connection, AWS Proton can use the associated IAM role to provision environment infrastructure resources in the associated environment account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM service role that's created in the environment account. AWS Proton uses this role to provision infrastructure resources in the associated environment account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the environment account connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "\u003cp\u003eAn optional list of metadata items that you can associate with the Proton environment account connection. A tag is a key-value pair.\u003c/p\u003e\n         \u003cp\u003eFor more information, see \u003ca href=\"https://docs.aws.amazon.com/proton/latest/userguide/resources.html\"\u003eProton resources and tagging\u003c/a\u003e in the\n        \u003ci\u003eProton User Guide\u003c/i\u003e.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "\u003cp\u003eThe key of the resource tag.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "\u003cp\u003eThe value of the resource tag.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Schema describing various properties for AWS Proton Environment Account Connections resources.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccProtonEnvironmentAccountConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccProtonEnvironmentAccountConnection), &result)
	return &result
}
