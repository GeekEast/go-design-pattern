package Observer

func removeFromSlice(observerList []IObserver, observerToRemove IObserver)  []IObserver{
	length := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getId() == observer.getId() {
			observerList[length - 1], observerList[i] = observerList[i], observerList[length - 1]
			return observerList[:length-1]
		}
	}
	return observerList
}
