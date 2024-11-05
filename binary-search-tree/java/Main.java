package java;

public class Main {
    public static void main(String[] args) {
        BST bst = new BST();
        bst.retrieve();
        bst.insert(5);
        bst.insert(4);
        bst.insert(6);
        bst.insert(3);
        bst.insert(8);
        bst.insert(2);
        bst.insert(7);
        bst.insert(1);
        bst.insert(53);
        bst.insert(-12);
        bst.retrieve();
    }
}