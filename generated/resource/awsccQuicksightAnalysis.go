package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightAnalysis = `{
  "block": {
    "attributes": {
      "analysis_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the analysis.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "\u003cp\u003eThe time that the analysis was created.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "data_set_arns": {
        "computed": true,
        "description": "\u003cp\u003eThe ARNs of the datasets of the analysis.\u003c/p\u003e",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "errors": {
        "computed": true,
        "description": "\u003cp\u003eErrors associated with the analysis.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "message": {
              "computed": true,
              "description": "\u003cp\u003eThe message associated with the analysis error.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "\u003cp\u003eThe time that the analysis was last updated.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "\u003cp\u003eThe descriptive name of the analysis.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description": "\u003cp\u003eA list of QuickSight parameters and the list's override values.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "date_time_parameters": {
              "computed": true,
              "description": "\u003cp\u003eDate-time parameters.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "\u003cp\u003eA display name for the date-time parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description": "\u003cp\u003eThe values for the date-time parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "decimal_parameters": {
              "computed": true,
              "description": "\u003cp\u003eDecimal parameters.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "\u003cp\u003eA display name for the decimal parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description": "\u003cp\u003eThe values for the decimal parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "integer_parameters": {
              "computed": true,
              "description": "\u003cp\u003eInteger parameters.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "\u003cp\u003eThe name of the integer parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description": "\u003cp\u003eThe values for the integer parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "string_parameters": {
              "computed": true,
              "description": "\u003cp\u003eString parameters.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "\u003cp\u003eA display name for a string parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description": "\u003cp\u003eThe values of a string parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "permissions": {
        "computed": true,
        "description": "\u003cp\u003eA structure that describes the principals and the resource-level permissions on an\n            analysis. You can use the \u003ccode\u003ePermissions\u003c/code\u003e structure to grant permissions by\n            providing a list of AWS Identity and Access Management (IAM) action information for each\n            principal listed by Amazon Resource Name (ARN). \u003c/p\u003e\n\n        \u003cp\u003eTo specify no permissions, omit \u003ccode\u003ePermissions\u003c/code\u003e.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "computed": true,
              "description": "\u003cp\u003eThe IAM action to grant or revoke permissions on.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "principal": {
              "computed": true,
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:\u003c/p\u003e\n        \u003cul\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an AWS account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across AWS accounts.\n                    (This is less common.) \u003c/p\u003e\n            \u003c/li\u003e\n         \u003c/ul\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "sheets": {
        "computed": true,
        "description": "\u003cp\u003eA list of the associated sheets with the unique identifier and name of each sheet.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "\u003cp\u003eThe name of a sheet. This name is displayed on the sheet's tab in the QuickSight\n            console.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "sheet_id": {
              "computed": true,
              "description": "\u003cp\u003eThe unique identifier associated with a sheet.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "source_entity": {
        "description": "\u003cp\u003eThe source entity of an analysis.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_template": {
              "computed": true,
              "description": "\u003cp\u003eThe source template of an analysis.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the source template of an analysis.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_set_references": {
                    "computed": true,
                    "description": "\u003cp\u003eThe dataset references of the source template of an analysis.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_set_arn": {
                          "computed": true,
                          "description": "\u003cp\u003eDataset Amazon Resource Name (ARN).\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "data_set_placeholder": {
                          "computed": true,
                          "description": "\u003cp\u003eDataset placeholder.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "\u003cp\u003eContains a map of the key-value pairs for the resource tag or tags assigned to the\n            analysis.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "\u003cp\u003eTag key.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "theme_arn": {
        "computed": true,
        "description": "\u003cp\u003eThe ARN of the theme of the analysis.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of the AWS::QuickSight::Analysis Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQuicksightAnalysisSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightAnalysis), &result)
	return &result
}
