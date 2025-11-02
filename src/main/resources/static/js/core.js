function colorizeLetters() {
    const letters = document.getElementById("raweceek")?.getElementsByTagName("span");
    if (!letters) return;

    for (let i = 0, color = 0; i < letters.length; i++) {
        const char = letters[i].textContent;
        if (!char || char.trim() === "") continue;

        letters[i].className = `color-${color++ % 6}`
    }
}

colorizeLetters();
