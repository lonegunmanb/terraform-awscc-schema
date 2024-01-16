package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodegurureviewerRepositoryAssociation = `{
  "block": {
    "attributes": {
      "association_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the repository association.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_name": {
        "computed": true,
        "description": "The name of the S3 bucket associated with an associated S3 repository. It must start with ` + "`" + `codeguru-reviewer-` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the repository to be associated.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description": "The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with a repository association.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description": "The type of repository to be associated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CodeGuruReviewer::RepositoryAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodegurureviewerRepositoryAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodegurureviewerRepositoryAssociation), &result)
	return &result
}
