function initializeMaps() {
    const mapContainers = document.querySelectorAll('.map-container');
    const geocoder = new google.maps.Geocoder();

    for (let i = 0; i < mapContainers.length; i++) {
        const mapContainer = mapContainers[i];
        const country = mapContainer.getAttribute('data-country');
        const map = new google.maps.Map(mapContainer.querySelector('.map'), {
            zoom: 5
        });

    geocoder.geocode({ 'address': country }, function(results, status) {
    if (status === 'OK') {
        const countryLocation = results[0].geometry.location;
        map.setCenter(countryLocation);

        const cities = mapContainer.querySelectorAll('.city');
        for (let j = 0; j < cities.length; j++) {
            const cityElement = cities[j];
            const cityName = cityElement.getAttribute('data-city');

            geocoder.geocode({ 'address': cityName + ', ' + country }, function(results, status) {
                if (status === 'OK') {
                    const cityLocation = results[0].geometry.location;
                    new google.maps.Marker({
                        position: cityLocation,
                        map: map,
                        title: cityName
                    });
                } else {
                    console.error('Geocode was not successful for the following reason: ' + status);
                }
            });
        }
    } else {
        console.error('Geocode was not successful for the following reason: ' + status);
    }
});
}
}