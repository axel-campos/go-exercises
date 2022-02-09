package nine

import "container/list"

func Concatenate(collection list.List, collection2 list.List) list.List {
	result := list.New()
	for element := collection.Back(); element != nil; element = collection.Back().Next() {
		result.PushBack(element)
	}
	for element := collection2.Back(); element != nil; element = collection2.Back().Next() {
		result.PushBack(element)
	}
	return *result
}

func ConcatenateBuildInFunction(collection list.List, collection2 list.List) list.List {
	collection.PushFrontList(&collection2)
	return collection
}
