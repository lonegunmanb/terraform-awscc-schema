package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccScnDataIntegrationFlow = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the DataIntegrationFlow.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The creation time of the DataIntegrationFlow.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The Amazon Web Services Supply Chain instance identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The last modified time of the DataIntegrationFlow.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the DataIntegrationFlow.",
        "description_kind": "plain",
        "type": "string"
      },
      "sources": {
        "computed": true,
        "description": "The source configurations for the DataIntegrationFlow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dataset_source": {
              "computed": true,
              "description": "The dataset source configuration parameters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dataset_identifier": {
                    "computed": true,
                    "description": "The ARN of the dataset.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "options": {
                    "computed": true,
                    "description": "The dataset options.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dedupe_records": {
                          "computed": true,
                          "description": "The option to perform deduplication on data records sharing same primary key values.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "dedupe_strategy": {
                          "computed": true,
                          "description": "The deduplication strategy.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_priority": {
                                "computed": true,
                                "description": "The field priority deduplication strategy configuration.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "fields": {
                                      "computed": true,
                                      "description": "The list of field names and their sort order for deduplication.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "The name of the deduplication field.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "sort_order": {
                                            "computed": true,
                                            "description": "The sort order.",
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
                              "type": {
                                "computed": true,
                                "description": "The deduplication strategy type.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "load_type": {
                          "computed": true,
                          "description": "The load type.",
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
            "s3_source": {
              "computed": true,
              "description": "The S3 source configuration parameters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "computed": true,
                    "description": "The S3 bucket name.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "options": {
                    "computed": true,
                    "description": "The Amazon S3 options.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "file_type": {
                          "computed": true,
                          "description": "The file type.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "prefix": {
                    "computed": true,
                    "description": "The S3 prefix.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "source_name": {
              "computed": true,
              "description": "The source name that can be used as table alias in SQL transformation query.",
              "description_kind": "plain",
              "type": "string"
            },
            "source_type": {
              "computed": true,
              "description": "The source type.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description": "The tags for the DataIntegrationFlow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "target": {
        "computed": true,
        "description": "The DataIntegrationFlow target parameters.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dataset_target": {
              "computed": true,
              "description": "The dataset target configuration parameters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dataset_identifier": {
                    "computed": true,
                    "description": "The dataset ARN.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "options": {
                    "computed": true,
                    "description": "The dataset options.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dedupe_records": {
                          "computed": true,
                          "description": "The option to perform deduplication on data records sharing same primary key values.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "dedupe_strategy": {
                          "computed": true,
                          "description": "The deduplication strategy.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_priority": {
                                "computed": true,
                                "description": "The field priority deduplication strategy configuration.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "fields": {
                                      "computed": true,
                                      "description": "The list of field names and their sort order for deduplication.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "The name of the deduplication field.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "sort_order": {
                                            "computed": true,
                                            "description": "The sort order.",
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
                              "type": {
                                "computed": true,
                                "description": "The deduplication strategy type.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "load_type": {
                          "computed": true,
                          "description": "The load type.",
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
            "target_type": {
              "computed": true,
              "description": "The target type.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "transformation": {
        "computed": true,
        "description": "The DataIntegrationFlow transformation parameters.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sql_transformation": {
              "computed": true,
              "description": "The SQL transformation configuration parameters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "query": {
                    "computed": true,
                    "description": "The transformation SQL query body based on SparkSQL.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "transformation_type": {
              "computed": true,
              "description": "The transformation type.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::SCN::DataIntegrationFlow",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccScnDataIntegrationFlowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccScnDataIntegrationFlow), &result)
	return &result
}
