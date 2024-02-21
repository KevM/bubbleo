package data

type Artist struct {
	Name        string
	Description string
	Paintings   []Painting
}

type Painting struct {
	Title       string
	Description string
	Colors      []Color
}

type Color struct {
	RGB    string
	Sample string
}

func GetArtists() []Artist {
	pp := Artist{
		Name:        "Pablo Picasso",
		Description: "Pablo Picasso was a Spanish painter, sculptor, printmaker, ceramicist and theatre designer who spent most of his adult life in France.",
		Paintings: []Painting{
			{
				Title:       "Guernica",
				Description: "Guernica is a large 1937 oil painting on canvas by Spanish artist Pablo Picasso. One of Picasso's best known works, Guernica is regarded by many art critics as one of the most moving and powerful anti-war paintings in history.",
				Colors: []Color{
					{
						RGB:    "#000000", //black
						Sample: "Black of the bull's eye facing down a matador it does not see.",
					},
					{
						RGB:    "#FFFFFF", //white
						Sample: "White of the background rendered against the sun.",
					},
					{
						RGB:    "#808080", //grey
						Sample: "So grey the Spanish Civil War returned to present day.",
					},
				},
			},
		},
	}

	vermeer := Artist{
		Name:        "Johannes Vermeer",
		Description: "His luminous paintings are celebrated for their exquisite portrayal of light, form, and serene dignity.",
		Paintings: []Painting{
			{
				Title:       "Girl With A Pearl Earring",
				Description: "A young woman wearing an exotic dress and turban by dutch master's standards looking side long at the viewer with a large pearlescent earning.",
				Colors: []Color{
					{
						RGB:    "#ffff00", //yellow
						Sample: "The coat and turban have yellow accents with a golden glow.",
					},
					{
						RGB:    "#0000ff", //blue
						Sample: "Hat's blue and stunning as the model's distain.",
					},
					{
						RGB:    "#cc7722",
						Sample: "Skin tones so ochre they burst with sun burn.",
					},
				},
			},
			{
				Title:       "Girl Reading a Letter at an Open Window",
				Description: "This captivating painting depicts a young woman standing by an open window, absorbed in reading a letter.",
				Colors: []Color{
					{
						RGB:    "#00ff00", //green
						Sample: "If only the green curtain could speak.",
					},
					{
						RGB:    "#ff0000", //red
						Sample: "The window drapes are the cheeryest thing in the room.",
					},
				},
			},
			{
				Title:       "Milkmaid",
				Description: "The scene exudes domestic tranquility and everyday beauty",
				Colors: []Color{
					{
						RGB:    "#ffff00",
						Sample: "Yellow was a popular color for the Dutch.",
					},
					{
						RGB:    "#fdfff5",
						Sample: "There is a milk the color of her bonet",
					},
				},
			},
		},
	}

	return []Artist{pp, vermeer}
}
