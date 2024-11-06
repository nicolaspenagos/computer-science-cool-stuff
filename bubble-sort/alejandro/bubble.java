package Algorithms.java.alejandro;

public class bubble {
    public static void main(String[] args) {
        int[] myArray = {3, 5, 6, 2, 1, 9, 7};

        for (int i = 0; i < myArray.length; i++) {
            for (int j = 0; j < myArray.length -1 - i; j++) {
                if ((myArray[j+1]) < myArray[j]) {
                    var temporalVar = myArray[j];
                    myArray[j] = myArray[j+1];
                    myArray[j+1] = temporalVar;
                }
            }
        }

        for (int i = 0; i < myArray.length; i++) {
            System.out.println(myArray[i]);
        }
    }
}
