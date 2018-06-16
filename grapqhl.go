package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type mapType struct {
	Universe   uint16
	Channel    uint16
	KeyboardID int
	Keycode    uint16
}

type sACNtype struct {
	Universe     uint16
	Multicast    bool
	Destinations []string
}

type deviceType struct {
	Name      string
	Id        int
	Listening bool
}

func initGraphql() {
	deviceType := graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "Device",
			Description: "a Device contains an ID and a Name of the USB Device",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.Int,
					Description: "the unique Id of the Device, mostly staring at 0",
				},
				"name": &graphql.Field{
					Type:        graphql.String,
					Description: "The Name of the Device",
				},
				"listening": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "If true, on the current device is a listener that handles events",
				},
			},
		},
	)

	sacnType := graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "sACN",
			Description: "sACN defines an sACN output with all the parameters",
			Fields: graphql.Fields{
				"universe": &graphql.Field{
					Type:        graphql.Int,
					Description: "the universe to which the sACN output is set",
				},
				"multicast": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "true if the output sends out via multicast",
				},
				"destinations": &graphql.Field{
					Type:        graphql.NewList(graphql.String),
					Description: "a list with all unicast destinations to which sACN is send out",
				},
			},
		},
	)

	mapType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "KeyMap",
			Description: `KeyMap stores information about a key and where it should be mapped to.
			It stores the key's keyboard id and key code. `,
			Fields: graphql.Fields{
				"keyboardID": &graphql.Field{
					Type:        graphql.Int,
					Description: "the keyboard ID of to which the keycode belongs. This is used to use multiple keyboards/inputdevices at the same time",
				},
				"keycode": &graphql.Field{
					Type:        graphql.Int,
					Description: "the keycode of the key that was pressed. Different on every keyboard",
				},
				"universe": &graphql.Field{
					Type:        graphql.Int,
					Description: "the sACN universe that is used to send out the key stroke",
				},
				"channel": &graphql.Field{
					Type:        graphql.Int,
					Description: "the DMX channel of the given universe. Between: [0-511]",
				},
			},
		},
	)

	rootQuery := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"Devices": &graphql.Field{
				Type:        graphql.NewList(deviceType),
				Description: "a list of Devices on the host at the current time",
				Resolve:     queryDevices,
			},
			"sACN": &graphql.Field{
				Type:        graphql.NewList(sacnType),
				Description: "a list of all sACN outputs",
				Resolve:     querySacnOutputs,
			},
			"Mapping": &graphql.Field{
				Type: graphql.NewList(mapType),
				Description: `A full list of all mapping information currently available. Every entry
				should be unique at least with keycode and keyboardID.`,
				Resolve: queryMapping,
			},
		}}

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"sACN": &graphql.Field{
				Type: sacnType, // the return type for this field
				Description: `starts a new sACN output with the given parameters. 
				If the output is already running, only the changed data (destinations, multicast) 
				will be set`,
				Args: graphql.FieldConfigArgument{
					"universe": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"destinations": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},
					"multicast": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
				},
				Resolve: mutateSacnOutput,
			},
			"stopSACN": &graphql.Field{
				Type: graphql.Boolean,
				Description: `tries to stop the output of the given universe. If something was stopped,
				the return value is true, otherwise false`,
				Args: graphql.FieldConfigArgument{
					"universe": &graphql.ArgumentConfig{
						Type:        graphql.Int,
						Description: "the universe to stop sending out sACN data",
					},
				},
				Resolve: mutateStopSacn,
			},
			"KeyMap": &graphql.Field{
				Type: mapType,
				Description: `Add or update a keyMap. If the keyMap does not exist already, it will
				automaticlyy created. The unique fileds are: keycode and keyboardID`,
				Args: graphql.FieldConfigArgument{
					"universe": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "the universe to use. Has to be [0-63999]",
					},
					"channel": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "the channel to which the DMX data is send. HAs to be [0-511]",
					},
					"keycode": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "the key on which the sending should be triggered. Is uint16",
					},
					"keyboardID": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "the keyboard on which the key is located",
					},
				},
				Resolve: mutateKeyMap,
			},
			"setListening": &graphql.Field{
				Type: deviceType,
				Description: `set wether to listen on the given device or not. Set the device id and 
				true to start listening. If the listening was stopped, the return value is null.`,
				Args: graphql.FieldConfigArgument{
					"deviceID": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "the device id of the device to change listening on",
					},
					"listen": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
						Description: `if true, start listening. If the device is already been 
						listened, nothing will happen.`,
					},
				},
				Resolve: mutateKeyListener,
			},
			"writeConfig": &graphql.Field{
				Type: graphql.Boolean,
				Description: `when this gets called, the current config (sACN, Mapping, DeviceListening)
				will be saved in the config.json next to the executable. If the writing was succesful,
				true will be returned, otherwise an error will be returned and false. It is
				recommended to call this separately and not together with other mutations.`,
				Resolve: mutateWriteConfig,
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: rootMutation,
	})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, //hosts a graphiql instance under the entrypoint
	})
	// serve HTTP
	http.Handle("/graphql", h)
}
