package strategy;

public class DescontoFixo implements Desconto {
    private double valor;

    public DescontoFixo(double valor) {
        this.valor = valor;
    }

    @Override
    public double aplicarDesconto(double preco) {
        return preco - valor;
    }
}