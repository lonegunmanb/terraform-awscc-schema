package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKafkaconnectWorkerConfiguration = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A summary description of the worker configuration.",
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
      "name": {
        "description": "The name of the worker configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "properties_file_content": {
        "description": "Base64 encoded contents of connect-distributed.properties file.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revision": {
        "computed": true,
        "description": "The description of a revision of the worker configuration.",
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "worker_configuration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the custom configuration.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The configuration of the workers, which are the processes that run the connector logic.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKafkaconnectWorkerConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKafkaconnectWorkerConfiguration), &result)
	return &result
}
