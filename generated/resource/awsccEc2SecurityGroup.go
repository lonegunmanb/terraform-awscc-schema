package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SecurityGroup = `{
  "block": {
    "attributes": {
      "group_description": {
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
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
              "optional": true,
              "type": "string"
            },
            "cidr_ipv_6": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_prefix_list_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_security_group_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "from_port": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ip_protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "to_port": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "security_group_ingress": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr_ip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cidr_ipv_6": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "from_port": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ip_protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_prefix_list_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_security_group_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_security_group_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_security_group_owner_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "to_port": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::SecurityGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SecurityGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SecurityGroup), &result)
	return &result
}
