function searchBarFunction(artists){
    const searchBar = document.getElementById('SearchBar');
    const suggestionsList = document.getElementById('suggestions');
    searchBar.addEventListener('input', function(event) {
        const searchTerm = event.target.value.toLowerCase();
        const matchingArtists = artists.filter(artist => artist.name.toLowerCase().includes(searchTerm));
        const matchingFirstAlbum   = artists.filter(artist => artist.firstAlbum.toString().toLowerCase().includes(searchTerm));
        const matchingCreationDate = artists.filter(artist => artist.creationDate.toString().toLowerCase().includes(searchTerm));
        suggestionsList.textContent = ''; // Clear previous suggestions
        matchingArtists.forEach(artist => {
            const option = document.createElement('a');
            option.innerText = `${artist.name} - artist/band`
            option.setAttribute('href', `artistid/`+artist.id)
            suggestionsList.appendChild(option);
        });
        artists.forEach(artist => { // Members
            artist.members.forEach(member => {
                if (!member.toLowerCase().includes(searchTerm)) {return}
                const option = document.createElement('a');
                option.innerText = `${member} - member of ${artist.name}`
                option.setAttribute('href', `artistid/`+artist.id)
                suggestionsList.appendChild(option);
            })
        });
        artists.forEach(artist => { // Locations
            artist.LocationData.locations.forEach(location => {
                if (!location.toLowerCase().includes(searchTerm)) {return}
                const option = document.createElement('a');
                option.innerText = `${location} - location for the artist: ${artist.name}`
                option.setAttribute('href', `artistid/`+artist.id)
                suggestionsList.appendChild(option);
            })
        });
        matchingFirstAlbum.forEach(artist => {
            const option = document.createElement('a');
            option.innerText = artist.name + ' ' + `${artist.firstAlbum} - First Album`
            option.setAttribute('href', `artistid/`+artist.id)
            suggestionsList.appendChild(option);
        });
        matchingCreationDate.forEach(artist => {
            const option = document.createElement('a');
            option.innerText = artist.name + ' ' + `${artist.creationDate} - Creation Date`
            option.setAttribute('href', `artistid/`+artist.id)
            suggestionsList.appendChild(option);
        });
        elements4boxshadowHover(document.querySelectorAll("null"),suggestionsList,0,"opacity 1s ease")
        if (searchTerm == "") {
            suggestions.style.opacity = '0';
            setTimeout(() => {
                suggestions.style.display = 'none';
            }, 500)
        } else {
            suggestions.style.display = 'block';
            suggestions.style.opacity = '1';
        }
    });
}

function getRandomColor(){
    var letters = '0123456789ABCDEF'; var color = '#';
    for (var i = 0; i < 6; i++){color += letters[Math.floor(Math.random() * 16)];}
    return color;
}

function elements4boxshadowHover(elems2reset,elem2apply,inset=1,extraTransition="") {
    elems2reset.forEach(element => {
        element.style.boxShadow="0px 0px 0px";
    });
    elem2apply.style.transition=`box-shadow 0.5s linear 0s ${(extraTransition)? ", " + extraTransition : ""}`
    elem2apply.style.boxShadow=`${(inset)? "inset" : ""} 0px 0px 16px 0.01px ${getRandomColor()}`
}

// window.addEventListener("load", function() {
//     document.querySelectorAll("*").forEach(function(elem1) {
//         var suggestions = document.getElementById('suggestions');
//         elem1.addEventListener("focus", function() {
//             if (elem1 == document.querySelector("#SearchBar")) {
//                 suggestions.style.display = 'block';
//                 suggestions.style.opacity = '1';console.log(elem1.tagName)
//             } else {
//                console.log(elem1.tagName)
//                 suggestions.style.opacity = '0';
//                 setTimeout(() => {
//                     suggestions.style.display = 'none';
//                 }, 500); // Wait for 2 seconds before hiding completely 
//             }
//         })
//     })
// });
