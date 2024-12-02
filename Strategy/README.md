## Strategy Pattern

O Strategy é um padrão de projeto comportamental que permite que você defina uma família de algoritmos, coloque-os em classes separadas, e faça os objetos deles intercambiáveis.

O padrão Strategy é muito comum no código Java e pode ser usado em situações em que o seu código terá muitas variações de algoritmo e condições, gerando uma corrente de IFs.

## Como implementar 

1) Na classe contexto, identifique um algoritmo que é sujeito a frequentes mudanças. Pode ser também uma condicional enorme que seleciona e executa uma variante do mesmo algoritmo durante a execução do programa.

2) Declare a interface da estratégia comum para todas as variantes do algoritmo.

3) Um por um, extraia todos os algoritmos para suas próprias classes. Elas devem todas implementar a interface estratégia.

4) Na classe contexto, adicione um campo para armazenar uma referência a um objeto estratégia. Forneça um setter para substituir valores daquele campo. O contexto deve trabalhar com o objeto estratégia somente através da interface estratégia. O contexto pode definir uma interface que deixa a estratégia acessar seus dados.

5) Os Clientes do contexto devem associá-lo com uma estratégia apropriada que coincide com a maneira que esperam que o contexto atue em seu trabalho primário.
