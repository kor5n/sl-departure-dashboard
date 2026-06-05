import pandas as pd
import json

routes = pd.read_csv("routes.txt")
trips = pd.read_csv("trips.txt")
stop_times = pd.read_csv("stop_times.txt")
stops = pd.read_csv("stops.txt")

sl_agency_id = 505000000000000001

routes = routes[
    routes["agency_id"] == sl_agency_id
]

# Keep only useful columns
routes = routes[[
    "route_id",
    "route_short_name",
    "route_long_name",
    "route_type"
]]

route_ids = set(routes["route_id"])

trips = trips[
    trips["route_id"].isin(route_ids)
]

trips = trips[[
    "trip_id",
    "route_id"
]]

trip_ids = set(trips["trip_id"])

stop_times = stop_times[
    stop_times["trip_id"].isin(trip_ids)
]

stop_times = stop_times[[
    "trip_id",
    "stop_id"
]]

stop_routes = stop_times.merge(
    trips,
    on="trip_id",
    how="inner"
)

stop_routes = stop_routes[[
    "stop_id",
    "route_id"
]].drop_duplicates()

stops = stops[[
    "stop_id",
    "stop_name"
]]

stop_routes = stop_routes.merge(
    stops,
    on="stop_id",
    how="left"
)

stop_routes = stop_routes.merge(
    routes,
    on="route_id",
    how="left"
)

result = {}

for _, row in stop_routes.iterrows():

    stop_id = str(row["stop_id"])

    if stop_id not in result:
        result[stop_id] = {
            "stop_name": row["stop_name"],
            "routes": []
        }

    result[stop_id]["routes"].append({
        "route_id": str(row["route_id"]),
        "line": row["route_short_name"],
        "name": row["route_long_name"],
        "type": int(row["route_type"])
    })

with open("stop_info.json", "w", encoding="utf-8") as f:
    json.dump(result, f, ensure_ascii=False)

print("Saved stop_info.json")