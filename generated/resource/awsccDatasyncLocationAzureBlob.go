package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationAzureBlob = `{
  "block": {
    "attributes": {
      "agent_arns": {
        "description": "The Amazon Resource Names (ARNs) of agents to use for an Azure Blob Location.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "azure_access_tier": {
        "computed": true,
        "description": "Specifies an access tier for the objects you're transferring into your Azure Blob Storage container.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "azure_blob_authentication_type": {
        "computed": true,
        "description": "The specific authentication type that you want DataSync to use to access your Azure Blob Container.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "azure_blob_container_url": {
        "computed": true,
        "description": "The URL of the Azure Blob container that was described.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "azure_blob_sas_configuration": {
        "computed": true,
        "description": "Specifies the shared access signature (SAS) that DataSync uses to access your Azure Blob Storage container.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "azure_blob_sas_token": {
              "computed": true,
              "description": "Specifies the shared access signature (SAS) token, which indicates the permissions DataSync needs to access your Azure Blob Storage container.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "azure_blob_type": {
        "computed": true,
        "description": "Specifies a blob type for the objects you're transferring into your Azure Blob Storage container.",
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
      "location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Azure Blob Location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the Azure Blob Location that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "subdirectory": {
        "computed": true,
        "description": "The subdirectory in the Azure Blob Container that is used to read data from the Azure Blob Source Location.",
        "description_kind": "plain",
        "optional": true,
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
              "description": "The key for an AWS resource tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for an AWS resource tag.",
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
    "description": "Resource schema for AWS::DataSync::LocationAzureBlob.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatasyncLocationAzureBlobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationAzureBlob), &result)
	return &result
}
