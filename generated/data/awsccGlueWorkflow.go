package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueWorkflow = `{
  "block": {
    "attributes": {
      "default_run_properties": {
        "computed": true,
        "description": "A collection of properties to be used as part of each execution of the workflow",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the workflow",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_concurrent_runs": {
        "computed": true,
        "description": "You can use this parameter to prevent unwanted multiple updates to data, to control costs, or in some cases, to prevent exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "The name of the workflow representing the flow",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to use with this workflow.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Glue::Workflow",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueWorkflowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueWorkflow), &result)
	return &result
}
