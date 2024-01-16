package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCustomerprofilesIntegration = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time of this integration got created",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The unique name of the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "flow_definition": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "flow_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "kms_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_flow_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connector_profile_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "connector_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "incremental_pull_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "datetime_type_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "source_connector_properties": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "marketo": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "object": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "s3": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "salesforce": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "enable_dynamic_field_update": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "include_deleted_records": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "object": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "service_now": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "object": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "zendesk": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "object": {
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
                "nesting_mode": "single"
              }
            },
            "tasks": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connector_operator": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "marketo": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "salesforce": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "service_now": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "zendesk": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "destination_field": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source_fields": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "task_properties": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "operator_property_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "property": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "task_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "trigger_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "trigger_properties": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "scheduled": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_pull_mode": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "first_execution_from": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "schedule_end_time": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "schedule_expression": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "schedule_offset": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "schedule_start_time": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "timezone": {
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
                  "trigger_type": {
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
      "last_updated_at": {
        "computed": true,
        "description": "The time of this integration got last updated at",
        "description_kind": "plain",
        "type": "string"
      },
      "object_type_name": {
        "computed": true,
        "description": "The name of the ObjectType defined for the 3rd party data in Profile Service",
        "description_kind": "plain",
        "type": "string"
      },
      "object_type_names": {
        "computed": true,
        "description": "The mapping between 3rd party event types and ObjectType names",
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
      "tags": {
        "computed": true,
        "description": "The tags (keys and values) associated with the integration",
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
      "uri": {
        "computed": true,
        "description": "The URI of the S3 bucket or any other type of data source.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CustomerProfiles::Integration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCustomerprofilesIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCustomerprofilesIntegration), &result)
	return &result
}
