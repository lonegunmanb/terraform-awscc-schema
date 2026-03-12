package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogStackSetConstraint = `{
  "block": {
    "attributes": {
      "accept_language": {
        "computed": true,
        "description": "The language code.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "account_list": {
        "description": "One or more AWS accounts that will have access to the provisioned product.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "admin_role": {
        "description": "AdminRole ARN.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "The description of the constraint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "execution_role": {
        "description": "ExecutionRole name.",
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
      "portfolio_id": {
        "description": "The portfolio identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_id": {
        "description": "The product identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region_list": {
        "description": "One or more AWS Regions where the provisioned product will be available.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "stack_instance_control": {
        "description": "Permission to create, update, and delete stack instances. Choose from ALLOWED and NOT_ALLOWED.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stack_set_constraint_id": {
        "computed": true,
        "description": "Unique identifier for the constraint",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ServiceCatalog::StackSetConstraint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicecatalogStackSetConstraintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogStackSetConstraint), &result)
	return &result
}
