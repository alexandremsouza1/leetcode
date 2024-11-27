function isBoolean(temp) {
    open = 0;
    for (let i = 0; i < temp.length; i++) {
        if (temp[i] === "(") {
            open++;
        }
        if (temp[i] === ")") {
            open--;
        }
        if(open < 0){
            return false;
        }
    }
    return open === 0;
}
/*
Escreva um método que recebe uma string e retorna uma booleana.
A booleana deve ser false se a string contiver parênteses inseridos de forma incorreta e true caso contrário.

Exemplo de strings válidas:
 - ""
 - "()"
 - "(sldjfs)"
 - "(())"
 - "()()"
 - "Hoje"
 - "Hoje (dia 7 de novembro (2024))"

Exemplo de strings inválidas:
 - "("
 - "(((("
 - "(())("

 */