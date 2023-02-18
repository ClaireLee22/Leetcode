from heapq import *
class MedianFinder(object):

    def __init__(self):
        self.minHeap = []
        self.maxHeap = []

    def addNum(self, num):
        """
        :type num: int
        :rtype: None
        """
        if not self.maxHeap or -self.maxHeap[0] >= num :
            heappush(self.maxHeap, -num)
        else:
            heappush(self.minHeap, num)
        
        self.balance()
    
    def balance(self):
        if len(self.maxHeap) > len(self.minHeap) + 1:
            heappush(self.minHeap, -heappop(self.maxHeap))
        elif len(self.maxHeap) < len(self.minHeap):
            heappush(self.maxHeap, -heappop(self.minHeap))

    def findMedian(self):
        """
        :rtype: float
        """
        if len(self.maxHeap) > len(self.minHeap):
            return -self.maxHeap[0]/1.0
        else:
            return -self.maxHeap[0]/2.0 + self.minHeap[0]/2.0

if __name__ == "__main__":
    dataStream = [2, 3, 7, 9]
    medianFinder = MedianFinder()

    for num in dataStream:
        print("add number", num)
        medianFinder.addNum(num)
        print("max heap", medianFinder.maxHeap)
        print("min heap", medianFinder.minHeap)
        print("median", medianFinder.findMedian())


"""
output:
('add number', 2)
('max heap', [-2])
('min heap', [])
('median', 2.0)

('add number', 3)
('max heap', [-2])
('min heap', [3])
('median', 2.5)

('add number', 7)
('max heap', [-3, -2])
('min heap', [7])
('median', 3.0)

('add number', 9)
('max heap', [-3, -2])
('min heap', [7, 9])
('median', 5.0)
"""