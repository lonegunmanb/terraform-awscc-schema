package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IpamResourceDiscoveryAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_arn": {
        "computed": true,
        "description": "Arn of the IPAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_id": {
        "description": "The Id of the IPAM this Resource Discovery is associated to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipam_region": {
        "computed": true,
        "description": "The home region of the IPAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_resource_discovery_association_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resource discovery association is a part of.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_resource_discovery_association_id": {
        "computed": true,
        "description": "Id of the IPAM Resource Discovery Association.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_resource_discovery_id": {
        "description": "The Amazon Resource Name (ARN) of the IPAM Resource Discovery Association.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_default": {
        "computed": true,
        "description": "If the Resource Discovery Association exists due as part of CreateIpam.",
        "description_kind": "plain",
        "type": "bool"
      },
      "owner_id": {
        "computed": true,
        "description": "The AWS Account ID for the account where the shared IPAM exists.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_discovery_status": {
        "computed": true,
        "description": "The status of the resource discovery.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The operational state of the Resource Discovery Association. Related to Create/Delete activities.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "Resource Schema of AWS::EC2::IPAMResourceDiscoveryAssociation Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2IpamResourceDiscoveryAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpamResourceDiscoveryAssociation), &result)
	return &result
}
