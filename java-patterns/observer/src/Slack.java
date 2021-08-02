public class Slack implements Subscriber{
    @Override
    public void update(String msg) {
        System.out.println("Enviando mensaje por slack: "+msg);
    }
}
