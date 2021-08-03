public class Cliente {

    public static void main(String[] args) {
        //levanto una config que me diga que el tipo de environment
        // y depende de eso es que platform levanto

        ITransporte t = Logistica.getConfig("aerea");
        t.realizarEnvio();
    }
}
