import unittest
from matching_delimeters import is_matched

class TestIsMatched(unittest.TestCase):

    def test_balanced_parentheses(self):
        self.assertTrue(is_matched("()"))
        self.assertTrue(is_matched("(())"))
        self.assertTrue(is_matched("((()))"))

    def test_unbalanced_parentheses(self):
        self.assertFalse(is_matched("("))
        self.assertFalse(is_matched(")"))
        self.assertFalse(is_matched("(()"))
        self.assertFalse(is_matched("())"))

    def test_balanced_mixed_symbols(self):
        self.assertTrue(is_matched("()[]{}"))
        self.assertTrue(is_matched("{[()]}"))
        self.assertTrue(is_matched("[({})]"))

    def test_unbalanced_mixed_symbols(self):
        self.assertFalse(is_matched("([)]"))
        self.assertFalse(is_matched("{[(])}"))
        self.assertFalse(is_matched("({[})]"))

    def test_empty_string(self):
        self.assertTrue(is_matched(""))

if __name__ == '__main__':
    unittest.main()