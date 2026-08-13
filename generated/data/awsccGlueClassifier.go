package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueClassifier = `{
  "block": {
    "attributes": {
      "csv_classifier": {
        "computed": true,
        "description": "A classifier for comma-separated values (CSV).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow_single_column": {
              "computed": true,
              "description": "Enables the processing of files that contain only one column.",
              "description_kind": "plain",
              "type": "bool"
            },
            "contains_custom_datatype": {
              "computed": true,
              "description": "Indicates whether the CSV file contains custom data types.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "contains_header": {
              "computed": true,
              "description": "Indicates whether the CSV file contains a header. A value of UNKNOWN specifies that the classifier will detect whether the CSV file contains headings. A value of PRESENT specifies that the CSV file contains headings. A value of ABSENT specifies that the CSV file does not contain headings.",
              "description_kind": "plain",
              "type": "string"
            },
            "custom_datatype_configured": {
              "computed": true,
              "description": "Enables the configuration of custom data types.",
              "description_kind": "plain",
              "type": "bool"
            },
            "delimiter": {
              "computed": true,
              "description": "A custom symbol to denote what separates each column entry in the row.",
              "description_kind": "plain",
              "type": "string"
            },
            "disable_value_trimming": {
              "computed": true,
              "description": "Specifies not to trim values before identifying the type of column values. The default value is true.",
              "description_kind": "plain",
              "type": "bool"
            },
            "header": {
              "computed": true,
              "description": "A list of strings representing column names.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "The name of the classifier.",
              "description_kind": "plain",
              "type": "string"
            },
            "quote_symbol": {
              "computed": true,
              "description": "A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "grok_classifier": {
        "computed": true,
        "description": "A classifier that uses grok.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "classification": {
              "computed": true,
              "description": "An identifier of the data format that the classifier matches, such as Twitter, JSON, Omniture logs, and so on.",
              "description_kind": "plain",
              "type": "string"
            },
            "custom_patterns": {
              "computed": true,
              "description": "Optional custom grok patterns defined by this classifier.",
              "description_kind": "plain",
              "type": "string"
            },
            "grok_pattern": {
              "computed": true,
              "description": "The grok pattern applied to a data store by this classifier.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the classifier.",
              "description_kind": "plain",
              "type": "string"
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
      "json_classifier": {
        "computed": true,
        "description": "A classifier for JSON content.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "json_path": {
              "computed": true,
              "description": "A JsonPath string defining the JSON data for the classifier to classify. AWS Glue supports a subset of JsonPath, as described in Writing JsonPath Custom Classifiers.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the classifier.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "One of XMLClassifier/Name, GrokClassifier/Name, JsonClassifier/Name or CsvClassifier/Name",
        "description_kind": "plain",
        "type": "string"
      },
      "xml_classifier": {
        "computed": true,
        "description": "A classifier for XML content.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "classification": {
              "computed": true,
              "description": "An identifier of the data format that the classifier matches.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the classifier.",
              "description_kind": "plain",
              "type": "string"
            },
            "row_tag": {
              "computed": true,
              "description": "The XML tag designating the element that contains each record in an XML document being parsed. This can't identify a self-closing element (closed by /\u003e). An empty row element that contains only attributes can be parsed as long as it ends with a closing tag (for example, \u003crow item_a=\"A\" item_b=\"B\"\u003e\u003c/row\u003e is okay, but \u003crow item_a=\"A\" item_b=\"B\" /\u003e is not).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Glue::Classifier",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueClassifierSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueClassifier), &result)
	return &result
}
