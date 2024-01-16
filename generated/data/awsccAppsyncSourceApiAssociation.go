package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncSourceApiAssociation = `{
  "block": {
    "attributes": {
      "association_arn": {
        "computed": true,
        "description": "ARN of the SourceApiAssociation.",
        "description_kind": "plain",
        "type": "string"
      },
      "association_id": {
        "computed": true,
        "description": "Id of the SourceApiAssociation.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the SourceApiAssociation.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_successful_merge_date": {
        "computed": true,
        "description": "Date of last schema successful merge.",
        "description_kind": "plain",
        "type": "string"
      },
      "merged_api_arn": {
        "computed": true,
        "description": "ARN of the Merged API in the association.",
        "description_kind": "plain",
        "type": "string"
      },
      "merged_api_id": {
        "computed": true,
        "description": "GraphQLApiId of the Merged API in the association.",
        "description_kind": "plain",
        "type": "string"
      },
      "merged_api_identifier": {
        "computed": true,
        "description": "Identifier of the Merged GraphQLApi to associate. It could be either GraphQLApi ApiId or ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "source_api_arn": {
        "computed": true,
        "description": "ARN of the source API in the association.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_api_association_config": {
        "computed": true,
        "description": "Customized configuration for SourceApiAssociation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "merge_type": {
              "computed": true,
              "description": "Configuration of the merged behavior for the association. For example when it could be auto or has to be manual.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "source_api_association_status": {
        "computed": true,
        "description": "Current status of SourceApiAssociation.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_api_association_status_detail": {
        "computed": true,
        "description": "Current SourceApiAssociation status details.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_api_id": {
        "computed": true,
        "description": "GraphQLApiId of the source API in the association.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_api_identifier": {
        "computed": true,
        "description": "Identifier of the Source GraphQLApi to associate. It could be either GraphQLApi ApiId or ARN",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AppSync::SourceApiAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppsyncSourceApiAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncSourceApiAssociation), &result)
	return &result
}
