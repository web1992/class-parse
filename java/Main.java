
import java.io.IOException;

public class Main extends AbstractMain implements InterfaceMain {

    public static final Integer INT_MAX = Integer.MAX_VALUE;
    public static final Integer INT_MIN = Integer.MIN_VALUE;

    public static final Long LONG_MAX = Long.MAX_VALUE;
    public static final Long LONG_MIN = Long.MIN_VALUE;

    public static final Double DOUBLE_MAX = Double.MAX_VALUE;
    public static final Double DOUBLE_MIN = Double.MIN_VALUE;

    public static final Float FLOAT_MAX = Float.MAX_VALUE;
    public static final Float FLOAT_MIN = Float.MIN_VALUE;

    public static final Short SHORT_MAX = Short.MAX_VALUE;
    public static final Short SHORT_MIN = Short.MIN_VALUE;

    public static void main(String[] args) throws IOException {

        System.out.println("Hello word!");

        Runnable r = () -> {
            System.out.println("run");
        };

        r.run();
    }

    @Override
    public void say() {
        System.out.println("ha ha ha ha");
    }
}


interface InterfaceMain {

    void say();
}

abstract class AbstractMain {

}

