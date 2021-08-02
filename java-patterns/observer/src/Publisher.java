public interface Publisher {
    void addSubscriber(Subscriber s);
    void removeSubscriber(Subscriber s);
    void notificar(String event);
}
