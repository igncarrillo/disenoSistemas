public abstract class Logistica {
    public abstract ITransporte asignarTransporte();

    public static Logistica getConfig(String s){
        switch (s){
            case "aerea":
                return new LogisticaAerea();
            case "maritima":
                return new LogisticaMaritima();
            default:
                throw new RuntimeException("Tipo de plataforma invalida");
        }
    }
}

