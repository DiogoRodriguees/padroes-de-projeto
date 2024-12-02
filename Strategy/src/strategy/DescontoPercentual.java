package strategy;

public class DescontoPercentual implements Desconto {
    private double percentual;

    public DescontoPercentual(double percentual) {
        this.percentual = percentual;
    }

    @Override
    public double aplicarDesconto(double preco) {
        return preco - (preco * percentual / 100);
    }
}