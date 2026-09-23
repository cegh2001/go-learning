package entrevista

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList invierte una lista enlazada in-place (reorientando punteros).
func ReverseList(head *ListNode) *ListNode {
	// TODO: Implementar inversión in-place O(1) memoria
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next // 1. Guardo el siguiente nodo para no perderlo
		curr.Next = prev  // 2. Doy vuelta la flecha (ahora apunto hacia atrás)
		prev = curr       // 3. Muevo 'prev' un paso adelante
		curr = next       // 4. Muevo 'curr' un paso adelante
	}
	return prev
}
