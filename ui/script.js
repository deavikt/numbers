let url = 'http://localhost:8081/data';
let numberInputForm = document.getElementById("number_input_form");
let numberInputField = document.getElementById("number_input_field");
let numbersSumField = document.getElementById("numbers_sum_field");
let countButton = document.getElementById("count_button");

numberInputForm.addEventListener('submit', (event) => {
    sendNumber(numberInputField.value);
});

countButton.addEventListener('click', (event) => {
    getNumbersSum();
})

function sendNumber(number) {
    let data = {value: Number(number)};

    fetch(
        url,
        {
            method: 'POST',
            body: JSON.stringify(data)
        }
    )
    .then(data => console.log('Success:', data))
    .catch(error => console.error('Error:', error));
}

function getNumbersSum() {
    fetch(url)
    .then(response => response.json())
    .then(data => {
        numbersSumField.textContent = data.sum;
        console.log('Success:', data)
    })
    .catch(error => console.error('Error:', error));
}
