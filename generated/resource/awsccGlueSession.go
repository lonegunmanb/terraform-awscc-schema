package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueSession = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the session.",
        "description_kind": "plain",
        "type": "string"
      },
      "command": {
        "description": "The SessionCommand that runs the job.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "Specifies the name of the SessionCommand. Can be 'glueetl' or 'gluestreaming'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "python_version": {
              "computed": true,
              "description": "Specifies the Python version. The Python version indicates the version supported for jobs of type Spark.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "connections": {
        "computed": true,
        "description": "Specifies the connections used by the session.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connections": {
              "computed": true,
              "description": "A list of connection names used by the session.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "created_on": {
        "computed": true,
        "description": "The time and date when the session was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_arguments": {
        "computed": true,
        "description": "A map array of key-value pairs. Max is 75 pairs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description": "The description of the session.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "glue_version": {
        "computed": true,
        "description": "The Glue version determines the versions of Apache Spark and Python that Glue supports. The GlueVersion must be greater than 2.0.",
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
      "idle_timeout": {
        "computed": true,
        "description": "The number of minutes when idle before session times out. Default is the value of Timeout.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_capacity": {
        "computed": true,
        "description": "The number of Glue data processing units (DPUs) that can be allocated when the job runs.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "number_of_workers": {
        "computed": true,
        "description": "The number of workers of a defined WorkerType to use for the session.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "progress": {
        "computed": true,
        "description": "The code execution progress of the session.",
        "description_kind": "plain",
        "type": "number"
      },
      "request_origin": {
        "computed": true,
        "description": "The origin of the request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role": {
        "description": "The IAM Role ARN.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_configuration": {
        "computed": true,
        "description": "The name of the SecurityConfiguration structure to be used with the session.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "session_id": {
        "description": "The ID of the session.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The session status.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags belonging to the session.",
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
      },
      "timeout": {
        "computed": true,
        "description": "The number of minutes before session times out.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "worker_type": {
        "computed": true,
        "description": "The type of predefined worker that is allocated when a session runs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Glue::Session. Sessions provide an on-demand, serverless Apache Spark runtime environment for building, testing, and running data preparation and analytics applications.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueSessionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueSession), &result)
	return &result
}
