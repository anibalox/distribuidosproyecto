package main

var ColaEspera [4]string //var ColaEspera = make([]string, 0)

/*
func enqueue(ColaEspera [4]string, element string) []string {
	ColaEspera = append(queue, element) // Simply append to enqueue.
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []string) (string, []string) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		var tmp = []string{}
		return element, tmp
	}
	return element, queue[1:] // Slice off the element once it is dequeued.
}
*/
func main() {
	println(ColaEspera)
	ColaEspera[0] = "1"
	println(ColaEspera)
}
