package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQbusinessDataSource = `{
  "block": {
    "attributes": {
      "application_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "configuration": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "document_enrichment_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "inline_configurations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "condition": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "operator": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "date_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
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
                  "document_content_operator": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "attribute_value_operator": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "date_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
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
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "post_extraction_hook_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "invocation_condition": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "operator": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "date_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
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
                  "lambda_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_bucket_name": {
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
            "pre_extraction_hook_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "invocation_condition": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "operator": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "date_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
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
                  "lambda_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_bucket_name": {
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
      "index_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "media_extraction_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "image_extraction_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "image_extraction_status": {
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
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sync_schedule": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
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
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::QBusiness::DataSource Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQbusinessDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQbusinessDataSource), &result)
	return &result
}
