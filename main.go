package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>Omoide Map</title>
    <style>
        #map { height: 90vh; width: 90%; }
        h1 { font-size: 16px; }
    </style>
    
</head>
<body>
    <h1>思い出マップ</h1>
    <div id="map"></div>
    <script>
        function initMap() {
            const center = { lat: 35.6910, lng: 139.6994 };
            const map = new google.maps.Map(document.getElementById("map"), {
                zoom: 12,
                center: center,
            });

            // 初期マーカー
            new google.maps.Marker({
                position: center,
                map: map,
                title: "ここに思い出！"
            });

            // 追加マーカーを管理する配列
            const markers = [];

            // クリックで新しいマーカーを追加
            map.addListener("click", (e) => {
                const latLng = e.latLng;
                
                // 例: ユーザーに名前と詳細を入力させる
                const title = prompt("思い出の名前を入力してください");
                if (!title) return;

                const description = prompt("詳細を入力してください");

                const marker = new google.maps.Marker({
                    position: latLng,
                    map: map,
                    title: "新しい思い出"
                });

                // InfoWindow 作成
                const infoWindow = new google.maps.InfoWindow({
                    content: "<strong>" + title + "</strong><br>" + description
                });

                // マーカークリックで InfoWindow を表示
                marker.addListener("click", () => {
                    infoWindow.open(map, marker);
                });

                // 配列に保存
                markers.push(marker);

                // マーカーをダブルクリックで削除
                marker.addListener("dblclick", () => {
                    marker.setMap(null); // 地図から削除
                    const index = markers.indexOf(marker);
                    if (index > -1) markers.splice(index, 1); // 配列からも削除
                });
            });
        }
    </script>
    <script async defer
        src="https://maps.googleapis.com/maps/api/js?key=AIzaSyAhpNDJo_fdZX8UwdyvaFF-DK58Bijnguw&callback=initMap">
    </script>
</body>
</html>`))
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
