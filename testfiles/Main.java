
import java.io.IOException;
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@FFF()
public class Main extends AbstractMain<String> implements InterfaceMain {

    @FFF
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

        int a = 1;
        String A = "";
        switch (a) {
            case 1: {
                A = "a";
            }
            case 2: {
                A = "c";
            }
            case 3: {
                A = "D";
            }
            default: {
                A = "Z";
            }
        }

        try {
            System.out.println("try");
        } catch (Exception e) {
            System.out.println("catch");
        } finally {
            System.out.println("finally");
        }
    }

    @Override
    @FFF("3FFF")
    @Name(name = "My Name", index = 666)
    public void say() {
        System.out.println("ha ha ha ha");
    }

    @Override
    String hi() {
        return "Hi,2021";
    }
}


interface InterfaceMain {
    void say();
}

abstract class AbstractMain<T> {
    abstract T hi();
}

@Target({ElementType.METHOD, ElementType.FIELD, ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
@interface FFF {
    String value() default "FFF_DEFAULT";
}

@Target({ElementType.METHOD, ElementType.FIELD, ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
@interface Name {
    String name() default "Name";

    int index() default -1;
}

