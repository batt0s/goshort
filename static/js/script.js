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
// Sends request to api (https://goshrt.herokuapp.com/api/shorten)
// Sends a JSON data with URL and gets a JSON data with shorten url or error
// When function gets the JSON Response, shows result area and change innerhtml to response
// Runs when user clickes the shorten button
function shortenUrl(userid) {
    let urlToShortObj = document.querySelector("#url");
    let customShortObj = document.querySelector("#custom");
    let shortResultDiv = document.querySelector("#shortDiv");
    let shortResult = document.querySelector("#short");

    var urlToShort = urlToShortObj.value.trim()

    var isCustom
    if (customShortObj.value.trim().length === 0) {
        isCustom = false
    } else {
        isCustom = true
        var customShort = customShortObj.value.trim()
    }

    let xhr = new XMLHttpRequest();
    let url = "api/v2/shorten";

    xhr.open("POST", url, true);

    xhr.setRequestHeader("ContentType","application/json");

    // When response comes check the status code and show results
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var response = JSON.parse(this.response);
            shortResultDiv.className = "px-8";
            var short = response["URL"];
            shortResult.innerHTML = short;
        } else if (xhr.readyState === 4 && xhr.status === 400 || xhr.status === 502) {
            var response = JSON.parse(this.response);
            shortResultDiv.className = "px-8";
            shortResult.innerHTML = response["error"];
        }
    }

    var data
    if (isCustom) {
        data = {
            "url": urlToShort,
            "is_custom": isCustom,
            "custom": customShort
        }
    } else {
        data = {"url": urlToShort}
    }
    if (userid.length !== 0) {
        data["author"] = userid;
        var today = new Date();
        var date = today.getFullYear()+"-"+(today.getMonth()+1)+"-"+today.getDate();
        var time = today.getHours()+":"+today.getMinutes()+":"+today.getSeconds();
        var datetime = date + " " + time
        var addToTable = "<tr class='bg-white border-b dark:bg-gray-800 dark:border-gray-700'><th scope='row' class='px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap'>" +
                        short +
                        "</th><td class='px-6 py-4'>" +
                        urlToShort +
                        "</td><td class='px-6 py-4'>0</td><td class='px-6 py-4'>" + 
                        datetime +
                        "</td></tr>";
        table = document.querySelector("#shorteneds");
        table.innerHTML += addToTable;
    }

    json = JSON.stringify(data)
    xhr.send(json);

}

function copyShort() {
    let shortObj = document.querySelector("#short");
    var short = shortObj.innerHTML.trim();
    navigator.clipboard.writeText(short);
    alert("Copied short url "+short+" to clipboard.");
}

// A function that gets Original URL of a Shortened URL
// Sends request to API (https://goshrt.herokuapp.com/api/getOrigin) 
// Sends Short URL with JSON and gets Original URL or an error with JSON
// When function gets the JSON Response, shows result area and change innerhtml to response
// Runs when user clickes the get origin button
function getOrigin() {
    let shortUrl = document.querySelector("#shorturl");
    let originResult = document.querySelector("#origin");
    let originResultDiv = document.querySelector("#originDiv");

    let xhr = new XMLHttpRequest();
    let url = "api/v2/getOrigin";

    xhr.open("POST", url, true);

    xhr.setRequestHeader("ContentType","application/json");

    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var response = JSON.parse(this.response);
            originResultDiv.className = "px-8";
            originResult.innerHTML = response["URL"];
        } else if (xhr.readyState === 4 && xhr.status === 400 || xhr.status === 502) {
            var response = JSON.parse(this.response);
            originResult.className = "px-8";
            originResult.innerHTML = response["error"];
        }
    }

    var data = JSON.stringify({"url":shortUrl.value});

    xhr.send(data);
}

function copyOrigin() {
    let originObj = document.querySelector("#origin");
    var origin = originObj.innerHTML.trim();
    navigator.clipboard.writeText(origin);
    alert("Copied original url "+origin+" to clipboard.");
}