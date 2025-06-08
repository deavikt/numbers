const url = 'http://localhost:8000/data';

let numberInputForm = document.getElementById("number_input_form");
let numberInputField = document.getElementById("number_input_field");
let numbersSumField = document.getElementById("numbers_sum_field");
let deleteNumbersButton = document.getElementById("delete_numbers_button");

document.addEventListener('DOMContentLoaded', function() {
    getNumbersSum();
});

numberInputForm.addEventListener('submit', (event) => {
    sendNumber(numberInputField.value);
});

deleteNumbersButton.addEventListener('click', (event) => {
    deleteNumbers();
    getNumbersSum();
});

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
        console.log('Success:', data);
        numbersSumField.textContent = data.sum;
    })
    .catch(error => console.error('Error:', error));
}

function deleteNumbers() {
    fetch(url,{method: 'DELETE'})
    .then(data => console.log('Success:', data))
    .catch(error => console.error('Error:', error));
}
