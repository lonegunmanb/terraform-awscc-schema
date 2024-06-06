package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubOrganizationConfiguration = `{
  "block": {
    "attributes": {
      "auto_enable": {
        "description": "Whether to automatically enable Security Hub in new member accounts when they join the organization.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "auto_enable_standards": {
        "computed": true,
        "description": "Whether to automatically enable Security Hub default standards in new member accounts when they join the organization.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "configuration_type": {
        "computed": true,
        "description": "Indicates whether the organization uses local or central configuration.",
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
      "member_account_limit_reached": {
        "computed": true,
        "description": "Whether the maximum number of allowed member accounts are already associated with the Security Hub administrator account.",
        "description_kind": "plain",
        "type": "bool"
      },
      "organization_configuration_identifier": {
        "computed": true,
        "description": "The identifier of the OrganizationConfiguration being created and assigned as the unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Describes whether central configuration could be enabled as the ConfigurationType for the organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description": "Provides an explanation if the value of Status is equal to FAILED when ConfigurationType is equal to CENTRAL.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The AWS::SecurityHub::OrganizationConfiguration resource represents the configuration of your organization in Security Hub. Only the Security Hub administrator account can create Organization Configuration resource in each region and can opt-in to Central Configuration only in the aggregation region of FindingAggregator.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubOrganizationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubOrganizationConfiguration), &result)
	return &result
}
