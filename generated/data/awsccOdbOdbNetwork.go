package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOdbOdbNetwork = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "computed": true,
        "description": "The AWS Availability Zone (AZ) where the ODB network is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description": "The AZ ID of the AZ where the ODB network is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_subnet_cidr": {
        "computed": true,
        "description": "The CIDR range of the backup subnet in the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "client_subnet_cidr": {
        "computed": true,
        "description": "The CIDR range of the client subnet in the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_dns_prefix": {
        "computed": true,
        "description": "The DNS prefix to the default DNS domain name. The default DNS domain name is oraclevcn.com.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_associated_resources": {
        "computed": true,
        "description": "Specifies whether to delete associated OCI networking resources along with the ODB network.",
        "description_kind": "plain",
        "type": "bool"
      },
      "display_name": {
        "computed": true,
        "description": "The user-friendly name of the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "oci_network_anchor_id": {
        "computed": true,
        "description": "The unique identifier of the OCI network anchor for the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor that's associated with the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_vcn_url": {
        "computed": true,
        "description": "The URL for the VCN that's associated with the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_id": {
        "computed": true,
        "description": "The unique identifier of the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the Odb Network.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, @, -, and \".",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that's 1 to 256 characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::ODB::OdbNetwork",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOdbOdbNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOdbOdbNetwork), &result)
	return &result
}
