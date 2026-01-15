package bubletea

import (
	"setup/models"
)

func initialModel() models.BubbleTea {
	return models.BubbleTea{
		// Our to-do list is a grocery list
		Choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		Selected: make(map[int]struct{}),
	}
}
