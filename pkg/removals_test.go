package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveComments(t *testing.T) {
	tests := []struct {
		name             string
		removeEmptyLines bool
		jsonc            string
		json             string
	}{
		{
			name:             "no change to straight json with spaces",
			removeEmptyLines: false,
			jsonc: `{
  "glossary": {
    "title": "example glossary",
    "GlossDiv": {
      "title": "S",
      "GlossList": {
        "GlossEntry": {
          "ID": "SGML",
          "SortAs": "SGML",
          "GlossTerm": "Standard Generalized Markup Language",
          "Acronym": "SGML",
          "Abbrev": "ISO 8879:1986",
          "GlossDef": {
            "para": "A meta-markup language, used to create markup languages such as DocBook.",
            "GlossSeeAlso": [
              "GML",
              "XML"
            ]
          },
          "GlossSee": "markup"
        }
      }
    }
  }
}`,
			json: `{
  "glossary": {
    "title": "example glossary",
    "GlossDiv": {
      "title": "S",
      "GlossList": {
        "GlossEntry": {
          "ID": "SGML",
          "SortAs": "SGML",
          "GlossTerm": "Standard Generalized Markup Language",
          "Acronym": "SGML",
          "Abbrev": "ISO 8879:1986",
          "GlossDef": {
            "para": "A meta-markup language, used to create markup languages such as DocBook.",
            "GlossSeeAlso": [
              "GML",
              "XML"
            ]
          },
          "GlossSee": "markup"
        }
      }
    }
  }
}`,
		},
		{
			name:             "no change to straight json with tabs",
			removeEmptyLines: false,
			jsonc: `{
	"glossary": {
		"title": "example glossary",
		"GlossDiv": {
			"title": "S",
			"GlossList": {
				"GlossEntry": {
					"ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
						"para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": [
							"GML",
							"XML"
						]
					},
					"GlossSee": "markup"
				}
			}
		}
	}
}`,
			json: `{
	"glossary": {
		"title": "example glossary",
		"GlossDiv": {
			"title": "S",
			"GlossList": {
				"GlossEntry": {
					"ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
						"para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": [
							"GML",
							"XML"
						]
					},
					"GlossSee": "markup"
				}
			}
		}
	}
}`,
		},
		{
			name:             "simple comment",
			removeEmptyLines: false,
			jsonc: `{
	"glossary": {
		// This is a glossay item
		"title": "example glossary",
		"GlossDiv": {
			"title": "S",
			"GlossList": {
				"GlossEntry": {
					"ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
						"para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": [
							"GML",
							"XML"
						]
					},
					"GlossSee": "markup"
				}
			}
		}
	}
}`,
			json: `{
	"glossary": {

		"title": "example glossary",
		"GlossDiv": {
			"title": "S",
			"GlossList": {
				"GlossEntry": {
					"ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
						"para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": [
							"GML",
							"XML"
						]
					},
					"GlossSee": "markup"
				}
			}
		}
	}
}`,
		},
		{
			name:             "simple comment remove empty lines",
			removeEmptyLines: true,
			jsonc: `{
	"glossary": {
		// This is a glossay item
		"title": "example glossary",
		"GlossDiv": {
			"title": "S",
			"GlossList": {
				"GlossEntry": {
					"ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
						"para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": [
							"GML",
							"XML"
						]
					},
					"GlossSee": "markup"
				}
			}
		}
	}
}`,
			json: `{
	"glossary": {
		"title": "example glossary",
		"GlossDiv": {
			"title": "S",
			"GlossList": {
				"GlossEntry": {
					"ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
						"para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": [
							"GML",
							"XML"
						]
					},
					"GlossSee": "markup"
				}
			}
		}
	}
}`,
		},

		{
			name:             "comment block single line",
			removeEmptyLines: false,
			jsonc: `{
	"glossary": {
		/* this is a block comment */
		"title": "example glossary",
		"GlossDiv": {
			"title": "S",
			"GlossList": {
				"GlossEntry": {
					"ID": "SGML",
					/* here's another */
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
						"para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": [
							"GML",
							"XML"
						]
					},
					"GlossSee": "markup"
				}
			}
		}
	}
}`,
			json: `{
	"glossary": {

		"title": "example glossary",
		"GlossDiv": {
			"title": "S",
			"GlossList": {
				"GlossEntry": {
					"ID": "SGML",

					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
						"para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": [
							"GML",
							"XML"
						]
					},
					"GlossSee": "markup"
				}
			}
		}
	}
}`,
		},
		{
			name:             "comment block multi line",
			removeEmptyLines: false,
			jsonc: `{
		"glossary": {
			"title": "example glossary",
			"GlossDiv": {
				"title": "S",
				/* this title
				  should be
				  more than 1 character
				  */
				"GlossList": {
					"GlossEntry": {
						"ID": "SGML",
						"SortAs": "SGML",
						"GlossTerm": "Standard Generalized Markup Language",
						"Acronym": "SGML",
						"Abbrev": "ISO 8879:1986",
						"GlossDef": {
							"para": "A meta-markup language, used to create markup languages such as DocBook.",
							"GlossSeeAlso": [
								"GML",
								"XML"
							]
						},
						"GlossSee": "markup"
					}
				}
			}
		}
	}`,
			json: `{
		"glossary": {
			"title": "example glossary",
			"GlossDiv": {
				"title": "S",

				"GlossList": {
					"GlossEntry": {
						"ID": "SGML",
						"SortAs": "SGML",
						"GlossTerm": "Standard Generalized Markup Language",
						"Acronym": "SGML",
						"Abbrev": "ISO 8879:1986",
						"GlossDef": {
							"para": "A meta-markup language, used to create markup languages such as DocBook.",
							"GlossSeeAlso": [
								"GML",
								"XML"
							]
						},
						"GlossSee": "markup"
					}
				}
			}
		}
	}`,
		},
		{
			name:             "everything",
			removeEmptyLines: false,
			jsonc: `{
		"glossary": {
			"title": "example glossary",
			// example comment
			"GlossDiv": {
				"title": "S",
				/* this title
				  should be
				  more than 1 character
				  */
				"GlossList": {
					"GlossEntry": {
						"ID": "SGML",
						// anotehr commeht
						"SortAs": "SGML",
						"GlossTerm": "Standard Generalized Markup Language",
						"Acronym": "SGML",
						/* soemthing asdasda */
						"Abbrev": "ISO 8879:1986",
						"GlossDef": {
							"para": "A meta-markup language, used to create markup languages such as DocBook.",
							"GlossSeeAlso": [
								"GML",
								"XML"
							]
						},
						"GlossSee": "markup"
					}
				}
			}
		}
	}`,
			json: `{
		"glossary": {
			"title": "example glossary",

			"GlossDiv": {
				"title": "S",

				"GlossList": {
					"GlossEntry": {
						"ID": "SGML",

						"SortAs": "SGML",
						"GlossTerm": "Standard Generalized Markup Language",
						"Acronym": "SGML",

						"Abbrev": "ISO 8879:1986",
						"GlossDef": {
							"para": "A meta-markup language, used to create markup languages such as DocBook.",
							"GlossSeeAlso": [
								"GML",
								"XML"
							]
						},
						"GlossSee": "markup"
					}
				}
			}
		}
	}`,
		},
		{
			name:             "everything remove empty lints",
			removeEmptyLines: true,
			jsonc: `{
		"glossary": {
			"title": "example glossary",
			// example comment
			"GlossDiv": {
				"title": "S",
				/* this title
				  should be
				  more than 1 character
				  */
				"GlossList": {
					"GlossEntry": {
						"ID": "SGML",
						// anotehr commeht
						"SortAs": "SGML",
						"GlossTerm": "Standard Generalized Markup Language",
						"Acronym": "SGML",
						/* soemthing asdasda */
						"Abbrev": "ISO 8879:1986",
						"GlossDef": {
							"para": "A meta-markup language, used to create markup languages such as DocBook.",
							"GlossSeeAlso": [
								"GML",
								"XML"
							]
						},
						"GlossSee": "markup"
					}
				}
			}
		}
	}`,
			json: `{
		"glossary": {
			"title": "example glossary",
			"GlossDiv": {
				"title": "S",
				"GlossList": {
					"GlossEntry": {
						"ID": "SGML",
						"SortAs": "SGML",
						"GlossTerm": "Standard Generalized Markup Language",
						"Acronym": "SGML",
						"Abbrev": "ISO 8879:1986",
						"GlossDef": {
							"para": "A meta-markup language, used to create markup languages such as DocBook.",
							"GlossSeeAlso": [
								"GML",
								"XML"
							]
						},
						"GlossSee": "markup"
					}
				}
			}
		}
	}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := RemoveComments(tt.jsonc, tt.removeEmptyLines)
			assert.Equal(t, tt.json, res)
		})
	}
}
