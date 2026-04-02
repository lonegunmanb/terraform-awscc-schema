package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityagentTargetDomain = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "Timestamp when the target domain was registered",
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
        "description": "Tags for the target domain",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_domain_id": {
        "computed": true,
        "description": "Unique identifier of the target domain",
        "description_kind": "plain",
        "type": "string"
      },
      "target_domain_name": {
        "computed": true,
        "description": "Domain name of the target domain",
        "description_kind": "plain",
        "type": "string"
      },
      "verification_details": {
        "computed": true,
        "description": "Verification details to verify registered target domain",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dns_txt": {
              "computed": true,
              "description": "Represents DNS TXT verification details",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dns_record_name": {
                    "computed": true,
                    "description": "Record name to be added in DNS for target domain",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "dns_record_type": {
                    "computed": true,
                    "description": "Type of record to be added in DNS for target domain",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "token": {
                    "computed": true,
                    "description": "Token used to verify domain ownership",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "http_route": {
              "computed": true,
              "description": "Represents HTTP route verification details",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "route_path": {
                    "computed": true,
                    "description": "Route path where verification token should be placed",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "token": {
                    "computed": true,
                    "description": "Token used to verify domain ownership",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "method": {
              "computed": true,
              "description": "Type of domain ownership verification method",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "verification_method": {
        "computed": true,
        "description": "Verification method for the target domain",
        "description_kind": "plain",
        "type": "string"
      },
      "verification_status": {
        "computed": true,
        "description": "Current verification status of the registered target domain",
        "description_kind": "plain",
        "type": "string"
      },
      "verified_at": {
        "computed": true,
        "description": "Timestamp when the target domain was last successfully verified",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityAgent::TargetDomain",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityagentTargetDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityagentTargetDomain), &result)
	return &result
}
