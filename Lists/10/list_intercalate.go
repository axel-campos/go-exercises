package ten

import "container/list"

func IntercalateLists(collection list.List, collection2 list.List) list.List {
	result := list.New()
	for element1, element2 := collection.Back(), collection2.Back(); element1 != nil && element2 != nil; element1, element2 = element1.Next(), element2.Next() {
		if element1 != nil {
			result.PushBack(element1)
		}
		if element2 != nil {
			result.PushBack(element2)
		}
	}
	return *result
}
