## Decorator Pattern
A descrição original do Padrão Decorator é: "O Padrão Decorator anexa responsabilidades adicionais a um objeto dinamicamente. Os decoradores fornecem uma alternativa flexível de subclasse para estender a funcionalidade".

O padrão Decorator usa a herança apenas para ter uma correspondência de tipo e não para obter o comportamento. Assim, quando compõe-se um decorador com um componente, adiciona-se um novo comportamento, nota-se que estamos adquirindo um novo comportamento e não herdando de alguma superclasse. Isso nos dá muito mais flexibilidade para compor mais objetos sem alterar uma linha de código.

## Implementação:
1. Determine a funcionalidade básica que pode ser estendida dinamicamente. Identifique o comportamento central que será decorado.
2. Defina uma interface ou classe abstrata que declare os métodos que serão implementados pelos componentes concretos e decoradores.
3. Crie uma ou mais classes concretas que implementem a funcionalidade básica definida na interface ou classe abstrata.
4. Defina uma classe abstrata decoradora que implemente a interface do componente e contenha uma referência a um objeto do tipo componente.
5. Crie classes concretas que herdem da classe decoradora e adicionem funcionalidades específicas.
6. No código use a composição para adicionar funcionalidades ao componente base ao encapsulá-lo com decoradores.
7. Certifique-se de que o comportamento decorado está funcionando conforme o esperado e que novos decoradores podem ser adicionados sem alterar os componentes existentes.

## Prós
- Você pode estender o comportamento de um objeto sem fazer um nova subclasse.
- Você pode adicionar ou remover responsabilidades de um objeto no momento da execução.
- Você pode combinar diversos comportamentos ao envolver o objeto com múltiplos decoradores.
- Princípio de responsabilidade única. Você pode dividir uma classe monolítica que implementa muitas possíveis variantes de um comportamento em diversas classes menores.

## Contras
- É difícil remover um invólucro de uma pilha de invólucros.
- É difícil implementar um decorador de tal maneira que seu comportamento não dependa da ordem do pilha de decoradores.
- A configuração inicial do código de camadas pode ficar bastante feia.