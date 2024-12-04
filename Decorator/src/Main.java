import decorator.*;

public class DecoratorExample {
    public static void main(String[] args) {
        // Pizza básica
        Pizza pizza = new PlainPizza();
        System.out.println(pizza.getDescription() + " $" + pizza.getCost());

        // Adicionando queijo
        pizza = new Cheese(pizza);
        System.out.println(pizza.getDescription() + " $" + pizza.getCost());

        // Adicionando pepperoni
        pizza = new Pepperoni(pizza);
        System.out.println(pizza.getDescription() + " $" + pizza.getCost());

        // Adicionando azeitonas
        pizza = new Olives(pizza);
        System.out.println(pizza.getDescription() + " $" + pizza.getCost());
    }
}