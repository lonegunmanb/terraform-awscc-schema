package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubPolicyAssociation = `{
  "block": {
    "attributes": {
      "association_identifier": {
        "computed": true,
        "description": "A unique identifier to indicates if the target has an association",
        "description_kind": "plain",
        "type": "string"
      },
      "association_status": {
        "computed": true,
        "description": "The current status of the association between the specified target and the configuration",
        "description_kind": "plain",
        "type": "string"
      },
      "association_status_message": {
        "computed": true,
        "description": "An explanation for a FAILED value for AssociationStatus",
        "description_kind": "plain",
        "type": "string"
      },
      "association_type": {
        "computed": true,
        "description": "Indicates whether the association between the specified target and the configuration was directly applied by the Security Hub delegated administrator or inherited from a parent",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_policy_id": {
        "computed": true,
        "description": "The universally unique identifier (UUID) of the configuration policy or a value of SELF_MANAGED_SECURITY_HUB for a self-managed configuration",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_id": {
        "computed": true,
        "description": "The identifier of the target account, organizational unit, or the root",
        "description_kind": "plain",
        "type": "string"
      },
      "target_type": {
        "computed": true,
        "description": "Indicates whether the target is an AWS account, organizational unit, or the organization root",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time, in UTC and ISO 8601 format, that the configuration policy association was last updated",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityHub::PolicyAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityhubPolicyAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubPolicyAssociation), &result)
	return &result
}
