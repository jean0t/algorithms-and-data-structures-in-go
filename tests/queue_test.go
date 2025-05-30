package tests

import (
    "reflect"
    "errors"
    "testing"
    ds "github.com/jean0t/algorithms-and-data-structures-in-go/data_structures"
)

func getQueueItems(q *ds.Queue) ([]int, error) {
    var itemsField reflect.Value = reflect.ValueOf(q).Elem().FieldByName("Items")

    if !itemsField.IsValid() {
        return nil, errors.New("Items field not valid")
    }

    items, ok := itemsField.Interface().([]int)
    if !ok {
        return nil, errors.New("queue.items field is not of type []int")
    }

    return items, nil
}

func TestNewQueue(t *testing.T) {
    var q *ds.Queue = ds.NewQueue()

    if q == nil {
        t.Error("NewQueue returned a nil queue, expected non nil.")
    }

    items, err := getQueueItems(q)
    if err != nil {
        t.Fatal(err)
    }

    if len(items) != 0 {
        t.Errorf("NewQueue() returned a queue with %d items, expected 0.", len(items))
    }
}


func TestEnqueue(t *testing.T) {
    tests := []struct {
        name string
        initialItems []int
        itemToEnqueue int
        expectedItems []int
    }{
        {
            name: "Enqueue to empty queue",
            initialItems: []int{},
            itemToEnqueue: 2,
            expectedItems: []int{2},
        },
        {
            name: "Enqueue to non-empty queue",
            initialItems: []int{10},
            itemToEnqueue: 2,
            expectedItems: []int{10, 2},
        },
        {
            name: "Enqueue zero",
            initialItems: []int{10, 2},
            itemToEnqueue: 0,
            expectedItems: []int{10, 2, 0},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var q *ds.Queue = ds.NewQueue()
            for _, item := range tt.initialItems {
                q.Enqueue(item)
            }

            q.Enqueue(tt.itemToEnqueue)
            actualItems, err := getQueueItems(q)
            if err != nil {
                t.Fatal(err)
            }

            if !reflect.DeepEqual(actualItems, tt.expectedItems) {
                t.Errorf("Enqueue method failed, got = %v; want %v\n", actualItems, tt.expectedItems)
            }
        })
    }
}


func TestDequeue(t *testing.T) {
    tests := []struct {
        name string
        initialItems []int
        expectedDequeue int
        expectedRemaining []int
        expectPanic bool
    }{
        {
            name: "Dequeue with multiple elements in queue",
            initialItems: []int{5, 3, 1},
            expectedDequeue: 5,
            expectedRemaining: []int{3, 1},
            expectPanic: false,
        },
        {
            name: "Dequeue with only one element in queue",
            initialItems: []int{7},
            expectedDequeue: 7,
            expectedRemaining: []int{},
            expectPanic: false,
        },
        {
            name: "Dequeue with empty queue",
            initialItems: []int{},
            expectedDequeue: 0,
            expectedRemaining: []int{},
            expectPanic: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.expectPanic {
                defer func() {
                    if r := recover(); r == nil {
                        t.Errorf("[%s] Dequeue didn't panic as expected", tt.name)
                    } else {}
                }()
            }

            var q *ds.Queue = ds.NewQueue()

            for _, item := range tt.initialItems {
                q.Enqueue(item)
            }

            var actualDequeued int = q.Dequeue()

            if tt.expectPanic {
                t.Error("Panic expected but not raised")
            }

            if actualDequeued != tt.expectedDequeue {
                t.Errorf("Dequeued failed. got = %d; wanted %d", actualDequeued, tt.expectedDequeue)
            }

            actualRemaining, err := getQueueItems(q)
            if err != nil {
                t.Fatal(err)
            }

            if !reflect.DeepEqual(actualRemaining, tt.expectedRemaining) {
                t.Errorf("Expected remaining %v; got = %v", tt.expectedRemaining, actualRemaining)
            }
        })
    }
}


func TestMixedOperations(t *testing.T) {
    t.Run("Mixed operations", func(t *testing.T) {
        var q *ds.Queue = ds.NewQueue()

        q.Enqueue(2)
        q.Enqueue(5)
        if d := q.Dequeue(); d != 2 {
            t.Errorf("Expected 2, got = %d", d)
        }

        q.Enqueue(3)
        if d := q.Dequeue(); d != 5 {
            t.Errorf("Expected 2, got = %d", d)
        }

        q.Enqueue(8)
        q.Enqueue(3)

        var actualItems, err = getQueueItems(q)
        if err != nil {
            t.Fatal(err)
        }

        if len(actualItems) != 3 {
            t.Errorf("Expected length of 3, got %d", len(actualItems))
        }
    })
}
