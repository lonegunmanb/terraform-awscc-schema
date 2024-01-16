package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SecurityGroup = `{
  "block": {
    "attributes": {
      "group_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_name": {
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
      "security_group_egress": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr_ip": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "cidr_ipv_6": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "destination_prefix_list_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "destination_security_group_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "from_port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "ip_protocol": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "to_port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "security_group_ingress": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr_ip": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "cidr_ipv_6": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "from_port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "ip_protocol": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_prefix_list_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_security_group_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_security_group_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_security_group_owner_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "to_port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
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
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::SecurityGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2SecurityGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SecurityGroup), &result)
	return &result
}
