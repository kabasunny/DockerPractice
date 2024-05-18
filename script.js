document.getElementById('pingButton').addEventListener('click', function() {
    fetch('http://localhost:8080/ping')
        .then(response => response.json())
        .then(data => {
            document.getElementById('result').innerText = JSON.stringify(data);
        })
        .catch(error => console.error('Error:', error));
});
