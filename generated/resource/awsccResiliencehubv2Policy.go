package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccResiliencehubv2Policy = `{
  "block": {
    "attributes": {
      "associated_service_count": {
        "computed": true,
        "description": "The number of services associated with this policy.",
        "description_kind": "plain",
        "type": "number"
      },
      "availability_slo": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "target": {
              "computed": true,
              "description": "Availability target percentage.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the policy was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_recovery": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "time_between_backups_in_minutes": {
              "computed": true,
              "description": "Time between backups in minutes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "description": {
        "computed": true,
        "description": "The description of the policy.",
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
      "kms_key_id": {
        "computed": true,
        "description": "The KMS key ID for encrypting policy data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_az": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "disaster_recovery_approach": {
              "computed": true,
              "description": "Multi-AZ disaster recovery approach.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rpo_in_minutes": {
              "computed": true,
              "description": "Recovery Point Objective in minutes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rto_in_minutes": {
              "computed": true,
              "description": "Recovery Time Objective in minutes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "multi_region": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "disaster_recovery_approach": {
              "computed": true,
              "description": "Multi-Region disaster recovery approach.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rpo_in_minutes": {
              "computed": true,
              "description": "Recovery Point Objective in minutes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rto_in_minutes": {
              "computed": true,
              "description": "Recovery Time Objective in minutes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "name": {
        "description": "The name of the policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_arn": {
        "computed": true,
        "description": "The ARN of the policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags assigned to the policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the policy was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates a resilience policy that defines availability and disaster recovery requirements.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccResiliencehubv2PolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResiliencehubv2Policy), &result)
	return &result
}
