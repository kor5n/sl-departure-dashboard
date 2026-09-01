export const CalculateTime = (departure) => {
    const currentDate = new Date;
    const depTime = departure.split("T")[1].split(":");

    //const curMonth = months.indexOf(currentDate[1]);
    //const curDate = currentDate.split(" ")[3] + "-" + curMonth + "-" + currentDate[2];
    //const depDate = departure.split("T")[0];

    const depSeconds = +depTime[0] * 3600 + +depTime[1]*60 + +depTime[2];
    const curSeconds = currentDate.getHours() * 3600 +currentDate.getMinutes() * 60 +currentDate.getSeconds();
    const timeDiff = depSeconds - curSeconds;
    let minutes = Math.round(timeDiff/60);

    if (minutes < 0){
	minutes += 1440;
    }

    return minutes.toString() + " minutes";
};
