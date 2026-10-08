package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectIntegrationAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "description": "Amazon Connect instance identifier",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integration_arn": {
        "description": "ARN of Integration being associated with the instance",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integration_association_id": {
        "computed": true,
        "description": "Identifier of the association with Connect Instance",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_type": {
        "description": "Specifies the integration type to be associated with the instance",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_application_name": {
        "computed": true,
        "description": "The name of the external application. This is only supported for the EVENT integration type",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_application_url": {
        "computed": true,
        "description": "The URL for the external application. This is only supported for the EVENT integration type",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_type": {
        "computed": true,
        "description": "The type of the data source. This is only supported for the EVENT integration type",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags used to organize, track, or control access for this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
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
    "description": "Resource Type definition for AWS::Connect::IntegrationAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectIntegrationAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectIntegrationAssociation), &result)
	return &result
}
