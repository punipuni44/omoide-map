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
    <style>#map { height: 100vh; width: 100%; }</style>
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

            // クリックで新しいマーカーを追加
            map.addListener("click", (e) => {
                const latLng = e.latLng;
                new google.maps.Marker({
                    position: latLng,
                    map: map,
                    title: "新しい思い出"
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
