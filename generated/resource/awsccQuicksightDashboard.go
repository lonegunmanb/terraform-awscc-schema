package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightDashboard = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
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
        "description": "\u003cp\u003eThe time that this dataset was created.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dashboard_publish_options": {
        "computed": true,
        "description": "\u003cp\u003eDashboard publish options.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ad_hoc_filtering_option": {
              "computed": true,
              "description": "\u003cp\u003eAd hoc (one-time) filtering option.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "export_to_csv_option": {
              "computed": true,
              "description": "\u003cp\u003eExport to .csv option.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "sheet_controls_option": {
              "computed": true,
              "description": "\u003cp\u003eSheet controls option.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "visibility_state": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_published_time": {
        "computed": true,
        "description": "\u003cp\u003eThe last time that this dataset was published.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "\u003cp\u003eThe last time that this dataset was updated.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "\u003cp\u003eThe display name of the dashboard.\u003c/p\u003e",
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
                    "description": "\u003cp\u003eA display name for the date-time parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description": "\u003cp\u003eThe values for the date-time parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
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
                    "description": "\u003cp\u003eA display name for the decimal parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description": "\u003cp\u003eThe values for the decimal parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
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
                    "description": "\u003cp\u003eThe name of the integer parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description": "\u003cp\u003eThe values for the integer parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
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
                    "description": "\u003cp\u003eA display name for a string parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
                    "description": "\u003cp\u003eThe values of a string parameter.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
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
        "description": "\u003cp\u003eA structure that contains the permissions of the dashboard. You can use this structure\n            for granting permissions by providing a list of IAM action information for each\n            principal ARN. \u003c/p\u003e\n\n        \u003cp\u003eTo specify no permissions, omit the permissions list.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "description": "\u003cp\u003eThe IAM action to grant or revoke permissions on.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "principal": {
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:\u003c/p\u003e\n        \u003cul\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an AWS account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across AWS accounts.\n                    (This is less common.) \u003c/p\u003e\n            \u003c/li\u003e\n         \u003c/ul\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "source_entity": {
        "description": "\u003cp\u003eDashboard source entity.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_template": {
              "computed": true,
              "description": "\u003cp\u003eDashboard source template.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "data_set_references": {
                    "description": "\u003cp\u003eDataset references.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_set_arn": {
                          "description": "\u003cp\u003eDataset Amazon Resource Name (ARN).\u003c/p\u003e",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "data_set_placeholder": {
                          "description": "\u003cp\u003eDataset placeholder.\u003c/p\u003e",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "required": true
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
      "tags": {
        "computed": true,
        "description": "\u003cp\u003eContains a map of the key-value pairs for the resource tag or tags assigned to the\n            dashboard.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "\u003cp\u003eTag key.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "theme_arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the theme that is being used for this dashboard. If\n            you add a value for this field, it overrides the value that is used in the source\n            entity. The theme ARN must exist in the same AWS account where you create the\n            dashboard.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "\u003cp\u003eDashboard version.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "created_time": {
              "computed": true,
              "description": "\u003cp\u003eThe time that this dashboard version was created.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "data_set_arns": {
              "computed": true,
              "description": "\u003cp\u003eThe Amazon Resource Numbers (ARNs) for the datasets that are associated with this\n            version of the dashboard.\u003c/p\u003e",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "description": {
              "computed": true,
              "description": "\u003cp\u003eDescription.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "errors": {
              "computed": true,
              "description": "\u003cp\u003eErrors associated with this dashboard version.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "message": {
                    "computed": true,
                    "description": "\u003cp\u003eMessage.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
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
            "source_entity_arn": {
              "computed": true,
              "description": "\u003cp\u003eSource entity ARN.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "theme_arn": {
              "computed": true,
              "description": "\u003cp\u003eThe ARN of the theme associated with a version of the dashboard.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "version_number": {
              "computed": true,
              "description": "\u003cp\u003eVersion number for this version of the dashboard.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "version_description": {
        "computed": true,
        "description": "\u003cp\u003eA description for the first version of the dashboard being created.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of the AWS::QuickSight::Dashboard Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQuicksightDashboardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightDashboard), &result)
	return &result
}
