import unittest
from Stack import Stack

class TestStack(unittest.TestCase):

    def setUp(self):
        self.stack = Stack()

    def test_is_empty(self):
        self.assertTrue(self.stack.is_empty())
        self.stack.push(1)
        self.assertFalse(self.stack.is_empty())

    def test_push(self):
        self.stack.push(1)
        self.assertEqual(self.stack.size(), 1)
        self.stack.push(2)
        self.assertEqual(self.stack.size(), 2)

    def test_pop(self):
        self.stack.push(1)
        self.stack.push(2)
        self.assertEqual(self.stack.pop(), 2)
        self.assertEqual(self.stack.pop(), 1)
        with self.assertRaises(IndexError):
            self.stack.pop()

    def test_peek(self):
        self.stack.push(1)
        self.stack.push(2)
        self.assertEqual(self.stack.peek(), 2)
        self.stack.pop()
        self.assertEqual(self.stack.peek(), 1)
        self.stack.pop()
        with self.assertRaises(IndexError):
            self.stack.peek()

    def test_size(self):
        self.assertEqual(self.stack.size(), 0)
        self.stack.push(1)
        self.assertEqual(self.stack.size(), 1)
        self.stack.push(2)
        self.assertEqual(self.stack.size(), 2)
        self.stack.pop()
        self.assertEqual(self.stack.size(), 1)

    def test_str(self):
        self.assertEqual(str(self.stack), "[]")
        self.stack.push(1)
        self.assertEqual(str(self.stack), "[1]")
        self.stack.push(2)
        self.assertEqual(str(self.stack), "[1, 2]")

    def test_len(self):
        self.assertEqual(len(self.stack), 0)
        self.stack.push(1)
        self.assertEqual(len(self.stack), 1)
        self.stack.push(2)
        self.assertEqual(len(self.stack), 2)
        self.stack.pop()
        self.assertEqual(len(self.stack), 1)

if __name__ == '__main__':
    unittest.main()