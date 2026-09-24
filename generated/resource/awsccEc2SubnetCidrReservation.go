package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SubnetCidrReservation = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the subnet CIDR reservation.",
        "description_kind": "plain",
        "type": "string"
      },
      "cidr": {
        "description": "The IPv4 or IPv6 CIDR range to reserve. Must lie inside the subnet's CIDR block and must not overlap an existing reservation. Supply an IPv6 range exactly as the service returns it: Amazon EC2 stores an IPv6 CIDR in RFC 5952 compressed form, and a noncanonical spelling of the same range is reported as drift because this property is create-only and read back.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the subnet CIDR reservation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "The identifier of the AWS account that owns the subnet CIDR reservation.",
        "description_kind": "plain",
        "type": "string"
      },
      "reservation_type": {
        "description": "The type of reservation. A prefix reservation is used for an IPv6 prefix delegated to a network interface; an explicit reservation is used for a range that Amazon EC2 must not assign automatically.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_cidr_reservation_id": {
        "computed": true,
        "description": "The identifier of the subnet CIDR reservation.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "description": "The identifier of the subnet the CIDR range is reserved in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the subnet CIDR reservation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag. Amazon EC2 reserves keys beginning with 'aws:' for internal use and rejects them for this resource. Amazon EC2 accepts a maximum of 128 Unicode characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag. Amazon EC2 accepts a maximum of 256 Unicode characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for a CIDR range reserved inside an Amazon EC2 subnet.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SubnetCidrReservationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SubnetCidrReservation), &result)
	return &result
}
