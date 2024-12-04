package strategy;

public class CalculadoraPreco {
    private Desconto desconto;

    public CalculadoraPreco(Desconto desconto) {
        this.desconto = desconto;
    }

    public double calcularPrecoFinal(double preco) {
        return desconto.aplicarDesconto(preco);
    }
}