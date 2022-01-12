Single Cycle Check

package main

func HasSingleCycle(array []int) bool {
	
	var visited =0
	var currentIdx =0
	
	for(visited < len(array)){
		if visited >0 && currentIdx ==0{return false}
		visited +=1
		currentIdx = getNext(currentIdx,array)
	}
	return currentIdx ==0
}

func getNext(currentIdx int, array []int )int{
	
	var jump = array[currentIdx]
	var nextIdx = (jump + currentIdx) % len(array)
	
	if(nextIdx>=0){return nextIdx}
	return nextIdx + len(array)
}
