package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotTopicRuleDestination = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN).",
        "description_kind": "plain",
        "type": "string"
      },
      "http_url_properties": {
        "computed": true,
        "description": "HTTP URL destination properties.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "confirmation_url": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
      "influx_db_properties": {
        "computed": true,
        "description": "InfluxDB destination properties.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint": {
              "computed": true,
              "description": "The endpoint URL of the InfluxDB database.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "influx_db_version": {
              "computed": true,
              "description": "The version of the InfluxDB database (for example, V2 or V3).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_id": {
              "computed": true,
              "description": "The ARN or name of the Secrets Manager secret containing the InfluxDB API token.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_key": {
              "computed": true,
              "description": "The key name within the secret that contains the InfluxDB token.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_type": {
              "computed": true,
              "description": "The type of the secret value (SecretString or SecretBinary).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "status": {
        "computed": true,
        "description": "The status of the TopicRuleDestination.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "The reasoning for the current status of the TopicRuleDestination.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_properties": {
        "computed": true,
        "description": "VPC destination properties.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "role_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_groups": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "vpc_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::IoT::TopicRuleDestination",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotTopicRuleDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotTopicRuleDestination), &result)
	return &result
}
