package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesVdmAttributes = `{
  "block": {
    "attributes": {
      "dashboard_attributes": {
        "computed": true,
        "description": "Preferences regarding the Dashboard feature.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "engagement_metrics": {
              "computed": true,
              "description": "Whether emails sent from this account have engagement tracking enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "guardian_attributes": {
        "computed": true,
        "description": "Preferences regarding the Guardian feature.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "optimized_shared_delivery": {
              "computed": true,
              "description": "Whether emails sent from this account have optimized delivery algorithm enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "vdm_attributes_resource_id": {
        "computed": true,
        "description": "Unique identifier for this resource",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SES::VdmAttributes",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesVdmAttributesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesVdmAttributes), &result)
	return &result
}
