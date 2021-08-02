public class LogisticaAerea extends Logistica{
    @Override
    public ITransporte asignarTransporte() {
        return new Avion();
    }
}
