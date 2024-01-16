package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the template.\u003c/p\u003e",
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
        "description": "\u003cp\u003eTime when this was created.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "\u003cp\u003eTime when this was last updated.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "\u003cp\u003eA display name for the template.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description": "\u003cp\u003eA list of resource permissions to be set on the template. \u003c/p\u003e",
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
        "description": "\u003cp\u003eThe source entity of the template.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_analysis": {
              "computed": true,
              "description": "\u003cp\u003eThe source analysis of the template.\u003c/p\u003e",
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
                    "description": "\u003cp\u003eA structure containing information about the dataset references used as placeholders\n            in the template.\u003c/p\u003e",
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
            },
            "source_template": {
              "computed": true,
              "description": "\u003cp\u003eThe source template of the template.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
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
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "\u003cp\u003eContains a map of the key-value pairs for the resource tag or tags assigned to the resource.\u003c/p\u003e",
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
      "template_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "\u003cp\u003eA version of a template.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "created_time": {
              "computed": true,
              "description": "\u003cp\u003eThe time that this template version was created.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "data_set_configurations": {
              "computed": true,
              "description": "\u003cp\u003eSchema of the dataset identified by the placeholder. Any dashboard created from this\n            template should be bound to new datasets matching the same schema described through this\n            API operation.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "column_group_schema_list": {
                    "computed": true,
                    "description": "\u003cp\u003eA structure containing the list of column group schemas.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "column_group_column_schema_list": {
                          "computed": true,
                          "description": "\u003cp\u003eA structure containing the list of schemas for column group columns.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "computed": true,
                                "description": "\u003cp\u003eThe name of the column group's column schema.\u003c/p\u003e",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "name": {
                          "computed": true,
                          "description": "\u003cp\u003eThe name of the column group schema.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "data_set_schema": {
                    "computed": true,
                    "description": "\u003cp\u003eDataset schema.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "column_schema_list": {
                          "computed": true,
                          "description": "\u003cp\u003eA structure containing the list of column schemas.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_type": {
                                "computed": true,
                                "description": "\u003cp\u003eThe data type of the column schema.\u003c/p\u003e",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "geographic_role": {
                                "computed": true,
                                "description": "\u003cp\u003eThe geographic role of the column schema.\u003c/p\u003e",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "\u003cp\u003eThe name of the column schema.\u003c/p\u003e",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "placeholder": {
                    "computed": true,
                    "description": "\u003cp\u003ePlaceholder.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "description": {
              "computed": true,
              "description": "\u003cp\u003eThe description of the template.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "errors": {
              "computed": true,
              "description": "\u003cp\u003eErrors associated with this template version.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "message": {
                    "computed": true,
                    "description": "\u003cp\u003eDescription of the error type.\u003c/p\u003e",
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
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of an analysis or template that was used to create this\n            template.\u003c/p\u003e",
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
              "description": "\u003cp\u003eThe ARN of the theme associated with this version of the template.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "version_number": {
              "computed": true,
              "description": "\u003cp\u003eThe version number of the template version.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "version_description": {
        "computed": true,
        "description": "\u003cp\u003eA description of the current template version being created. This API operation creates the\n\t\t\tfirst version of the template. Every time \u003ccode\u003eUpdateTemplate\u003c/code\u003e is called, a new\n\t\t\tversion is created. Each version of the template maintains a description of the version\n\t\t\tin the \u003ccode\u003eVersionDescription\u003c/code\u003e field.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of the AWS::QuickSight::Template Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQuicksightTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightTemplate), &result)
	return &result
}
