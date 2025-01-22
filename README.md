# algorithm practice #

###  Search Algorithms
1. Linear Search (O(n)) <!--线性搜索，暴力查找-->
2. Binary Search (O(log n)) <!--二分查找,适用于有序数组-->
3. Interpolation Search (best: O(log log n), worst: O(n)) <!--插值查找，类似二分，同样需要有序数组，不同于二分的是中间值的选取，有特殊计算公式-->
4. Jump Search (O(&#8730;n)) <!--跳跃查找，适用于有序数组，根号n大小分组-->
5. Fibonacci Search (O(log n)) <!--斐波那契查找，适用于有序数组-->
6. Hash Search (ave: O(1), worst: O(n)) <!--哈希查找，如果产生哈希冲突，最坏时间复杂度会退化到O(n),具体却决于处理冲突的方法，常见的有链式哈希法，开放定址法等-->
7. Block Search (O(&#8730;n)) <!--分块查找，适用于有序数组-->
8. B-tree Search (O(log n)) <!--B树查找，适用于需要频繁CURD的系统，常用于数据库个文件系统-->
9. AVL Tree Search (O(log n)) <!--AVL树查找，适用于动态数据集-->
10. Red-Black Tree Search (O(log n)) <!--红黑树查找，适用于动态数据集-->
11. Breadth-First Search, BFS (O(V+E) v:node e:edge) <!--广度优先搜索，v是节点数，e是边数-->
12. Depth-First Search, DFS (O(V+E) v:node e:edge) <!--深度优先搜索，v是节点数，e是边数-->