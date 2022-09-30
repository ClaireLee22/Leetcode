class LRUCache(object):

    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.capacity = capacity
        self.cache = {}
        self.linkedList = LinkedList()  

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        if key not in self.cache:
            return -1
        
        value = self.cache[key].value
        self.linkedList.removeNode(self.cache[key])
        self.linkedList.insertAtTail(key, value)
        self.cache[key] = self.linkedList.getTail()
        
        return value
        

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: None
        """
        if key not in self.cache:
            if self.linkedList.size >= self.capacity:
                self.linkedList.insertAtTail(key, value)
                self.cache[key] = self.linkedList.getTail()
                del self.cache[self.linkedList.getHead().key]
                self.linkedList.removeHead()
            else:
                self.linkedList.insertAtTail(key, value)
                self.cache[key] = self.linkedList.getTail()
        else:
            self.linkedList.removeNode(self.cache[key])
            self.linkedList.insertAtTail(key, value)
            self.cache[key] = self.linkedList.getTail()
    
    def printLinkedList(self):
        current = self.linkedList.head
        elements = []
        while current is not None:
            elements.append(str(current.value))
            current = current.next
        return " -> ".join(elements)
        
class LinkedListNode:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.next = None
        self.prev = None

class LinkedList:
    def __init__(self):
        self.size = 0
        self.head = None
        self.tail = None
    
    def removeNode(self, node):
        if node is None:
            return
        if node.prev is not None:
            node.prev.next = node.next
        if node.next is not None:
            node.next.prev = node.prev
        if node == self.head:
            self.head = self.head.next
        if node == self.tail:
            self.tail = self.tail.prev
        
        self.size -= 1
    
    def insertAtTail(self, key, value):
        newNode = LinkedListNode(key, value)
        if self.tail is None:
            self.head = newNode
            self.tail = newNode
        else:
            self.tail.next = newNode
            newNode.prev = self.tail
            self.tail = newNode
        
        self.size += 1
    
    def removeHead(self):
        self.removeNode(self.head)
    
    def getTail(self):
        return self.tail
    
    def getHead(self):
        return self.head


if __name__ == "__main__":
    lruCache = LRUCache(2)
    lruCache.put(1, 1)
    print("put(1, 1)")
    print("cache: ", lruCache.cache)
    print("doubly linked list: ", lruCache.printLinkedList())

    lruCache.put(2, 2)
    print("\nput(2, 2)")
    print("cache: ", lruCache.cache)
    print("doubly linked list: ", lruCache.printLinkedList())

    print("\nget(1)")
    print("value: ", lruCache.get(1))
    print("cache: ", lruCache.cache)
    print("doubly linked list: ", lruCache.printLinkedList())

    lruCache.put(3, 3)
    print("\nput(3, 3)")
    print("cache: ", lruCache.cache)
    print("doubly linked list: ", lruCache.printLinkedList())

    print("\nget(2)")
    print("value: ", lruCache.get(2))
    print("cache: ", lruCache.cache)
    print("doubly linked list: ", lruCache.printLinkedList())

    lruCache.put(4, 4)
    print("\nput(4, 4)")
    print("cache: ", lruCache.cache)
    print("doubly linked list: ", lruCache.printLinkedList())

    print("\nget(1)")
    print("value: ", lruCache.get(1))
    print("cache: ", lruCache.cache)
    print("doubly linked list: ", lruCache.printLinkedList())

    print("\nget(3)")
    print("value: ", lruCache.get(3))
    print("cache: ", lruCache.cache)
    print("doubly linked list: ", lruCache.printLinkedList())

    print("\nget(4)")
    print("value: ", lruCache.get(4))
    print("cache: ", lruCache.cache)
    print("doubly linked list: ", lruCache.printLinkedList())

"""
output:
put(1, 1)
('cache: ', {1: <__main__.LinkedListNode instance at 0x107789cf8>})
('doubly linked list: ', '1')

put(2, 2)
('cache: ', {1: <__main__.LinkedListNode instance at 0x107789cf8>, 2: <__main__.LinkedListNode instance at 0x107789d40>})
('doubly linked list: ', '1 -> 2')

get(1)
('value: ', 1)
('cache: ', {1: <__main__.LinkedListNode instance at 0x107789d88>, 2: <__main__.LinkedListNode instance at 0x107789d40>})
('doubly linked list: ', '2 -> 1')

put(3, 3)
('cache: ', {1: <__main__.LinkedListNode instance at 0x107789d88>, 3: <__main__.LinkedListNode instance at 0x107789cf8>})
('doubly linked list: ', '1 -> 3')

get(2)
('value: ', -1)
('cache: ', {1: <__main__.LinkedListNode instance at 0x107789d88>, 3: <__main__.LinkedListNode instance at 0x107789cf8>})
('doubly linked list: ', '1 -> 3')

put(4, 4)
('cache: ', {3: <__main__.LinkedListNode instance at 0x107789cf8>, 4: <__main__.LinkedListNode instance at 0x107789d40>})
('doubly linked list: ', '3 -> 4')

get(1)
('value: ', -1)
('cache: ', {3: <__main__.LinkedListNode instance at 0x107789cf8>, 4: <__main__.LinkedListNode instance at 0x107789d40>})
('doubly linked list: ', '3 -> 4')

get(3)
('value: ', 3)
('cache: ', {3: <__main__.LinkedListNode instance at 0x107789d88>, 4: <__main__.LinkedListNode instance at 0x107789d40>})
('doubly linked list: ', '4 -> 3')

get(4)
('value: ', 4)
('cache: ', {3: <__main__.LinkedListNode instance at 0x107789d88>, 4: <__main__.LinkedListNode instance at 0x107789cf8>})
('doubly linked list: ', '3 -> 4')

"""

