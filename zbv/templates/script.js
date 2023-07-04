document.getElementById('forgotPasswordForm').addEventListener('submit', function(event) {
  event.preventDefault();
  var email = document.getElementById('email').value;

  // Prepare the data to be sent in the request body
  var data = {
    email: email
  };

  // Send the POST request using fetch to the backend.
  fetch('/reset-password', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  })
  .then(response => response.json())
  .then(result => {
    console.log(result); // Log the response from the backend server
    // Clear the input field
    document.getElementById('email').value = '';
    // respond with a success message to the user.
    var successMessage = document.createElement('p');
    successMessage.textContent = 'Reset password email sent successfully!';
    successMessage.style.color = 'green';
    this.appendChild(successMessage);
  })
  .catch(error => {
    console.error(error); // Log any error that occurred in our logging that you made
  });
});
