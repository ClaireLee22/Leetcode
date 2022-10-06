class Solution(object):
    def canFinish(self, numCourses, prerequisites):
        """
        :type numCourses: int
        :type prerequisites: List[List[int]]
        :rtype: bool
        """
        graph = self.buildGraph(numCourses, prerequisites)
        print("adjacency list", graph)
        print("")
        visiting = set()
        visited = set()
        for node in range(numCourses):
            print("----------------------------------------------")
            print("dfs node", node)
            if self.detectCycle(graph, node, visited, visiting):
                return False
        
        return len(visited) == numCourses
            
    def detectCycle(self, graph, node, visited, visiting):
        if node in visited:
            print("node ", node, "is already visited(black) -> skip")
            return False
        
        if node in visiting:
            print("node ", node, "is in visiting(gray) -> a cycle is detected")
            print("")
            return True
        
        visiting.add(node)
        print("nodes in visiting(gray)", visiting)
        print("nodes in visited(black)", visited)
        print("")
        for descendant in graph[node]:
            print("current node:", node)
            print("visit descendant: node", descendant)
            if self.detectCycle(graph, descendant, visited, visiting):
                return True
            print("")
    
        visiting.remove(node)
        visited.add(node)
        print("finish explored node", node)
        print("nodes in visiting(gray)", visiting)
        print("nodes in visited(black)", visited)
        return False
        
        
    def buildGraph(self, numNodes, edges):
        graph = {x: [] for x in range(numNodes)}
        for edge in edges:
            a, b = edge
            graph[b].append(a)
        return graph
        

if __name__ == "__main__":
    solution = Solution()
    numCourses = 6
    prerequisites = [[1, 0], [3, 0], [5, 0], [2, 1], [3, 1], [4, 3], [5, 4], [3, 5]]
    print("can finish courses?", solution.canFinish(numCourses, prerequisites))

"""
output:
('adjacency list', {0: [1, 3, 5], 1: [2, 3], 2: [], 3: [4], 4: [5], 5: [3]})

----------------------------------------------
('dfs node', 0)
('nodes in visiting(gray)', set([0]))
('nodes in visited(black)', set([]))

('current node:', 0)
('visit descendant: node', 1)
('nodes in visiting(gray)', set([0, 1]))
('nodes in visited(black)', set([]))

('current node:', 1)
('visit descendant: node', 2)
('nodes in visiting(gray)', set([0, 1, 2]))
('nodes in visited(black)', set([]))

('finish explored node', 2)
('nodes in visiting(gray)', set([0, 1]))
('nodes in visited(black)', set([2]))

('current node:', 1)
('visit descendant: node', 3)
('nodes in visiting(gray)', set([0, 1, 3]))
('nodes in visited(black)', set([2]))

('current node:', 3)
('visit descendant: node', 4)
('nodes in visiting(gray)', set([0, 1, 3, 4]))
('nodes in visited(black)', set([2]))

('current node:', 4)
('visit descendant: node', 5)
('nodes in visiting(gray)', set([0, 1, 3, 4, 5]))
('nodes in visited(black)', set([2]))

('current node:', 5)
('visit descendant: node', 3)
('node ', 3, 'is in visiting(gray) -> a cycle is detected')

('can finish courses?', False)
"""