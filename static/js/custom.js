var counter = 0

document.getElementById("counter").innerText = counter

function increaseCounter() {
  console.log('it is happening')
  counter++
  document.getElementById("counter").innerText = counter
}
function decreaseCounter() {
  console.log('it is happening')
  counter--
  document.getElementById("counter").innerText = counter
}