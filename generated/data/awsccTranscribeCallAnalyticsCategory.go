package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTranscribeCallAnalyticsCategory = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Call Analytics category.",
        "description_kind": "plain",
        "type": "string"
      },
      "category_name": {
        "computed": true,
        "description": "A unique name, chosen by you, for your Call Analytics category.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The date and time the Call Analytics category was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "input_type": {
        "computed": true,
        "description": "The input type associated with the specified category.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_update_time": {
        "computed": true,
        "description": "The date and time the Call Analytics category was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description": "Rules define a Call Analytics category.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "interruption_filter": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "absolute_time_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end_time": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "first": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "last": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "start_time": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "negate": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "participant_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "relative_time_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end_percentage": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "first": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "last": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "start_percentage": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "threshold": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "non_talk_time_filter": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "absolute_time_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end_time": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "first": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "last": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "start_time": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "negate": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "relative_time_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end_percentage": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "first": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "last": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "start_percentage": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "threshold": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "sentiment_filter": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "absolute_time_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end_time": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "first": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "last": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "start_time": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "negate": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "participant_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "relative_time_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end_percentage": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "first": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "last": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "start_percentage": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "sentiments": {
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
            },
            "transcript_filter": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "absolute_time_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end_time": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "first": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "last": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "start_time": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "negate": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "participant_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "relative_time_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end_percentage": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "first": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "last": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "start_percentage": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "targets": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "transcript_filter_type": {
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
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the Call Analytics category.",
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
    "description": "Data Source schema for AWS::Transcribe::CallAnalyticsCategory",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccTranscribeCallAnalyticsCategorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTranscribeCallAnalyticsCategory), &result)
	return &result
}
