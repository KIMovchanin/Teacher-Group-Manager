const tbody = document.getElementById('student-tbody');
const form = document.getElementById('student-form');

form.addEventListener('submit', async (event) => {
    event.preventDefault();

    const first_name = form.querySelector('input[name="first_name"]').value;
    const last_name = form.querySelector('input[name="last_name"]').value;

    const response = await fetch('/students', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ first_name: first_name, last_name: last_name}),
    });

    if (!response.ok) return;
    
    logAction(`Добавлен ученик: ${first_name} ${last_name}`);

    form.reset();
    loadStudents();
})

async function loadStudents() {
    const response = await fetch('/students');
    const students = await response.json();
    console.log('Ученики:', students);

    tbody.innerHTML = '';
    for (const student of students) {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${student.id}</td>
            <td>${student.first_name}</td>
            <td>${student.last_name}</td>
            <td><button class="delete-btn">Удалить</button></td>
        `;
        tbody.appendChild(row);

        const deleteBtn = row.querySelector('button');
        deleteBtn.addEventListener('click', async () => {
            await fetch(`/students/${student.id}`, { method: 'DELETE'});

            logAction(`Удалён ученик: #${student.id}`);

            loadStudents();
        })
    }
}

loadStudents();