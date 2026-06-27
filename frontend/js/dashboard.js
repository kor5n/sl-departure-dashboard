const dashboardContainer = document.querySelector("#dashboard-container");

const LoadDashboard = async () =>{
    const params = new URLSearchParams(window.location.search);
    const index = params.get("index");
    const response = await fetch("http://127.0.0.1:8080/api/dashboard/"+index);
    if (response.ok){
        const data = await response.json();
        document.title = data.name + " Dashboard";
        document.querySelector("#stop-name").textContent = data.name;
        UpdateDashboard(data["stopid"]);
        setInterval(() => UpdateDashboard(data["stopid"]), 20000);

    }else{
        const error = await response.text();
        window.alert(error);
    }
};

const UpdateDashboard = async (id) => {
    const resp = await fetch("http://127.0.0.1:8080/api/departures/"+ +id);
    if (resp.ok){
        const data = await resp.json();
        dashboardContainer.replaceChildren();

        for (let i = 0; i<data.length; i++){
            if (data[i]["canceled"]){
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

const CalculateTime = (departure) => {
    const currentDate = new Date;
    const depTime = departure.split("T")[1].split(":");

    //const curMonth = months.indexOf(currentDate[1]);
    //const curDate = currentDate.split(" ")[3] + "-" + curMonth + "-" + currentDate[2];
    //const depDate = departure.split("T")[0];

    const depSeconds = +depTime[0] * 3600 + +depTime[1]*60 + +depTime[2];
    const curSeconds = currentDate.getHours() * 3600 +currentDate.getMinutes() * 60 +currentDate.getSeconds();
    const timeDiff = depSeconds - curSeconds;
    const minutes = Math.round(timeDiff/60);

    return minutes.toString() + " minutes";
};

LoadDashboard();