from Stack import Stack

def is_matched(expreression):
    left_symbols = '({['
    right_symbols = ')}]'
    S = Stack()
    for symbol in expreression:
        if symbol in left_symbols:
            S.push(symbol)
        elif symbol in right_symbols:
            if S.is_empty():
                return False
            if right_symbols.index(symbol) != left_symbols.index(S.pop()):
                return False
    return S.is_empty()

def remove_non_delimiters(expression):
    delimiters = '(){}[]'
    return ''.join([char for char in expression if char in delimiters])


if __name__ == "__main__":
    input = "(a+b[5]*c)[{2x+3y}+4]"
    print(remove_non_delimiters(input)) # ()[][{x+y}+]
    print(is_matched(remove_non_delimiters(input))) # True