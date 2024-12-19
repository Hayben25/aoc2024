import java.io.File;
import java.io.IOException;
import java.util.Scanner;

public class Day3 {

    public static void main(String[] args) throws IOException {
        Scanner f = new Scanner(new File("test.txt"));
        int tot = 0;
        while (f.hasNextLine()) {
            tot += check(f.nextLine());
        }
        f.close();
        System.out.println(tot);
    }

    public static int check(String s) {
        String part;
        int tot = 0;
        while (s.indexOf("mul(") != -1) {
            s = s.substring(s.indexOf("mul(") + 4);
            System.out.println(s + " " + s.indexOf(",") + " " + s.indexOf(")"));
            if (
                s.indexOf(")") <= 8 &&
                s.indexOf(",") <= 4 &&
                s.indexOf(")") >= 3 &&
                s.indexOf(",") >= 1
            ) {
                part = s.substring(0, s.indexOf(")"));
                tot +=
                    Integer.parseInt(part.substring(0, part.indexOf(","))) *
                    Integer.parseInt(part.substring(part.indexOf(",") + 1));
            }
        }
        return tot;
    }
}
