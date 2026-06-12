const recentList = document.getElementById("recent-list")

const list = getActions();

if (list.length === 0){
    const li = document.createElement("li");
    li.textContent = "Действий пока нет.";
    recentList.appendChild(li);
} else {
    let index = 0;
    for (const action of list) {
        const li = document.createElement("li");
        li.textContent = `${action.time} - ${action.text}`;
        li.style.animationDelay = `${index*0.1}s`;
        recentList.appendChild(li);

        index++;
    }
};

