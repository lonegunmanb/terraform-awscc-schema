package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRolesanywhereTrustAnchor = `{
  "block": {
    "attributes": {
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "channel": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "event": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "threshold": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "source": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_data": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "acm_pca_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "x509_certificate_data": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "source_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "trust_anchor_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "trust_anchor_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::RolesAnywhere::TrustAnchor Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRolesanywhereTrustAnchorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRolesanywhereTrustAnchor), &result)
	return &result
}
