public class Cliente {

    public static void main(String[] args) {
        Database db = Database.getInstance();
        System.out.println(db.toString());

        Database other = Database.getInstance();
        System.out.println(other.toString());

        //VEO QUE SIEMPRE ME DEVUELVE LA MISMA INSTANCIA
    }
}
