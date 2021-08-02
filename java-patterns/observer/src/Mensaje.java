import java.util.ArrayList;

public class Mensaje implements Publisher{
    private String msg;
    private ArrayList<Subscriber> subscribers = new ArrayList<>();

    @Override
    public void addSubscriber(Subscriber s) {
        this.subscribers.add(s);
    }

    @Override
    public void removeSubscriber(Subscriber s) {
        this.subscribers.remove(s);
    }

    @Override
    public void notificar(String event) {
        this.msg=event;
        for (Subscriber s : subscribers){
            s.update(this.msg);
        }
    }
}
