const inputField = document.querySelector("#station-input");
const dropDown = document.querySelector("#dropdown");

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
            }
        }
    }  
});