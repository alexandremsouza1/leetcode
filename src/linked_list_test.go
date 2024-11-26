package src


import (
    "testing"
)

func TestIsLinkedList(t *testing.T) {
    tests := []struct {
        name     string
        input    *ListNode
        expected bool
    }{
        {
            name:     "Lista vazia",
            input:    nil,
            expected: true,
        },
        {
            name: "Lista válida sem ciclo",
            input: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
            expected: true,
        },
        {
            name: "Lista com ciclo",
            input: func() *ListNode {
                node1 := &ListNode{1, nil}
                node2 := &ListNode{2, nil}
                node3 := &ListNode{3, nil}
                node1.Next = node2
                node2.Next = node3
                node3.Next = node1 // Ciclo aqui
                return node1
            }(),
            expected: false,
        },
    }

    for _, test := range tests {
        result := IsLinkedList(test.input)
        if result != test.expected {
            t.Errorf("%s: isLinkedList(%v) = %v; want %v", test.name, test.input, result, test.expected)
        }
    }
}