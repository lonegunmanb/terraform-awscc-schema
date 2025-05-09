package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneDomainUnit = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp at which the domain unit was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the domain unit.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The ID of the domain where the domain unit was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "description": "The ID of the domain where you want to create a domain unit.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_unit_id": {
        "computed": true,
        "description": "The ID of the domain unit.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identifier": {
        "computed": true,
        "description": "The identifier of the domain unit that you want to get.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The timestamp at which the domain unit was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the domain unit.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent_domain_unit_id": {
        "computed": true,
        "description": "The ID of the parent domain unit.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent_domain_unit_identifier": {
        "description": "The ID of the parent domain unit.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "A domain unit enables you to easily organize your assets and other domain entities under specific business units and teams.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazoneDomainUnitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneDomainUnit), &result)
	return &result
}
