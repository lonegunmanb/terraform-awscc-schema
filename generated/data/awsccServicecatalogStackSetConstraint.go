package data

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
        "type": "string"
      },
      "account_list": {
        "computed": true,
        "description": "One or more AWS accounts that will have access to the provisioned product.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "admin_role": {
        "computed": true,
        "description": "AdminRole ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the constraint.",
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role": {
        "computed": true,
        "description": "ExecutionRole name.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "portfolio_id": {
        "computed": true,
        "description": "The portfolio identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_id": {
        "computed": true,
        "description": "The product identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "region_list": {
        "computed": true,
        "description": "One or more AWS Regions where the provisioned product will be available.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "stack_instance_control": {
        "computed": true,
        "description": "Permission to create, update, and delete stack instances. Choose from ALLOWED and NOT_ALLOWED.",
        "description_kind": "plain",
        "type": "string"
      },
      "stack_set_constraint_id": {
        "computed": true,
        "description": "Unique identifier for the constraint",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServiceCatalog::StackSetConstraint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicecatalogStackSetConstraintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogStackSetConstraint), &result)
	return &result
}
