function getCoordinate(c){
    return [c.latitude,c.longitude]
}
function getCoordinates(cs){
    var coordinates = []
    for(let c of cs){
        coordinates.push(getCoordinate(c))
    }
    if(coordinates.length==1){
        coordinates.push(coordinates[0])
    }
    return coordinates
}
function getFeature(element){
    return {
        "type": "Feature",
        "geometry": {
            "type": "LineString",
            "coordinates": getCoordinates(element.coordinates)
        }
    }
}
function GetFeatures(elements){
    var features = []
    for(let e of elements){
        features.push(getFeature(e))
    }
    return {
        "type": "FeatureCollection",
        "features": features
    }
}
function GetBusstops(elements, ops){
    var busstops = []
    for(let op of ops){
        for(let e of elements){
            if(op.elementID == e.id){
                coordinate = getCoordinate(e.coordinates[0])
                busstops.push({
                    "coordinate": [coordinate[1],coordinate[0]],
                    "name": op.name
                })
            }
        }
    }
    return busstops
}