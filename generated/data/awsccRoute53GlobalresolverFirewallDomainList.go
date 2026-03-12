package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53GlobalresolverFirewallDomainList = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "domain_file_url": {
        "computed": true,
        "description": "S3 URL to import domains from.",
        "description_kind": "plain",
        "type": "string"
      },
      "domains": {
        "computed": true,
        "description": "An inline list of domains to use for this domain list.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "firewall_domain_list_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "global_resolver_id": {
        "computed": true,
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
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
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
    "description": "Data Source schema for AWS::Route53GlobalResolver::FirewallDomainList",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53GlobalresolverFirewallDomainListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53GlobalresolverFirewallDomainList), &result)
	return &result
}
