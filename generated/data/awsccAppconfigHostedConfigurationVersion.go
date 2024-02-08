package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppconfigHostedConfigurationVersion = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The application ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_profile_id": {
        "computed": true,
        "description": "The configuration profile ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "content": {
        "computed": true,
        "description": "The content of the configuration or the configuration data.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_type": {
        "computed": true,
        "description": "A standard MIME type describing the format of the configuration content.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the hosted configuration version.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "latest_version_number": {
        "computed": true,
        "description": "An optional locking token used to prevent race conditions from overwriting configuration updates when creating a new version. To ensure your data is not overwritten when creating multiple hosted configuration versions in rapid succession, specify the version number of the latest hosted configuration version.",
        "description_kind": "plain",
        "type": "number"
      },
      "version_label": {
        "computed": true,
        "description": "A user-defined label for an AWS AppConfig hosted configuration version.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_number": {
        "computed": true,
        "description": "Current version number of hosted configuration version.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AppConfig::HostedConfigurationVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppconfigHostedConfigurationVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppconfigHostedConfigurationVersion), &result)
	return &result
}
