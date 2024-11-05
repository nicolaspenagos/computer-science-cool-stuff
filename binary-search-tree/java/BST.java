package java;

public class BST {
    private Node node = null;

    public BST insert(int data) {

        if (this.node == null) {
            node = new Node(data);
            return this;
        }

        insert(node, data);
        return this;
    }

    private Node insert(Node node, int data) {
        if (node == null) {
            return new Node(data);
        }

        if (node.getData() > data) {
            node.setLeft(insert(node.getLeft(), data));
            return node;
        } else {
            node.setRight(insert(node.getRight(), data));
            return node;
        }
    }

    public void retrieve() {
        if (this.node != null) {
            if (this.node.getLeft() != null) {
                retrieve(this.node.getLeft());
            }
            System.out.println(this.node.getData());
            if (this.node.getRight() != null) {
                retrieve(this.node.getRight());
            }
        } else {
            System.out.println("Arbol vacio");
        }
    }

    private void retrieve(Node node) {
        if (node != null) {
            if (node.getLeft() != null) {
                retrieve(node.getLeft());
            }
            System.out.println(node.getData());
            if (node.getRight() != null) {
                retrieve(node.getRight());
            }
        }
    }


}
