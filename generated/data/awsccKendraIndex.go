package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKendraIndex = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_units": {
        "computed": true,
        "description": "Capacity units",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "query_capacity_units": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "storage_capacity_units": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A description for the index",
        "description_kind": "plain",
        "type": "string"
      },
      "document_metadata_configurations": {
        "computed": true,
        "description": "Document metadata configurations",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "relevance": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "duration": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "freshness": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "importance": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "rank_order": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value_importance_items": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "search": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "displayable": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "facetable": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "searchable": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "sortable": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
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
      "edition": {
        "computed": true,
        "description": "Edition of index",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of index",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "Role Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "server_side_encryption_configuration": {
        "computed": true,
        "description": "Server side encryption configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "Tags for labeling the index",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string used to identify this tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string containing the value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "user_context_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_token_configurations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "json_token_type_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "group_attribute_field": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "user_name_attribute_field": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "jwt_token_type_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "claim_regex": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "group_attribute_field": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "issuer": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_location": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secret_manager_arn": {
                    "computed": true,
                    "description": "Role Arn",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "user_name_attribute_field": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
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
    "description": "Data Source schema for AWS::Kendra::Index",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKendraIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKendraIndex), &result)
	return &result
}
