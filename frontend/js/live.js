import { CalculateTime } from "./modules.js";

const blobContainer = document.querySelector("#dashboard-blobs");
let time = new Date()
const week = ["sun","mon", "tue", "wed", "thu", "fri", "sat"];

const LoadBoards = async () =>{
    const resp = await fetch("/api/dashboards");
    if (resp.ok){
        const data = await resp.json();
        for(let i =0;i<data.length; i++){
            console.log("day");
            if(!data[i]["times"].some(item => item.startsWith(week[time.getDay()]))){
                continue;
            }
            const index = data[i]["times"].findIndex(item => item.startsWith(week[time.getDay()]));
            const startTime = data[i]["times"][index].split("-")[1];
            const endTime = data[i]["times"][index].split("-")[2];
            console.log("hours");
            if(!(time.getHours() >= +startTime.split(":")[0] && time.getHours() <= +endTime.split(":")[0])){
                continue;
            }

            console.log("minutes");

            if (time.getHours() == +startTime.split(":")[0] || time.getHours() == +endTime.split(":")[0]){
                if(!(time.getMinutes() >= +startTime.split(":")[1] && time.getMinutes() <= +endTime.split(":")[1])){
                    continue;
                }
            }
            
            const stopName = data[i]["name"];
            const stopId = data[i]["stopid"];
            const req = await fetch("/api/departures/"+stopId);
            if(req.ok){
                const deps = await req.json();
                blobContainer.replaceChildren();
                for (let j = 0; j<deps.length;j++){

                    if (!data[i]["routes"].includes(deps[j]["route"])){
                        continue;
                    }
                    const blob = document.createElement("div");
                    blob.classList.add("blob");

                    const stopTitle = document.createElement("h2");
                    stopTitle.classList.add("stop-title");
                    stopTitle.textContent = stopName;
                    blob.appendChild(stopTitle);

                    const nextStop = document.createElement("div");
                    nextStop.classList.add("next-stop");

                    const depTitle = document.createElement("span");
                    depTitle.textContent = deps[j]["route"] + " " + deps[j]["direction"];
                    nextStop.appendChild(depTitle);

                    const br = document.createElement("br");
                    nextStop.appendChild(br);

                    const timeLabel = document.createElement("span");
                    const em = document.createElement("em");
                    em.textContent = CalculateTime(deps[j]["departure"]);
                    timeLabel.appendChild(em);
                    nextStop.appendChild(timeLabel);

                    blob.appendChild(nextStop);
                    blobContainer.appendChild(blob);
                    break;
                }
            }else{
                const error = await req.json();
                window.alert(error);
            }
        }
    }else{
        const error = await resp.json();
        window.alert(error);
    }
};

LoadBoards();
setInterval(() => 
    LoadBoards()
, 100000);