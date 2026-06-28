document.addEventListener("click", (e)=>{
    const { target } = e;
    if (!target.matches("nav a") && !target.className.includes("navigation")){
        return;
    }
    e.preventDefault();
    urlRoute();
});

const urlRoutes = {
    404: {
        "html": "html/404.html",
        "js": "",
        "title": "404"
    },
    "/":{
        "html": "html/home.html",
        "js": "js/main.js",
        "title": "SL Avgång"
    },
    "/create":{
        "html":"html/create.html",
        "js": "js/create.js",
        "title": "Create SL Dashboard"
    },//Add index field to dashboard page somehow
    "/dashboard":{
        "html":"html/dashboard.html",
        "js":"js/dashboard.js",
        "title": "{Name} Dashboard"
    }
};

const urlRoute = (event) =>{
    event = event || window.event;
    event.preventDefault();
    const href = event.target.href || event.target.getAttribute("href");
    window.history.pushState({}, "", href);
    urlLocationHandler();
};

const urlLocationHandler = async () => {
    let location = window.location.pathname;
    if (location.length === 0){
        location = "/";
    }

    const route =  urlRoutes[location] || urlRoutes[404];
    const html = await fetch(route.html).then((response) => 
    response.text());
    document.querySelector("#app").innerHTML = html;
    document.title = route.title;

    document.getElementById("page-script")?.remove();

    if (route.js){
        const script = document.createElement("script");
        script.src = route.js
        script.id = "page-script"
        document.body.appendChild(script)
    }
 
};

window.onpopstate = urlLocationHandler;
window.route = urlRoute;

urlLocationHandler();