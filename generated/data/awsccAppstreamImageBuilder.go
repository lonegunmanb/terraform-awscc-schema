package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamImageBuilder = `{
  "block": {
    "attributes": {
      "access_endpoints": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vpce_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "appstream_agent_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_join_info": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "directory_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "organizational_unit_distinguished_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "enable_default_internet_access": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "iam_role_arn": {
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
      "image_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "streaming_url": {
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
      "vpc_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::AppStream::ImageBuilder",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppstreamImageBuilderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamImageBuilder), &result)
	return &result
}
