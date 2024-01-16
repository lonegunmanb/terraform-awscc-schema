package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOrganizationsOrganization = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "feature_set": {
        "computed": true,
        "description": "Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The unique identifier (ID) of an organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "management_account_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the account that is designated as the management account for the organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "management_account_email": {
        "computed": true,
        "description": "The email address that is associated with the AWS account that is designated as the management account for the organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "management_account_id": {
        "computed": true,
        "description": "The unique identifier (ID) of the management account of an organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "root_id": {
        "computed": true,
        "description": "The unique identifier (ID) for the root.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Organizations::Organization",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOrganizationsOrganizationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOrganizationsOrganization), &result)
	return &result
}
