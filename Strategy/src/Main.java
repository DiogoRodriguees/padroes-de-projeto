import strategy.*;

public class Main {
    public static void main(String[] args) {
        CalculadoraDePreco calculadora;

        calculadora = new CalculadoraDePreco(new SemDesconto());
        System.out.println("Preço final (Sem Desconto): " + calculadora.calcularPrecoFinal(100));

        calculadora = new CalculadoraDePreco(new DescontoPercentual(10));
        System.out.println("Preço final (Desconto Percentual): " + calculadora.calcularPrecoFinal(100));

        calculadora = new CalculadoraDePreco(new DescontoFixo(15));
        System.out.println("Preço final (Desconto Fixo): " + calculadora.calcularPrecoFinal(100));
    }
}
