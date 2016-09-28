public class Fruta {
  public static void main(String[] args) {

    String fruta = "cereza";
    switch (fruta){
        case "platano":
            System.out.println("Color amarillo");
            break;
        case "naranja":
            System.out.println("Color naranja");
            break;
        case "fresa":
            System.out.println("Color rojo");
            break;
        case "pera":
            System.out.println("Color verde");
            break;
        case "cereza":
            System.out.println("Color rojo");
            break;
        default:
            System.out.println("Color indefinido");
    }
  }
}
