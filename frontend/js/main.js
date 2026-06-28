const containerDashboard = document.querySelector("#dashboard-container");

async function LoadDashboards() {
    const response = await fetch("/api/dashboards");
    if (response.ok){
        const data = await response.json();
        containerDashboard.replaceChildren();

        for (let i=0; i<data.length;i++){
            const toggleContainer = document.createElement("div");
            toggleContainer.className = "toggle-container navigation";
            toggleContainer.setAttribute("href", "/dashboard?index="+i.toString());

            const dashboardTitle = document.createElement("h3")
            dashboardTitle.textContent = data[i].name;

            const filterLabel = document.createElement("span")
            filterLabel.classList.add("filter-container");
            const filterp = document.createElement("p");
            filterp.classList.add("filter-label");
            filterp.textContent = data[i].routes.join(",");
            filterLabel.appendChild(filterp);

            toggleContainer.appendChild(dashboardTitle);
            toggleContainer.appendChild(filterLabel);
            containerDashboard.appendChild(toggleContainer);
        }
    }else{
        const error = await response.text();
        window.alert(error);
    }
}

LoadDashboards();
