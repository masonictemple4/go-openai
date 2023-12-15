package main

type MenuItem struct {
	Index int64
	Name  string
}

// TODO: We'll define some menu items here.
func ShowSettingsMenu() {
	items := []MenuItem{
		{Index: 0, Name: "model"},
	}
	// TODO: Generate option index.
	for i := range items {
		item := items[i]
		println(item.Name)
	}
}
