package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicediscoveryPublicDnsNamespace = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the public namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the namespace.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hosted_zone_id": {
        "computed": true,
        "description": "The ID for the Route 53 hosted zone that AWS Cloud Map creates when you create a namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name that you want to assign to this namespace.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "properties": {
        "computed": true,
        "description": "Properties for the public DNS namespace.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dns_properties": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "soa": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ttl": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "public_dns_namespace_id": {
        "computed": true,
        "description": "The ID of the public namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the namespace. Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.",
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
      }
    },
    "description": "Resource Type definition for AWS::ServiceDiscovery::PublicDnsNamespace",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicediscoveryPublicDnsNamespaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicediscoveryPublicDnsNamespace), &result)
	return &result
}
