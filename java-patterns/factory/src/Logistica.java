public abstract class Logistica {
    public abstract ITransporte asignarTransporte();

    public static ITransporte getConfig(String s){
        switch (s){
            case "aerea":
                return new LogisticaAerea().asignarTransporte();
            case "maritima":
                return new LogisticaMaritima().asignarTransporte();
            default:
                throw new RuntimeException("Tipo de plataforma invalida");
        }
    }
}

