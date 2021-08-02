public class Client {
    public static void main(String[] args) {
        Publisher m = new Mensaje();
        Subscriber e = new Email();
        Subscriber s = new Slack();

        m.addSubscriber(e);
        m.addSubscriber(s);

        m.notificar("Nuevo mensaje");

        m.removeSubscriber(e);

        m.notificar("Este mensaje solo va por slack");

    }
}
