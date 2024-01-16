package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFrauddetectorDetector = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the detector.",
        "description_kind": "plain",
        "type": "string"
      },
      "associated_models": {
        "computed": true,
        "description": "The models to associate with this detector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "created_time": {
        "computed": true,
        "description": "The time when the detector was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the detector.",
        "description_kind": "plain",
        "type": "string"
      },
      "detector_id": {
        "computed": true,
        "description": "The ID of the detector",
        "description_kind": "plain",
        "type": "string"
      },
      "detector_version_id": {
        "computed": true,
        "description": "The active version ID of the detector",
        "description_kind": "plain",
        "type": "string"
      },
      "detector_version_status": {
        "computed": true,
        "description": "The desired detector version status for the detector",
        "description_kind": "plain",
        "type": "string"
      },
      "event_type": {
        "computed": true,
        "description": "The event type to associate this detector with.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "The ARN of the event type.",
              "description_kind": "plain",
              "type": "string"
            },
            "created_time": {
              "computed": true,
              "description": "The time when the event type was created.",
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description of the event type.",
              "description_kind": "plain",
              "type": "string"
            },
            "entity_types": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "created_time": {
                    "computed": true,
                    "description": "The time when the entity type was created.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "The description.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "inline": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "last_updated_time": {
                    "computed": true,
                    "description": "The time when the entity type was last updated.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tags": {
                    "computed": true,
                    "description": "Tags associated with this entity type.",
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
                  }
                },
                "nesting_mode": "list"
              }
            },
            "event_variables": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "created_time": {
                    "computed": true,
                    "description": "The time when the event variable was created.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "data_source": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "data_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "default_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "The description.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "inline": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "last_updated_time": {
                    "computed": true,
                    "description": "The time when the event variable was last updated.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tags": {
                    "computed": true,
                    "description": "Tags associated with this event variable.",
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
                  "variable_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "inline": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "labels": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "created_time": {
                    "computed": true,
                    "description": "The time when the label was created.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "The description.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "inline": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "last_updated_time": {
                    "computed": true,
                    "description": "The time when the label was last updated.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tags": {
                    "computed": true,
                    "description": "Tags associated with this label.",
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
                  }
                },
                "nesting_mode": "list"
              }
            },
            "last_updated_time": {
              "computed": true,
              "description": "The time when the event type was last updated.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name for the event type",
              "description_kind": "plain",
              "type": "string"
            },
            "tags": {
              "computed": true,
              "description": "Tags associated with this event type.",
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
      "last_updated_time": {
        "computed": true,
        "description": "The time when the detector was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_execution_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "created_time": {
              "computed": true,
              "description": "The time when the event type was created.",
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description.",
              "description_kind": "plain",
              "type": "string"
            },
            "detector_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "expression": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "language": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "last_updated_time": {
              "computed": true,
              "description": "The time when the event type was last updated.",
              "description_kind": "plain",
              "type": "string"
            },
            "outcomes": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "created_time": {
                    "computed": true,
                    "description": "The time when the outcome was created.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "The description.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "inline": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "last_updated_time": {
                    "computed": true,
                    "description": "The time when the outcome was last updated.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tags": {
                    "computed": true,
                    "description": "Tags associated with this outcome.",
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
                  }
                },
                "nesting_mode": "list"
              }
            },
            "rule_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "rule_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tags": {
              "computed": true,
              "description": "Tags associated with this event type.",
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
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with this detector.",
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
      }
    },
    "description": "Data Source schema for AWS::FraudDetector::Detector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFrauddetectorDetectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFrauddetectorDetector), &result)
	return &result
}
