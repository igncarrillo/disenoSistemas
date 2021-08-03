public class Email implements Subscriber{
    public void update(String msg) {
        System.out.println("Enviando por email: "+msg);
    }
}
