// you can also use imports, for example:
// import java.util.*;

// you can write to stdout for debugging purposes, e.g.
// System.out.println("this is a debug message");

class Solution {
    public int solution(int A, int B, int C) {
        int result = 0;
        // (A + B + C) - ((A|B) + (A|C) + (B|C)) + (A|B|C)
        result = conformsSum(A) + conformsSum(B) + conformsSum(C)
                - conformsSum(A | B) - conformsSum(A | C) - conformsSum(B | C)
                + conformsSum(A | B | C);

        return result;
    }

    private int conformsSum(int integer) {
        // Turn the integer into zero filled binary string. 30 bits.
        String bitString = String.format("%30s", Integer.toBinaryString(integer)).replace(' ', '0');
        char[] bitArray = bitString.toCharArray();
        int zeros = 0;

        for (int i = 0; i < bitArray.length; i++) {
            // Only zero bits cause the conforming.
            if (bitArray[i] == '0') {
                zeros++;
            }
        }
        // Possiblilty is 2 to the power of number of zeros.
        return (int) Math.pow(2, zeros);
    }
}