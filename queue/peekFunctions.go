package queue

//PeekElem returns a copy of the elem that would be returned on a call to Remove().
//Returns an error of type EmptyListError when the list is empty.
func (q *Queue) PeekElem() (float64, interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.numElems == 0 {
		return 0, nil, EmptyListError{}
	}
	elem := *q.queSlice[q.numElems-1] //dereference is a copy
	return elem.priority, elem.content, nil
}

//PeekElemAtIndex returns a copy of the elem at index.
//Returns an error of type EmtyListError when the list is empty.
//Returns an error of type IndexOutOfBoundsError when the provided index is out of bounds.
func (q *Queue) PeekElemAtIndex(index int) (float64, interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.numElems == 0 {
		return 0, nil, EmptyListError{}
	}

	realIndex := (q.numElems - 1) - index
	if realIndex < 0 {
		return 0, nil, IndexOutOfBoundsError{}
	}

	elem := *q.queSlice[realIndex] //dereference is a copy
	return elem.priority, elem.content, nil
}
