package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncDomainName = `{
  "block": {
    "attributes": {
      "app_sync_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the Domain Name.",
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_id": {
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
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this Domain Name.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string used to identify this tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string containing the value for this tag. You can specify a maximum of 256 characters for a tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::AppSync::DomainName",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppsyncDomainNameSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncDomainName), &result)
	return &result
}
