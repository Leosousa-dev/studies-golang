# **Verificador de palíndromo**
Este é um algoritmo simples escrito em Golang para verificar se uma palavra ou frase é um palíndromo, ou seja, pode ser lida da mesma forma tanto da esquerda para a direita quanto da direita para a esquerda.
## **📜 Passo a passo**
Aqui está o passo a passo para o algoritmo:
- **[ ✔️ ] Receber a palavra que o usuário digitar.**
- **[ ✔️ ] Armazenar essa palavra em uma variável.**
- **[ ✔️ ] Inverter a palavra, fazendo ela ficar de trás para frente.**
- **[  ] Verificar se a palavra digitada é igual à palavra invertida (de trás para frente).**
## **✨ Melhorias**
**Aqui estão algumas sugestões de melhorias que você pode considerar para aprimorar o algoritmo:**

- ### **✨ Tratamento de entrada:**
    Atualmente, o algoritmo considera apenas palavras ou frases com letras minúsculas e sem caracteres especiais. Você pode adicionar etapas adicionais de tratamento de entrada para lidar com diferentes casos, como letras maiúsculas, caracteres especiais e espaços em branco excessivos.
- ### **✨ Otimização:**
    O algoritmo atualmente utiliza um loop para inverter a palavra. No entanto, você pode explorar métodos mais eficientes para realizar essa tarefa, como o uso de slices ou a conversão para um array de bytes.
- ### **✨ Considerar acentos:**
    Se desejar tratar corretamente palavras com acentos, você pode adicionar etapas para remover ou substituir os acentos antes de realizar a verificação do palíndromo. Existem bibliotecas em Golang, como a "golang.org/x/text", que podem auxiliar nesse processo.

Essas são apenas algumas sugestões de melhorias. Sinta-se à vontade para explorar outras funcionalidades e aprimoramentos que possam tornar o algoritmo mais robusto e eficiente.




