package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGameliftScript = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift script resource and uniquely identifies it. ARNs are unique across all Regions. In a GameLift script ARN, the resource ID matches the Id value.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example \"1469498468.057\").",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A descriptive label that is associated with a script. Script names do not need to be unique.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "script_id": {
        "computed": true,
        "description": "A unique identifier for the Realtime script",
        "description_kind": "plain",
        "type": "string"
      },
      "size_on_disk": {
        "computed": true,
        "description": "The file size of the uploaded Realtime script, expressed in bytes. When files are uploaded from an S3 location, this value remains at \"0\".",
        "description_kind": "plain",
        "type": "number"
      },
      "storage_location": {
        "description": "The location of the Amazon S3 bucket where a zipped file containing your Realtime scripts is stored. The storage location must specify the Amazon S3 bucket name, the zip file name (the \"key\"), and a role ARN that allows Amazon GameLift to access the Amazon S3 storage location. The S3 bucket must be in the same Region where you want to create a new script. By default, Amazon GameLift uploads the latest version of the zip file; if you have S3 object versioning turned on, you can use the ObjectVersion parameter to specify an earlier version.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket": {
              "description": "An Amazon S3 bucket identifier. This is the name of the S3 bucket.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key": {
              "description": "The name of the zip file that contains the script files.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "object_version": {
              "computed": true,
              "description": "The version of the file, if object versioning is turned on for the bucket. Amazon GameLift uses this information when retrieving files from your S3 bucket. To retrieve a specific version of the file, provide an object version. To retrieve the latest version of the file, do not set this parameter.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description": "The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access the S3 bucket.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "version": {
        "computed": true,
        "description": "The version that is associated with a script. Version strings do not need to be unique.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The AWS::GameLift::Script resource creates a new script record for your Realtime Servers script. Realtime scripts are JavaScript that provide configuration settings and optional custom game logic for your game. The script is deployed when you create a Realtime Servers fleet to host your game sessions. Script logic is executed during an active game session.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGameliftScriptSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGameliftScript), &result)
	return &result
}
