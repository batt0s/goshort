// A function to show the custom shorten input area.
// Hides shorten button and shows a custom url input div (custom shorten input area with custom shorten button)
// Runs when user check the checkbox above the shorten button (Do you want to assign a cutsom short url?)
function showCustomBox() {
    // get checkbox, customboxdiv (custom short url input area) and shortenbutton
    var checkbox = document.getElementById("customcheck");
    var customboxdiv = document.getElementById("customdiv");
    var shortenbutton = document.getElementById("shortenbtn");
    // if checkbox is checked then hide default shorten button and show the custom url input div
    if (checkbox.checked == true) {
        customboxdiv.style.display = "block";
        shortenbutton.style.display = "none";

    } else {
        customboxdiv.style.display = "none";
        shortenbutton.style.display = "initial";
    }
}

// A function to shorten the given url
// Sends request to api (https://goshrt.herokuapp.com/api/shorten)
// Sends a JSON data with URL and gets a JSON data with shorten url or error
// When function gets the JSON Response, shows result area and change innerhtml to response
// Runs when user clickes the shorten button
function shortenUrl() {
    let urlToShort = document.querySelector("#url");
    let shortResult = document.querySelector("#short");

    let xhr = new XMLHttpRequest();
    let url = "api/latest/shorten";

    xhr.open("POST", url, true);

    xhr.setRequestHeader("ContentType","application/json");

    // When response comes check the status code and show results
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var response = JSON.parse(this.response);
            shortResult.style.display = "block";
            shortResult.innerHTML = response["URL"];
        } else if (xhr.readyState === 4 && xhr.status === 400 || xhr.status === 502) {
            var response = JSON.parse(this.response);
            shortResult.style.display = "block";
            shortResult.innerHTML = response["error"];
        }
    }

    var data = JSON.stringify({"url":urlToShort.value});

    xhr.send(data);

}

// A function that shortens the given URL with a custom short url
// Almost same with shortenUrl()
// https://goshrt.herokuapp.com/api/shorten/custom
// Sends url and custom with json and gets the short url or error with json
// When function gets the JSON Response, shows result area and change innerhtml to response
// Runs when user clickes the custom shorten button
function customShorten() {
    let urlToShort = document.querySelector("#url");
    let customShortUrl = document.querySelector("#customurl");
    let shortResult = document.querySelector("#short");

    let xhr = new XMLHttpRequest();
    let url = "api/latest/customShorten";
    xhr.open("POST", url, true);

    xhr.setRequestHeader("ContentType","application/json");

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var response = JSON.parse(this.response);
            shortResult.style.display = "block";
            shortResult.innerHTML = response["URL"];
        } else if (xhr.readyState === 4 && xhr.status === 400 || xhr.status === 502) {
            var response = JSON.parse(this.response);
            shortResult.style.display = "block";
            shortResult.innerHTML = response["error"];
        }
    }

    var data = JSON.stringify({"url":urlToShort.value, "custom":customShortUrl.value});

    xhr.send(data);
}

// A function that gets Original URL of a Shortened URL
// Sends request to API (https://goshrt.herokuapp.com/api/getOrigin) 
// Sends Short URL with JSON and gets Original URL or an error with JSON
// When function gets the JSON Response, shows result area and change innerhtml to response
// Runs when user clickes the get origin button
function getOrigin() {
    let shortUrl = document.querySelector("#shorturl");
    let originResult = document.querySelector("#origin");

    let xhr = new XMLHttpRequest();
    let url = "api/latest/getOrigin";

    xhr.open("POST", url, true);

    xhr.setRequestHeader("ContentType","application/json");

    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var response = JSON.parse(this.response);
            originResult.style.display = "block";
            originResult.innerHTML = response["URL"];
        } else if (xhr.readyState === 4 && xhr.status === 400 || xhr.status === 502) {
            var response = JSON.parse(this.response);
            originResult.style.display = "block";
            originResult.innerHTML = response["error"];
        }
    }

    var data = JSON.stringify({"url":shortUrl.value});

    xhr.send(data);
}