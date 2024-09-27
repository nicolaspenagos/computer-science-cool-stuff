def slow_recursive_fibonacci(n:int)->int:
    if n == 0:
        return 0
    if n == 1:
        return 1
    return slow_recursive_fibonacci(n - 1) + slow_recursive_fibonacci(n - 2)

def fast_recursive_fibonacci(n:int)->int:
    if n == 0:
        return 0
    if n == 1:
        return 1
    
    results = {}
    
    if n not in results:
        results[n] = fast_recursive_fibonacci(n - 1) + fast_recursive_fibonacci(n - 2)
    
    return results[n]

def iterative_fibonacci(n:int)->int:
    a1 = 0
    a2 = 1
    
    if n == 0:
        return a1
    if n == 1:
        return a2
    
    a_n = a1
    for _ in range(2, n + 1):
        a_n = a1 + a2
        a1 = a2
        a2 = a_n  
        
    return a_n
            

