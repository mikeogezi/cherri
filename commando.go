package main

func commandoActions() {
	actions["askCommando"] = &actionDefinition{
		identifier: "co.getcommando.CommandoAskIntent",
		parameters: []parameterDefinition{
			{
				name:      "prompt",
				validType: String,
				key:       "prompt",
			},
		},
		addParams: func(args []actionArgument) []plistData {
			return []plistData{
				{
					key:      "ShowWhenRun",
					dataType: Boolean,
					value:    false,
				},
			}
		},
	}
	actions["end"] = &actionDefinition{
		identifier: "co.getcommando.CommandoEndIntent",
		parameters: []parameterDefinition{
			{
				name:      "sessionId",
				validType: String,
				key:       "sessionId",
			},
			{
				name:      "finalResults",
				validType: String,
				key:       "finalResults",
			},
		},
	}
	actions["continue"] = &actionDefinition{
		identifier: "co.getcommando.CommandoContinueIntent",
		parameters: []parameterDefinition{
			{
				name:      "sessionId",
				validType: String,
				key:       "sessionId",
			},
			{
				name:      "intermediateResults",
				validType: String,
				key:       "intermediateResults",
			},
		},
	}
}
