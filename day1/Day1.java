import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Scanner;

public class Day1 {

    public static void main(String[] args) throws IOException {
        Scanner f = new Scanner(new File("test.txt"));

        ArrayList<Integer> list1 = new ArrayList<Integer>();
        ArrayList<Integer> list2 = new ArrayList<Integer>();
        while (f.hasNextLine()) {
            list1.add(f.nextInt());
            list2.add(f.nextInt());
            f.nextLine();
        }
        f.close();
        list1.sort(null);
        list2.sort(null);
        System.out.println(list1 + "\n" + list2);
        int sum = 0;
        for (int i = 0; i < list1.size(); i++) {
            sum +=
                Math.max(list1.get(i), list2.get(i)) -
                Math.min(list1.get(i), list2.get(i));
        }
        System.out.print(sum);
    }
}
