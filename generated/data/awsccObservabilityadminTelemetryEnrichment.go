package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccObservabilityadminTelemetryEnrichment = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope": {
        "computed": true,
        "description": "Scope of the Telemetry Enrichment",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Current status of the resource tags for telemetry feature (Running, Stopped, or Impaired).",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ObservabilityAdmin::TelemetryEnrichment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccObservabilityadminTelemetryEnrichmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccObservabilityadminTelemetryEnrichment), &result)
	return &result
}
