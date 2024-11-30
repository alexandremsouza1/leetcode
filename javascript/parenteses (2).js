/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    let arr = [0, 0, 0]; // Contadores para parênteses, chaves e colchetes
    const open = ["(", "{", "["];
    const close = [")", "}", "]"];
  
    for (let index = 0; index < s.length; index++) {
        let char = s[index];
        
        // Verifica se é uma abertura
        let openIndex = open.indexOf(char);
        if (openIndex !== -1) { // Se estiver no array 'open'
            arr[openIndex] += 1;
            continue;
        }
        
        // Verifica se é um fechamento
        let closeIndex = close.indexOf(char);
        if (closeIndex !== -1) { // Se estiver no array 'close'
            arr[closeIndex] -= 1;
            
            // Verifica se há um fechamento sem abertura correspondente
            if (arr[closeIndex] < 0) {
                return false;
            }
        }
    }
    
    // Verifica se todos os contadores estão equilibrados
    return arr.every(e => e === 0);
};

var isValid = function(s) {
    const stack = [];
    const matchingPairs = {
        ')': '(',
        '}': '{',
        ']': '['
    };

    for (let char of s) {
        // Se for uma abertura, empilha
        if (['(', '{', '['].includes(char)) {
            stack.push(char);
        }
        // Se for um fechamento, verifica correspondência com o topo da pilha
        else if ([')', '}', ']'].includes(char)) {
            if (stack.pop() !== matchingPairs[char]) {
                return false; // Fechamento não corresponde à última abertura
            }
        }
    }

    // Se a pilha estiver vazia no final, é válida
    return stack.length === 0;
};

// console.log(isValid("()"));      // true
// console.log(isValid("({[]})"));  // true
// console.log(isValid("({[})"));   // false
// console.log(isValid("())))"));   // false
console.log(isValid("([)]"));   // false
