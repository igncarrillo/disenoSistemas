public class Slack implements Subscriber{
    @Override
    public void update(String msg) {
        System.out.println("Enviando por slack: "+msg);
    }
}
