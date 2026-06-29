const group_form = document.getElementById('group-form');
const group_tbody = document.getElementById('group-tbody');


group_form.addEventListener('submit', async (event) => {
    event.preventDefault();

    const name = group_form.querySelector('input[name="name"]').value;

    const response = await fetch('/groups', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json'},
        body: JSON.stringify({name}),
    });

    if (!response.ok) return;

    logAction(`Добавлена группа: ${name}`);

    group_form.reset();
    loadGroups();
})


async function loadGroups() {
    const response = await fetch('/groups');
    const groups = await response.json();
    console.log('Группы:', groups);

    group_tbody.innerHTML = "";
    for (const group of groups) {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${group.id}</td>
            <td>${group.name}</td>
            <td><button class="delete-btn">Удалить</button></td>
        `;
        group_tbody.appendChild(row);

        const deleteBtn = row.querySelector('button');
        deleteBtn.addEventListener('click', async () => {
            await fetch(`/groups/${group.id}`, {method: 'DELETE'});

            logAction(`Удалена группа: #${group.id}`);

            loadGroups();
        })
    }
}

loadGroups();
