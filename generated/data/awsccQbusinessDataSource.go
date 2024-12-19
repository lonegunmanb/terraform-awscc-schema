package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQbusinessDataSource = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description_kind": "plain",
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
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
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
                          "type": "string"
                        },
                        "operator": {
                          "computed": true,
                          "description_kind": "plain",
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
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "document_content_operator": {
                    "computed": true,
                    "description_kind": "plain",
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
                          "type": "string"
                        },
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
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
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
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
                          "type": "string"
                        },
                        "operator": {
                          "computed": true,
                          "description_kind": "plain",
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
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "lambda_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_bucket_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
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
                          "type": "string"
                        },
                        "operator": {
                          "computed": true,
                          "description_kind": "plain",
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
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "lambda_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_bucket_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "index_id": {
        "computed": true,
        "description_kind": "plain",
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
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
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
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::QBusiness::DataSource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQbusinessDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQbusinessDataSource), &result)
	return &result
}
