package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApsWorkspace = `{
  "block": {
    "attributes": {
      "alert_manager_definition": {
        "computed": true,
        "description": "The AMP Workspace alert manager definition data",
        "description_kind": "plain",
        "type": "string"
      },
      "alias": {
        "computed": true,
        "description": "AMP Workspace alias.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Workspace arn.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "KMS Key ARN used to encrypt and decrypt AMP workspace data.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging_configuration": {
        "computed": true,
        "description": "Logging configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_group_arn": {
              "computed": true,
              "description": "CloudWatch log group ARN",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "prometheus_endpoint": {
        "computed": true,
        "description": "AMP Workspace prometheus endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "query_logging_configuration": {
        "computed": true,
        "description": "Query logging configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destinations": {
              "computed": true,
              "description": "The destinations configuration for query logging",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_logs": {
                    "computed": true,
                    "description": "Represents a cloudwatch logs destination for query logging",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "log_group_arn": {
                          "computed": true,
                          "description": "The ARN of the CloudWatch Logs log group",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "filters": {
                    "computed": true,
                    "description": "Filters for logging",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "qsp_threshold": {
                          "computed": true,
                          "description": "Query logs with QSP above this limit are vended",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "workspace_configuration": {
        "computed": true,
        "description": "Workspace configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "limits_per_label_sets": {
              "computed": true,
              "description": "An array of label set and associated limits",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "label_set": {
                    "computed": true,
                    "description": "An array of series labels",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "Name of the label",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "Value of the label",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
                    }
                  },
                  "limits": {
                    "computed": true,
                    "description": "Limits that can be applied to a label set",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_series": {
                          "computed": true,
                          "description": "The maximum number of active series that can be ingested for this label set",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "set"
              }
            },
            "retention_period_in_days": {
              "computed": true,
              "description": "How many days that metrics are retained in the workspace",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "workspace_id": {
        "computed": true,
        "description": "Required to identify a specific APS Workspace.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::APS::Workspace",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApsWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApsWorkspace), &result)
	return &result
}
