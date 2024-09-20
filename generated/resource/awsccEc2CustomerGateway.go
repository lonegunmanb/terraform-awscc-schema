package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2CustomerGateway = `{
  "block": {
    "attributes": {
      "bgp_asn": {
        "computed": true,
        "description": "For customer gateway devices that support BGP, specify the device's ASN. You must specify either ` + "`" + `` + "`" + `BgpAsn` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `BgpAsnExtended` + "`" + `` + "`" + ` when creating the customer gateway. If the ASN is larger than ` + "`" + `` + "`" + `2,147,483,647` + "`" + `` + "`" + `, you must use ` + "`" + `` + "`" + `BgpAsnExtended` + "`" + `` + "`" + `.\n Default: 65000\n Valid values: ` + "`" + `` + "`" + `1` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `2,147,483,647` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "bgp_asn_extended": {
        "computed": true,
        "description": "For customer gateway devices that support BGP, specify the device's ASN. You must specify either ` + "`" + `` + "`" + `BgpAsn` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `BgpAsnExtended` + "`" + `` + "`" + ` when creating the customer gateway. If the ASN is larger than ` + "`" + `` + "`" + `2,147,483,647` + "`" + `` + "`" + `, you must use ` + "`" + `` + "`" + `BgpAsnExtended` + "`" + `` + "`" + `.\n Valid values: ` + "`" + `` + "`" + `2,147,483,648` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `4,294,967,295` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "certificate_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the customer gateway certificate.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "customer_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "device_name": {
        "computed": true,
        "description": "The name of customer gateway device.",
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
      "ip_address": {
        "description": "IPv4 address for the customer gateway device's outside interface. The address must be static. If ` + "`" + `` + "`" + `OutsideIpAddressType` + "`" + `` + "`" + ` in your VPN connection options is set to ` + "`" + `` + "`" + `PrivateIpv4` + "`" + `` + "`" + `, you can use an RFC6598 or RFC1918 private IPv4 address. If ` + "`" + `` + "`" + `OutsideIpAddressType` + "`" + `` + "`" + ` is set to ` + "`" + `` + "`" + `PublicIpv4` + "`" + `` + "`" + `, you can use a public IPv4 address.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags for the customer gateway.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "type": {
        "description": "The type of VPN connection that this customer gateway supports (` + "`" + `` + "`" + `ipsec.1` + "`" + `` + "`" + `).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Specifies a customer gateway.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2CustomerGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2CustomerGateway), &result)
	return &result
}
