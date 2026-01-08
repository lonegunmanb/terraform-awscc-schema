package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCasesLayout = `{
  "block": {
    "attributes": {
      "content": {
        "description": "Defines the layout structure and field organization for the case interface. Specifies which fields appear in the top panel and More Info tab, and their display order.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "basic": {
              "computed": true,
              "description": "Defines the field layout for the agent's case interface. Configures which fields appear in the top panel (immediately visible) and More Info tab (expandable section) of the case view, allowing customization of the agent experience.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "more_info": {
                    "computed": true,
                    "description": "Sections within a panel or tab of the page layout.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "sections": {
                          "computed": true,
                          "description": "Defines the sections within a panel or tab. Contains field groups that organize related fields together.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_group": {
                                "computed": true,
                                "description": "Consists of a group of fields and associated properties.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "fields": {
                                      "computed": true,
                                      "description": "An ordered list of fields to display in this group. The order determines the sequence in which fields appear in the agent interface. Each field is referenced by its unique field ID.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "id": {
                                            "computed": true,
                                            "description": "The unique identifier of a field.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      },
                                      "optional": true
                                    },
                                    "name": {
                                      "computed": true,
                                      "description": "A descriptive name for the field group. Helps organize related fields together in the layout interface.",
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
                            "nesting_mode": "list"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "top_panel": {
                    "computed": true,
                    "description": "Sections within a panel or tab of the page layout.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "sections": {
                          "computed": true,
                          "description": "Defines the sections within a panel or tab. Contains field groups that organize related fields together.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_group": {
                                "computed": true,
                                "description": "Consists of a group of fields and associated properties.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "fields": {
                                      "computed": true,
                                      "description": "An ordered list of fields to display in this group. The order determines the sequence in which fields appear in the agent interface. Each field is referenced by its unique field ID.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "id": {
                                            "computed": true,
                                            "description": "The unique identifier of a field.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      },
                                      "optional": true
                                    },
                                    "name": {
                                      "computed": true,
                                      "description": "A descriptive name for the field group. Helps organize related fields together in the layout interface.",
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
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "created_time": {
        "computed": true,
        "description": "The time at which the layout was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The unique identifier of the Cases domain.",
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
      "last_modified_time": {
        "computed": true,
        "description": "The time at which the layout was created or last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "layout_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the layout.",
        "description_kind": "plain",
        "type": "string"
      },
      "layout_id": {
        "computed": true,
        "description": "The unique identifier of the layout.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "A descriptive name for the layout. Must be unique within the Cases domain and should clearly indicate the layout's purpose and field organization.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags that you attach to this layout.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "A layout in the Cases domain. Layouts define the following configuration in the top section and More Info tab of the Cases user interface: Fields to display to the users and Field ordering.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCasesLayoutSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCasesLayout), &result)
	return &result
}
