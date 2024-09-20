package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectFlowEntitlement = `{
  "block": {
    "attributes": {
      "data_transfer_subscriber_fee_percent": {
        "computed": true,
        "description": "Percentage from 0-100 of the data transfer cost to be billed to the subscriber.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description": "A description of the entitlement.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encryption": {
        "computed": true,
        "description": "The type of encryption that will be used on the output that is associated with this entitlement.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "algorithm": {
              "computed": true,
              "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "constant_initialization_vector": {
              "computed": true,
              "description": "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_id": {
              "computed": true,
              "description": "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_type": {
              "computed": true,
              "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region": {
              "computed": true,
              "description": "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_id": {
              "computed": true,
              "description": "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "computed": true,
              "description": "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "entitlement_arn": {
        "computed": true,
        "description": "The ARN of the entitlement.",
        "description_kind": "plain",
        "type": "string"
      },
      "entitlement_status": {
        "computed": true,
        "description": " An indication of whether the entitlement is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flow_arn": {
        "description": "The ARN of the flow.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the entitlement.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subscribers": {
        "description": "The AWS account IDs that you want to share your content with. The receiving accounts (subscribers) will be allowed to create their own flow using your content as the source.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Resource schema for AWS::MediaConnect::FlowEntitlement",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectFlowEntitlementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectFlowEntitlement), &result)
	return &result
}
