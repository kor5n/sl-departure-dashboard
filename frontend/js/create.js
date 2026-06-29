const inputField = document.querySelector("#station-input");
const dropDown = document.querySelector("#dropdown");
const linesField = document.querySelector("#dashboard-filter")
const form = document.querySelector("form");
let savedStation;
let lines = [];

inputField.addEventListener("input", async () =>{
    if (inputField.value.length > 4){
        //make an API call to search for stations
        const resp = await fetch("/api/search-stop/" + inputField.value);
        if (resp.ok){
            const data = await resp.json();
            dropDown.replaceChildren();
            for(let i = 0 ; i<data.length; i++){
                const dropElement = document.createElement("li");
                dropElement.classList.add("dropdown-op");
                dropElement.textContent = data[i].split("|")[0];
                dropDown.appendChild(dropElement);
                dropElement.addEventListener("click", async () =>{
                    savedStation = data[i];
                    const call = await fetch("/api/departures/" + data[i].split("|")[1]);
                    lines = [];
                    if (call.ok){
                        const dep = await call.json();
                        for (let i=0; i<dep.length;i++){
                            if(lines.includes(dep[i]["route"])){
                                continue;
                            }
                            lines.push(dep[i]["route"]);
                        }
                        linesField.replaceChildren();
                        for (let i = 0; i<lines.length; i++){
                            const lineBox = document.createElement("li");
                            lineBox.classList.add("lines-list");

                            const checkbox = document.createElement("input");
                            checkbox.setAttribute("type", "checkbox");
                            checkbox.setAttribute("name", lines[i]);
                            lineBox.appendChild(checkbox);

                            const label = document.createElement("label");
                            label.classList.add("line-num");
                            label.textContent = lines[i];
                            lineBox.appendChild(label);

                            linesField.appendChild(lineBox);
                        }
                        
                    }else{
                        const error = await call.text();
                        window.alert(error);
                    }
                });
            }
        }else{
            window.alert(await resp.text);
        }
    }  
});

form.addEventListener("submit", async (e)=>{
    e.preventDefault();

    const formData = new FormData(form);

    const data = Object.fromEntries(formData.entries());

    let keys = Object.keys(data)

    let week = ["mon", "tue", "wed", "thu", "fri", "sat", "sun"];
    let times = [];
    for (let i = 0; i<week.length;i++){
        if (keys.includes(week[i])){
            times.push(week[i] + "-" +data[week[i]+"1"]+"-"+data[week[i]+"2"]);
        }
        delete data[week[i]+"1"];
        delete data[week[i]+"2"];
        delete data[week[i]];
    }

    console.log(data);
    keys = Object.keys(data);

    console.log(times);
    const req = await fetch("/api/add-dashboard/", {
        method: "POST",
        headers: {
            "Content-Type":"application/json; charset=UTF-8"
        },
        body: JSON.stringify({
            name:savedStation.split("|")[0],
            stopid: savedStation.split("|")[1],
            routes: keys,
            times: times
        })
    });

    const resp = await req.text();
    if (req.ok){
        window.location.assign("/");
    }else{
        window.alert(resp);
    }
});