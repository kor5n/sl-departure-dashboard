import { CalculateTime } from "./modules.js";

const dashboardContainer = document.querySelector("#dashboard-container");
const params = new URLSearchParams(window.location.search);

if (params.size == 0){
    window.alert("No dashboard specified");
    window.history.back();
}
const index = params.get("index");
let dashboard;

const LoadDashboard = async () =>{
    const response = await fetch("/api/dashboard/"+index);
    if (response.ok){
        dashboard = await response.json();
        document.title = dashboard.name + " Dashboard";
        document.querySelector("#stop-name").textContent = dashboard.name;
        UpdateDashboard(dashboard["stopid"]);
        setInterval(() => UpdateDashboard(dashboard["stopid"]), 20000);

    }else{
        const error = await response.text();
        window.alert(error);
        window.history.back();
    }
};

const UpdateDashboard = async () => {
    const resp = await fetch("/api/departures/"+ +dashboard.stopid);
    if (resp.ok){
        const data = await resp.json();
        dashboardContainer.replaceChildren();

        for (let i = 0; i<data.length; i++){
            if (data[i]["canceled"]){
                continue;
            }

            if (!dashboard.routes.includes(data[i]["route"])){
                continue;
            }
            const departure = document.createElement("div");
            departure.classList.add("departure");

            //const image = document.createElement("img");
            //image.src = "dunno";
            //departure.appendChild(image);

            const destContainer = document.createElement("span");
            destContainer.classList.add("dest-container");
            const lineLabel = document.createElement("p");
            lineLabel.classList.add("line-label");
            lineLabel.textContent = data[i]["route"];
            const em = document.createElement("em");
            const destLabel = document.createElement("p");
            destLabel.classList.add("dest-label");
            destLabel.textContent = data[i]["direction"];
            em.appendChild(destLabel);
            destContainer.appendChild(lineLabel);
            destContainer.appendChild(em);
            departure.appendChild(destContainer);

            const strong = document.createElement("strong");
            const depTime = document.createElement("p");
            depTime.classList.add("dep-time");
            depTime.textContent = CalculateTime(data[i]["departure"]);
            strong.appendChild(depTime)
            departure.appendChild(strong);

            if (data[i]["alerts"].length > 0){
                let alert = "";
                for (let j = 0; j<data[i]["alerts"].length;j++){
                    alert += data[i]["alerts"][j];
                }
                const alerts = document.createElement("small");
                alerts.classList.add("alerts");
                alerts.textContent = alert;
                departure.appendChild(alerts);
            }

            dashboardContainer.appendChild(departure);
        }
    }else{
        const error = await resp.json();
        window.alert(error);
    }
};

const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Dec"];

LoadDashboard();

document.querySelector("#no").addEventListener("click", ()=>{
    document.querySelector("#confirm-panel").style.display = "none";
});

document.querySelector("#rm-btn").addEventListener("click", ()=>{
    document.querySelector("#confirm-panel").style.display = "flex";
});

document.querySelector("#yes").addEventListener("click", async () =>{
    const req = await fetch("/api/delete-dashboard/" + index, {method: "DELETE"});
    const msg = await req.text();
    window.alert(msg);
    if (req.ok){
        window.location.assign("/");
    }
});
