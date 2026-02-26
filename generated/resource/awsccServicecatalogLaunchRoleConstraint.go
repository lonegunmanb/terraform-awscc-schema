package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogLaunchRoleConstraint = `{
  "block": {
    "attributes": {
      "accept_language": {
        "computed": true,
        "description": "The language code for the constraint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the launch role constraint.",
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
      "launch_role_constraint_id": {
        "computed": true,
        "description": "The unique identifier for the launch role constraint.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_role_name": {
        "computed": true,
        "description": "The local IAM role name to use in the launch constraint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "portfolio_id": {
        "description": "The ID of the portfolio to which this launch role constraint applies.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_id": {
        "description": "The ID of the product to which this launch role constraint applies.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role used for the launch constraint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ServiceCatalog::LaunchRoleConstraint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicecatalogLaunchRoleConstraintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogLaunchRoleConstraint), &result)
	return &result
}
