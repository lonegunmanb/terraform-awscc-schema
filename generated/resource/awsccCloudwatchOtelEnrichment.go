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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
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
