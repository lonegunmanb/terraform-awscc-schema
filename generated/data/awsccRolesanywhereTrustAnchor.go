package data

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
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
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
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "event": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "threshold": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "source": {
        "computed": true,
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
                    "type": "string"
                  },
                  "x509_certificate_data": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "source_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
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
    "description": "Data Source schema for AWS::RolesAnywhere::TrustAnchor",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRolesanywhereTrustAnchorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRolesanywhereTrustAnchor), &result)
	return &result
}
