package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IpamInternetRegistryAssociation = `{
  "block": {
    "attributes": {
      "child_request_xml": {
        "computed": true,
        "description": "The XML content for the child request to be submitted to the internet registry to complete the BPKI setup.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the internet registry association.",
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
      "ipam_id": {
        "description": "The ID of the IPAM to associate with the internet registry.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipam_internet_registry_association_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the internet registry association.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_internet_registry_association_id": {
        "computed": true,
        "description": "The ID of the internet registry association.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_region": {
        "computed": true,
        "description": "The Amazon Web Services Region of the IPAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "organization_handle": {
        "description": "The organization handle at the internet registry (for example, a RIPE NCC organization ID or ARIN Org ID).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "The ID of the Amazon Web Services account that owns the internet registry association.",
        "description_kind": "plain",
        "type": "string"
      },
      "rir": {
        "description": "The Regional Internet Registry to associate with. The service returns these values in upper case, so the enum uses the round-tripping form.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the internet registry association.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the internet registry association.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
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
    "description": "An association between an IPAM and a Regional Internet Registry (RIR) for delegated Resource Public Key Infrastructure (RPKI) management. Use it to have IPAM create Route Origin Authorizations (ROAs) for IP address prefixes registered with the internet registry. The IPAM must be in the Advanced tier.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2IpamInternetRegistryAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpamInternetRegistryAssociation), &result)
	return &result
}
