package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicediscoveryPrivateDnsNamespace = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the private namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_id": {
        "computed": true,
        "description": "The ID for the Route 53 hosted zone that AWS Cloud Map creates when you create a namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name that you want to assign to this namespace. When you create a private DNS namespace, AWS Cloud Map automatically creates an Amazon Route 53 private hosted zone that has the same name as the namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns_namespace_id": {
        "computed": true,
        "description": "The ID of the private namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "properties": {
        "computed": true,
        "description": "Properties of the private DNS namespace.",
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
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
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
      "vpc": {
        "computed": true,
        "description": "The ID of the Amazon VPC that you want to associate the namespace with.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServiceDiscovery::PrivateDnsNamespace",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicediscoveryPrivateDnsNamespaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicediscoveryPrivateDnsNamespace), &result)
	return &result
}
