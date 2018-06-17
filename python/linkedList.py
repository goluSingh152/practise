
import os



class Node:
    def __init__(self, data, next = None):
        self.num = data
        self.next = next

    def NodePrint(self):
        print(self.num)

    def SetNode(self, data):
        self.num = data

    def SetNext(self, next):
        self.next = next

    def NextNode(self):
        return self.next

class LinkedList(Node):

    def __init__(self):
        self.head = None

    def PrintNode(self):
        while self.head.next != None:
            print(self.head.num)
            self.head = self.head.next

    def InsertNode(self, data):
        temp = self.head
        if temp == None:
            self.head = Node(data)
        else:
            while temp.next != None:
                temp = temp.next
            temp.next = Node(data)
def main():
    head = LinkedList()
    head.InsertNode(2)
    head.head.next = Node(1)
    head.head.next.SetNext(None)
    head.InsertNode(3)
    head.InsertNode(4)
    head.InsertNode(5)
    head.InsertNode(6)
    head.PrintNode()

if __name__ == '__main__':
    main()
