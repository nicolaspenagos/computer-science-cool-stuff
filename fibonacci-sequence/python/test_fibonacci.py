import unittest
from fibonacci import (
    slow_recursive_fibonacci,
    fast_recursive_fibonacci,
    iterative_fibonacci)

class TestFibonacciMethods(unittest.TestCase):
    
    def test_slow_recursive_fibonacci(self):
        self.assertEqual(slow_recursive_fibonacci(0), 0)
        self.assertEqual(slow_recursive_fibonacci(1), 1)
        self.assertEqual(slow_recursive_fibonacci(5), 5)
        self.assertEqual(slow_recursive_fibonacci(10), 55)
        self.assertEqual(slow_recursive_fibonacci(15), 610)
    
    def test_fast_recursive_fibonacci(self):
        self.assertEqual(fast_recursive_fibonacci(0), 0)
        self.assertEqual(fast_recursive_fibonacci(1), 1)
        self.assertEqual(fast_recursive_fibonacci(5), 5)
        self.assertEqual(fast_recursive_fibonacci(10), 55)
        self.assertEqual(fast_recursive_fibonacci(15), 610)
    
    def test_iterative_fibonacci(self):
        self.assertEqual(iterative_fibonacci(0), 0)
        self.assertEqual(iterative_fibonacci(1), 1)
        self.assertEqual(iterative_fibonacci(5), 5)
        self.assertEqual(iterative_fibonacci(10), 55)
        self.assertEqual(iterative_fibonacci(15), 610)

if __name__ == '__main__':
    unittest.main()
