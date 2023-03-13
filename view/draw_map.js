var map = L.map('map').setView([37.440888,138.869906], 14);
L.tileLayer('https://mt1.google.com/vt/lyrs=r&x={x}&y={y}&z={z}', {
    attribution: "<a href='https://developers.google.com/maps/documentation' target='_blank'>Google Map</a>"
}).addTo(map);

DrawLine("悠久山線")

// 再描画
function ReDrawMap(center,line){
    // 地図を消す
    if(map){
        map.remove();
        map=null;
    }
    // 再描画
    map = L.map('map').setView(center, 14);
    L.tileLayer('https://mt1.google.com/vt/lyrs=r&x={x}&y={y}&z={z}', {
        attribution: "<a href='https://developers.google.com/maps/documentation' target='_blank'>Google Map</a>"
    }).addTo(map);
    DrawLine(line)
}
// 路線を地図に書く
function DrawLine(line){
    fetch('http://localhost:8000/line/'+line)
        .then(response => {
            return response.json()
        }).then(res => {
            // 線(FeatureCollection)オブジェクト生成
            var line = GetFeatures(res.elements)
            // 線を描画
            L.geoJSON(line,{
                onEachFeature: function onEachFeature(
                    feature,
                    layer
                ){
                    if(feature.properties && feature.properties.popupContent){
                        layer.bindPopup(feature.properties.popupContent);
                    }
                }
            }).addTo(map)
            // バス停情報の配列生成
            var busstops = GetBusstops(res.elements,res.operationalPoints)
            // バス停の表示
            var popup = L.popup();
            for(let busstop of busstops){
                L.marker(busstop.coordinate).addTo(map).on("click",function(e){
                    popup.setLatLng(e.latlng).setContent(busstop.name).openOn(map);
                })
            }
        })
        .catch(error => {
            console.error(error);
        });
}    