package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccObservabilityadminTelemetryEnrichment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "scope": {
        "computed": true,
        "description": "Scope of the Telemetry Enrichment",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Current status of the resource tags for telemetry feature (Running, Stopped, or Impaired).",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "AWS::ObservabilityAdmin::TelemetryEnrichment cloudformation resource enables the resource tags for telemetry feature in CloudWatch to enrich infrastructure metrics with AWS resource tags. For more details: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/resource-tags-for-telemetry.html",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccObservabilityadminTelemetryEnrichmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccObservabilityadminTelemetryEnrichment), &result)
	return &result
}
