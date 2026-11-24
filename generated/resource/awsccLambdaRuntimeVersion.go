package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaRuntimeVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the runtime version.",
        "description_kind": "plain",
        "type": "string"
      },
      "compatibility": {
        "computed": true,
        "description": "Compatibility matrix for the runtime version: supported integration models, architectures, and feature policies.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "architectures": {
              "computed": true,
              "description": "Supported CPU architectures for the runtime version.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "integration_models": {
              "computed": true,
              "description": "Supported integration models for the runtime version.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "managed_instances": {
              "computed": true,
              "description": "Whether the runtime version supports managed instances.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "created": {
        "computed": true,
        "description": "The timestamp when the runtime version was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "image_uri": {
        "computed": true,
        "description": "The URI of the OCI image to publish as this runtime version. Not returned in Read.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified": {
        "computed": true,
        "description": "The timestamp when the runtime version was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The version identifier (1-128 chars, alphanumeric plus _.-).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "platforms": {
        "computed": true,
        "description": "The OCI platforms resolved from the image manifest.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "architecture": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "digest": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "os": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "release_notes": {
        "computed": true,
        "description": "Optional release notes for this runtime version. Updatable; removing it from the template clears the release notes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resolved_uri": {
        "computed": true,
        "description": "The fully resolved (digest-pinned) image URI for this runtime version.",
        "description_kind": "plain",
        "type": "string"
      },
      "runtime_name": {
        "description": "The name of the parent runtime.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The lifecycle state of the runtime version.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lambda::RuntimeVersion. An immutable, OCI-image-backed version of an AWS::Lambda::Runtime.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaRuntimeVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaRuntimeVersion), &result)
	return &result
}
