package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchOtelEnrichment = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The AWS account ID. This is the primary identifier for this singleton resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "exclude_filters": {
        "computed": true,
        "description": "Removes metrics from the include set. Absent or empty means nothing is removed. Evaluated after IncludeFilters, so ExcludeFilters always wins.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "metric_names": {
              "computed": true,
              "description": "Absent or empty means every metric in this namespace, in whichever direction this selector appears. Present means only these metric names.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "namespace": {
              "computed": true,
              "description": "The CloudWatch namespace this selector applies to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "include_filters": {
        "computed": true,
        "description": "Scopes enrichment to a subset of the account's telemetry. Absent or empty means all namespaces are in scope. Present means only these are. The service enforces a combined cap of 100 selectors across IncludeFilters and ExcludeFilters, and rejects more than one selector for the same namespace within a direction.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "metric_names": {
              "computed": true,
              "description": "Absent or empty means every metric in this namespace, in whichever direction this selector appears. Present means only these metric names.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "namespace": {
              "computed": true,
              "description": "The CloudWatch namespace this selector applies to.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "status": {
        "computed": true,
        "description": "Current status of OTel enrichment (RUNNING or STOPPED).",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "AWS::CloudWatch::OTelEnrichment enables OTel metric enrichment in CloudWatch, allowing CloudWatch vended metrics to be available for PromQL querying enriched with AWS resource tags and metadata.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchOtelEnrichmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchOtelEnrichment), &result)
	return &result
}
