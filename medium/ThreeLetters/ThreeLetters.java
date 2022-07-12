// you can also use imports, for example:
// import java.util.*;

// you can write to stdout for debugging purposes, e.g.
// System.out.println("this is a debug message");

class Solution {
    public String solution(int A, int B) {
        // The one with more letters shall be the first letter.
        // Which will have 1 additional gap for it to put in.
        String firstLetter = "a";
        String secondLetter = "b";
        int firstLetterCount = A;
        int secondLetterCount = B;
        int extraLetters = A - B;
        int loops = B;

        if (B > A) {
            firstLetter = "b";
            secondLetter = "a";
            firstLetterCount = B;
            secondLetterCount = A;
            extraLetters = B - A;
            loops = A;
        }
        // Loop through all the possible gaps, which is the length of the shorter
        // letters plus 1.
        // e.g. _a_a_ <- Those underscopes indicates the gap we want to loop through.
        loops++;

        String answer = "";
        extraLetters--;
        for (int i = 0; i < loops; i++) {
            if (firstLetterCount > 0) {
                answer += firstLetter;
                firstLetterCount--;
                // Consume those extra letters to form a double letter e.g. 'aa'
                if (extraLetters > 0) {
                    answer += firstLetter;
                    extraLetters--;
                    firstLetterCount--;
                }
            }
            if (secondLetterCount > 0) {
                answer += secondLetter;
                secondLetterCount--;
            }
        }

        return answer;
    }
}