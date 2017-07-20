/*
 * All classes in the java.lang package are imported by default. The classes
 * may then be referenced without their fully qualified names.
 *
 * import java.lang.*;
 * import java.lang.System;
 */

public class HelloWorld {

    /*
     * Being public, the method may be accessed by its class, package,
     * subclasses, and world. The JVM is accessing the main method from
     * outside this package.
     * Being static, the method belongs at the class level. The JVM does not
     * know how to construct class instances.
     * Being void, the method does not have a return type. The JVM is not
     * expecting a return type. However, the JVM does expect to be able to
     * pass command-line arguments.
     */
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }

}
