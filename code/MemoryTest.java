public class MemoryTest {
    public static void main(String[] args) {
        System.out.println("Total: " + (double) Runtime.getRuntime().totalMemory() 
                                     / (1024*1024) + " MiB");
        System.out.println("Used:  " + (double) (Runtime.getRuntime().totalMemory() 
                                     - Runtime.getRuntime().freeMemory()) 
                                     / (1024*1024) + " MiB");
    }
}