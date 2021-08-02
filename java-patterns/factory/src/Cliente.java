public class Cliente {

    public static void main(String[] args) {
        //levanto una config que me diga que el tipo de environment
        // y depende de eso es que platform levanto

        //Logistica la = new LogisticaAerea();

        Logistica l = new LogisticaMaritima();
        ITransporte t = l.asignarTransporte();
        t.realizarEnvio();
    }
}
