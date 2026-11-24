package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53GlobalresolverCoreNetworkEndpoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the core network endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "available_regions": {
        "computed": true,
        "description": "Regions where the endpoint can serve DNS: the Global Resolver regions that are also core network edge locations.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "client_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_arn": {
        "computed": true,
        "description": "The ARN of the AWS Cloud WAN core network to connect. The core network must be in the same account and AVAILABLE. A core network can have one core network endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_endpoint_id": {
        "computed": true,
        "description": "The ID of the core network endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "global_resolver_id": {
        "computed": true,
        "description": "The ID of the Global Resolver to connect. A Global Resolver can have one core network endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_addresses": {
        "computed": true,
        "description": "The private DNS IP addresses of the endpoint (1 to 5). Each IP address can appear once.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dns_label": {
              "computed": true,
              "description": "Lowercase DNS label for the IP address. The service appends the endpoint ID to build the DNS name. The label of an existing IP address cannot be changed; to use a new label, change the IP address too.",
              "description_kind": "plain",
              "type": "string"
            },
            "ipv_4_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ipv_6_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "region_scope": {
              "computed": true,
              "description": "ALL to serve the IP address in every available region of the endpoint, or REGIONAL to serve it only in Regions.",
              "description_kind": "plain",
              "type": "string"
            },
            "regions": {
              "computed": true,
              "description": "Regions that serve the IP address. Required when RegionScope is REGIONAL and not allowed when it is ALL. Must be a subset of AvailableRegions.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "name": {
        "computed": true,
        "description": "A name for the core network endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "CREATING, OPERATIONAL, UPDATING or DELETING. The service can also report FAILURE or ACTION_NEEDED, so this is not an enum.",
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
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53GlobalResolver::CoreNetworkEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53GlobalresolverCoreNetworkEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53GlobalresolverCoreNetworkEndpoint), &result)
	return &result
}
