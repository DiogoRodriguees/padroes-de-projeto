package strategy;

public class SemDesconto implements Desconto {
    @Override
    public double aplicarDesconto(double preco) {
        return preco;
    }
}