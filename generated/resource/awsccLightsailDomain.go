package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailDomain = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the domain (read-only).",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the domain was created (read-only).",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_entries": {
        "computed": true,
        "description": "An array of key-value pairs containing information about the domain entries.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "The ID of the domain recordset entry.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "is_alias": {
              "computed": true,
              "description": "When true, specifies whether the domain entry is an alias used by the Lightsail load balancer, Lightsail container service, Lightsail content delivery network (CDN) distribution, or another AWS resource. You can include an alias (A type) record in your request, which points to the DNS name of a load balancer, container service, CDN distribution, or other AWS resource and routes traffic to that resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "computed": true,
              "description": "The name of the domain entry.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target": {
              "computed": true,
              "description": "The target AWS name server (e.g., ns-111.awsdns-11.com).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of domain entry (e.g., A, CNAME, MX, NS, SOA, SRV, TXT).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "domain_name": {
        "description": "The name of the domain to manage in Lightsail.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description": "The AWS Region and Availability Zone where the domain was created (read-only).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description": "The Availability Zone.",
              "description_kind": "plain",
              "type": "string"
            },
            "region_name": {
              "computed": true,
              "description": "The AWS Region name.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_type": {
        "computed": true,
        "description": "The Lightsail resource type (read-only).",
        "description_kind": "plain",
        "type": "string"
      },
      "support_code": {
        "computed": true,
        "description": "The support code. Include this code in your email to support when you have questions (read-only).",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
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
    "description": "Resource Type definition for AWS::Lightsail::Domain",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLightsailDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailDomain), &result)
	return &result
}
