// document.addEventListener('DOMContentLoaded', () => {
//     const themeCss = document.getElementById("linkStyle");
//     const storedTheme = localStorage.getItem('theme');
//     if (storedTheme) {
//         themeCss.href = storedTheme;
//     }
//     const toggleTheme = document.getElementById("toggleTheme");
//     toggleTheme.addEventListener('click', () => {
//         if (themeCss.href.includes('dark')) {
//             themeCss.href = '/static/css/style.css';
//             toggleTheme.innerText = "Switch to Dark Mode";
//         } else {
//             themeCss.href = '/static/css/dark-style.css';
//             toggleTheme.innerText = "Switch to Light Mode";
//         }
//         localStorage.setItem('theme', themeCss.href)
//     });
// })


// A function to shorten the given url
// Sends request to api (/api/v3/shorten)
// Sends a JSON data with URL and gets a JSON data with shorten url or error
// When function gets the JSON Response, shows result area and change inner-html to response
// Runs when user clicks the shorten button
function shortenUrl() {
    let urlToShortObj = document.querySelector("#url");
    let customShortObj = document.querySelector("#custom");
    let shortResultDiv = document.querySelector("#shortDiv");
    let shortResult = document.querySelector("#short");

    const urlToShort = urlToShortObj.value.trim();
    const customShort = customShortObj.value.trim();

    let xhr = new XMLHttpRequest();
    let url = "api/v3/shorten";

    xhr.open("POST", url, true);

    xhr.setRequestHeader("Content-Type","application/json");

    // When response comes check the status code and show results
    xhr.onreadystatechange = function () {
        let response;
        if (xhr.readyState === 4 && xhr.status === 200) {
            response = JSON.parse(this.response);
            shortResultDiv.className = "px-8";
            shortResult.innerHTML = response["URL"];
        } else if (xhr.readyState === 4 && xhr.status === 400 || xhr.status === 502) {
            response = JSON.parse(this.response);
            shortResultDiv.className = "px-8";
            shortResult.innerHTML = response["error"];
        }
    }

    let data = {
        "url": urlToShort,
        "custom": customShort
    }

    let json = JSON.stringify(data)
    xhr.send(json);

}

function copyShort() {
    let shortObj = document.querySelector("#short");
    const short = shortObj.innerHTML.trim();
    navigator.clipboard.writeText(short);
    alert("Copied short url "+short+" to clipboard.");
}

// A function that gets Original URL of a Shortened URL
// Sends request to API (/api/v3/getOrigin)
// Sends Short URL with JSON and gets Original URL or an error with JSON
// When function gets the JSON Response, shows result area and change inner-html to response
// Runs when user clicks the get origin button
function getOrigin() {
    let shortUrl = document.querySelector("#shorturl");
    let originResult = document.querySelector("#origin");
    let originResultDiv = document.querySelector("#originDiv");

    let xhr = new XMLHttpRequest();
    let url = "api/v3/getOrigin";

    xhr.open("POST", url, true);

    xhr.setRequestHeader("Content-Type","application/json");

    xhr.onreadystatechange = function() {
        let response;
        if (xhr.readyState === 4 && xhr.status === 200) {
            response = JSON.parse(this.response);
            originResultDiv.className = "px-8";
            originResult.innerHTML = response["URL"];
        } else if (xhr.readyState === 4 && xhr.status === 400 || xhr.status === 502) {
            response = JSON.parse(this.response);
            originResult.className = "px-8";
            originResult.innerHTML = response["error"];
        }
    }

    const data = JSON.stringify({
        "url": shortUrl.value
    });

    xhr.send(data);
}

function copyOrigin() {
    let originObj = document.querySelector("#origin");
    const origin = originObj.innerHTML.trim();
    navigator.clipboard.writeText(origin);
    alert("Copied original url "+origin+" to clipboard.");
}