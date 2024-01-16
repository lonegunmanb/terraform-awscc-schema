package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverFirewallDomainList = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "Rfc3339TimeString",
        "description_kind": "plain",
        "type": "string"
      },
      "creator_request_id": {
        "computed": true,
        "description": "The id of the creator request.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_count": {
        "computed": true,
        "description": "Count",
        "description_kind": "plain",
        "type": "number"
      },
      "domain_file_url": {
        "computed": true,
        "description": "S3 URL to import domains from.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domains": {
        "computed": true,
        "description": "An inline list of domains to use for this domain list.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "ResourceId",
        "description_kind": "plain",
        "type": "string"
      },
      "managed_owner_name": {
        "computed": true,
        "description": "ServicePrincipal",
        "description_kind": "plain",
        "type": "string"
      },
      "modification_time": {
        "computed": true,
        "description": "Rfc3339TimeString",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "FirewallDomainListName",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "ResolverFirewallDomainList, possible values are COMPLETE, DELETING, UPDATING, COMPLETE_IMPORT_FAILED, IMPORTING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description": "FirewallDomainListAssociationStatus",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::Route53Resolver::FirewallDomainList.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53ResolverFirewallDomainListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverFirewallDomainList), &result)
	return &result
}
