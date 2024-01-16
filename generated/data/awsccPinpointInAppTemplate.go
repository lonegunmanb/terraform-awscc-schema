package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPinpointInAppTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "background_color": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "body_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alignment": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "body": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "text_color": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "header_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alignment": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "header": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "text_color": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "image_url": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "primary_btn": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "android": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "button_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "link": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "default_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "background_color": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "border_radius": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "button_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "link": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text_color": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "ios": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "button_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "link": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "web": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "button_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "link": {
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
            "secondary_btn": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "android": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "button_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "link": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "default_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "background_color": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "border_radius": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "button_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "link": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text_color": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "ios": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "button_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "link": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "web": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "button_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "link": {
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
      "custom_config": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "layout": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Pinpoint::InAppTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPinpointInAppTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPinpointInAppTemplate), &result)
	return &result
}
