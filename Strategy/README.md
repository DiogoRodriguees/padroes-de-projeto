## Strategy Pattern

O Strategy é um padrão de projeto comportamental que permite que você defina uma família de algoritmos, coloque-os em classes separadas, e faça os objetos deles intercambiáveis.

O padrão Strategy é muito comum no código Java e pode ser usado em situações em que o seu código terá muitas variações de algoritmo e condições, gerando uma corrente de IFs.

## Como implementar 

1) Na classe contexto, identifique um algoritmo que é sujeito a frequentes mudanças. Pode ser também uma condicional enorme que seleciona e executa uma variante do mesmo algoritmo durante a execução do programa.

2) Declare a interface da estratégia comum para todas as variantes do algoritmo.

3) Um por um, extraia todos os algoritmos para suas próprias classes. Elas devem todas implementar a interface estratégia.

4) Na classe contexto, adicione um campo para armazenar uma referência a um objeto estratégia. Forneça um setter para substituir valores daquele campo. O contexto deve trabalhar com o objeto estratégia somente através da interface estratégia. O contexto pode definir uma interface que deixa a estratégia acessar seus dados.

5) Os Clientes do contexto devem associá-lo com uma estratégia apropriada que coincide com a maneira que esperam que o contexto atue em seu trabalho primário.

## Prós 

- Você pode trocar algoritmos usados dentro de um objeto durante a execução.
- Você pode isolar os detalhes de implementação de um algoritmo do código que usa ele.
- Você pode substituir a herança por composição.
- Princípio aberto/fechado. Você pode introduzir novas estratégias sem mudar o contexto.

## Contras

- Se você só tem um par de algoritmos e eles raramente mudam, não há motivo real para deixar o programa mais complicado com novas classes e interfaces que vêm junto com o padrão.
- Os Clientes devem estar cientes das diferenças entre as estratégias para serem capazes de selecionar a adequada.
- Muitas linguagens de programação modernas tem suporte do tipo funcional que permite que você implemente diferentes versões de um algoritmo dentro de um conjunto de funções anônimas. Então você poderia usar essas funções exatamente como se estivesse usando objetos estratégia, mas sem inchar seu código com classes e interfaces adicionais.

## Compilar e Executar

Para compilar e executar o exemplo, basta digitar o comando:

```bash
javac -d . src/Main.java
```
Assim o exemplo será compilado, e então poderá ser executado digitando:

```bash
java Main
```

## Referências

- [Refactoring Guru](https://refactoring.guru/pt-br/design-patterns/strategy)
- [Grupo SoftPlan](https://www.softplan.com.br/tech-writers/descomplicando-o-strategy/)