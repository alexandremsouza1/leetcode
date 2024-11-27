function isAnagram(s, t) {
    if (s.length !== t.length) {
        return false;
    }

    let countS = [];
    let countT = [];

    for (let i = 0; i < s.length; i++) {
        countS[s[i]] = (countS[s[i]] || 0) + 1;
    }

    for (let i = 0; i < t.length; i++) {
        countT[t[i]] = (countT[t[i]] || 0) + 1;
    }

    for (let char in countS) {
        if (countS[char] !== countT[char]) {
            return false;
        }
    }

    return true;
}

console.log(isAnagram("hello", "world"))