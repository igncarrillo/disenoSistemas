public class LogisticaMaritima extends Logistica{
    @Override
    public ITransporte asignarTransporte() {
        return new Barco();
    }
}
