public class Cliente {
    private static Logistica l;

    public static void main(String[] args) {
        //levanto una config que me diga que el tipo de environment
        // y depende de eso es que platform levanto

        l = Logistica.getConfig("aerea");
        l.asignarTransporte().realizarEnvio();
    }
}
