const inputField = document.querySelector("#station-input");
const dropDown = document.querySelector("#dropdown");
const linesField = document.querySelector("#dashboard-filter")

inputField.addEventListener("input", async () =>{
    if (inputField.value.length > 4){
        //make an API call to search for stations
        const resp = await fetch("http://127.0.0.1:8080/api/search-stop/" + inputField.value);
        if (resp.ok){
            const data = await resp.json();
            dropDown.replaceChildren();
            for(let i = 0 ; i<data.length; i++){
                const dropElement = document.createElement("li");
                dropElement.classList.add("dropdown-op");
                dropElement.textContent = data[i].split("|")[0];
                dropDown.appendChild(dropElement);
                dropElement.addEventListener("click", async () =>{
                    const call = await fetch("http://127.0.0.1:8080/api/get-lines/" + data[i].split("|")[0]);
                    if (call.ok){
                        const lines = await call.json();
                        linesField.replaceChildren();
                        for (let i = 0; i<lines.length; i++){
                            const lineBox = document.createElement("li");
                            lineBox.classList.add("lines-list");

                            const checkbox = document.createElement("input");
                            checkbox.setAttribute("type", "checkbox");
                            checkbox.setAttribute("name", "line");
                            lineBox.appendChild(checkbox);

                            const label = document.createElement("label");
                            label.classList.add("line-num");
                            label.textContent = lines[i];
                            lineBox.appendChild(label);

                            linesField.appendChild(lineBox);
                        }
                        
                    }else{
                        window.alert(await call.text);
                    }
                });
            }
        }else{
            window.alert(await resp.text);
        }
    }  
});