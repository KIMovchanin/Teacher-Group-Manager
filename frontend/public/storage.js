function logAction(text)  {
    const actions = getActions();
    actions.unshift({ text: text, time: new Date().toLocaleString() });
    localStorage.setItem('recentActions', JSON.stringify(actions.slice(0, 10)));
}

function getActions() {
    const raw = localStorage.getItem('recentActions');
    return raw ? JSON.parse(raw) : [];
}