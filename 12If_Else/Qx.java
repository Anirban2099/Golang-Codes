import java.util.*;
public class Qx{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        System.out.println("What is your age?");
        int age = sc.nextInt();

        if (age < 18)
        {
            System.out.println("You are not eligible for voting");
        }
        else {
            System.out.println("You are eligible for voting");
        }
    }}
