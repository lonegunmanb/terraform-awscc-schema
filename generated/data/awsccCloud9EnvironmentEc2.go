package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloud9EnvironmentEc2 = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the development environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "automatic_stop_time_minutes": {
        "computed": true,
        "description": "The number of minutes until the running instance is shut down after the environment was last used.",
        "description_kind": "plain",
        "type": "number"
      },
      "connection_type": {
        "computed": true,
        "description": "The connection type used for connecting to an Amazon EC2 environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "The ID of the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_id": {
        "computed": true,
        "description": "The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description": "The type of instance to connect to the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the environment owner.",
        "description_kind": "plain",
        "type": "string"
      },
      "repositories": {
        "computed": true,
        "description": "Any AWS CodeCommit source code repositories to be cloned into the development environment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "path_component": {
              "computed": true,
              "description": "The path within the development environment's default file system location to clone the AWS CodeCommit repository into.",
              "description_kind": "plain",
              "type": "string"
            },
            "repository_url": {
              "computed": true,
              "description": "The clone URL of the AWS CodeCommit repository to be cloned.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "subnet_id": {
        "computed": true,
        "description": "The ID of the subnet in Amazon VPC that AWS Cloud9 will use.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs that will be associated with the new AWS Cloud9 development environment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Cloud9::EnvironmentEC2",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloud9EnvironmentEc2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloud9EnvironmentEc2), &result)
	return &result
}
