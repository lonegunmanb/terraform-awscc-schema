package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLexBot = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_build_bot_locales": {
        "computed": true,
        "description": "Specifies whether to build the bot locales after bot creation completes.",
        "description_kind": "plain",
        "type": "bool"
      },
      "bot_file_s3_location": {
        "computed": true,
        "description": "S3 location of bot definitions zip file, if it's not defined inline in CloudFormation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket": {
              "computed": true,
              "description": "An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_object_key": {
              "computed": true,
              "description": "The Amazon S3 key of the deployment package.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_object_version": {
              "computed": true,
              "description": "For versioned objects, the version of the deployment package object to use. If not specified, the current object version will be used.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "bot_locales": {
        "computed": true,
        "description": "List of bot locales",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_vocabulary": {
              "computed": true,
              "description": "A custom vocabulary is a list of specific phrases that you want Amazon Lex V2 to recognize in the audio input.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "custom_vocabulary_items": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "phrase": {
                          "computed": true,
                          "description": "Phrase that should be recognized.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "weight": {
                          "computed": true,
                          "description": "The degree to which the phrase recognition is boosted.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "set"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "description": {
              "computed": true,
              "description": "A description of the resource",
              "description_kind": "plain",
              "type": "string"
            },
            "intents": {
              "computed": true,
              "description": "List of intents",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description": "A description of the resource",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "dialog_code_hook": {
                    "computed": true,
                    "description": "Settings that determine the Lambda function that Amazon Lex uses for processing user responses.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "fulfillment_code_hook": {
                    "computed": true,
                    "description": "Settings that determine if a Lambda function should be invoked to fulfill a specific intent.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "fulfillment_updates_specification": {
                          "computed": true,
                          "description": "Provides information for updating the user on the progress of fulfilling an intent.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "active": {
                                "computed": true,
                                "description": "Determines whether fulfillment updates are sent to the user. When this field is true, updates are sent.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "start_response": {
                                "computed": true,
                                "description": "Provides settings for a message that is sent to the user when a fulfillment Lambda function starts running.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "allow_interrupt": {
                                      "computed": true,
                                      "description": "Determines whether the user can interrupt the start message while it is playing.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "delay_in_seconds": {
                                      "computed": true,
                                      "description": "The delay between when the Lambda fulfillment function starts running and the start message is played. If the Lambda function returns before the delay is over, the start message isn't played.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "message_groups": {
                                      "computed": true,
                                      "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "message": {
                                            "computed": true,
                                            "description": "The primary message that Amazon Lex should send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                          "variations": {
                                            "computed": true,
                                            "description": "Message variations to send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                        "nesting_mode": "list"
                                      }
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "timeout_in_seconds": {
                                "computed": true,
                                "description": "The length of time that the fulfillment Lambda function should run before it times out.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "update_response": {
                                "computed": true,
                                "description": "Provides settings for a message that is sent periodically to the user while a fulfillment Lambda function is running.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "allow_interrupt": {
                                      "computed": true,
                                      "description": "Determines whether the user can interrupt an update message while it is playing.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "frequency_in_seconds": {
                                      "computed": true,
                                      "description": "The frequency that a message is sent to the user. When the period ends, Amazon Lex chooses a message from the message groups and plays it to the user. If the fulfillment Lambda returns before the first period ends, an update message is not played to the user.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "message_groups": {
                                      "computed": true,
                                      "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "message": {
                                            "computed": true,
                                            "description": "The primary message that Amazon Lex should send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                          "variations": {
                                            "computed": true,
                                            "description": "Message variations to send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                        "nesting_mode": "list"
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
                        "post_fulfillment_status_specification": {
                          "computed": true,
                          "description": "Provides information for updating the user on the progress of fulfilling an intent.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "failure_response": {
                                "computed": true,
                                "description": "A list of message groups that Amazon Lex uses to respond the user input.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "allow_interrupt": {
                                      "computed": true,
                                      "description": "Indicates whether the user can interrupt a speech prompt from the bot.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "message_groups_list": {
                                      "computed": true,
                                      "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "message": {
                                            "computed": true,
                                            "description": "The primary message that Amazon Lex should send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                          "variations": {
                                            "computed": true,
                                            "description": "Message variations to send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                        "nesting_mode": "list"
                                      }
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "success_response": {
                                "computed": true,
                                "description": "A list of message groups that Amazon Lex uses to respond the user input.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "allow_interrupt": {
                                      "computed": true,
                                      "description": "Indicates whether the user can interrupt a speech prompt from the bot.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "message_groups_list": {
                                      "computed": true,
                                      "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "message": {
                                            "computed": true,
                                            "description": "The primary message that Amazon Lex should send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                          "variations": {
                                            "computed": true,
                                            "description": "Message variations to send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                        "nesting_mode": "list"
                                      }
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "timeout_response": {
                                "computed": true,
                                "description": "A list of message groups that Amazon Lex uses to respond the user input.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "allow_interrupt": {
                                      "computed": true,
                                      "description": "Indicates whether the user can interrupt a speech prompt from the bot.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "message_groups_list": {
                                      "computed": true,
                                      "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "message": {
                                            "computed": true,
                                            "description": "The primary message that Amazon Lex should send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                          "variations": {
                                            "computed": true,
                                            "description": "Message variations to send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                        "nesting_mode": "list"
                                      }
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
                  "input_contexts": {
                    "computed": true,
                    "description": "The list of input contexts specified for the intent.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the context.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "intent_closing_setting": {
                    "computed": true,
                    "description": "Response that Amazon Lex sends to the user when the intent is closed.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "closing_response": {
                          "computed": true,
                          "description": "A list of message groups that Amazon Lex uses to respond the user input.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allow_interrupt": {
                                "computed": true,
                                "description": "Indicates whether the user can interrupt a speech prompt from the bot.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "message_groups_list": {
                                "computed": true,
                                "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "message": {
                                      "computed": true,
                                      "description": "The primary message that Amazon Lex should send to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "custom_payload": {
                                            "computed": true,
                                            "description": "A message in a custom format defined by the client application.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The string that is sent to your application.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "image_response_card": {
                                            "computed": true,
                                            "description": "A message that defines a response card that the client application can show to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "buttons": {
                                                  "computed": true,
                                                  "description": "A list of buttons that should be displayed on the response card.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "text": {
                                                        "computed": true,
                                                        "description": "The text that appears on the button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "list"
                                                  }
                                                },
                                                "image_url": {
                                                  "computed": true,
                                                  "description": "The URL of an image to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "subtitle": {
                                                  "computed": true,
                                                  "description": "The subtitle to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "title": {
                                                  "computed": true,
                                                  "description": "The title to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "plain_text_message": {
                                            "computed": true,
                                            "description": "A message in plain text format.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The message to send to the user.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "ssml_message": {
                                            "computed": true,
                                            "description": "A message in Speech Synthesis Markup Language (SSML).",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The SSML text that defines the prompt.",
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
                                    "variations": {
                                      "computed": true,
                                      "description": "Message variations to send to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "custom_payload": {
                                            "computed": true,
                                            "description": "A message in a custom format defined by the client application.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The string that is sent to your application.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "image_response_card": {
                                            "computed": true,
                                            "description": "A message that defines a response card that the client application can show to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "buttons": {
                                                  "computed": true,
                                                  "description": "A list of buttons that should be displayed on the response card.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "text": {
                                                        "computed": true,
                                                        "description": "The text that appears on the button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "list"
                                                  }
                                                },
                                                "image_url": {
                                                  "computed": true,
                                                  "description": "The URL of an image to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "subtitle": {
                                                  "computed": true,
                                                  "description": "The subtitle to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "title": {
                                                  "computed": true,
                                                  "description": "The title to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "plain_text_message": {
                                            "computed": true,
                                            "description": "A message in plain text format.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The message to send to the user.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "ssml_message": {
                                            "computed": true,
                                            "description": "A message in Speech Synthesis Markup Language (SSML).",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The SSML text that defines the prompt.",
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
                                  "nesting_mode": "list"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "is_active": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "intent_confirmation_setting": {
                    "computed": true,
                    "description": "Prompts that Amazon Lex sends to the user to confirm the completion of an intent.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "declination_response": {
                          "computed": true,
                          "description": "A list of message groups that Amazon Lex uses to respond the user input.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allow_interrupt": {
                                "computed": true,
                                "description": "Indicates whether the user can interrupt a speech prompt from the bot.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "message_groups_list": {
                                "computed": true,
                                "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "message": {
                                      "computed": true,
                                      "description": "The primary message that Amazon Lex should send to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "custom_payload": {
                                            "computed": true,
                                            "description": "A message in a custom format defined by the client application.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The string that is sent to your application.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "image_response_card": {
                                            "computed": true,
                                            "description": "A message that defines a response card that the client application can show to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "buttons": {
                                                  "computed": true,
                                                  "description": "A list of buttons that should be displayed on the response card.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "text": {
                                                        "computed": true,
                                                        "description": "The text that appears on the button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "list"
                                                  }
                                                },
                                                "image_url": {
                                                  "computed": true,
                                                  "description": "The URL of an image to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "subtitle": {
                                                  "computed": true,
                                                  "description": "The subtitle to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "title": {
                                                  "computed": true,
                                                  "description": "The title to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "plain_text_message": {
                                            "computed": true,
                                            "description": "A message in plain text format.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The message to send to the user.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "ssml_message": {
                                            "computed": true,
                                            "description": "A message in Speech Synthesis Markup Language (SSML).",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The SSML text that defines the prompt.",
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
                                    "variations": {
                                      "computed": true,
                                      "description": "Message variations to send to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "custom_payload": {
                                            "computed": true,
                                            "description": "A message in a custom format defined by the client application.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The string that is sent to your application.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "image_response_card": {
                                            "computed": true,
                                            "description": "A message that defines a response card that the client application can show to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "buttons": {
                                                  "computed": true,
                                                  "description": "A list of buttons that should be displayed on the response card.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "text": {
                                                        "computed": true,
                                                        "description": "The text that appears on the button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "list"
                                                  }
                                                },
                                                "image_url": {
                                                  "computed": true,
                                                  "description": "The URL of an image to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "subtitle": {
                                                  "computed": true,
                                                  "description": "The subtitle to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "title": {
                                                  "computed": true,
                                                  "description": "The title to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "plain_text_message": {
                                            "computed": true,
                                            "description": "A message in plain text format.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The message to send to the user.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "ssml_message": {
                                            "computed": true,
                                            "description": "A message in Speech Synthesis Markup Language (SSML).",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The SSML text that defines the prompt.",
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
                                  "nesting_mode": "list"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "is_active": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "prompt_specification": {
                          "computed": true,
                          "description": "Prompts the user to confirm the intent.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allow_interrupt": {
                                "computed": true,
                                "description": "Indicates whether the user can interrupt a speech prompt from the bot.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "max_retries": {
                                "computed": true,
                                "description": "The maximum number of times the bot tries to elicit a resonse from the user using this prompt.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "message_groups_list": {
                                "computed": true,
                                "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "message": {
                                      "computed": true,
                                      "description": "The primary message that Amazon Lex should send to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "custom_payload": {
                                            "computed": true,
                                            "description": "A message in a custom format defined by the client application.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The string that is sent to your application.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "image_response_card": {
                                            "computed": true,
                                            "description": "A message that defines a response card that the client application can show to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "buttons": {
                                                  "computed": true,
                                                  "description": "A list of buttons that should be displayed on the response card.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "text": {
                                                        "computed": true,
                                                        "description": "The text that appears on the button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "list"
                                                  }
                                                },
                                                "image_url": {
                                                  "computed": true,
                                                  "description": "The URL of an image to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "subtitle": {
                                                  "computed": true,
                                                  "description": "The subtitle to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "title": {
                                                  "computed": true,
                                                  "description": "The title to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "plain_text_message": {
                                            "computed": true,
                                            "description": "A message in plain text format.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The message to send to the user.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "ssml_message": {
                                            "computed": true,
                                            "description": "A message in Speech Synthesis Markup Language (SSML).",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The SSML text that defines the prompt.",
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
                                    "variations": {
                                      "computed": true,
                                      "description": "Message variations to send to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "custom_payload": {
                                            "computed": true,
                                            "description": "A message in a custom format defined by the client application.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The string that is sent to your application.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "image_response_card": {
                                            "computed": true,
                                            "description": "A message that defines a response card that the client application can show to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "buttons": {
                                                  "computed": true,
                                                  "description": "A list of buttons that should be displayed on the response card.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "text": {
                                                        "computed": true,
                                                        "description": "The text that appears on the button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "list"
                                                  }
                                                },
                                                "image_url": {
                                                  "computed": true,
                                                  "description": "The URL of an image to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "subtitle": {
                                                  "computed": true,
                                                  "description": "The subtitle to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "title": {
                                                  "computed": true,
                                                  "description": "The title to display on the response card.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "plain_text_message": {
                                            "computed": true,
                                            "description": "A message in plain text format.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The message to send to the user.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "ssml_message": {
                                            "computed": true,
                                            "description": "A message in Speech Synthesis Markup Language (SSML).",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "computed": true,
                                                  "description": "The SSML text that defines the prompt.",
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
                                  "nesting_mode": "list"
                                }
                              },
                              "message_selection_strategy": {
                                "computed": true,
                                "description": "Indicates how a message is selected from a message group among retries.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "prompt_attempts_specification": {
                                "computed": true,
                                "description": "Specifies the advanced settings on each attempt of the prompt.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "allow_interrupt": {
                                      "computed": true,
                                      "description": "Indicates whether the user can interrupt a speech prompt attempt from the bot.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "allowed_input_types": {
                                      "computed": true,
                                      "description": "Specifies the allowed input types.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "allow_audio_input": {
                                            "computed": true,
                                            "description": "Indicates whether audio input is allowed.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "allow_dtmf_input": {
                                            "computed": true,
                                            "description": "Indicates whether DTMF input is allowed.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "audio_and_dtmf_input_specification": {
                                      "computed": true,
                                      "description": "Specifies the audio and DTMF input specification.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "audio_specification": {
                                            "computed": true,
                                            "description": "Specifies the audio input specifications.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "end_timeout_ms": {
                                                  "computed": true,
                                                  "description": "Time for which a bot waits after the customer stops speaking to assume the utterance is finished.",
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                },
                                                "max_length_ms": {
                                                  "computed": true,
                                                  "description": "Time for how long Amazon Lex waits before speech input is truncated and the speech is returned to application.",
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "dtmf_specification": {
                                            "computed": true,
                                            "description": "Specifies the settings on DTMF input.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "deletion_character": {
                                                  "computed": true,
                                                  "description": "The DTMF character that clears the accumulated DTMF digits and immediately ends the input.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "end_character": {
                                                  "computed": true,
                                                  "description": "The DTMF character that immediately ends input. If the user does not press this character, the input ends after the end timeout.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "end_timeout_ms": {
                                                  "computed": true,
                                                  "description": "How long the bot should wait after the last DTMF character input before assuming that the input has concluded.",
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                },
                                                "max_length": {
                                                  "computed": true,
                                                  "description": "The maximum number of DTMF digits allowed in an utterance.",
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "start_timeout_ms": {
                                            "computed": true,
                                            "description": "Time for which a bot waits before assuming that the customer isn't going to speak or press a key. This timeout is shared between Audio and DTMF inputs.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "text_input_specification": {
                                      "computed": true,
                                      "description": "Specifies the text input specifications.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "start_timeout_ms": {
                                            "computed": true,
                                            "description": "Time for which a bot waits before re-prompting a customer for text input.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    }
                                  },
                                  "nesting_mode": "map"
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
                  "kendra_configuration": {
                    "computed": true,
                    "description": "Configuration for searching a Amazon Kendra index specified for the intent.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "kendra_index": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the Amazon Kendra index that you want the AMAZON.KendraSearchIntent intent to search.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "query_filter_string": {
                          "computed": true,
                          "description": "A query filter that Amazon Lex sends to Amazon Kendra to filter the response from a query.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "query_filter_string_enabled": {
                          "computed": true,
                          "description": "Determines whether the AMAZON.KendraSearchIntent intent uses a custom query string to query the Amazon Kendra index.",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "name": {
                    "computed": true,
                    "description": "Unique name for a resource.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_contexts": {
                    "computed": true,
                    "description": "A list of contexts that the intent activates when it is fulfilled.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "Unique name for a resource.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "time_to_live_in_seconds": {
                          "computed": true,
                          "description": "The amount of time, in seconds, that the output context should remain active.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "turns_to_live": {
                          "computed": true,
                          "description": "The number of conversation turns that the output context should remain active.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "parent_intent_signature": {
                    "computed": true,
                    "description": "A unique identifier for the built-in intent to base this intent on.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "sample_utterances": {
                    "computed": true,
                    "description": "An array of sample utterances",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "utterance": {
                          "computed": true,
                          "description": "The sample utterance that Amazon Lex uses to build its machine-learning model to recognize intents/slots.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "slot_priorities": {
                    "computed": true,
                    "description": "List for slot priorities",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "priority": {
                          "computed": true,
                          "description": "The priority that a slot should be elicited.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "slot_name": {
                          "computed": true,
                          "description": "The name of the slot.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "slots": {
                    "computed": true,
                    "description": "List of slots",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "description": {
                          "computed": true,
                          "description": "A description of the resource",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "multiple_values_setting": {
                          "computed": true,
                          "description": "Indicates whether a slot can return multiple values.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allow_multiple_values": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "name": {
                          "computed": true,
                          "description": "Unique name for a resource.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "obfuscation_setting": {
                          "computed": true,
                          "description": "Determines whether Amazon Lex obscures slot values in conversation logs.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "obfuscation_setting_type": {
                                "computed": true,
                                "description": "Value that determines whether Amazon Lex obscures slot values in conversation logs. The default is to obscure the values.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "slot_type_name": {
                          "computed": true,
                          "description": "The slot type name that is used in the slot. Allows for custom and built-in slot type names",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value_elicitation_setting": {
                          "computed": true,
                          "description": "Settings that you can use for eliciting a slot value.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "default_value_specification": {
                                "computed": true,
                                "description": "A list of default values for a slot.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "default_value_list": {
                                      "computed": true,
                                      "description": "A list of slot default values",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "default_value": {
                                            "computed": true,
                                            "description": "The default value to use when a user doesn't provide a value for a slot.",
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
                              "prompt_specification": {
                                "computed": true,
                                "description": "The prompt that Amazon Lex uses to elicit the slot value from the user.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "allow_interrupt": {
                                      "computed": true,
                                      "description": "Indicates whether the user can interrupt a speech prompt from the bot.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "max_retries": {
                                      "computed": true,
                                      "description": "The maximum number of times the bot tries to elicit a resonse from the user using this prompt.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "message_groups_list": {
                                      "computed": true,
                                      "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "message": {
                                            "computed": true,
                                            "description": "The primary message that Amazon Lex should send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                          "variations": {
                                            "computed": true,
                                            "description": "Message variations to send to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_payload": {
                                                  "computed": true,
                                                  "description": "A message in a custom format defined by the client application.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The string that is sent to your application.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "image_response_card": {
                                                  "computed": true,
                                                  "description": "A message that defines a response card that the client application can show to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "buttons": {
                                                        "computed": true,
                                                        "description": "A list of buttons that should be displayed on the response card.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "text": {
                                                              "computed": true,
                                                              "description": "The text that appears on the button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "list"
                                                        }
                                                      },
                                                      "image_url": {
                                                        "computed": true,
                                                        "description": "The URL of an image to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "subtitle": {
                                                        "computed": true,
                                                        "description": "The subtitle to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "title": {
                                                        "computed": true,
                                                        "description": "The title to display on the response card.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "plain_text_message": {
                                                  "computed": true,
                                                  "description": "A message in plain text format.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The message to send to the user.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "ssml_message": {
                                                  "computed": true,
                                                  "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "value": {
                                                        "computed": true,
                                                        "description": "The SSML text that defines the prompt.",
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
                                        "nesting_mode": "list"
                                      }
                                    },
                                    "message_selection_strategy": {
                                      "computed": true,
                                      "description": "Indicates how a message is selected from a message group among retries.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "prompt_attempts_specification": {
                                      "computed": true,
                                      "description": "Specifies the advanced settings on each attempt of the prompt.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "allow_interrupt": {
                                            "computed": true,
                                            "description": "Indicates whether the user can interrupt a speech prompt attempt from the bot.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "allowed_input_types": {
                                            "computed": true,
                                            "description": "Specifies the allowed input types.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "allow_audio_input": {
                                                  "computed": true,
                                                  "description": "Indicates whether audio input is allowed.",
                                                  "description_kind": "plain",
                                                  "type": "bool"
                                                },
                                                "allow_dtmf_input": {
                                                  "computed": true,
                                                  "description": "Indicates whether DTMF input is allowed.",
                                                  "description_kind": "plain",
                                                  "type": "bool"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "audio_and_dtmf_input_specification": {
                                            "computed": true,
                                            "description": "Specifies the audio and DTMF input specification.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "audio_specification": {
                                                  "computed": true,
                                                  "description": "Specifies the audio input specifications.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "end_timeout_ms": {
                                                        "computed": true,
                                                        "description": "Time for which a bot waits after the customer stops speaking to assume the utterance is finished.",
                                                        "description_kind": "plain",
                                                        "type": "number"
                                                      },
                                                      "max_length_ms": {
                                                        "computed": true,
                                                        "description": "Time for how long Amazon Lex waits before speech input is truncated and the speech is returned to application.",
                                                        "description_kind": "plain",
                                                        "type": "number"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "dtmf_specification": {
                                                  "computed": true,
                                                  "description": "Specifies the settings on DTMF input.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "deletion_character": {
                                                        "computed": true,
                                                        "description": "The DTMF character that clears the accumulated DTMF digits and immediately ends the input.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "end_character": {
                                                        "computed": true,
                                                        "description": "The DTMF character that immediately ends input. If the user does not press this character, the input ends after the end timeout.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "end_timeout_ms": {
                                                        "computed": true,
                                                        "description": "How long the bot should wait after the last DTMF character input before assuming that the input has concluded.",
                                                        "description_kind": "plain",
                                                        "type": "number"
                                                      },
                                                      "max_length": {
                                                        "computed": true,
                                                        "description": "The maximum number of DTMF digits allowed in an utterance.",
                                                        "description_kind": "plain",
                                                        "type": "number"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "start_timeout_ms": {
                                                  "computed": true,
                                                  "description": "Time for which a bot waits before assuming that the customer isn't going to speak or press a key. This timeout is shared between Audio and DTMF inputs.",
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "text_input_specification": {
                                            "computed": true,
                                            "description": "Specifies the text input specifications.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "start_timeout_ms": {
                                                  "computed": true,
                                                  "description": "Time for which a bot waits before re-prompting a customer for text input.",
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          }
                                        },
                                        "nesting_mode": "map"
                                      }
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "sample_utterances": {
                                "computed": true,
                                "description": "If you know a specific pattern that users might respond to an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "utterance": {
                                      "computed": true,
                                      "description": "The sample utterance that Amazon Lex uses to build its machine-learning model to recognize intents/slots.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "slot_constraint": {
                                "computed": true,
                                "description": "Specifies whether the slot is required or optional.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "wait_and_continue_specification": {
                                "computed": true,
                                "description": "Specifies the prompts that Amazon Lex uses while a bot is waiting for customer input.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "continue_response": {
                                      "computed": true,
                                      "description": "The response that Amazon Lex sends to indicate that the bot is ready to continue the conversation.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "allow_interrupt": {
                                            "computed": true,
                                            "description": "Indicates whether the user can interrupt a speech prompt from the bot.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "message_groups_list": {
                                            "computed": true,
                                            "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "message": {
                                                  "computed": true,
                                                  "description": "The primary message that Amazon Lex should send to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "custom_payload": {
                                                        "computed": true,
                                                        "description": "A message in a custom format defined by the client application.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The string that is sent to your application.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "image_response_card": {
                                                        "computed": true,
                                                        "description": "A message that defines a response card that the client application can show to the user.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "buttons": {
                                                              "computed": true,
                                                              "description": "A list of buttons that should be displayed on the response card.",
                                                              "description_kind": "plain",
                                                              "nested_type": {
                                                                "attributes": {
                                                                  "text": {
                                                                    "computed": true,
                                                                    "description": "The text that appears on the button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  },
                                                                  "value": {
                                                                    "computed": true,
                                                                    "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  }
                                                                },
                                                                "nesting_mode": "list"
                                                              }
                                                            },
                                                            "image_url": {
                                                              "computed": true,
                                                              "description": "The URL of an image to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "subtitle": {
                                                              "computed": true,
                                                              "description": "The subtitle to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "title": {
                                                              "computed": true,
                                                              "description": "The title to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "plain_text_message": {
                                                        "computed": true,
                                                        "description": "A message in plain text format.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The message to send to the user.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "ssml_message": {
                                                        "computed": true,
                                                        "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The SSML text that defines the prompt.",
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
                                                "variations": {
                                                  "computed": true,
                                                  "description": "Message variations to send to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "custom_payload": {
                                                        "computed": true,
                                                        "description": "A message in a custom format defined by the client application.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The string that is sent to your application.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "image_response_card": {
                                                        "computed": true,
                                                        "description": "A message that defines a response card that the client application can show to the user.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "buttons": {
                                                              "computed": true,
                                                              "description": "A list of buttons that should be displayed on the response card.",
                                                              "description_kind": "plain",
                                                              "nested_type": {
                                                                "attributes": {
                                                                  "text": {
                                                                    "computed": true,
                                                                    "description": "The text that appears on the button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  },
                                                                  "value": {
                                                                    "computed": true,
                                                                    "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  }
                                                                },
                                                                "nesting_mode": "list"
                                                              }
                                                            },
                                                            "image_url": {
                                                              "computed": true,
                                                              "description": "The URL of an image to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "subtitle": {
                                                              "computed": true,
                                                              "description": "The subtitle to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "title": {
                                                              "computed": true,
                                                              "description": "The title to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "plain_text_message": {
                                                        "computed": true,
                                                        "description": "A message in plain text format.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The message to send to the user.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "ssml_message": {
                                                        "computed": true,
                                                        "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The SSML text that defines the prompt.",
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
                                              "nesting_mode": "list"
                                            }
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "is_active": {
                                      "computed": true,
                                      "description": "Specifies whether the bot will wait for a user to respond.",
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "still_waiting_response": {
                                      "computed": true,
                                      "description": "The response that Amazon Lex sends periodically to the user to indicate that the bot is still waiting for input from the user.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "allow_interrupt": {
                                            "computed": true,
                                            "description": "Indicates whether the user can interrupt a speech prompt from the bot.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "frequency_in_seconds": {
                                            "computed": true,
                                            "description": "How often a message should be sent to the user in seconds.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "message_groups_list": {
                                            "computed": true,
                                            "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "message": {
                                                  "computed": true,
                                                  "description": "The primary message that Amazon Lex should send to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "custom_payload": {
                                                        "computed": true,
                                                        "description": "A message in a custom format defined by the client application.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The string that is sent to your application.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "image_response_card": {
                                                        "computed": true,
                                                        "description": "A message that defines a response card that the client application can show to the user.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "buttons": {
                                                              "computed": true,
                                                              "description": "A list of buttons that should be displayed on the response card.",
                                                              "description_kind": "plain",
                                                              "nested_type": {
                                                                "attributes": {
                                                                  "text": {
                                                                    "computed": true,
                                                                    "description": "The text that appears on the button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  },
                                                                  "value": {
                                                                    "computed": true,
                                                                    "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  }
                                                                },
                                                                "nesting_mode": "list"
                                                              }
                                                            },
                                                            "image_url": {
                                                              "computed": true,
                                                              "description": "The URL of an image to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "subtitle": {
                                                              "computed": true,
                                                              "description": "The subtitle to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "title": {
                                                              "computed": true,
                                                              "description": "The title to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "plain_text_message": {
                                                        "computed": true,
                                                        "description": "A message in plain text format.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The message to send to the user.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "ssml_message": {
                                                        "computed": true,
                                                        "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The SSML text that defines the prompt.",
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
                                                "variations": {
                                                  "computed": true,
                                                  "description": "Message variations to send to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "custom_payload": {
                                                        "computed": true,
                                                        "description": "A message in a custom format defined by the client application.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The string that is sent to your application.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "image_response_card": {
                                                        "computed": true,
                                                        "description": "A message that defines a response card that the client application can show to the user.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "buttons": {
                                                              "computed": true,
                                                              "description": "A list of buttons that should be displayed on the response card.",
                                                              "description_kind": "plain",
                                                              "nested_type": {
                                                                "attributes": {
                                                                  "text": {
                                                                    "computed": true,
                                                                    "description": "The text that appears on the button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  },
                                                                  "value": {
                                                                    "computed": true,
                                                                    "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  }
                                                                },
                                                                "nesting_mode": "list"
                                                              }
                                                            },
                                                            "image_url": {
                                                              "computed": true,
                                                              "description": "The URL of an image to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "subtitle": {
                                                              "computed": true,
                                                              "description": "The subtitle to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "title": {
                                                              "computed": true,
                                                              "description": "The title to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "plain_text_message": {
                                                        "computed": true,
                                                        "description": "A message in plain text format.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The message to send to the user.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "ssml_message": {
                                                        "computed": true,
                                                        "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The SSML text that defines the prompt.",
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
                                              "nesting_mode": "list"
                                            }
                                          },
                                          "timeout_in_seconds": {
                                            "computed": true,
                                            "description": "If Amazon Lex waits longer than this length of time in seconds for a response, it will stop sending messages.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "waiting_response": {
                                      "computed": true,
                                      "description": "The response that Amazon Lex sends to indicate that the bot is waiting for the conversation to continue.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "allow_interrupt": {
                                            "computed": true,
                                            "description": "Indicates whether the user can interrupt a speech prompt from the bot.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "message_groups_list": {
                                            "computed": true,
                                            "description": "One to 5 message groups that contain update messages. Amazon Lex chooses one of the messages to play to the user.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "message": {
                                                  "computed": true,
                                                  "description": "The primary message that Amazon Lex should send to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "custom_payload": {
                                                        "computed": true,
                                                        "description": "A message in a custom format defined by the client application.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The string that is sent to your application.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "image_response_card": {
                                                        "computed": true,
                                                        "description": "A message that defines a response card that the client application can show to the user.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "buttons": {
                                                              "computed": true,
                                                              "description": "A list of buttons that should be displayed on the response card.",
                                                              "description_kind": "plain",
                                                              "nested_type": {
                                                                "attributes": {
                                                                  "text": {
                                                                    "computed": true,
                                                                    "description": "The text that appears on the button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  },
                                                                  "value": {
                                                                    "computed": true,
                                                                    "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  }
                                                                },
                                                                "nesting_mode": "list"
                                                              }
                                                            },
                                                            "image_url": {
                                                              "computed": true,
                                                              "description": "The URL of an image to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "subtitle": {
                                                              "computed": true,
                                                              "description": "The subtitle to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "title": {
                                                              "computed": true,
                                                              "description": "The title to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "plain_text_message": {
                                                        "computed": true,
                                                        "description": "A message in plain text format.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The message to send to the user.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "ssml_message": {
                                                        "computed": true,
                                                        "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The SSML text that defines the prompt.",
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
                                                "variations": {
                                                  "computed": true,
                                                  "description": "Message variations to send to the user.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "custom_payload": {
                                                        "computed": true,
                                                        "description": "A message in a custom format defined by the client application.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The string that is sent to your application.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "image_response_card": {
                                                        "computed": true,
                                                        "description": "A message that defines a response card that the client application can show to the user.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "buttons": {
                                                              "computed": true,
                                                              "description": "A list of buttons that should be displayed on the response card.",
                                                              "description_kind": "plain",
                                                              "nested_type": {
                                                                "attributes": {
                                                                  "text": {
                                                                    "computed": true,
                                                                    "description": "The text that appears on the button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  },
                                                                  "value": {
                                                                    "computed": true,
                                                                    "description": "The value returned to Amazon Lex when the user chooses this button.",
                                                                    "description_kind": "plain",
                                                                    "type": "string"
                                                                  }
                                                                },
                                                                "nesting_mode": "list"
                                                              }
                                                            },
                                                            "image_url": {
                                                              "computed": true,
                                                              "description": "The URL of an image to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "subtitle": {
                                                              "computed": true,
                                                              "description": "The subtitle to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            },
                                                            "title": {
                                                              "computed": true,
                                                              "description": "The title to display on the response card.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "plain_text_message": {
                                                        "computed": true,
                                                        "description": "A message in plain text format.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The message to send to the user.",
                                                              "description_kind": "plain",
                                                              "type": "string"
                                                            }
                                                          },
                                                          "nesting_mode": "single"
                                                        }
                                                      },
                                                      "ssml_message": {
                                                        "computed": true,
                                                        "description": "A message in Speech Synthesis Markup Language (SSML).",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "value": {
                                                              "computed": true,
                                                              "description": "The SSML text that defines the prompt.",
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
                                              "nesting_mode": "list"
                                            }
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
                        }
                      },
                      "nesting_mode": "set"
                    }
                  }
                },
                "nesting_mode": "set"
              }
            },
            "locale_id": {
              "computed": true,
              "description": "The identifier of the language and locale that the bot will be used in.",
              "description_kind": "plain",
              "type": "string"
            },
            "nlu_confidence_threshold": {
              "computed": true,
              "description": "The specified confidence threshold for inserting the AMAZON.FallbackIntent and AMAZON.KendraSearchIntent intents.",
              "description_kind": "plain",
              "type": "number"
            },
            "slot_types": {
              "computed": true,
              "description": "List of SlotTypes",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description": "A description of the resource",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "external_source_setting": {
                    "computed": true,
                    "description": "Provides information about the external source of the slot type's definition.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "grammar_slot_type_setting": {
                          "computed": true,
                          "description": "Settings required for a slot type based on a grammar that you provide.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "source": {
                                "computed": true,
                                "description": "Describes the Amazon S3 bucket name and location for the grammar that is the source for the slot type.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "kms_key_arn": {
                                      "computed": true,
                                      "description": "The Amazon KMS key required to decrypt the contents of the grammar, if any.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "s3_bucket_name": {
                                      "computed": true,
                                      "description": "The name of the S3 bucket that contains the grammar source.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "s3_object_key": {
                                      "computed": true,
                                      "description": "The path to the grammar in the S3 bucket.",
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
                  "name": {
                    "computed": true,
                    "description": "Unique name for a resource.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "parent_slot_type_signature": {
                    "computed": true,
                    "description": "The built-in slot type used as a parent of this slot type.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "slot_type_values": {
                    "computed": true,
                    "description": "A List of slot type values",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "sample_value": {
                          "computed": true,
                          "description": "Defines one of the values for a slot type.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "value": {
                                "computed": true,
                                "description": "The value that can be used for a slot type.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "synonyms": {
                          "computed": true,
                          "description": "Additional values related to the slot type entry.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "value": {
                                "computed": true,
                                "description": "The value that can be used for a slot type.",
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
                  "value_selection_setting": {
                    "computed": true,
                    "description": "Contains settings used by Amazon Lex to select a slot value.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "advanced_recognition_setting": {
                          "computed": true,
                          "description": "Provides settings that enable advanced recognition settings for slot values.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "audio_recognition_strategy": {
                                "computed": true,
                                "description": "Enables using slot values as a custom vocabulary when recognizing user utterances.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "regex_filter": {
                          "computed": true,
                          "description": "A regular expression used to validate the value of a slot.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "pattern": {
                                "computed": true,
                                "description": "Regex pattern",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "resolution_strategy": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "set"
              }
            },
            "voice_settings": {
              "computed": true,
              "description": "Settings for using an Amazon Polly voice to communicate with a user.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "engine": {
                    "computed": true,
                    "description": "Indicates the type of Amazon Polly voice that Amazon Lex should use for voice interaction with the user. For more information, see the engine parameter of the SynthesizeSpeech operation in the Amazon Polly developer guide.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "voice_id": {
                    "computed": true,
                    "description": "The Amazon Polly voice ID that Amazon Lex uses for voice interaction with the user.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "set"
        }
      },
      "bot_tags": {
        "computed": true,
        "description": "A list of tags to add to the bot, which can only be added at bot creation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "data_privacy": {
        "computed": true,
        "description": "Data privacy setting of the Bot.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "child_directed": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of the resource",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "idle_session_ttl_in_seconds": {
        "computed": true,
        "description": "IdleSessionTTLInSeconds of the resource",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "Unique name for a resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM role that has permission to access the bot.",
        "description_kind": "plain",
        "type": "string"
      },
      "test_bot_alias_settings": {
        "computed": true,
        "description": "Configuring the test bot alias settings for a given bot",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bot_alias_locale_settings": {
              "computed": true,
              "description": "A list of bot alias locale settings to add to the bot alias.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bot_alias_locale_setting": {
                    "computed": true,
                    "description": "You can use this parameter to specify a specific Lambda function to run different functions in different locales.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "code_hook_specification": {
                          "computed": true,
                          "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "lambda_code_hook": {
                                "computed": true,
                                "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "code_hook_interface_version": {
                                      "computed": true,
                                      "description": "The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "lambda_arn": {
                                      "computed": true,
                                      "description": "The Amazon Resource Name (ARN) of the Lambda function.",
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
                        "enabled": {
                          "computed": true,
                          "description": "Whether the Lambda code hook is enabled",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "locale_id": {
                    "computed": true,
                    "description": "A string used to identify the locale",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "conversation_log_settings": {
              "computed": true,
              "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "audio_log_settings": {
                    "computed": true,
                    "description": "List of audio log settings that pertain to the conversation log settings for the bot's TestBotAlias.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "destination": {
                          "computed": true,
                          "description": "The location of audio log files collected when conversation logging is enabled for a bot.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "s3_bucket": {
                                "computed": true,
                                "description": "Specifies an Amazon S3 bucket for logging audio conversations",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "kms_key_arn": {
                                      "computed": true,
                                      "description": "The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an S3 bucket.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "log_prefix": {
                                      "computed": true,
                                      "description": "The Amazon S3 key of the deployment package.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "s3_bucket_arn": {
                                      "computed": true,
                                      "description": "The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.",
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
                        "enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "set"
                    }
                  },
                  "text_log_settings": {
                    "computed": true,
                    "description": "List of text log settings that pertain to the conversation log settings for the bot's TestBotAlias",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "destination": {
                          "computed": true,
                          "description": "Defines the Amazon CloudWatch Logs destination log group for conversation text logs.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "cloudwatch": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "cloudwatch_log_group_arn": {
                                      "computed": true,
                                      "description": "A string used to identify the groupArn for the Cloudwatch Log Group",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "log_prefix": {
                                      "computed": true,
                                      "description": "A string containing the value for the Log Prefix",
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
                        "enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "set"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "description": {
              "computed": true,
              "description": "A description of the resource",
              "description_kind": "plain",
              "type": "string"
            },
            "sentiment_analysis_settings": {
              "computed": true,
              "description": "Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "detect_sentiment": {
                    "computed": true,
                    "description": "Enable to call Amazon Comprehend for Sentiment natively within Lex",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "test_bot_alias_tags": {
        "computed": true,
        "description": "A list of tags to add to the test alias for a bot, , which can only be added at bot/bot alias creation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Lex::Bot",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLexBotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLexBot), &result)
	return &result
}
