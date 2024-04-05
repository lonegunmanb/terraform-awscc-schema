package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApsScraper = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description": "Scraper alias.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Scraper ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination": {
        "description": "Scraper metrics destination",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "amp_configuration": {
              "computed": true,
              "description": "Configuration for Amazon Managed Prometheus metrics destination",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "workspace_arn": {
                    "description": "ARN of an Amazon Managed Prometheus workspace",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "IAM role ARN for the scraper.",
        "description_kind": "plain",
        "type": "string"
      },
      "scrape_configuration": {
        "description": "Scraper configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configuration_blob": {
              "computed": true,
              "description": "Prometheus compatible scrape configuration in base64 encoded blob format",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "scraper_id": {
        "computed": true,
        "description": "Required to identify a specific scraper.",
        "description_kind": "plain",
        "type": "string"
      },
      "source": {
        "description": "Scraper metrics source",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "eks_configuration": {
              "computed": true,
              "description": "Configuration for EKS metrics source",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cluster_arn": {
                    "description": "ARN of an EKS cluster",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "security_group_ids": {
                    "computed": true,
                    "description": "List of security group IDs",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_ids": {
                    "description": "List of subnet IDs",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "Resource Type definition for AWS::APS::Scraper",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApsScraperSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApsScraper), &result)
	return &result
}
