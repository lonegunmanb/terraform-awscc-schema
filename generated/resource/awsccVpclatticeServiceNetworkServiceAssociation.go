package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVpclatticeServiceNetworkServiceAssociation = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_entry": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "domain_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "hosted_zone_id": {
              "computed": true,
              "description_kind": "plain",
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
      "service_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_network_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_network_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_network_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_network_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_network_service_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Associates a service with a service network.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccVpclatticeServiceNetworkServiceAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVpclatticeServiceNetworkServiceAssociation), &result)
	return &result
}
