const dashboardContainer = document.querySelector("#dashboard-container");

async function LoadDashboards() {
    const response = await fetch("http://127.0.0.1:8080/api/dashboards");
    if (response.ok){
        const data = await response.json();
        dashboardContainer.replaceChildren();

        for (let i=0; i<data.length;i++){
            console.log(data);
            const toggleContainer = document.createElement("div");
            toggleContainer.classList.add("toggle-container");
            toggleContainer.setAttribute("onclick", "window.location.assign('dashboard.html')");

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
            dashboardContainer.appendChild(toggleContainer);
        }
    }else{
        const error = await response.text();
        window.alert(error);
    }
}

LoadDashboards()